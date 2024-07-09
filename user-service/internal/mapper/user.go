package mapper

import (
	"user-service/internal/model/dto"
	"user-service/internal/model/entity"
	"user-service/pkg/user_v1"
)

func UserToUserV1(user entity.User) *user_v1.User {
	return &user_v1.User{
		Id:              user.Id.String(),
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		IsEmailVerified: user.IsEmailVerified,
	}
}

func UserToUserDto(user entity.User) *dto.UserDto {
	return &dto.UserDto{
		Id:              user.Id,
		Email:           user.Email,
		PhoneNumber:     user.PhoneNumber,
		IsEmailVerified: user.IsEmailVerified,
	}
}
