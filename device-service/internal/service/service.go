package service

import (
	"device-service/internal/model/dto"
	"device-service/internal/model/entity"
	"github.com/google/uuid"
)

type CodeService interface {
	Create(code entity.ActivationCode) (entity.ActivationCode, error)
	Get(id uuid.UUID) (entity.ActivationCode, error)
	GetByDevice(deviceId uuid.UUID) (entity.ActivationCode, error)
	Delete(id uuid.UUID) error
}

type DeviceService interface {
	Create(device entity.Device) (entity.Device, error)
	Get(id uuid.UUID) (entity.Device, error)
	Verify(id uuid.UUID) error
	SetPin(id uuid.UUID, pin string) error
	SetUser(id, userId uuid.UUID) error
}

type AuthService interface {
	SingUp(req dto.SingUpDeviceRequestDto) (dto.SingUpDeviceResponseDto, error)
	SetPin(req dto.SetPinRequestDto, deviceId uuid.UUID) error
	LoginUser(req dto.LoginUserRequestDto, deviceId uuid.UUID) (dto.LoginUserResponseDto, error)
	VerifyDevice(code string, codeId uuid.UUID) (dto.VerifyDeviceResponseDto, error)
	BindUserToDevice(req dto.BindUserToDeviceDtoRequest, deviceId uuid.UUID) error
}
