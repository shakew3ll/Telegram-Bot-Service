package apperrors

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Err              error  `json:"-"`
	StatusCode       int    `json:"code,omitempty"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

var (
	ErrBadRequest      = NewAppError(http.StatusBadRequest, "Bad Request", "")
	ErrNotFound        = NewAppError(http.StatusNotFound, "Not Found", "Not Found :(")
	ErrConflict        = NewAppError(http.StatusConflict, "This email or phone_number already exists", "Conflict")
	ErrTeapot          = NewAppError(http.StatusTeapot, "I'm teapot", "")
	ErrTooManyRequests = NewAppError(http.StatusTooManyRequests, "Too many requests", "")
	ErrInternal        = NewAppError(http.StatusInternalServerError, "Internal server error", "")
	ErrBadGateway      = NewAppError(http.StatusBadGateway, "Bad gateway", "")
)

func NewAppError(statusCode int, message, developerMessage string) *AppError {
	return &AppError{
		StatusCode:       statusCode,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func BadRequestError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusBadRequest, message, developerMessage)
}

func ConflictError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusConflict, message, developerMessage)
}

func NotFoundError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusNotFound, message, developerMessage)
}

func TooManyRequestsError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusTooManyRequests, message, developerMessage)
}

func InternalServerError(message, developerMessage string) *AppError {
	return NewAppError(http.StatusInternalServerError, message, developerMessage)
}
