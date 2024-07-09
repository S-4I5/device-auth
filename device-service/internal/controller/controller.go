package controller

import (
	"context"
	"net/http"
)

type AuthController interface {
	SignIn(ctx context.Context) http.HandlerFunc
	SetPin(ctx context.Context) http.HandlerFunc
	LoginUser(ctx context.Context) http.HandlerFunc
	VerifyDevice(ctx context.Context) http.HandlerFunc
}
