package main

import (
	"fmt"
	"net/http"
	"github.com/aaronlyy/go-api-example/internal/handler"
	"github.com/aaronlyy/go-api-example/internal/middleware"
)


func main() {

	mux := http.NewServeMux()

	mux.Handle("GET /hello", &handler.HelloGetHandler{})

	handler := middleware.LoggerMiddleware(mux)

	var err error = http.ListenAndServe("localhost:3000", handler)
	if err != nil {
		fmt.Println("An error occured")
	}
}