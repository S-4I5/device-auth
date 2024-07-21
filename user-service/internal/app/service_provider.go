package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	auth3 "user-service/internal/api/auth"
	"user-service/internal/api/interceptor"
	"user-service/internal/api/scope"
	user4 "user-service/internal/api/user"
	"user-service/internal/config"
	"user-service/internal/controller"
	"user-service/internal/controller/auth"
	"user-service/internal/controller/middleware"
	user3 "user-service/internal/controller/user"
	"user-service/internal/email"
	"user-service/internal/httperr"
	"user-service/internal/jwt"
	"user-service/internal/repository"
	"user-service/internal/repository/client"
	"user-service/internal/repository/code"
	"user-service/internal/repository/user"
	"user-service/internal/service"
	auth2 "user-service/internal/service/auth"
	client2 "user-service/internal/service/client"
	code2 "user-service/internal/service/code"
	user2 "user-service/internal/service/user"
)

type serviceProvider struct {
	authServerImpl         *auth3.AuthenticationServiceImplementation
	userServerImpl         *user4.UserV1ServiceImplementation
	errMessageSource       httperr.ErrorMessageSource
	errHandler             httperr.ErrorHandler
	tokenService           jwt.TokenService
	config                 config.Config
	pool                   *pgxpool.Pool
	scopeValidator         scope.Validator
	codeRepository         repository.CodeRepository
	userRepository         repository.UserRepository
	userService            service.UserService
	codeService            service.CodeService
	authService            service.AuthService
	userController         controller.UserController
	authController         controller.AuthController
	authInterceptor        interceptor.AuthInterceptor
	clientService          service.ClientService
	clientRepository       repository.ClientRepository
	emailSender            email.Sender
	authMiddlewareProvider middleware.AuthMiddlewareProvider
}

func NewServiceProvider(cfg config.Config) *serviceProvider {
	return &serviceProvider{
		config: cfg,
	}
}

func (p *serviceProvider) AuthMiddlewareProvider() middleware.AuthMiddlewareProvider {
	if p.authMiddlewareProvider == nil {
		p.authMiddlewareProvider = middleware.NewAuthMiddlewareProvider(p.TokenService(), p.ErrorHandler())
	}
	return p.authMiddlewareProvider
}

func (p *serviceProvider) ScopeValidator() scope.Validator {
	if p.scopeValidator == nil {
		p.scopeValidator = scope.NewValidator(
			p.AuthenticationServiceImplementation().GetMethodScopeMap(),
			p.UserV1ServiceImplementation().GetMethodScopeMap(),
		)
	}
	return p.scopeValidator
}

func (p *serviceProvider) AuthInterceptor() interceptor.AuthInterceptor {
	if p.authInterceptor == nil {
		p.authInterceptor = interceptor.NewAuthInterceptor(p.ClientService(), p.ScopeValidator())
	}
	return p.authInterceptor
}

func (p *serviceProvider) ClientService() service.ClientService {
	if p.clientService == nil {
		p.clientService = client2.NewService(p.ClientRepository())
	}
	return p.clientService
}

func (p *serviceProvider) ClientRepository() repository.ClientRepository {
	if p.clientRepository == nil {
		p.clientRepository = client.NewRepository()
	}
	return p.clientRepository
}

func (p *serviceProvider) ErrorMessageSource() httperr.ErrorMessageSource {
	if p.errMessageSource == nil {
		p.errMessageSource = httperr.NewMessageSource()
	}
	return p.errMessageSource
}

func (p *serviceProvider) ErrorHandler() httperr.ErrorHandler {
	if p.errHandler == nil {
		p.errHandler = httperr.NewHandler(p.ErrorMessageSource())
	}
	return p.errHandler
}

func (p *serviceProvider) PgxPool() *pgxpool.Pool {
	if p.pool == nil {
		ctx := context.TODO()
		pool, err := pgxpool.New(ctx, p.config.PostgresConnectionString())
		if err != nil {
			panic(err.Error())
		}
		p.pool = pool
	}
	return p.pool
}

func (p *serviceProvider) CodeRepository() repository.CodeRepository {
	if p.codeRepository == nil {
		p.codeRepository = code.NewRepository(p.PgxPool())
	}
	return p.codeRepository
}

func (p *serviceProvider) UserRepository() repository.UserRepository {
	if p.userRepository == nil {
		p.userRepository = user.NewRepository(p.PgxPool())
	}
	return p.userRepository
}

func (p *serviceProvider) CodeService() service.CodeService {
	if p.userService == nil {
		p.codeService = code2.NewService(p.CodeRepository())
	}
	return p.codeService
}

func (p *serviceProvider) UserService() service.UserService {
	if p.userService == nil {
		p.userService = user2.NewService(p.UserRepository(), p.CodeService())
	}
	return p.userService
}

func (p *serviceProvider) UserController() controller.UserController {
	if p.userController == nil {
		p.userController = user3.NewController(p.UserService(), p.ErrorHandler(), p.TokenService())
	}
	return p.userController
}

func (p *serviceProvider) AuthController() controller.AuthController {
	if p.authController == nil {
		p.authController = auth.NewController(p.AuthService(), p.ErrorHandler())
	}
	return p.authController
}

func (p *serviceProvider) AuthService() service.AuthService {
	if p.authService == nil {
		p.authService = auth2.NewService(p.TokenService(), p.UserService(), p.CodeService(), p.EmailSender())
	}
	return p.authService
}

func (p *serviceProvider) TokenService() jwt.TokenService {
	if p.tokenService == nil {
		p.tokenService = jwt.NewService(p.config.Secret)
	}
	return p.tokenService
}

func (p *serviceProvider) AuthenticationServiceImplementation() *auth3.AuthenticationServiceImplementation {
	if p.authServerImpl == nil {
		p.authServerImpl = auth3.NewImplementation(p.AuthService())
	}
	return p.authServerImpl
}

func (p *serviceProvider) UserV1ServiceImplementation() *user4.UserV1ServiceImplementation {
	if p.userServerImpl == nil {
		p.userServerImpl = user4.NewImplementation(p.UserService())
	}
	return p.userServerImpl
}

func (p *serviceProvider) EmailSender() email.Sender {
	if p.emailSender == nil {
		p.emailSender = email.NewSender(p.config.Mail)
	}
	return p.emailSender
}
