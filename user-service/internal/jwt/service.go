package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
	"user-service/internal/model/entity"
)

type TokenParser interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type TokenCreator interface {
	CreateToken(user entity.User, issuer string) (string, error)
}

type TokenService interface {
	TokenParser
	TokenCreator
}

type tokenService struct {
	secret string
}

func NewService(secret string) *tokenService {
	return &tokenService{secret: secret}
}

func (s *tokenService) CreateToken(user entity.User, issuer string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   user.Id.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tkn, _ := token.SignedString([]byte(s.secret))

	return tkn, nil
}

func (s *tokenService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})
}

func GetSubjectIdFromToken(token jwt.Token) (uuid.UUID, error) {
	targetId, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	subjectId, err := uuid.Parse(targetId)
	if err != nil {
		return uuid.Nil, err
	}

	return subjectId, nil
}
