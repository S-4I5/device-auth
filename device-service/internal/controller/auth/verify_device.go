package auth

import (
	"context"
	"device-service/internal/jwt"
	"encoding/json"
	"net/http"
)

func (c *controller) VerifyDevice(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		code := request.FormValue("code")
		if len(code) == 0 {
			c.errorHandler.ReturnIncorrectReqParamError(writer)
			return
		}

		token, err := c.tokenVerifier.VerifyRow(request.Header.Get("Authorization"))
		if err != nil {
			c.errorHandler.ReturnUnauthenticatedError(writer, err)
			return
		}

		deviceUuid, err := jwt.GetSubjectIdFromToken(*token)
		if err != nil {
			c.errorHandler.ReturnUnauthenticatedError(writer, err)
			return
		}

		resp, err := c.authService.VerifyDevice(code, deviceUuid)
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err)
			return
		}

		if err = json.NewEncoder(writer).Encode(resp); err != nil {
			c.errorHandler.ReturnProcessingResponseError(writer, err)
			return
		}
	}
}
