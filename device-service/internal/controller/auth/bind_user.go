package auth

import (
	"context"
	"device-service/internal/jwt"
	"device-service/internal/model/dto"
	"encoding/json"
	"net/http"
)

func (c *controller) BindUser(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		token, err := c.tokenVerifier.VerifyRow(request.Header.Get("Authorization"))
		if err != nil {
			c.errorHandler.ReturnUnauthenticatedError(writer, err)
			return
		}

		deviceUuid, err := jwt.GetSubjectIdFromToken(*token)
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err)
			return
		}

		var req dto.BindUserToDeviceDtoRequest
		if err = json.NewDecoder(request.Body).Decode(&req); err != nil {
			c.errorHandler.ReturnUnprocessableEntityError(writer, err)
			return
		}

		if err = c.authService.BindUserToDevice(req, deviceUuid); err != nil {
			c.errorHandler.ReturnServiceError(writer, err)
			return
		}
	}
}
