package auth

import (
	"context"
	"user-service/pkg/auth_v1"
)

func (a *AuthenticationServiceImplementation) ValidateToken(ctx context.Context, req *auth_v1.ValidateTokenRequest) (*auth_v1.ValidateTokenResponse, error) {

	claims, err := a.authService.VerifyUserToken(req.Token)
	if err != nil {
		return nil, err
	}

	userId, _ := claims.GetSubject()

	return &auth_v1.ValidateTokenResponse{UserId: userId}, nil
}
