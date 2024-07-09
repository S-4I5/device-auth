package entity

import "github.com/google/uuid"

type ActivationCode struct {
	Id       uuid.UUID `db:"id"`
	DeviceId uuid.UUID `dv:"device_id"`
	Code     string    `db:"code"`
}

func NewActivationCode(id uuid.UUID, deviceId uuid.UUID, code string) ActivationCode {
	return ActivationCode{
		Id:       id,
		DeviceId: deviceId,
		Code:     code,
	}
}

func ActivationCodeNil() ActivationCode {
	return ActivationCode{}
}
