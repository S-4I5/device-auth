package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"strings"
	"user-service/pkg/auth_v1"
)

func (a *AuthenticationServiceImplementation) AuthenticateUser(ctx context.Context, req *auth_v1.AuthenticateUserRequest) (*auth_v1.AuthenticateUserResponse, error) {

	fmt.Println("XD")

	userUuid, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, err
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, err
	}

	issuerToken, _ := strings.CutPrefix(values[0], "Bearer ")

	tok, err := a.authService.AuthenticateUserBySideApp(userUuid, issuerToken)
	if err != nil {
		return nil, err
	}

	fmt.Println(tok)

	return &auth_v1.AuthenticateUserResponse{Token: tok}, nil
}
