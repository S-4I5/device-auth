package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"user-service/internal/httperr"
	"user-service/internal/jwt"
)

type AuthMiddlewareProvider interface {
	GetAuthMiddleware(next http.Handler) http.Handler
}

const (
	userKey = "user"
)

var (
	errCannotParseExtractedValue = fmt.Errorf("cannot parse extracted value")
	errNoValueInContext          = fmt.Errorf("no value in context")
)

type provider struct {
	handler     httperr.ErrorHandler
	tokenParser jwt.TokenParser
}

func NewAuthMiddlewareProvider(tokenParser jwt.TokenParser, handler httperr.ErrorHandler) *provider {
	return &provider{
		tokenParser: tokenParser,
		handler:     handler,
	}
}

func (p *provider) GetAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		token, err := p.tokenParser.ValidateToken(strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer "))
		if err != nil {
			p.handler.ReturnAuthenticationError(writer, err, request.RequestURI)
			return
		}

		id, err := jwt.GetSubjectIdFromToken(*token)
		if err != nil {
			p.handler.ReturnAuthenticationError(writer, err, request.RequestURI)
			return
		}

		newCtx := context.WithValue(context.Background(), userKey, id)

		next.ServeHTTP(writer, request.WithContext(newCtx))
	})
}

func GetUserUUIDFromContext(ctx context.Context) (uuid.UUID, error) {
	rowId := ctx.Value(userKey)

	if rowId == nil {
		return uuid.Nil, errNoValueInContext
	}

	fmt.Println(rowId)

	id, ok := rowId.(uuid.UUID)
	if !ok {
		return uuid.Nil, errCannotParseExtractedValue
	}

	return id, nil
}
