package middleware

import (
	"context"
	"device-service/internal/httperr"
	"device-service/internal/jwt"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type AuthMiddlewareProvider interface {
	GetAuthMiddlewareProvider(next http.Handler) http.Handler
}

const subjectIdKey = "subjectId"

var (
	errCannotParseExtractedValue = fmt.Errorf("cannot parse extracted value")
	errNoValueInContext          = fmt.Errorf("no value in context")
)

type provider struct {
	tokenVerifier jwt.TokenVerifier
	errorHandler  httperr.ErrorHandler
}

func NewAuthMiddlewareProvider(tokenVerifier jwt.TokenVerifier, errorHandler httperr.ErrorHandler) *provider {
	return &provider{
		tokenVerifier: tokenVerifier,
		errorHandler:  errorHandler,
	}
}

func (p *provider) GetAuthMiddlewareProvider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token, err := p.tokenVerifier.VerifyRow(request.Header.Get("Authorization"))
		if err != nil {
			p.errorHandler.ReturnUnauthenticatedError(writer, err, request.RequestURI)
			return
		}

		subject, err := jwt.GetSubjectIdFromToken(*token)
		if err != nil {
			p.errorHandler.ReturnUnauthenticatedError(writer, err, request.RequestURI)
			return
		}

		newCtx := context.WithValue(context.Background(), subjectIdKey, subject)

		next.ServeHTTP(writer, request.WithContext(newCtx))
	})
}

func GetSubjectUUIDFromContext(ctx context.Context) (uuid.UUID, error) {
	rowId := ctx.Value(subjectIdKey)

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
