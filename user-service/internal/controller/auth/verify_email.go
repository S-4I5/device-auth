package auth

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func (c *controller) VerifyUserEmail(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		codeId, err := uuid.Parse(request.PathValue("id"))
		if err != nil {
			c.errHandler.ReturnIncorrectPathValueError(writer, err, request.RequestURI)
			return
		}

		if err = c.authService.VerifyUserEmail(codeId); err != nil {
			c.errHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}
	}
}
