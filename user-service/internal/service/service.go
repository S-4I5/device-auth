package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"user-service/internal/model/dto"
	"user-service/internal/model/entity"
)

type UserService interface {
	Create(email, password string, phoneNumber string) (entity.User, error)
	Get(id uuid.UUID) (entity.User, error)
	SetEmailVerified(email string) error
	GetByEmail(email string) (entity.User, error)
	GetByPhoneNumber(phoneNumber string) (entity.User, error)
}

type CodeService interface {
	Create(email string) (uuid.UUID, error)
	GetEmailById(id uuid.UUID) (string, error)
	Delete(id uuid.UUID) error
}

type AuthService interface {
	SignInUser(dto dto.SignInUserRequestDto) (dto.SignInUserResponseDto, error)
	LoginUser(dto dto.LoginUserRequestDto) (dto.LoginUserResponseDto, error)
	VerifyUserEmail(codeId uuid.UUID) error
	AuthenticateUserBySideApp(userId uuid.UUID, issuerId string) (string, error)
	VerifyUserToken(token string) (jwt.Claims, error)
}
