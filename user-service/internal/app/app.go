package app

import (
	"context"
	"crypto/tls"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
	"user-service/internal/config"
	"user-service/pkg/auth_v1"
	"user-service/pkg/user_v1"
)

type App struct {
	provider         *serviceProvider
	config           config.Config
	httpServer       http.Server
	grpcAuthServer   *grpc.Server
	grpcUserV1Server *grpc.Server
}

func NewApp(ctx context.Context, cfg config.Config) (*App, error) {

	app := &App{
		config:   cfg,
		provider: NewServiceProvider(cfg),
	}

	if err := app.runMigrations(); err != nil {
		return nil, err
	}

	if err := app.setupHttp(ctx); err != nil {
		return nil, err
	}

	if err := app.setupAuthGrpc(); err != nil {
		return nil, err
	}

	if err := app.setupUserV1Grpc(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Start() error {

	group := &sync.WaitGroup{}
	group.Add(2)

	go func() {
		defer group.Done()

		if err := a.RunHttpServer(); err != nil {
			log.Fatal("http fall:", err)
		}
	}()

	go func() {
		defer group.Done()

		if err := a.RunAuthGrpcServer(); err != nil {
			log.Fatal("auth grpc fall:", err)
		}
	}()

	go func() {
		defer group.Done()

		if err := a.RunUserV1GrpcServer(); err != nil {
			log.Fatal("user v1 grpc fall:", err)
		}
	}()

	group.Wait()

	return nil
}

func (a *App) RunHttpServer() error {
	return a.httpServer.ListenAndServe()
}

func (a *App) RunAuthGrpcServer() error {
	fmt.Printf(fmt.Sprintf(":%s", a.config.Grpc.AuthPort))

	list, err := net.Listen("tcp", fmt.Sprintf(":%s", a.config.Grpc.AuthPort))
	if err != nil {
		return err
	}

	fmt.Println(list.Addr().String())

	return a.grpcAuthServer.Serve(list)
}

func (a *App) setupAuthGrpc() error {

	cert, err := a.getCertificate()
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	a.grpcAuthServer = grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	reflection.Register(a.grpcAuthServer)

	auth_v1.RegisterAuthV1Server(a.grpcAuthServer, a.provider.AuthenticationServiceImplementation())

	return nil
}

func (a *App) RunUserV1GrpcServer() error {
	fmt.Printf(fmt.Sprintf("localhost:%s", a.config.Grpc.UserPort))

	list, err := net.Listen("tcp", fmt.Sprintf(":%s", a.config.Grpc.UserPort))
	if err != nil {
		return err
	}

	return a.grpcUserV1Server.Serve(list)
}

func (a *App) setupUserV1Grpc() error {

	cert, err := a.getCertificate()
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	a.grpcUserV1Server = grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))

	reflection.Register(a.grpcUserV1Server)

	user_v1.RegisterUserV1Server(a.grpcUserV1Server, a.provider.UserV1ServiceImplementation())

	return nil
}

func (a *App) setupHttp(ctx context.Context) error {

	authRouter := http.NewServeMux()

	authRouter.HandleFunc("POST /sign-up", a.provider.AuthController().SignUpUser(ctx))
	authRouter.HandleFunc("PATCH /verify-email/{id}", a.provider.AuthController().VerifyUserEmail(ctx))
	authRouter.HandleFunc("POST /login", a.provider.AuthController().LoginUser(ctx))

	userRouter := http.NewServeMux()

	userRouter.HandleFunc("GET /me", a.provider.UserController().GetMe(ctx))
	userRouter.HandleFunc("GET /{id}", a.provider.UserController().Get(ctx))
	userRouter.HandleFunc("GET /search", a.provider.UserController().GetByPhoneNumber(ctx))

	version := http.NewServeMux()

	version.Handle(
		fmt.Sprintf("/%s/%s/", a.config.Version, "user"),
		http.StripPrefix(
			fmt.Sprintf("/%s/%s", a.config.Version, "user"),
			userRouter),
	)
	version.Handle(
		fmt.Sprintf("/%s/%s/", a.config.Version, "auth"),
		http.StripPrefix(
			fmt.Sprintf("/%s/%s", a.config.Version, "auth"),
			authRouter),
	)

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", version))

	a.httpServer = http.Server{
		Addr:         ":" + a.config.Http.Port,
		Handler:      api,
		IdleTimeout:  a.config.Http.IdleTimeout,
		ReadTimeout:  a.config.Http.Timeout,
		WriteTimeout: a.config.Http.Timeout,
	}

	return nil
}

func (a *App) runMigrations() error {
	fmt.Printf(a.config.PostgresConnectionString())

	db, err := goose.OpenDBWithDriver("postgres", a.config.PostgresConnectionString())
	if err != nil {
		return fmt.Errorf("%s: %s", "migr", err.Error())
	}
	defer db.Close()

	err = goose.Up(db, "./migrations")
	if err != nil {
		return fmt.Errorf("%s: %s", "migr", err.Error())
	}

	return nil
}

func (a *App) getCertificate() (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair("./server-cert.pem", "./server-key.pem")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	return cert, nil
}
