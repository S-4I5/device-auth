package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"user-service/internal/mapper"
)

func (c *controller) GetByPhoneNumber(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		phone := strings.Replace(request.FormValue("phone"), " ", "+", 1)

		user, err := c.userService.GetByPhoneNumber(phone)
		if err != nil {
			c.errHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		if err = json.NewEncoder(writer).Encode(mapper.UserToUserDto(user)); err != nil {
			c.errHandler.ReturnProcessingResponseError(writer, err, request.RequestURI)
			return
		}
	}
}
