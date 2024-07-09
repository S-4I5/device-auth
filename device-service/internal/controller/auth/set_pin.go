package auth

import (
	"context"
	"device-service/internal/model/dto"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (c *controller) SetPin(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		token, err := c.tokenVerifier.VerifyRow(request.Header.Get("Authorization"))
		if err != nil {
			c.errorHandler.ReturnUnauthenticatedError(writer, err)
			return
		}

		targetId, err := token.Claims.GetSubject()
		if err != nil {
			c.errorHandler.ReturnUnauthenticatedError(writer, err)
		}

		var req dto.SetPinRequestDto
		if err = json.NewDecoder(request.Body).Decode(&req); err != nil {
			c.errorHandler.ReturnUnprocessableEntityError(writer, err)
			return
		}

		targetUuid, err := uuid.Parse(targetId)
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err)
			return
		}

		if err = c.authService.SetPin(req, targetUuid); err != nil {
			c.errorHandler.ReturnProcessingResponseError(writer, err)
			return
		}
	}
}
