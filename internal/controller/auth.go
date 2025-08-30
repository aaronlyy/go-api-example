package controller

import (
	"net/http"
	"github.com/aaronlyy/go-api-example/internal/response"
)

// create a new controller struct
type Auth struct {}

// attach method to controller, needs ResponseWriter and Request
func (c *Auth) Authenticate(w http.ResponseWriter, r *http.Request) {
	// prepare response object
	var res = response.NewResponse(200, "Authorized", nil)
	res.Send(w)
}