package handler

import (
	"encoding/json"
	"log"
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
		log.Println("Missing ip in request")
		ErrBadRequest(api.IpLookupResponse_INVALID_IP.String(), "Invalid IP", "Empty IP in request").WriteError(w)
		return
	}

	log.Printf("Handling request. ip=%s\n", ip)

	if !model.IsValidIpAddress(ip) {
		log.Printf("Invalid ip in request: %s\n", ip)
		ErrBadRequest(api.IpLookupResponse_INVALID_IP.String(), "Invalid IP","Invalid IP in request").WriteError(w)
		return
	}

	geoLocationData, err := h.GeoLocationRepository.GetGeoLocation(ip)
	if err != nil {
		log.Printf("Error getting geolocation data for %s: %s\n", ip, err.Error())
		ErrInternalServerError(api.IpLookupResponse_INTERNAL_ERROR.String(), "Internal Server Error,", "Something went wrong").WriteError(w)
		return
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
		log.Printf("Error marshalling response for %s: %s\n", ip, err.Error())
		ErrInternalServerError(api.IpLookupResponse_INTERNAL_ERROR.String(), "Internal Server Error,", "Something went wrong").WriteError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}
