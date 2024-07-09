package entity

import "github.com/google/uuid"

type Device struct {
	Id          uuid.UUID `db:"id"`
	PhoneNumber string    `db:"phone_number"`
	PinCode     string    `db:"pin_code"`
	UserId      uuid.UUID `db:"user_id"`
	IsVerified  bool      `db:"is_verified"`
}

func NewDevice(id uuid.UUID, phoneNumber, pin string, userId uuid.UUID, verified bool) Device {
	return Device{
		Id:          id,
		PhoneNumber: phoneNumber,
		PinCode:     pin,
		UserId:      userId,
		IsVerified:  verified,
	}
}

func DeviceNil() Device {
	return Device{}
}
