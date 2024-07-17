package auth

import (
	"user-service/internal/api/scope"
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

func (*AuthenticationServiceImplementation) GetMethodScopeMap() scope.MethodMap {
	return scope.NewMapBasedMethodMap(map[string]scope.Scope{
		auth_v1.AuthV1_AuthenticateUser_FullMethodName: scope.AuthV1Authenticate,
		auth_v1.AuthV1_ValidateToken_FullMethodName:    scope.AuthV1Validation,
	})
}
