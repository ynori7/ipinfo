package application

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/ynori7/ipinfo/api"
	"github.com/ynori7/ipinfo/internal/model"
	"github.com/ynori7/ipinfo/internal/repository"
)

var (
	ErrInvalidIp     = errors.New("invalid ip")
	ErrInternalError = errors.New("internal error")
)

func LookupIpData(ip string, logger log.FieldLogger, geoIpRepo *repository.GeoLocationRepository) (*api.IpLookupResponse, error) {
	if !model.IsValidIpAddress(ip) {
		logger.Debug("Invalid ip in request")
		return nil, ErrInvalidIp
	}

	geoLocationData, err := geoIpRepo.GetGeoLocation(ip)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting geolocation data")
		return nil, ErrInternalError
	}

	return &api.IpLookupResponse{
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
	}, nil
}
