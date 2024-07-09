package jwt

import (
	"device-service/internal/model/entity"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strings"
	"time"
)

type TokenCreator interface {
	CreateDeviceToken(device entity.Device) (string, error)
	CreateCodeToken(code entity.ActivationCode) (string, error)
}

type TokenVerifier interface {
	Verify(token string) (*jwt.Token, error)
	VerifyRow(token string) (*jwt.Token, error)
}

type TokenService interface {
	TokenCreator
	TokenVerifier
}

type serivce struct {
	secret string
}

func NewService(secret string) *serivce {
	return &serivce{secret: secret}
}

const (
	issuerSelf          = "device-service"
	audCodeVerification = "code-verification"
	audDeviceOperations = "device-operations"
	bearerPrefix        = "Bearer "
)

func (s *serivce) CreateDeviceToken(device entity.Device) (string, error) {
	expAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 1024))
	return s.CreateToken(issuerSelf, device.Id, jwt.ClaimStrings{audDeviceOperations}, expAt)
}

func (s *serivce) CreateCodeToken(code entity.ActivationCode) (string, error) {
	expAt := jwt.NewNumericDate(time.Now().Add(time.Hour))
	return s.CreateToken(issuerSelf, code.Id, jwt.ClaimStrings{audCodeVerification}, expAt)
}

func (s *serivce) VerifyRow(rowToken string) (*jwt.Token, error) {
	return s.Verify(strings.TrimPrefix(rowToken, bearerPrefix))
}

func (s *serivce) Verify(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secret), nil
	})
}

func (s *serivce) CreateToken(issuer string, subject uuid.UUID, aud jwt.ClaimStrings, expiresAt *jwt.NumericDate) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   subject.String(),
		Audience:  aud,
		ExpiresAt: expiresAt,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secret))
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
