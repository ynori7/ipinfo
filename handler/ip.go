package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

func (h *IpHandler) Lookup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip, ok := vars["ipAddress"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"missing ip"}`)) //TODO: errors
		return
	}

	if !model.IsValidIpAddress(ip) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid ip"}`)) //TODO: errors
		return
	}

	geoLocationData, err := h.GeoLocationRepository.GetGeoLocation(ip)
	if err != nil {
		//TODO: logging
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"invalid json"}`)) //TODO: errors
	}


	ipData := &api.IpLookupResponse{
		Ip:          ip,
		Hostnames:   model.ResolveHost(ip),
		Geolocation: &api.Location{
			CountryCode: geoLocationData.CountryCode,
			Country: geoLocationData.Country,
			City: geoLocationData.City,
			Timezone: geoLocationData.Timezone,
			Lat: geoLocationData.Location.Latitude,
			Long: geoLocationData.Location.Longitude,
		},
	}

	jsonRes, err := json.Marshal(ipData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"invalid json"}`)) //TODO: errors
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
