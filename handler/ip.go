package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ynori7/ipinfo/api"
	"github.com/ynori7/ipinfo/internal/application"
	"github.com/ynori7/ipinfo/internal/model"
	"github.com/ynori7/ipinfo/internal/repository"
	"github.com/ynori7/lilypad/handler"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
)

type IpHandler struct {
	GeoLocationRepository *repository.GeoLocationRepository
}

func NewIpHandler(geoIpRepo *repository.GeoLocationRepository) *IpHandler {
	h := &IpHandler{GeoLocationRepository: geoIpRepo}

	routing.RegisterRoutes([]routing.Route{
		{
			Method: "GET",
			Path: "/ip",
			Handler: h.WhatsMyIp,
		},
		{
			Method: "GET",
			Path: "/ip/{ipAddress}",
			Handler: h.LookupIp,
		},
		{
			Method: "GET",
			Path: "/host/{hostname}",
			Handler: h.LookupHost,
		},
	}...)

	return h
}

func (h *IpHandler) LookupIp(r *http.Request) handler.Response {
	logger := log.WithFields(log.Fields{"Handler": "LookupIp"})

	vars := routing.Vars(r)
	ip, ok := vars["ipAddress"]
	if !ok {
		logger.Debug("Missing ip in request")
		return handler.ErrorResponse(GetMappedError(LOOKUP_IP, MissingIp))
	}

	logger = logger.WithFields(log.Fields{"IP": ip})
	logger.Info("Handling request")

	ipData, err := application.LookupIpData(ip, logger, h.GeoLocationRepository)
	if err != nil {
		switch err {
		case application.ErrInvalidIp:
			return handler.ErrorResponse(GetMappedError(LOOKUP_IP, InvalidIp))
		default:
			return handler.ErrorResponse(GetMappedError(LOOKUP_IP, InternalError))
		}
	}

	jsonRes, err := json.Marshal(ipData)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error marshaling json")
		return handler.ErrorResponse(GetMappedError(LOOKUP_IP, InternalError))
	}

	return handler.SuccessResponse(jsonRes)
}

func (h *IpHandler) WhatsMyIp(r *http.Request) handler.Response {
	ip := model.GetIpFromRequest(r)
	logger := log.WithFields(log.Fields{"Handler": "WhatsMyIp", "IP": ip.IpAddress, "ForwardedFor": ip.ForwardedFor})

	logger.Info("Handling request")

	ipData, err := application.LookupIpData(ip.IpAddress, logger, h.GeoLocationRepository)
	if err != nil {
		switch err {
		case application.ErrInvalidIp:
			return handler.ErrorResponse(GetMappedError(WHATS_MY_IP, InvalidIp))
		default:
			return handler.ErrorResponse(GetMappedError(WHATS_MY_IP, InternalError))
		}
	}

	resp := &api.WhatsMyIpResponse{
		Ip:           ipData.Ip,
		ForwardedFor: ip.ForwardedFor,
		Hostnames:    ipData.Hostnames,
		Geolocation:  ipData.Geolocation,
	}

	jsonRes, err := json.Marshal(resp)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error marshaling json")
		return handler.ErrorResponse(GetMappedError(WHATS_MY_IP, InternalError))
	}

	return handler.SuccessResponse(jsonRes)
}

func (h *IpHandler) LookupHost(r *http.Request) handler.Response {
	logger := log.WithFields(log.Fields{"Handler": "LookupHost"})

	vars := routing.Vars(r)
	hostname, ok := vars["hostname"]
	if !ok {
		logger.Debug("Missing hostname in request")
		return handler.ErrorResponse(GetMappedError(LOOKUP_HOST, MissingIp))
	}

	logger = logger.WithFields(log.Fields{"Hostname": hostname})
	logger.Info("Handling request")

	ipData, err := application.LookupHostData(hostname, logger, h.GeoLocationRepository)
	if err != nil {
		switch err {
		case application.ErrInvalidIp:
			return handler.ErrorResponse(GetMappedError(LOOKUP_HOST, InvalidIp))
		case application.ErrNotFound:
			return handler.ErrorResponse(GetMappedError(LOOKUP_HOST, NotFound))
		default:
			return handler.ErrorResponse(GetMappedError(LOOKUP_HOST, InternalError))
		}
	}

	jsonRes, err := json.Marshal(ipData)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error marshaling json")
		return handler.ErrorResponse(GetMappedError(LOOKUP_HOST, InternalError))
	}

	return handler.SuccessResponse(jsonRes)
}