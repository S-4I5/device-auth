package httperr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ErrorHandler interface {
	ReturnUnprocessableEntityError(writer http.ResponseWriter, err error, path string)
	ReturnIncorrectPathValueError(writer http.ResponseWriter, err error, path string)
	ReturnProcessingResponseError(writer http.ResponseWriter, err error, path string)
	ReturnServiceError(writer http.ResponseWriter, err error, path string)
	ReturnAuthenticationError(writer http.ResponseWriter, err error, path string)
}

type handler struct {
	source ErrorMessageSource
}

func NewHandler(source ErrorMessageSource) *handler {
	return &handler{source: source}
}

func (h *handler) ReturnIncorrectPathValueError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, incorrectPathValue, 400, path)
}

func (h *handler) ReturnAuthenticationError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, authorizationError, 401, path)
}

func (h *handler) ReturnServiceError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, serviceError, 400, path)
}

func (h *handler) ReturnUnprocessableEntityError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, unprocessableEntity, 422, path)
}

func (h *handler) ReturnProcessingResponseError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, cannotProcessResponseEntity, 500, path)
}

func (h *handler) returnError(writer http.ResponseWriter, err error, message string, status int, path string) {
	response := Dto{
		Error:       err.Error(),
		Message:     message,
		MessageCode: h.source.GetErrorMessage(message),
		Path:        path,
		Timestamp:   time.Now().String(),
		Status:      status,
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(status)
	if err = json.NewEncoder(writer).Encode(&response); err != nil {
		fmt.Println(err)
	}
}
