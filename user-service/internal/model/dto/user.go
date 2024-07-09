package dto

import "github.com/google/uuid"

type SignInUserRequestDto struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type SignInUserResponseDto struct {
	Id    uuid.UUID `yaml:"id"`
	Token string    `yaml:"token"`
}

type LoginUserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponseDto struct {
	Token string `yaml:"token"`
}

type UserDto struct {
	Id              uuid.UUID `json:"id"`
	PhoneNumber     string    `json:"phoneNumber"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"isEmailVerified"`
}
