package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ynori7/ipinfo/handler"
	"github.com/ynori7/ipinfo/internal/repository"
)

func main() {
	var (
		geoDbPath string
		serverHost string
	)

	flag.StringVar(&geoDbPath, "geo-db-path", "./GeoLite2-City.mmdb", "The path to the GeoLite2 city database")
	flag.StringVar(&serverHost, "server-host", ":8081", "The hostname and port that this API should run on")

	flag.Parse()

	geoLocationDb := repository.NewGeoLocationRepository(geoDbPath)

	ipHandler := handler.NewIpHandler(geoLocationDb)

	router := handler.NewRouter(ipHandler)

	if err := http.ListenAndServe(serverHost, router); err != nil {
		log.Fatal(err)
	}
}
