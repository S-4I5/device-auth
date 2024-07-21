package auth

import (
	"context"
	"device-service/internal/controller/middleware"
	"encoding/json"
	"net/http"
)

func (c *controller) VerifyDevice(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		code := request.FormValue("code")
		if len(code) == 0 {
			c.errorHandler.ReturnIncorrectReqParamError(writer, request.RequestURI)
			return
		}

		id, err := middleware.GetSubjectUUIDFromContext(request.Context())
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		resp, err := c.authService.VerifyDevice(code, id)
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		if err = json.NewEncoder(writer).Encode(resp); err != nil {
			c.errorHandler.ReturnProcessingResponseError(writer, err, request.RequestURI)
			return
		}
	}
}
