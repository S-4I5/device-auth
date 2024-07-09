package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"user-service/internal/mapper"
	"user-service/pkg/user_v1"
)

func (u *UserV1ServiceImplementation) GetUserById(ctx context.Context, req *user_v1.GetUserByIdRequest) (*user_v1.GetUserByIdResponse, error) {
	userUuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	user, err := u.userService.Get(userUuid)
	if err != nil {
		return nil, err
	}

	return &user_v1.GetUserByIdResponse{User: mapper.UserToUserV1(user)}, nil
}

func (u *UserV1ServiceImplementation) GetUserByPhoneNumber(ctx context.Context, req *user_v1.GetUserByPhoneNumberRequest) (*user_v1.GetUserByPhoneNumberResponse, error) {
	user, err := u.userService.GetByPhoneNumber(req.PhoneNumber)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user_v1.GetUserByPhoneNumberResponse{User: mapper.UserToUserV1(user)}, nil
}
