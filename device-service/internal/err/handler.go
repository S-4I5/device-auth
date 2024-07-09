package err

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorHandler interface {
	ReturnUnprocessableEntityError(writer http.ResponseWriter, err error)
	ReturnIncorrectPathValueError(writer http.ResponseWriter, err error)
	ReturnProcessingResponseError(writer http.ResponseWriter, err error)
	ReturnServiceError(writer http.ResponseWriter, err error)
	ReturnIncorrectReqParamError(writer http.ResponseWriter)
	ReturnUnauthenticatedError(writer http.ResponseWriter, err error)
}

type handler struct {
	source ErrorMessageSource
}

func NewHandler(source ErrorMessageSource) *handler {
	return &handler{source: source}
}

func (h *handler) ReturnIncorrectPathValueError(writer http.ResponseWriter, err error) {
	h.returnError(writer, err, incorrectPathValue, 400)
}

func (h *handler) ReturnServiceError(writer http.ResponseWriter, err error) {
	h.returnError(writer, err, serviceError, 400)
}

func (h *handler) ReturnUnprocessableEntityError(writer http.ResponseWriter, err error) {
	h.returnError(writer, err, unprocessableEntity, 422)
}

func (h *handler) ReturnProcessingResponseError(writer http.ResponseWriter, err error) {
	h.returnError(writer, err, cannotProcessResponseEntity, 500)
}

func (h *handler) ReturnIncorrectReqParamError(writer http.ResponseWriter) {
	h.returnError(writer, fmt.Errorf(incorrectReqParam), incorrectReqParam, 400)
}

func (h *handler) ReturnUnauthenticatedError(writer http.ResponseWriter, err error) {
	h.returnError(writer, err, unauthenticated, 401)
}

func (h *handler) returnError(writer http.ResponseWriter, err error, message string, status int) {
	response := Dto{
		Message:     message,
		MessageCode: h.source.GetErrorMessage(message),
		Error:       err.Error(),
	}

	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(&response)
}
