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
	ReturnIncorrectReqParamError(writer http.ResponseWriter, path string)
	ReturnUnauthenticatedError(writer http.ResponseWriter, err error, path string)
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

func (h *handler) ReturnServiceError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, serviceError, 400, path)
}

func (h *handler) ReturnUnprocessableEntityError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, unprocessableEntity, 422, path)
}

func (h *handler) ReturnProcessingResponseError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, cannotProcessResponseEntity, 500, path)
}

func (h *handler) ReturnIncorrectReqParamError(writer http.ResponseWriter, path string) {
	h.returnError(writer, fmt.Errorf(incorrectReqParam), incorrectReqParam, 400, path)
}

func (h *handler) ReturnUnauthenticatedError(writer http.ResponseWriter, err error, path string) {
	h.returnError(writer, err, unauthenticated, 401, path)
}

func (h *handler) returnError(writer http.ResponseWriter, err error, message string, status int, path string) {
	response := Dto{
		Message:     message,
		MessageCode: h.source.GetErrorMessage(message),
		Error:       err.Error(),
		Path:        path,
		Status:      status,
		Timestamp:   time.Now().String(),
	}

	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(&response)
}
