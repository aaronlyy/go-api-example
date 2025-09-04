package middleware

import (
	"net/http"
	"fmt"
	"github.com/aaronlyy/go-api-example/internal/response"
)


func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {

		var uid string
		var roles []string
	  var ok bool

		if uid, ok = r.Context().Value(userIdKey).(string); !ok {
			response.NewResponse(500, "missing context in authorize", nil).Send(w)
			return
		}

		if roles, ok = r.Context().Value(rolesKey).([]string); !ok {
			response.NewResponse(500, "missing context in authorize", nil).Send(w)
			return
		}

		fmt.Printf("uid: %s, roles: %v", uid, roles)

		next.ServeHTTP(w, r)
	})
}