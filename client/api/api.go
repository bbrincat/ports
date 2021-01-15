package api

import "net/http"

type Api struct {
	//TODO add dependencies (So they can be mocked in tests)
}

func GetPort(w http.ResponseWriter, r *http.Request) {
	//TODO implement fetching of port for service.
	w.WriteHeader(200)
}
