package repository

import (
	"log"
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
	"github.com/ynori7/ipinfo/internal/model"
)

type GeoLocationRepository struct {
	db *geoip2.Reader
}

func NewGeoLocationRepository(dbPath string) *GeoLocationRepository {
	db, err := geoip2.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return &GeoLocationRepository{db: db}
}

func (r *GeoLocationRepository) GetGeoLocation(ip string) (*model.GeoLocation, error) {
	parsedIp := net.ParseIP(ip)

	record, err := r.db.City(parsedIp)
	if err != nil {
		return nil, err
	}

	return &model.GeoLocation{
		CountryCode: record.Country.IsoCode,
		Country:     record.Country.Names["en"],
		City:        record.City.Names["en"],
		Timezone:    record.Location.TimeZone,
		Location: &model.Coordinates{
			Latitude:  record.Location.Latitude,
			Longitude: record.Location.Longitude,
		},
	}, nil
}
