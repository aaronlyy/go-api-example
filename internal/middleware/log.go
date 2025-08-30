package middleware

import (
	"fmt"
	"net/http"
)

// funktion bekommt den nächsten http handler welcher ausgeführt werden soll am ende
// muss in dem fall wrapped werden in http.HandlerFunc da die handler nur normale funktionen sind
// funktion returned den neuen handler in welcher next dann gewrapped ist, welcher im ende an .Handle kommt
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		// logik der middleware
		fmt.Printf("[gae] %s %s\n", r.Method, r.URL.Path)

		// call des nächsten handlers
		next.ServeHTTP(w, r)
	})
}