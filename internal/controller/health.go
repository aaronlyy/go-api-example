package controller

import (
	"net/http"
	"github.com/aaronlyy/go-api-example/internal/response"
)

// create a new controller struct
type Health struct {}

// attach method to controller, needs ResponseWriter and Request
func (c *Health) Health(w http.ResponseWriter, r *http.Request) {
	// prepare response object
	var res = response.NewResponse(200, "Service is healty", nil)
	res.Send(w)
}