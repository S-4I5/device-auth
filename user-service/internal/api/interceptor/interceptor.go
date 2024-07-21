package interceptor

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
	"user-service/internal/api/scope"
	"user-service/internal/model/entity"
	"user-service/internal/service"
)

var (
	errMissingMetadata                  = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidCredentials               = status.Errorf(codes.Unauthenticated, "invalid credentials")
	errNoClientInContext                = status.Errorf(codes.Internal, "no client in context")
	errCannotCastValueInContextToClient = status.Errorf(codes.Internal, "cannot cast value in context to client")
)

const (
	userKey     = "user"
	headerBasic = "Basic "
)

type AuthInterceptor interface {
	GetInterceptor() grpc.UnaryServerInterceptor
}

type interceptor struct {
	clientService service.ClientService
	validator     scope.Validator
}

func NewAuthInterceptor(clientService service.ClientService, validator scope.Validator) *interceptor {
	return &interceptor{
		clientService: clientService,
		validator:     validator,
	}
}

func (i *interceptor) GetInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errMissingMetadata
		}

		user, ok := i.getUserFromAuth(md["authorization"])
		if !ok {
			return nil, errInvalidCredentials
		}

		err := i.validator.IsAllowed(info.FullMethod, user.Scopes)
		if err != nil {
			return nil, err
		}

		m, err := handler(context.WithValue(ctx, userKey, user), req)

		return m, err
	}
}

func (i *interceptor) getUserFromAuth(headers []string) (entity.Client, bool) {
	if len(headers) == 0 {
		return entity.Client{}, false
	}

	authHeader := headers[0]
	if !strings.HasPrefix(authHeader, headerBasic) {
		return entity.Client{}, false
	}

	encodedCredentials := strings.TrimPrefix(authHeader, headerBasic)
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return entity.Client{}, false
	}

	parts := strings.SplitN(string(decodedCredentials), ":", 2)
	if len(parts) != 2 {
		return entity.Client{}, false
	}

	username, password := parts[0], parts[1]

	user, err := i.clientService.Get(username)
	if err != nil {
		return entity.Client{}, false
	}

	if user.Secret != password {
		return entity.Client{}, false
	}

	return user, true
}

func GetUserFromContext(ctx context.Context) (entity.Client, error) {
	rowClient := ctx.Value(userKey)

	if rowClient == nil {
		return entity.Client{}, errNoClientInContext
	}

	client, ok := rowClient.(entity.Client)
	if !ok {
		return entity.Client{}, errCannotCastValueInContextToClient
	}
	return client, nil
}
