package auth

import (
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"user-service/internal/email"
	"user-service/internal/jwt"
	dtoPack "user-service/internal/model/dto"
	service2 "user-service/internal/service"
)

type service struct {
	tokenService jwt.TokenService
	codeService  service2.CodeService
	userService  service2.UserService
	emailSender  email.Sender
}

func NewService(tokenService jwt.TokenService, userService service2.UserService, codeService service2.CodeService, emailSender email.Sender) *service {
	return &service{
		codeService:  codeService,
		tokenService: tokenService,
		userService:  userService,
		emailSender:  emailSender,
	}
}

func (s *service) SignInUser(dto dtoPack.SignInUserRequestDto) (dtoPack.SignInUserResponseDto, error) {
	user, err := s.userService.Create(dto.Email, dto.Password, dto.PhoneNumber)
	if err != nil {
		return dtoPack.SignInUserResponseDto{}, err
	}

	code, _ := s.codeService.Create(dto.Email)

	s.emailSender.SendEmailVerification(user, code)

	tok, err := s.tokenService.CreateToken(user, "self")
	if err != nil {
		return dtoPack.SignInUserResponseDto{}, err
	}

	return dtoPack.SignInUserResponseDto{
		Id:    user.Id,
		Token: tok,
	}, err
}

func (s *service) LoginUser(dto dtoPack.LoginUserRequestDto) (dtoPack.LoginUserResponseDto, error) {
	user, err := s.userService.GetByEmail(dto.Email)
	if err != nil {
		return dtoPack.LoginUserResponseDto{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)) != nil {
		return dtoPack.LoginUserResponseDto{}, fmt.Errorf("incorrect password")
	}

	tok, err := s.tokenService.CreateToken(user, "self")
	if err != nil {
		return dtoPack.LoginUserResponseDto{}, err
	}

	return dtoPack.LoginUserResponseDto{Token: tok}, nil
}

func (s *service) AuthenticateUserBySideApp(userId uuid.UUID, issuerId string) (string, error) {
	if issuerId != "device_app_secret_id" {
		return "", fmt.Errorf("unknown issuer: %s", issuerId)
	}

	user, err := s.userService.Get(userId)
	if err != nil {
		return "", err
	}

	tok, err := s.tokenService.CreateToken(user, issuerId)
	if err != nil {
		return "", err
	}

	return tok, err
}

func (s *service) VerifyUserEmail(codeId uuid.UUID) error {
	mail, err := s.codeService.GetEmailById(codeId)
	if err != nil {
		return err
	}

	return s.userService.SetEmailVerified(mail)
}

func (s *service) VerifyUserToken(token string) (jwt2.Claims, error) {
	parsedToken, err := s.tokenService.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	userId, err := parsedToken.Claims.GetSubject()
	if err != nil {
		return nil, err
	}

	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	if _, err = s.userService.Get(userUuid); err != nil {
		return nil, err
	}

	return parsedToken.Claims, nil
}
