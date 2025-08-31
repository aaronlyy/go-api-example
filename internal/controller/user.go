package controller

import (
	"fmt"
	"net/http"

	"github.com/aaronlyy/go-api-example/internal/response"
	"github.com/aaronlyy/go-api-example/internal/util"
)

// structs to parse body in
type CreateUserRequestBody struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// create a new controller struct
type User struct {}

// attach method to controller, needs ResponseWriter and Request
func (c *User) Register(w http.ResponseWriter, r *http.Request) {

	// create new struct for parsing request body
	var rb CreateUserRequestBody

	// parse request body to struct
	if err := util.ParseBody(r.Body, &rb); err != nil {
		response.NewResponse(500, "error parsing body or missing data", nil).Send(w)
		return
	}

	fmt.Printf("New user created: %s", rb.Username)

	response.NewResponse(201, "New user created", nil).Send(w)
}

// attach method to controller, needs ResponseWriter and Request
func (c *User) Delete(w http.ResponseWriter, r *http.Request) {
	// prepare response object
	response.NewResponse(200, "User deleted", nil).Send(w)
}