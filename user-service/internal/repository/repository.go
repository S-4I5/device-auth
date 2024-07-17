package repository

import (
	"github.com/google/uuid"
	"user-service/internal/model/entity"
)

type UserRepository interface {
	Save(email, password string, phoneNumber string) (entity.User, error)
	Get(id uuid.UUID) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetByPhoneNumber(phoneNumber string) (entity.User, error)
	VerifyEmail(email string) error
}

type CodeRepository interface {
	Save(email string) (entity.Code, error)
	Get(id uuid.UUID) (entity.Code, error)
	Delete(id uuid.UUID) error
}

type ClientRepository interface {
	Save(client entity.Client) error
	Get(id string) (entity.Client, error)
}
