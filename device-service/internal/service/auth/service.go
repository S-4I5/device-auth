package auth

import (
	"context"
	"device-service/internal/jwt"
	"device-service/internal/model/dto"
	"device-service/internal/model/entity"
	service2 "device-service/internal/service"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"s4i5.xyz/user-service/pkg/auth_v1"
	"s4i5.xyz/user-service/pkg/user_v1"
)

type service struct {
	codeService   service2.CodeService
	deviceService service2.DeviceService
	tokenCreator  jwt.TokenCreator
	authClient    auth_v1.AuthV1Client
}

func NewService(codeService service2.CodeService, deviceService service2.DeviceService, tokenCreator jwt.TokenCreator, authClient auth_v1.AuthV1Client, userClient user_v1.UserV1Client) *service {
	return &service{
		codeService:   codeService,
		deviceService: deviceService,
		tokenCreator:  tokenCreator,
		authClient:    authClient,
	}
}

func (s *service) SingUp(req dto.SingUpDeviceRequestDto) (dto.SingUpDeviceResponseDto, error) {
	device, err := s.deviceService.Create(entity.Device{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		return dto.SingUpDeviceResponseDto{}, err
	}

	code, err := s.codeService.Create(entity.ActivationCode{
		DeviceId: device.Id,
	})
	if err != nil {
		return dto.SingUpDeviceResponseDto{}, err
	}

	//TODO: send code in other wayXD
	log.Println(code.Code)

	tkn, err := s.tokenCreator.CreateCodeToken(code)
	if err != nil {
		return dto.SingUpDeviceResponseDto{}, err
	}

	return dto.SingUpDeviceResponseDto{CodeToken: tkn}, nil
}

func (s *service) SetPin(req dto.SetPinRequestDto, deviceId uuid.UUID) error {
	return s.deviceService.SetPin(deviceId, req.PinCode)
}

func (s *service) LoginUser(req dto.LoginUserRequestDto, deviceId uuid.UUID) (dto.LoginUserResponseDto, error) {

	device, err := s.deviceService.Get(deviceId)
	if err != nil {
		return dto.LoginUserResponseDto{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(device.PinCode), []byte(req.Pin)); err != nil {
		return dto.LoginUserResponseDto{}, fmt.Errorf("incorrect pin")
	}

	ctx := context.TODO()
	resp, err := s.authClient.AuthenticateUser(ctx, &auth_v1.AuthenticateUserRequest{UserId: device.UserId.String()})
	if err != nil {
		return dto.LoginUserResponseDto{}, err
	}

	return dto.LoginUserResponseDto{Token: resp.GetToken()}, nil
}

func (s *service) VerifyDevice(code string, codeId uuid.UUID) (dto.VerifyDeviceResponseDto, error) {
	actCode, err := s.codeService.Get(codeId)
	if err != nil {
		return dto.VerifyDeviceResponseDto{}, err
	}

	if actCode.Code != code {
		return dto.VerifyDeviceResponseDto{}, fmt.Errorf("incorrect code")
	}

	if err = s.deviceService.Verify(actCode.DeviceId); err != nil {
		return dto.VerifyDeviceResponseDto{}, err
	}

	if err = s.codeService.Delete(codeId); err != nil {
		return dto.VerifyDeviceResponseDto{}, err
	}

	device, err := s.deviceService.Get(actCode.DeviceId)
	if err != nil {
		return dto.VerifyDeviceResponseDto{}, err
	}

	tkn, err := s.tokenCreator.CreateDeviceToken(device)
	if err != nil {
		return dto.VerifyDeviceResponseDto{}, err
	}

	return dto.VerifyDeviceResponseDto{Token: tkn}, nil
}

func (s *service) BindUserToDevice(req dto.BindUserToDeviceDtoRequest, deviceId uuid.UUID) error {
	ctx := context.TODO()
	fmt.Printf(req.UserToken)

	authResp, err := s.authClient.ValidateToken(ctx, &auth_v1.ValidateTokenRequest{Token: req.UserToken})
	if err != nil {
		return err
	}

	userId, err := uuid.Parse(authResp.GetUserId())
	if err != nil {
		return err
	}

	return s.deviceService.SetUser(deviceId, userId)
}
