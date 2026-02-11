package httperrors

import "fmt"

type APIError struct {
	Code    int
	Message string
	Source  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Source, e.Message)
}

func NotFoundError(source, message string) *APIError {
	return &APIError{
		Code:    404,
		Message: message,
		Source:  source,
	}
}

func BadRequestError(source, message string) *APIError {
	return &APIError{
		Code:    400,
		Message: message,
		Source:  source,
	}
}

func InternalServerError(source, message string) *APIError {
	return &APIError{
		Code:    500,
		Message: message,
		Source:  source,
	}
}

func MethodNotAllowedError(source, message string) *APIError {
	return &APIError{
		Code:    405,
		Message: message,
		Source:  source,
	}
}
