package entity

import "github.com/google/uuid"

type Code struct {
	Id    uuid.UUID `db:"id"`
	Email string    `db:"email"`
}

func CodeNil() Code {
	return Code{}
}
