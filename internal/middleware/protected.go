package middleware

import (
	"net/http"
)


func Protected(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// get cookie and verify jwt
		

		// if verified: add user id and role to request
		// if not veriefied, return error
		// continue
	})
}