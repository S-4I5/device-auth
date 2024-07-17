package auth

import (
	errPack "user-service/internal/httperr"
	"user-service/internal/service"
)

type controller struct {
	authService service.AuthService
	errHandler  errPack.ErrorHandler
}

func NewController(authService service.AuthService, errHandler errPack.ErrorHandler) *controller {
	return &controller{authService: authService, errHandler: errHandler}
}
