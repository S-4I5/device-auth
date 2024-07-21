package auth

import (
	"context"
	"device-service/internal/controller/middleware"
	"device-service/internal/model/dto"
	"encoding/json"
	"net/http"
)

func (c *controller) SetPin(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		id, err := middleware.GetSubjectUUIDFromContext(request.Context())
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		var req dto.SetPinRequestDto
		if err = json.NewDecoder(request.Body).Decode(&req); err != nil {
			c.errorHandler.ReturnUnprocessableEntityError(writer, err, request.RequestURI)
			return
		}

		if err = c.authService.SetPin(req, id); err != nil {
			c.errorHandler.ReturnProcessingResponseError(writer, err, request.RequestURI)
			return
		}
	}
}
