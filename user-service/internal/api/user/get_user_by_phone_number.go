package user

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user-service/internal/mapper"
	"user-service/pkg/user_v1"
)

func (u *UserV1ServiceImplementation) GetUserByPhoneNumber(ctx context.Context, req *user_v1.GetUserByPhoneNumberRequest) (*user_v1.GetUserByPhoneNumberResponse, error) {
	user, err := u.userService.GetByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &user_v1.GetUserByPhoneNumberResponse{User: mapper.UserToUserV1(user)}, nil
}
