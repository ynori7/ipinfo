package handler

import (
	"github.com/gorilla/mux"
)

func NewRouter(
	ipHandler *IpHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ip", ipHandler.WhatsMyIp).
		Methods("GET")

	r.HandleFunc("/ip/{ipAddress}", ipHandler.LookupIp).
		Methods("GET")

	return r
}
