package app

import (
	"context"
	"device-service/internal/config"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"net/http"
)

type App struct {
	provider *serviceProvider
	config   config.Config
	server   http.Server
}

func NewApp(config config.Config) (*App, error) {
	app := App{
		config:   config,
		provider: newServiceProvider(config),
	}

	if err := app.runMigrations(); err != nil {
		return nil, err
	}

	if err := app.setupHttp(); err != nil {
		return nil, err
	}

	return &app, nil
}

func (a *App) Start() error {
	return a.runHttpServer()
}

func (a *App) runHttpServer() error {
	return a.server.ListenAndServe()
}

func (a *App) setupHttp() error {
	authRouter := http.NewServeMux()

	ctx := context.TODO()

	authRouter.HandleFunc("POST /sign-up", a.provider.AuthController().SignUp(ctx))

	verifyDeviceRoute := http.NewServeMux()
	verifyDeviceRoute.HandleFunc("/", a.provider.AuthController().VerifyDevice(ctx))
	authRouter.Handle("PATCH /verify-device", a.provider.AuthMiddlewareProvider().GetAuthMiddlewareProvider(verifyDeviceRoute))

	setPinRoute := http.NewServeMux()
	setPinRoute.HandleFunc("/", a.provider.AuthController().SetPin(ctx))
	authRouter.Handle("PATCH /set-pin", a.provider.AuthMiddlewareProvider().GetAuthMiddlewareProvider(setPinRoute))

	loginRoute := http.NewServeMux()
	loginRoute.HandleFunc("/", a.provider.AuthController().LoginUser(ctx))
	authRouter.Handle("POST /login", a.provider.AuthMiddlewareProvider().GetAuthMiddlewareProvider(loginRoute))

	bindUserRoute := http.NewServeMux()
	bindUserRoute.HandleFunc("/", a.provider.AuthController().BindUser(ctx))
	authRouter.Handle("PATCH /bind-user", a.provider.AuthMiddlewareProvider().GetAuthMiddlewareProvider(bindUserRoute))

	version := http.NewServeMux()

	version.Handle(
		fmt.Sprintf("/%s/%s/", a.config.Version, "auth"),
		http.StripPrefix(
			fmt.Sprintf("/%s/%s", a.config.Version, "auth"),
			authRouter),
	)

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", version))

	a.server = http.Server{
		Addr:         ":" + a.config.Http.Port,
		Handler:      api,
		IdleTimeout:  a.config.Http.IdleTimeout,
		ReadTimeout:  a.config.Http.Timeout,
		WriteTimeout: a.config.Http.Timeout,
	}
	return nil
}

func (a *App) runMigrations() error {
	fmt.Println(a.config.PostgresConnectionString())

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
