package api

import (
	"encoding/json"
	"net/http"
)

func ErrNotFound(errorCode string, title, message string) *ErrorResponse {
	return NewErrorResponse(http.StatusNotFound, errorCode, title, message, false)
}

func ErrBadRequest(errorCode string, title, message string) *ErrorResponse {
	return NewErrorResponse(http.StatusBadRequest, errorCode, title, message, false)
}

func ErrInternalServerError(errorCode string, title, message string) *ErrorResponse {
	return NewErrorResponse(http.StatusInternalServerError, errorCode, title, message, true)
}

// WithMessage returns a copy of a given error and overrides its messages
func (e *ErrorResponse) WithMessage(message string) *ErrorResponse {
	return NewErrorResponse(e.StatusCode, e.ErrorCode, e.Title, message, e.Retriable)
}

// WithErrorCode returns a copy of a given error and overrides its messages
func (e *ErrorResponse) WithErrorCode(errorCode string) *ErrorResponse {
	return NewErrorResponse(e.StatusCode, errorCode, e.Title, e.Message, e.Retriable)
}

// WithTitle returns a copy of a given error and overrides its title
func (e *ErrorResponse) WithTitle(title string) *ErrorResponse {
	return NewErrorResponse(e.StatusCode, e.ErrorCode, title, e.Message, e.Retriable)
}

func (err *ErrorResponse) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

func (err *ErrorResponse) WriteResponse(w http.ResponseWriter) error {
	jsonRes, _ := json.Marshal(err)
	w.WriteHeader(int(err.StatusCode))
	_, e := w.Write(jsonRes)
	return e
}

func NewErrorResponse(statusCode int32, errorCode string, title string, message string, retriable bool) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Title:      title,
		Message:    message,
		Retriable:  retriable,
	}
}
