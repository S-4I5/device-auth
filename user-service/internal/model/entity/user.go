package entity

import (
	"github.com/google/uuid"
)

type User struct {
	Id              uuid.UUID `db:"id"`
	Email           string    `db:"email"`
	PhoneNumber     string    `db:"phone_number"`
	Password        string    `db:"password"`
	IsEmailVerified bool      `db:"is_email_verified"`
}

func UserNil() User {
	return User{}
}

func NewUser(id uuid.UUID, email, password string, emailVerified bool) (User, error) {
	return User{
		Id:              id,
		Password:        password,
		Email:           email,
		IsEmailVerified: emailVerified,
	}, nil
}
