package repository

import (
	"device-service/internal/model/entity"
	"github.com/google/uuid"
)

type CodeRepository interface {
	Create(code entity.ActivationCode) (entity.ActivationCode, error)
	Get(id uuid.UUID) (entity.ActivationCode, error)
	GetByDeviceId(deviceId uuid.UUID) (entity.ActivationCode, error)
	Delete(id uuid.UUID) error
}

type DeviceRepository interface {
	Create(device entity.Device) (entity.Device, error)
	Get(id uuid.UUID) (entity.Device, error)
	SetVerified(id uuid.UUID) error
	SetPin(id uuid.UUID, pin string) error
	SetUser(id, userId uuid.UUID) error
}
