package device

import (
	"device-service/internal/model/entity"
	"device-service/internal/repository"
	"github.com/google/uuid"
)

type service struct {
	deviceRepository repository.DeviceRepository
}

func NewService(deviceRepository repository.DeviceRepository) *service {
	return &service{deviceRepository: deviceRepository}
}

func (s *service) Create(device entity.Device) (entity.Device, error) {
	return s.deviceRepository.Create(device)
}

func (s *service) Get(id uuid.UUID) (entity.Device, error) {
	return s.deviceRepository.Get(id)
}

func (s *service) SetPin(id uuid.UUID, pin string) error {
	return s.deviceRepository.SetPin(id, pin)
}

func (s *service) SetUser(id, userId uuid.UUID) error {
	return s.deviceRepository.SetUser(id, userId)
}

func (s *service) Verify(id uuid.UUID) error {
	return s.deviceRepository.SetVerified(id)
}
