package controller

import (
	"context"
	"net/http"
)

type UserController interface {
	Get(cxt context.Context) http.HandlerFunc
	GetMe(ctx context.Context) http.HandlerFunc
	GetByPhoneNumber(ctx context.Context) http.HandlerFunc
}

type AuthController interface {
	VerifyUserEmail(cxt context.Context) http.HandlerFunc
	SignUpUser(cxt context.Context) http.HandlerFunc
	LoginUser(cxt context.Context) http.HandlerFunc
}
