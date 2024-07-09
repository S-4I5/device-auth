package user

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"user-service/internal/model/entity"
	"user-service/internal/repository"
	def "user-service/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	repository  repository.UserRepository
	codeService def.CodeService
}

func NewService(userRepository repository.UserRepository, codeService def.CodeService) *service {
	return &service{repository: userRepository, codeService: codeService}
}

func (s *service) Create(email, password string, phoneNumber string) (entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entity.UserNil(), err
	}

	user, err := s.repository.Create(email, string(hashedPassword), phoneNumber)
	return user, err
}

func (s *service) Get(id uuid.UUID) (entity.User, error) {
	user, err := s.repository.Get(id)

	return user, err
}

func (s *service) GetByEmail(email string) (entity.User, error) {
	user, err := s.repository.GetByEmail(email)

	return user, err
}

func (s *service) SetEmailVerified(email string) error {
	return s.repository.VerifyEmail(email)
}

func (s *service) GetByPhoneNumber(phoneNumber string) (entity.User, error) {
	return s.repository.GetByPhoneNumber(phoneNumber)
}
