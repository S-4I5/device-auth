package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"user-service/internal/model/dto"
)

func (c *controller) LoginUser(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var requestDto dto.LoginUserRequestDto
		if err := json.NewDecoder(request.Body).Decode(&requestDto); err != nil {
			c.errHandler.ReturnUnprocessableEntityError(writer, err, request.RequestURI)
			return
		}

		response, err := c.authService.LoginUser(requestDto)
		if err != nil {
			c.errHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		if err = json.NewEncoder(writer).Encode(&response); err != nil {
			c.errHandler.ReturnProcessingResponseError(writer, err, request.RequestURI)
			return
		}
	}
}
