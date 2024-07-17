package code

import (
	"github.com/google/uuid"
	"user-service/internal/repository"
	def "user-service/internal/service"
)

var _ def.CodeService = (*service)(nil)

type service struct {
	repository repository.CodeRepository
}

func NewService(codeRepository repository.CodeRepository) *service {
	return &service{repository: codeRepository}
}

func (s *service) Create(email string) (uuid.UUID, error) {
	code, err := s.repository.Save(email)
	if err != nil {
		return uuid.Nil, err
	}

	return code.Id, nil
}

func (s *service) GetEmailById(id uuid.UUID) (string, error) {
	code, err := s.repository.Get(id)
	if err != nil {
		return "", err
	}

	return code.Email, nil
}

func (s *service) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}
