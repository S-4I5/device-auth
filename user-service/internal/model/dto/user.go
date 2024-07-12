package dto

import "github.com/google/uuid"

type SignUpUserRequestDto struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type SignUpUserResponseDto struct {
	Id    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}

type LoginUserRequestDto struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
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
