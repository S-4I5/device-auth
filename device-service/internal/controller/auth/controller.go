package auth

import (
	"device-service/internal/err"
	"device-service/internal/jwt"
	"device-service/internal/service"
)

type controller struct {
	authService   service.AuthService
	tokenVerifier jwt.TokenVerifier
	errorHandler  err.ErrorHandler
}

func NewController(authService service.AuthService, tokenVerifier jwt.TokenVerifier, errorHandler err.ErrorHandler) *controller {
	return &controller{
		authService:   authService,
		tokenVerifier: tokenVerifier,
		errorHandler:  errorHandler,
	}
}
