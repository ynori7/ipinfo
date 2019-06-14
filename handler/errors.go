package handler

import (
	"github.com/ynori7/ipinfo/api"
)

type ApiEndpoint int

const (
	UNKNOWN_ENDPOINT ApiEndpoint = iota
	LOOKUP_IP
	WHATS_MY_IP
)

const (
	InternalErrorMessage = "Internal Server Error"
)

const (
	InternalError = "InternalError"
	MissingIp     = "MissingIp"
	InvalidIp     = "InvalidIp"
)

var ErrorMapping = map[ApiEndpoint]map[string]*api.ErrorResponse{
	LOOKUP_IP: {
		MissingIp:     api.ErrBadRequest(api.IpLookupResponse_INVALID_IP.String(), "Invalid IP", "Empty IP in request"),
		InvalidIp:     api.ErrBadRequest(api.IpLookupResponse_INVALID_IP.String(), "Invalid IP", "Invalid IP in request"),
		InternalError: api.ErrInternalServerError(api.IpLookupResponse_INTERNAL_ERROR.String(), "Internal Server Error", "Something went wrong"),
	},
	WHATS_MY_IP: {
		InvalidIp:     api.ErrBadRequest(api.WhatsMyIpResponse_INVALID_IP.String(), "Invalid IP", "Invalid IP in request"),
		InternalError: api.ErrInternalServerError(api.WhatsMyIpResponse_INTERNAL_ERROR.String(), "Internal Server Error", "Something went wrong"),
	},
}

func GetMappedError(endpoint ApiEndpoint, errorKey string) *api.ErrorResponse {
	err, ok := ErrorMapping[endpoint][errorKey]
	if !ok {
		errorKey = InternalError
		err, ok = ErrorMapping[endpoint][errorKey]
		if !ok {
			return api.ErrInternalServerError("", "", InternalErrorMessage)
		}
	}

	return err
}
