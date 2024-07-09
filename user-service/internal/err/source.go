package err

type ErrorMessageSource interface {
	GetErrorMessage(key string) string
}

type source struct {
	messages map[string]string
}

func NewMessageSource() *source {
	return &source{messages: map[string]string{
		unprocessableEntity:         "cannot process given json",
		cannotProcessResponseEntity: "cannot process response entity",
		serviceError:                "error while processing request",
		incorrectPathValue:          "incorrect path value",
	}}
}

func (s *source) GetErrorMessage(key string) string {
	return s.messages[key]
}
