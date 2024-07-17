package auth

import (
	"device-service/internal/httperr"
	"device-service/internal/jwt"
	"device-service/internal/service"
)

type controller struct {
	authService   service.AuthService
	tokenVerifier jwt.TokenVerifier
	errorHandler  httperr.ErrorHandler
}

func NewController(authService service.AuthService, tokenVerifier jwt.TokenVerifier, errorHandler httperr.ErrorHandler) *controller {
	return &controller{
		authService:   authService,
		tokenVerifier: tokenVerifier,
		errorHandler:  errorHandler,
	}
}
