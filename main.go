package main

import (
	"github.com/ynori7/ipinfo/internal/repository"
	"log"
	"net/http"

	"github.com/ynori7/ipinfo/handler"
)

func main() {
	geoLocationDb := repository.NewGeoLocationRepository("/home/sfinlay/Downloads/GeoLite2-City_20190611/GeoLite2-City.mmdb")

	ipHandler := handler.NewIpHandler(geoLocationDb)

	router := handler.NewRouter(ipHandler)

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}
