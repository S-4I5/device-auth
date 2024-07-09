package err

const (
	unprocessableEntity         = "api.unprocessable_entity"
	cannotProcessResponseEntity = "api.cannot_process_responseEntity"
	serviceError                = "api.service_error"
	incorrectPathValue          = "api.incorrect_path_value"
	incorrectReqParam           = "api.incorrect_request_param"
	unauthenticated             = "api.unauthenticated"
)

type Dto struct {
	Error       string `json:"err"`
	Message     string `json:"message"`
	MessageCode string `json:"messageCode"`
}
