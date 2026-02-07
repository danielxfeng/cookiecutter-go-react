package apierror

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(status int, message string) *ApiError {
	return &ApiError{
		Status:  status,
		Message: message,
	}
}
