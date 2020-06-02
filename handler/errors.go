package handler

import (
	"github.com/ynori7/ipinfo/api"
	"github.com/ynori7/lilypad/errors"
)

type ApiEndpoint int

const (
	UNKNOWN_ENDPOINT ApiEndpoint = iota
	LOOKUP_IP
	WHATS_MY_IP
	LOOKUP_HOST
)

var GenericError = errors.InternalServerError(InternalErrorMessage).WithTitle(InternalErrorMessage)

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

var ErrorMapping = map[ApiEndpoint]map[string]errors.HttpError{
	LOOKUP_IP: {
		MissingIp:     errors.BadRequestError("Empty IP in request").WithTitle("Invalid IP").WithCode(api.LookupIpResponse_INVALID_IP.String()),
		InvalidIp:     errors.BadRequestError("Invalid IP in request").WithTitle("Invalid IP").WithCode(api.LookupIpResponse_INVALID_IP.String()),
		InternalError: errors.InternalServerError("Something went wrong").WithCode(api.LookupIpResponse_INTERNAL_ERROR.String()).WithTitle(InternalErrorMessage),
	},
	WHATS_MY_IP: {
		InvalidIp:     errors.BadRequestError("Invalid IP in request").WithCode(api.WhatsMyIpResponse_INVALID_IP.String()).WithTitle("Invalid IP"),
		InternalError: errors.InternalServerError("Something went wrong").WithCode(api.WhatsMyIpResponse_INTERNAL_ERROR.String()).WithTitle(InternalErrorMessage),
	},
	LOOKUP_HOST: {
		MissingHost:   errors.BadRequestError("Empty hostname in request").WithCode(api.LookupHostResponse_INVALID_HOSTNAME.String()).WithTitle("Invalid Hostname"),
		InvalidHost:   errors.BadRequestError("Invalid hostname in request").WithCode(api.LookupHostResponse_INVALID_HOSTNAME.String()).WithTitle("Invalid hostname"),
		NotFound:      errors.NotFoundError("The hostname was not found").WithCode(api.LookupHostResponse_NOT_FOUND.String()).WithTitle("Hostname not found"),
		InternalError: errors.InternalServerError("Something went wrong").WithCode(api.LookupHostResponse_INTERNAL_ERROR.String()).WithTitle(InternalErrorMessage),
	},
}

func GetMappedError(endpoint ApiEndpoint, errorKey string) errors.HttpError {
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
