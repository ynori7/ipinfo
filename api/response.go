package api

import (
	"encoding/json"
	"net/http"
)

type HttpResponseWriter interface {
	WriteResponse(w http.ResponseWriter) error
}

type HttpResponse struct {
	Status  int
	Payload interface{}
}

func (resp *HttpResponse) WriteResponse(w http.ResponseWriter) error {
	jsonRes, err := json.Marshal(resp.Payload)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonRes)
	return err
}
