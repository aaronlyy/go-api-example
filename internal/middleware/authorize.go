package middleware

import (
	"net/http"
	"github.com/aaronlyy/go-api-example/internal/response"
	"github.com/aaronlyy/go-api-example/internal/util"
)

func Authorize(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			var roles []string
	  	var ok bool

			if roles, ok = r.Context().Value(rolesKey).([]string); !ok {
				response.NewResponse(500, "missing context in authorize", nil).Send(w)
				return
			}

			if !util.AnyMatchSlices(roles, requiredRoles) {
				response.NewResponse(401, "unauthorized", nil).Send(w)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}