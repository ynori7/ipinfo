package handler

import (
	"github.com/ynori7/ipinfo/api"
)

type ApiEndpoint int

const (
	UNKNOWN_ENDPOINT ApiEndpoint = iota
	LOOKUP_IP
	WHATS_MY_IP
	LOOKUP_HOST
)

var GenericError = api.ErrInternalServerError("", InternalErrorMessage, InternalErrorMessage)

const (
	InternalErrorMessage = "Internal Server Error"
)

const (
	InternalError = "InternalError"
	MissingIp     = "MissingIp"
	InvalidIp     = "InvalidIp"
	MissingHost   = "MissingHost"
	InvalidHost   = "InvalidHost"
	NotFound      = "NotFound"
)

var ErrorMapping = map[ApiEndpoint]map[string]*api.ErrorResponse{
	LOOKUP_IP: {
		MissingIp:     api.ErrBadRequest(api.LookupIpResponse_INVALID_IP.String(), "Invalid IP", "Empty IP in request"),
		InvalidIp:     api.ErrBadRequest(api.LookupIpResponse_INVALID_IP.String(), "Invalid IP", "Invalid IP in request"),
		InternalError: api.ErrInternalServerError(api.LookupIpResponse_INTERNAL_ERROR.String(), InternalErrorMessage, "Something went wrong"),
	},
	WHATS_MY_IP: {
		InvalidIp:     api.ErrBadRequest(api.WhatsMyIpResponse_INVALID_IP.String(), "Invalid IP", "Invalid IP in request"),
		InternalError: api.ErrInternalServerError(api.WhatsMyIpResponse_INTERNAL_ERROR.String(), InternalErrorMessage, "Something went wrong"),
	},
	LOOKUP_HOST: {
		MissingHost:   api.ErrBadRequest(api.LookupHostResponse_INVALID_HOSTNAME.String(), "Invalid Hostname", "Empty hostname in request"),
		InvalidHost:   api.ErrBadRequest(api.LookupHostResponse_INVALID_HOSTNAME.String(), "Invalid Hostname", "Invalid hostname in request"),
		NotFound:      api.ErrNotFound(api.LookupHostResponse_NOT_FOUND.String(), "Hostname Not Found", "The hostname was not found"),
		InternalError: api.ErrInternalServerError(api.LookupHostResponse_INTERNAL_ERROR.String(), InternalErrorMessage, "Something went wrong"),
	},
}

func GetMappedError(endpoint ApiEndpoint, errorKey string) *api.ErrorResponse {
	err, ok := ErrorMapping[endpoint][errorKey]
	if !ok {
		errorKey = InternalError
		err, ok = ErrorMapping[endpoint][errorKey]
		if !ok {
			return GenericError
		}
	}

	return err
}
