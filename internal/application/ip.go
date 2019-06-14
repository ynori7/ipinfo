package application

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/ynori7/ipinfo/api"
	"github.com/ynori7/ipinfo/internal/model"
	"github.com/ynori7/ipinfo/internal/repository"
)

var (
	ErrInvalidIp       = errors.New("invalid ip")
	ErrInvalidHostname = errors.New("invalid hostname")
	ErrInternalError   = errors.New("internal error")
	ErrNotFound        = errors.New("not found")
)

func LookupIpData(ip string, logger log.FieldLogger, geoIpRepo *repository.GeoLocationRepository) (*api.LookupIpResponse, error) {
	if !model.IsValidIpAddress(ip) {
		logger.Debug("Invalid ip in request")
		return nil, ErrInvalidIp
	}

	geoLocationData, err := geoIpRepo.GetGeoLocation(ip)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting geolocation data")
		return nil, ErrInternalError
	}

	return &api.LookupIpResponse{
		Ip:        ip,
		Hostnames: model.GetHostnamesByIp(ip),
		Geolocation: &api.Location{
			CountryCode: geoLocationData.CountryCode,
			Country:     geoLocationData.Country,
			City:        geoLocationData.City,
			Timezone:    geoLocationData.Timezone,
			Lat:         geoLocationData.Location.Latitude,
			Long:        geoLocationData.Location.Longitude,
		},
	}, nil
}

func LookupHostData(hostname string, logger log.FieldLogger, geoIpRepo *repository.GeoLocationRepository) (*api.LookupHostResponse, error) {
	if !model.IsValidHostname(hostname) {
		logger.Debug("Invalid hostname in request")
		return nil, ErrInvalidHostname
	}

	ips := model.GetIpsByHostname(hostname)

	if len(ips) == 0 {
		logger.Debug("Hostname not found")
		return nil, ErrNotFound
	}

	geoLocationData, err := geoIpRepo.GetGeoLocation(ips[0])
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting geolocation data")
		return nil, ErrInternalError
	}

	return &api.LookupHostResponse{
		Ips:       ips,
		Hostnames: model.GetHostnamesByIpList(ips),
		Geolocation: &api.Location{
			CountryCode: geoLocationData.CountryCode,
			Country:     geoLocationData.Country,
			City:        geoLocationData.City,
			Timezone:    geoLocationData.Timezone,
			Lat:         geoLocationData.Location.Latitude,
			Long:        geoLocationData.Location.Longitude,
		},
	}, nil
}
