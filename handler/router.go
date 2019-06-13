package handler

import (
	"github.com/gorilla/mux"
)

func NewRouter(
	ipHandler *IpHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ip/{ipAddress}", ipHandler.Lookup).
		Methods("GET")

	return r
}
