package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/ynori7/ipinfo/api"
	"github.com/ynori7/ipinfo/internal/model"
	"github.com/ynori7/ipinfo/internal/repository"
)

type IpHandler struct {
	GeoLocationRepository *repository.GeoLocationRepository
}

func NewIpHandler(geoIpRepo *repository.GeoLocationRepository) *IpHandler {
	return &IpHandler{GeoLocationRepository: geoIpRepo}
}

func (h *IpHandler) LookupIp(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"Handler": "LookupIp"})

	vars := mux.Vars(r)
	ip, ok := vars["ipAddress"]
	if !ok {
		logger.Debug("Missing ip in request")
		GetMappedError(LOOKUP_IP, MissingIp).WriteError(w)
		return
	}

	logger = logger.WithFields(log.Fields{"IP": ip})
	logger.Info("Handling request")

	if !model.IsValidIpAddress(ip) {
		logger.Debug("Invalid ip in request")
		GetMappedError(LOOKUP_IP, InvalidIp).WriteError(w)
		return
	}

	geoLocationData, err := h.GeoLocationRepository.GetGeoLocation(ip)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting geolocation data")
		GetMappedError(LOOKUP_IP, InternalError).WriteError(w)
		return
	}

	ipData := &api.IpLookupResponse{
		Ip:        ip,
		Hostnames: model.ResolveHost(ip),
		Geolocation: &api.Location{
			CountryCode: geoLocationData.CountryCode,
			Country:     geoLocationData.Country,
			City:        geoLocationData.City,
			Timezone:    geoLocationData.Timezone,
			Lat:         geoLocationData.Location.Latitude,
			Long:        geoLocationData.Location.Longitude,
		},
	}

	jsonRes, err := json.Marshal(ipData)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error marshalling response")
		GetMappedError(LOOKUP_IP, InternalError).WriteError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func (h *IpHandler) WhatsMyIp(w http.ResponseWriter, r *http.Request) {
	ip := model.GetIpFromRequest(r)
	logger := log.WithFields(log.Fields{"Handler": "WhatsMyIp", "IP": ip.IpAddress})

	logger.Info("Handling request")

	if !model.IsValidIpAddress(ip.IpAddress) {
		logger.Debug("Invalid ip in request")
		GetMappedError(WHATS_MY_IP, InvalidIp).WriteError(w)
		return
	}

	geoLocationData, err := h.GeoLocationRepository.GetGeoLocation(ip.IpAddress)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting geolocation data")
		GetMappedError(WHATS_MY_IP, InternalError).WriteError(w)
		return
	}

	ipData := &api.WhatsMyIpResponse{
		Ip:           ip.IpAddress,
		ForwardedFor: ip.ForwardedFor,
		Hostnames:    model.ResolveHost(ip.IpAddress),
		Geolocation: &api.Location{
			CountryCode: geoLocationData.CountryCode,
			Country:     geoLocationData.Country,
			City:        geoLocationData.City,
			Timezone:    geoLocationData.Timezone,
			Lat:         geoLocationData.Location.Latitude,
			Long:        geoLocationData.Location.Longitude,
		},
	}

	jsonRes, err := json.Marshal(ipData)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error marshalling response")
		GetMappedError(WHATS_MY_IP, InternalError).WriteError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
