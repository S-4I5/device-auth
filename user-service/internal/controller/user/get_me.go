package user

import (
	"context"
	"encoding/json"
	"net/http"
	"user-service/internal/controller/middleware"
	"user-service/internal/mapper"
)

func (c *controller) GetMe(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		id, err := middleware.GetUserUUIDFromContext(request.Context())
		if err != nil {
			c.errHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		user, err := c.userService.Get(id)
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
