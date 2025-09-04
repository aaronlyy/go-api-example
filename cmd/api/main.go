package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/aaronlyy/go-api-example/internal/controller"
	"github.com/aaronlyy/go-api-example/internal/middleware"
	"github.com/aaronlyy/go-api-example/internal/util"
)

func main() {

	// load env vars
	_ = godotenv.Load()
	var PORT string = util.GetEnv("PORT")
	var ENV string = util.GetEnv("ENV")

	// create all needed subrouters
	muxMain := http.NewServeMux() // create main router
	muxAuth := http.NewServeMux() // router for authentication and authorization stuff
	muxUser := http.NewServeMux() // user creation and deletion

	// create controller structs
	var auth = controller.Auth{}
	var user = controller.User{}
	var health = controller.Health{}

	// main router handlers, no authentication required
	muxMain.HandleFunc("GET /health", health.Health)

	// auth handlers, no authentication required
	muxAuth.HandleFunc("POST /login", auth.Login)
	muxAuth.HandleFunc("POST /logout", auth.Logout)

	// user handlers
	muxUser.HandleFunc("POST /register", user.Register)
	muxUser.Handle(
		"PUT /deactivate/{uid}",
		middleware.Chain(
			http.HandlerFunc(user.Deactivate),
			middleware.Authenticate,
			middleware.Authorize("admin", "member"),
			),
		) // TODO: create ChainHandler & ChainFunc

	// add subrouters to main router
	muxMain.Handle("/auth/", http.StripPrefix("/auth", muxAuth))
	muxMain.Handle("/user/", http.StripPrefix("/user", muxUser))

	// chain main router
	muxMainChained := middleware.Chain(muxMain,middleware.Recover, middleware.Log, )

	// start server
	var addr string

	if (ENV == "DEV") {
		addr =  fmt.Sprintf("localhost:%s", PORT)
	} else {
		addr = fmt.Sprintf(":%s", PORT)
	}
	fmt.Printf("Server listening on\u001B[1;32m http://%s \u001B[0m\n", addr)
	log.Fatal(http.ListenAndServe(addr, muxMainChained))
}