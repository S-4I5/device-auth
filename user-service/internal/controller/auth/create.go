package auth

import (
	"context"
	"encoding/json"
	"net/http"
	dto2 "user-service/internal/model/dto"
)

func (c *controller) SignInUser(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var dto dto2.SignInUserRequestDto
		if err := json.NewDecoder(request.Body).Decode(&dto); err != nil {
			c.errHandler.ReturnUnprocessableEntityError(writer, err)
			return
		}

		resp, err := c.authService.SignInUser(dto)
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
