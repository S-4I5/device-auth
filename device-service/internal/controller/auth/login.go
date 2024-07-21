package auth

import (
	"context"
	"device-service/internal/controller/middleware"
	"device-service/internal/model/dto"
	"encoding/json"
	"net/http"
)

// LoginUser allows to get user token using device token
//
//	@Summary      List accounts
//	@Description  get accounts
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Success      200  {array}   dto.LoginUserResponseDto
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /auth/login [post]
func (c *controller) LoginUser(ctx context.Context) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		id, err := middleware.GetSubjectUUIDFromContext(request.Context())
		if err != nil {
			c.errorHandler.ReturnServiceError(writer, err, request.RequestURI)
			return
		}

		var req dto.LoginUserRequestDto
		err = json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			c.errorHandler.ReturnUnprocessableEntityError(writer, err, request.RequestURI)
			return
		}

		resp, err := c.authService.LoginUser(req, id)
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
