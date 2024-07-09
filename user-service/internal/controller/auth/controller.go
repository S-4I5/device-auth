package auth

import (
	errPack "user-service/internal/err"
	"user-service/internal/service"
)

type controller struct {
	authService service.AuthService
	errHandler  errPack.ErrorHandler
}

func NewController(authService service.AuthService, errHandler errPack.ErrorHandler) *controller {
	return &controller{authService: authService, errHandler: errHandler}
}
