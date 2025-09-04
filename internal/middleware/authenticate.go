package middleware

import (
	"context"
	"net/http"
	"github.com/aaronlyy/go-api-example/internal/auth"
	"github.com/aaronlyy/go-api-example/internal/response"
)

type ctxKey string

const userIdKey ctxKey = "uid"
const rolesKey ctxKey = "roles"

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get access token from cookie
		cookie, err := r.Cookie("access_token")

		if err != nil {
			response.NewResponse(401, "missing access token", nil).Send(w)
			return
		}

		// verify access token
		claims, err := auth.ParseAccessToken(cookie.Value)

		if err != nil {
			response.NewResponse(401, "could not verify access token", nil).Send(w)
			return
		}

		// get user id and roles out of token
		roles := claims.Roles
		uid := claims.Subject

		ctx := context.WithValue(
			context.WithValue(r.Context(), userIdKey, uid),
			rolesKey, roles,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
