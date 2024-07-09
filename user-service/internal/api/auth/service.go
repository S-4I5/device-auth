package auth

import (
	"user-service/internal/service"
	"user-service/pkg/auth_v1"
)

type AuthenticationServiceImplementation struct {
	auth_v1.UnimplementedAuthV1Server
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *AuthenticationServiceImplementation {
	return &AuthenticationServiceImplementation{
		authService: authService,
	}
}
