package dto

type SingInDeviceRequestDto struct {
	PhoneNumber string `json:"phoneNumber"`
}

type SingInDeviceResponseDto struct {
	CodeToken string `json:"codeToken"`
}

type VerifyDeviceResponseDto struct {
	Token string `json:"token"`
}

type SetPinRequestDto struct {
	PinCode string
}

type LoginUserRequestDto struct {
	Pin string
}

type LoginUserResponseDto struct {
	Token string
}
