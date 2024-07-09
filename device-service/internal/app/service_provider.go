package app

import (
	"context"
	"device-service/internal/config"
	"device-service/internal/controller"
	auth2 "device-service/internal/controller/auth"
	"device-service/internal/err"
	"device-service/internal/jwt"
	"device-service/internal/repository"
	code2 "device-service/internal/repository/code"
	device2 "device-service/internal/repository/device"
	"device-service/internal/service"
	"device-service/internal/service/auth"
	"device-service/internal/service/code"
	"device-service/internal/service/device"
	"device-service/internal/util"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"s4i5.xyz/user-service/pkg/auth_v1"
	"s4i5.xyz/user-service/pkg/user_v1"
)

type serviceProvider struct {
	authController     controller.AuthController
	config             config.Config
	authV1Client       auth_v1.AuthV1Client
	userV1Client       user_v1.UserV1Client
	pool               *pgxpool.Pool
	tokenService       jwt.TokenService
	authService        service.AuthService
	deviceService      service.DeviceService
	codeService        service.CodeService
	errorHandler       err.ErrorHandler
	errorMessageSource err.ErrorMessageSource
	codeRepository     repository.CodeRepository
	deviceRepository   repository.DeviceRepository
}

func newServiceProvider(config config.Config) *serviceProvider {
	return &serviceProvider{
		config: config,
	}
}

func (s *serviceProvider) AuthController() controller.AuthController {
	if s.authController == nil {
		s.authController = auth2.NewController(s.AuthService(), s.TokenService(), s.ErrorHandler())
	}
	return s.authController
}

func (s *serviceProvider) Config() config.Config {
	return s.config
}

func (s *serviceProvider) ErrorMessageSource() err.ErrorMessageSource {
	if s.errorMessageSource == nil {
		s.errorMessageSource = err.NewMessageSource()
	}

	return s.errorMessageSource
}

func (s *serviceProvider) ErrorHandler() err.ErrorHandler {
	if s.errorHandler == nil {
		s.errorHandler = err.NewHandler(s.ErrorMessageSource())
	}

	return s.errorHandler
}

func (s *serviceProvider) AuthV1Client() auth_v1.AuthV1Client {
	if s.authV1Client == nil {
		creds, err := credentials.NewClientTLSFromFile("./ca-cert.pem", "0.0.0.0")
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(s.config.Grpc.AuthUrl)

		conn, err := grpc.NewClient(s.config.Grpc.AuthUrl,
			grpc.WithTransportCredentials(creds),
			grpc.WithPerRPCCredentials(util.TokenAuth{Token: s.config.Grpc.AuthToken}),
		)
		if err != nil {
			panic("cannot setup AuthV1 grpc conn" + err.Error())
		}

		s.authV1Client = auth_v1.NewAuthV1Client(conn)
	}

	return s.authV1Client
}

func (s *serviceProvider) UserV1Client() user_v1.UserV1Client {
	if s.userV1Client == nil {
		creds, err := credentials.NewClientTLSFromFile("./ca-cert.pem", "0.0.0.0")
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(s.config.Grpc.UserUrl)

		conn, err := grpc.NewClient(s.config.Grpc.UserUrl,
			grpc.WithTransportCredentials(creds),
			grpc.WithPerRPCCredentials(util.TokenAuth{Token: s.config.Grpc.UserToken}),
		)
		if err != nil {
			panic("cannot setup UserV1 grpc conn" + err.Error())
		}

		s.userV1Client = user_v1.NewUserV1Client(conn)
	}

	return s.userV1Client
}

func (s *serviceProvider) Pool() *pgxpool.Pool {
	if s.pool == nil {
		ctx := context.TODO()
		pool, err := pgxpool.New(ctx, s.config.PostgresConnectionString())
		if err != nil {
			panic(err.Error())
		}
		s.pool = pool
	}
	return s.pool
}
func (s *serviceProvider) TokenService() jwt.TokenService {
	if s.tokenService == nil {
		s.tokenService = jwt.NewService(s.Config().Secret)
	}
	return s.tokenService
}

func (s *serviceProvider) AuthService() service.AuthService {
	if s.authService == nil {
		s.authService = auth.NewService(s.CodeService(), s.DeviceService(), s.TokenService(), s.AuthV1Client(), s.UserV1Client())
	}
	return s.authService
}

func (s *serviceProvider) DeviceService() service.DeviceService {
	if s.deviceService == nil {
		s.deviceService = device.NewService(s.DeviceRepository())
	}
	return s.deviceService
}

func (s *serviceProvider) CodeService() service.CodeService {
	if s.codeService == nil {
		s.codeService = code.NewService(s.CodeRepository())
	}
	return s.codeService
}

func (s *serviceProvider) CodeRepository() repository.CodeRepository {
	if s.codeRepository == nil {
		s.codeRepository = code2.NewRepository(s.Pool())
	}
	return s.codeRepository
}

func (s *serviceProvider) DeviceRepository() repository.DeviceRepository {
	if s.deviceRepository == nil {
		s.deviceRepository = device2.NewRepository(s.Pool())
	}
	return s.deviceRepository
}
