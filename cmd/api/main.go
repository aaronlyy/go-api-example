package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aaronlyy/go-api-example/internal/controller"
	"github.com/aaronlyy/go-api-example/internal/middleware"
	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return "n/a"
	}
	return value
}

func main() {

	// load env vars
	_ = godotenv.Load()
	var PORT string = getEnv("PORT")

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

	// wrap main mux in middleware
	var muxMainChained = middleware.Chain(
		muxMain,
		middleware.Recover,
		middleware.Log,
	)

	// start server
	var addr string = fmt.Sprintf("localhost:%s", PORT)
	fmt.Printf("Server listening on\u001B[1;32m http://%s \u001B[0m\n", addr)
	log.Fatal(http.ListenAndServe(addr, muxMainChained))

}
