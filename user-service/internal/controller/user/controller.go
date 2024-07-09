package user

import (
	def "user-service/internal/controller"
	errPack "user-service/internal/err"
	"user-service/internal/jwt"
	"user-service/internal/service"
)

var _ def.UserController = (*controller)(nil)

type controller struct {
	userService service.UserService
	errHandler  errPack.ErrorHandler
	tokenParser jwt.TokenParser
}

func NewController(userService service.UserService, errHandler errPack.ErrorHandler, tokenParser jwt.TokenParser) *controller {
	return &controller{userService: userService, errHandler: errHandler, tokenParser: tokenParser}
}
