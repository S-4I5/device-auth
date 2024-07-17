package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Secret   string         `yaml:"secret"`
	Version  string         `yaml:"version" env:"VERSION" env-default:"v1"`
	Http     HttpConfig     `yaml:"http"`
	Postgres PostgresConfig `yaml:"postgres"`
	UserGrpc UserGrpcConfig `yaml:"user-grpc"`
}

type HttpConfig struct {
	Port        string        `yaml:"port" env:"HTTP-PORT" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle-timeout"`
}

type PostgresConfig struct {
	User     string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"password"`
	Database string `yaml:"database" env:"POSTGRES_DB" env-default:"postgres"`
	Url      string `yaml:"url" env:"POSTGRES_URL" env-default:"localhost:5434"`
}

type UserGrpcConfig struct {
	ClientId     string `yaml:"client-id" env:"GRPC-CLIENT-ID"`
	ClientSecret string `yaml:"client-secret" env:"GRPC-CLIENT-SECRET"`
	AuthUrl      string `yaml:"auth-url" env:"GRPC-AUTH-URL" env-default:":50051"`
	UserUrl      string `yaml:"user-url" env:"GRPC-USER-URL" env-default:":50052"`
}

func MustRead(path string) Config {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err.Error())
	}

	fmt.Println(cfg)

	return cfg
}

func (c *Config) PostgresConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s",
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Url,
		c.Postgres.Database,
	)
}
