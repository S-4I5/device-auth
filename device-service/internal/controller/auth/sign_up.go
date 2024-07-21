package auth

import (
	"context"
	"device-service/internal/model/dto"
	"encoding/json"
	"net/http"
)

func (c *controller) SignUp(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var req dto.SingUpDeviceRequestDto
		err := json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			c.errorHandler.ReturnUnprocessableEntityError(writer, err, request.RequestURI)
			return
		}

		resp, err := c.authService.SingUp(req)
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		err = json.NewEncoder(writer).Encode(resp)
		if err != nil {
			c.errorHandler.ReturnProcessingResponseError(writer, err, request.RequestURI)
			return
		}
	}
}
