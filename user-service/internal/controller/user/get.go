package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (c *controller) Get(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		id, err := uuid.Parse(request.PathValue("id"))
		if err != nil {
			c.errHandler.ReturnIncorrectPathValueError(writer, err, request.RequestURI)
			return
		}

		user, err := c.userService.Get(id)
		if err != nil {
			c.errHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		if err = json.NewEncoder(writer).Encode(user); err != nil {
			c.errHandler.ReturnProcessingResponseError(writer, err, request.RequestURI)
			return
		}
	}
}
