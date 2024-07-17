package util

import (
	"context"
	"encoding/base64"
)

type TokenAuth struct {
	Username string
	Password string
}

func (t TokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := t.Username + ":" + t.Password
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": authHeader,
	}, nil
}

func (TokenAuth) RequireTransportSecurity() bool {
	return true
}
