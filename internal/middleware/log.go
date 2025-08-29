package middleware

import (
	"fmt"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}