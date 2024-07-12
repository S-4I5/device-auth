package auth

import (
	"context"
	"encoding/json"
	"net/http"
	dto2 "user-service/internal/model/dto"
)

func (c *controller) SignUpUser(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var dto dto2.SignUpUserRequestDto
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			c.errHandler.ReturnUnprocessableEntityError(writer, err)
			return
		}

		resp, err := c.authService.SignUpUser(dto)
		if err != nil {
			c.errHandler.ReturnServiceError(writer, err)
			return
		}

		if err = json.NewEncoder(writer).Encode(&resp); err != nil {
			c.errHandler.ReturnProcessingResponseError(writer, err)
			return
		}
	}
}
