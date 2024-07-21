package auth

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user-service/internal/api/interceptor"
	"user-service/pkg/auth_v1"
)

func (a *AuthenticationServiceImplementation) AuthenticateUser(ctx context.Context, req *auth_v1.AuthenticateUserRequest) (*auth_v1.AuthenticateUserResponse, error) {
	userUuid, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	client, err := interceptor.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tok, err := a.authService.AuthenticateUserBySideApp(userUuid, client.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &auth_v1.AuthenticateUserResponse{Token: tok}, nil
}
