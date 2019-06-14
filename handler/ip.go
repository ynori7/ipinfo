package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/ynori7/ipinfo/api"
	"github.com/ynori7/ipinfo/internal/application"
	"github.com/ynori7/ipinfo/internal/model"
	"github.com/ynori7/ipinfo/internal/repository"
)

type IpHandler struct {
	GeoLocationRepository *repository.GeoLocationRepository
}

func NewIpHandler(geoIpRepo *repository.GeoLocationRepository) *IpHandler {
	return &IpHandler{GeoLocationRepository: geoIpRepo}
}

func (h *IpHandler) LookupIp(w http.ResponseWriter, r *http.Request) api.HttpResponseWriter {
	logger := log.WithFields(log.Fields{"Handler": "LookupIp"})

	vars := mux.Vars(r)
	ip, ok := vars["ipAddress"]
	if !ok {
		logger.Debug("Missing ip in request")
		return GetMappedError(LOOKUP_IP, MissingIp)
	}

	logger = logger.WithFields(log.Fields{"IP": ip})
	logger.Info("Handling request")

	ipData, err := application.LookupIpData(ip, logger, h.GeoLocationRepository)
	if err != nil {
		switch err {
		case application.ErrInvalidIp:
			return GetMappedError(LOOKUP_IP, InvalidIp)
		default:
			return GetMappedError(LOOKUP_IP, InternalError)
		}
	}

	return &api.HttpResponse{Payload: ipData, Status: http.StatusOK}
}

func (h *IpHandler) WhatsMyIp(w http.ResponseWriter, r *http.Request) api.HttpResponseWriter {
	ip := model.GetIpFromRequest(r)
	logger := log.WithFields(log.Fields{"Handler": "WhatsMyIp", "IP": ip.IpAddress})

	logger.Info("Handling request")

	ipData, err := application.LookupIpData(ip.IpAddress, logger, h.GeoLocationRepository)
	if err != nil {
		switch err {
		case application.ErrInvalidIp:
			return GetMappedError(WHATS_MY_IP, InvalidIp)
		default:
			return GetMappedError(WHATS_MY_IP, InternalError)
		}
	}

	resp := &api.WhatsMyIpResponse{
		Ip:           ipData.Ip,
		ForwardedFor: ip.ForwardedFor,
		Hostnames:    ipData.Hostnames,
		Geolocation:  ipData.Geolocation,
	}

	return &api.HttpResponse{Payload: resp, Status: http.StatusOK}
}

func (h *IpHandler) LookupHost(w http.ResponseWriter, r *http.Request) api.HttpResponseWriter {
	logger := log.WithFields(log.Fields{"Handler": "LookupHost"})

	vars := mux.Vars(r)
	hostname, ok := vars["hostname"]
	if !ok {
		logger.Debug("Missing hostname in request")
		return GetMappedError(LOOKUP_HOST, MissingIp)
	}

	logger = logger.WithFields(log.Fields{"Hostname": hostname})
	logger.Info("Handling request")

	ipData, err := application.LookupHostData(hostname, logger, h.GeoLocationRepository)
	if err != nil {
		switch err {
		case application.ErrInvalidIp:
			return GetMappedError(LOOKUP_HOST, InvalidIp)
		case application.ErrNotFound:
			return GetMappedError(LOOKUP_HOST, NotFound)
		default:
			return GetMappedError(LOOKUP_HOST, InternalError)
		}
	}

	return &api.HttpResponse{Payload: ipData, Status: http.StatusOK}
}