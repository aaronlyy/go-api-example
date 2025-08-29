package handler

import (
	"net/http"
	"fmt"
)


type HelloGetHandler struct {}

func (handler *HelloGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Hello stranger"}`)
}