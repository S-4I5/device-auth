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
	Grpc     GrpcConfig     `yaml:"grpc"`
	Mail     MailConfig     `yaml:"mail"`
	Clients  []Client       `yaml:"clients"`
}

type HttpConfig struct {
	Port        string        `yaml:"port" env:"HTTP-PORT" env-default:"8080"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type PostgresConfig struct {
	User     string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:""`
	Database string `yaml:"database" env:"POSTGRES_DB" env-default:"postgres"`
	Url      string `yaml:"url" env:"POSTGRES_URL" env-default:"localhost:5432"`
}

type GrpcConfig struct {
	AuthPort string `yaml:"auth_port" env:"AUTH-GRPC-URL" env-default:"50051"`
	UserPort string `yaml:"user_port" env:"USER-GRPC-URL" env-default:"50052"`
}

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Client struct {
	ClientId     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	Scopes       []string `yaml:"scopes"`
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
