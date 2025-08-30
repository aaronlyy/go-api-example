package controller

import (
	"net/http"
	"github.com/aaronlyy/go-api-example/internal/response"
)

// create a new controller struct
type User struct {}

// attach method to controller, needs ResponseWriter and Request
func (c *User) Register(w http.ResponseWriter, r *http.Request) {
	// prepare response object
	var res = response.NewResponse(200, "New user created", nil)
	res.Send(w)
}

// attach method to controller, needs ResponseWriter and Request
func (c *User) Delete(w http.ResponseWriter, r *http.Request) {
	// prepare response object
	var res = response.NewResponse(200, "User deleted", nil)
	res.Send(w)
}