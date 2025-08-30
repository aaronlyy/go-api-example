package main

import (
	"net/http"
	"log"
	"github.com/aaronlyy/go-api-example/internal/controller"
	"github.com/aaronlyy/go-api-example/internal/middleware"
)

func main() {

	muxMain := http.NewServeMux() // create main router
	muxMisc := http.NewServeMux() // router for health, status
	muxAuth := http.NewServeMux() // router for authentication and authorization stuff
	muxUser := http.NewServeMux() // user creation and deletion

	// add subroutes to main router
	muxMain.Handle("/misc/", http.StripPrefix("/misc", muxMisc))
	muxMain.Handle("/auth/", http.StripPrefix("/auth", muxAuth))
	muxMain.Handle("/user/", http.StripPrefix("/user", muxUser))

	// create controller structs
	var misc = controller.Misc{}
	var auth = controller.Auth{}
	var user = controller.User{}

	// register misc handlers
	muxMisc.HandleFunc("GET /health", misc.Health)

	// register auth handlers
	muxAuth.HandleFunc("POST /authenticate", auth.Authenticate)

	// register user handlers
	muxUser.HandleFunc("POST /register", user.Register)
	muxUser.HandleFunc("POST /delete", user.Delete)

	
	var muxMainChained = middleware.Chain(muxMain, middleware.Log)

	// start server
	var err error = http.ListenAndServe("localhost:3000", muxMainChained)
	if err != nil {
		log.Fatal(http.ListenAndServe(":3000", muxMainChained))
	}
}