package model

import (
	"database/sql"
	"device-service/internal/model/entity"
	"github.com/google/uuid"
)

type Device struct {
	Id          uuid.UUID      `db:"id"`
	PhoneNumber string         `db:"phone_number"`
	PinCode     sql.NullString `db:"pin_code"`
	UserId      uuid.UUID      `db:"user_id"`
	IsVerified  bool           `db:"is_verified"`
}

func DbDeviceToDevice(device Device) entity.Device {
	return entity.Device{
		Id:          device.Id,
		PhoneNumber: device.PhoneNumber,
		PinCode:     device.PinCode.String,
		UserId:      device.UserId,
		IsVerified:  device.IsVerified,
	}
}
