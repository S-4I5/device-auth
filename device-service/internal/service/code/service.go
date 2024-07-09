package code

import (
	"device-service/internal/model/entity"
	"device-service/internal/repository"
	"github.com/google/uuid"
	"math/rand"
)

type service struct {
	codeRepository repository.CodeRepository
}

func NewService(codeRepository repository.CodeRepository) *service {
	return &service{codeRepository: codeRepository}
}

func (s *service) Create(code entity.ActivationCode) (entity.ActivationCode, error) {
	code.Code = generateCode(5)

	return s.codeRepository.Create(code)
}

func (s *service) GetByDevice(deviceId uuid.UUID) (entity.ActivationCode, error) {
	return s.codeRepository.GetByDeviceId(deviceId)
}

func (s *service) Get(id uuid.UUID) (entity.ActivationCode, error) {
	return s.codeRepository.Get(id)
}

func (s *service) Delete(id uuid.UUID) error {
	return s.codeRepository.Delete(id)
}

var allowedLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func generateCode(length int) string {
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = allowedLetters[rand.Intn(len(allowedLetters))]
	}
	return string(result)
}
