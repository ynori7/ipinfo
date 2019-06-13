package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ynori7/ipinfo/api"
)

type Error api.ErrorResponse

func ErrNotFound(errorCode string, title, message string) *Error {
	return NewErrorResponse(http.StatusNotFound, "", "", message, false)
}

func ErrBadRequest(errorCode string, title, message string) *Error {
	return NewErrorResponse(http.StatusBadRequest, errorCode, "", message, false)
}

func ErrInternalServerError(errorCode string, title, message string) *Error {
	return NewErrorResponse(http.StatusInternalServerError, "", "", message, true)
}

// WithMessage returns a copy of a given error and overrides its messages
func (e *Error) WithMessage(message string) *Error {
	return NewErrorResponse(e.StatusCode, e.ErrorCode, e.Title, message, e.Retriable)
}

// WithErrorCode returns a copy of a given error and overrides its messages
func (e *Error) WithErrorCode(errorCode string) *Error {
	return NewErrorResponse(e.StatusCode, errorCode, e.Title, e.Message, e.Retriable)
}

// WithTitle returns a copy of a given error and overrides its title
func (e *Error) WithTitle(title string) *Error {
	return NewErrorResponse(e.StatusCode, e.ErrorCode, title, e.Message, e.Retriable)
}

func (err *Error) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

func (err *Error) WriteError(w http.ResponseWriter) {
	jsonRes, _ := json.Marshal(err)
	w.WriteHeader(int(err.StatusCode))
	w.Write(jsonRes)
}

func NewErrorResponse(statusCode int32, errorCode string, title string, message string, retriable bool) *Error {
	return &Error{
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Title:      title,
		Message:    message,
		Retriable:  retriable,
	}
}
