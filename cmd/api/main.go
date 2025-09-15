package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/aaronlyy/go-api-example/internal/controller"
	"github.com/aaronlyy/go-api-example/internal/middleware"
	"github.com/aaronlyy/go-api-example/internal/util"
)

func main() {

	ctx := context.Background()

	// load env vars
	_ = godotenv.Load()
	var PORT string = util.GetEnv("PORT")
	var ENV string = util.GetEnv("ENV")
	var DBURL string = util.GetEnv("DBURL")

	// connect to database
	cfg, err := pgxpool.ParseConfig(DBURL)
	if err != nil {
		log.Fatalf("pgxpool config creation failed: %s\n", err.Error())
	}
	cfg.MaxConns = 10
	cfg.MinConns = 1
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("connection to database failed: %s\n", err.Error())
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("ping to database failed: %s\n", err.Error())
	}

	// create all needed subrouters
	muxMain := http.NewServeMux() // create main router
	muxAuth := http.NewServeMux() // router for authentication and authorization stuff
	muxUser := http.NewServeMux() // user creation and deletion

	// create controller structs
	var authController = controller.NewAuthController(pool)
	var usersController = controller.NewUserController(pool)
	var healthController = controller.NewHealthController()

	// --- main router handlers, no authentication required ---
	muxMain.HandleFunc("GET /health", healthController.Health)

	// --- auth handlers, no authentication required ---
	// login user with username and pw, save jwt as cookie
	muxAuth.HandleFunc("POST /login", authController.Login)
	// remove cookie
	muxAuth.HandleFunc("POST /logout", authController.Logout)

	// --- user handlers ---
	// register a new user
	muxUser.HandleFunc("POST /register", usersController.Register)

	// get all users
	muxUser.Handle(
		"GET /",
		middleware.Chain(
			http.HandlerFunc(usersController.GetAll),
			middleware.Authenticate,
			middleware.Authorize("admin"),
		),
	)
	// get one user
	muxUser.Handle(
		"GET /{username}",
		middleware.Chain(
			http.HandlerFunc(usersController.GetOneByUsername),
			middleware.Authenticate,
			middleware.Authorize("admin"),
		),
	)

	// deactivate a user
	muxUser.Handle(
		"PUT /deactivate/{uid}",
		middleware.Chain(
			http.HandlerFunc(usersController.Deactivate),
			middleware.Authenticate,
			middleware.Authorize("admin", "member"),
		),
	) // TODO: create ChainHandler & ChainFunc

	// add subrouters to main router
	muxMain.Handle("/auth/", http.StripPrefix("/auth", muxAuth))
	muxMain.Handle("/users/", http.StripPrefix("/users", muxUser))

	// chain main router
	muxMainChained := middleware.Chain(muxMain, middleware.Recover, middleware.Log)

	// start server
	var addr string

	if ENV == "DEV" {
		addr = fmt.Sprintf("localhost:%s", PORT)
	} else {
		addr = fmt.Sprintf(":%s", PORT)
	}
	fmt.Printf("Server listening on\u001B[1;32m http://%s \u001B[0m\n", addr)

	log.Fatal(http.ListenAndServe(addr, muxMainChained))
}
