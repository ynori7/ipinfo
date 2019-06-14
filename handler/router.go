package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/ynori7/ipinfo/api"
)

type HttpEndpoint func(http.ResponseWriter, *http.Request) api.HttpResponseWriter

func NewRouter(
	ipHandler *IpHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ip", HandlerWrapper(ipHandler.WhatsMyIp)).
		Methods("GET")

	r.HandleFunc("/ip/{ipAddress}", HandlerWrapper(ipHandler.LookupIp)).
		Methods("GET")

	r.HandleFunc("/host/{hostname}", HandlerWrapper(ipHandler.LookupHost)).
		Methods("GET")

	return r
}

func HandlerWrapper(f HttpEndpoint) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		respWriter := f(writer, req)
		if err := respWriter.WriteResponse(writer); err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error writing response")
			GenericError.WriteResponse(writer)
		}
	}
}
