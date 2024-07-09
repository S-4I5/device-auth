package auth

import (
	"context"
	"device-service/internal/model/dto"
	"encoding/json"
	"net/http"
)

func (c *controller) SignIn(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var req dto.SingInDeviceRequestDto
		err := json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			c.errorHandler.ReturnUnprocessableEntityError(writer, err)
			return
		}

		resp, err := c.authService.SingIn(req)
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err)
		}

		err = json.NewEncoder(writer).Encode(resp)
		if err != nil {
			c.errorHandler.ReturnProcessingResponseError(writer, err)
			return
		}
	}
}
