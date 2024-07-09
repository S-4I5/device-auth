package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"user-service/internal/jwt"
	"user-service/internal/mapper"
)

func (c *controller) GetMe(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		token, err := c.tokenParser.ValidateToken(strings.TrimPrefix(request.Header.Get("Authorization"), "Bearer "))
		if err != nil {
			fmt.Println("XD", err)
			return
		}

		id, err := jwt.GetSubjectIdFromToken(*token)
		if err != nil {
			fmt.Println("XD", err)
			return
		}

		user, err := c.userService.Get(id)
		if err != nil {
			c.errHandler.ReturnServiceError(writer, err)
			return
		}

		if err = json.NewEncoder(writer).Encode(mapper.UserToUserDto(user)); err != nil {
			c.errHandler.ReturnProcessingResponseError(writer, err)
			return
		}
	}
}
