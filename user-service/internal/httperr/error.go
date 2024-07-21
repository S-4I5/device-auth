package httperr

const (
	unprocessableEntity         = "api.unprocessable_entity"
	cannotProcessResponseEntity = "api.cannot_process_responseEntity"
	serviceError                = "api.service_error"
	incorrectPathValue          = "api.incorrect_path_value"
	authorizationError          = "api.authorization_error"
)

type Dto struct {
	Error       string `json:"error"`
	Message     string `json:"message"`
	MessageCode string `json:"messageCode"`
	Path        string `json:"path"`
	Timestamp   string `json:"timestamp"`
	Status      int    `json:"status"`
}
