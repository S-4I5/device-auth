package user

import (
	"user-service/internal/api/scope"
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

func (*UserV1ServiceImplementation) GetMethodScopeMap() scope.MethodMap {
	return scope.NewMapBasedMethodMap(map[string]scope.Scope{
		user_v1.UserV1_GetUserById_FullMethodName:          scope.UserV1GetUserById,
		user_v1.UserV1_GetUserByPhoneNumber_FullMethodName: scope.UserV1GetUserByPhoneNumber,
	})
}
