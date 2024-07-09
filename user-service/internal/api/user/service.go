package user

import (
	"user-service/internal/service"
	"user-service/pkg/user_v1"
)

type UserV1ServiceImplementation struct {
	user_v1.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *UserV1ServiceImplementation {
	return &UserV1ServiceImplementation{
		userService: userService,
	}
}
