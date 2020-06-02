package main

import (
	"flag"

	"github.com/ynori7/ipinfo/handler"
	"github.com/ynori7/ipinfo/internal/repository"
	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
)

func main() {
	var (
		geoDbPath  string
		serverHost string
	)

	errors.UseJsonErrors()

	flag.StringVar(&geoDbPath, "geo-db-path", "./GeoLite2-City.mmdb", "The path to the GeoLite2 city database")
	flag.StringVar(&serverHost, "server-host", ":8081", "The hostname and port that this API should run on")

	flag.Parse()

	geoLocationDb := repository.NewGeoLocationRepository(geoDbPath)

	handler.NewIpHandler(geoLocationDb)

	log.Info("Starting service")
	routing.ServeHttp(serverHost)

}
