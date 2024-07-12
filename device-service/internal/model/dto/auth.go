package dto

type SingUpDeviceRequestDto struct {
	PhoneNumber string `json:"phoneNumber"`
}

type SingUpDeviceResponseDto struct {
	CodeToken string `json:"codeToken"`
}

type VerifyDeviceResponseDto struct {
	Token string `json:"token"`
}

type SetPinRequestDto struct {
	PinCode string `json:"pinCode"`
}

type LoginUserRequestDto struct {
	Pin string `json:"pin"`
}

type LoginUserResponseDto struct {
	Token string `json:"token"`
}

type BindUserToDeviceDtoRequest struct {
	UserToken string `json:"userToken"`
}
