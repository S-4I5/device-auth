package auth

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user-service/pkg/auth_v1"
)

func (a *AuthenticationServiceImplementation) ValidateToken(ctx context.Context, req *auth_v1.ValidateTokenRequest) (*auth_v1.ValidateTokenResponse, error) {

	claims, err := a.authService.VerifyUserToken(req.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userId, _ := claims.GetSubject()

	return &auth_v1.ValidateTokenResponse{UserId: userId}, nil
}
