package response

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func NewResponse(status int, message string, data any) *Response {
	return &Response {
		status,
		message,
		data,
	}
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	var err = json.NewEncoder(w).Encode(r)
	if err != nil {
		fmt.Printf("error encoding response struct\n")
	}
}