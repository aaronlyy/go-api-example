package controller

import (
	"net/http"
	"time"

	"github.com/aaronlyy/go-api-example/internal/auth"
	"github.com/aaronlyy/go-api-example/internal/response"
	"github.com/aaronlyy/go-api-example/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthController struct {
	DB *pgxpool.Pool
}

func NewAuthController(db *pgxpool.Pool) AuthController {
	return AuthController{DB: db}
}

// first auth with username and password
type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// attach method to controller, needs ResponseWriter and Request
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// get username and password from request
	var rb LoginRequestBody

	if err := util.ParseBody(r.Body, &rb); err != nil {
		response.NewResponse(500, "error parsing body or missing data", nil).Send(w)
		return
	}

	// load hash from db
	hash, err := auth.HashPassword("a1sdf234", 10)

	if err != nil {
		response.NewResponse(500, "error loading pw hash", nil)
		return
	}

	// verify password and username
	if rb.Username != "aaron" || auth.VerifyPassword(rb.Password, hash) != nil {
		response.NewResponse(500, "wrong username or password", nil).Send(w)
		return
	}

	// create new jwt with userid and role
	token, exp, err := auth.SignAccessToken("1", []string{"admin", "member", "guest"}, 15*time.Minute)

	if err != nil {
		response.NewResponse(500, "error signing token", nil).Send(w)
		return
	}

	// set cookie with jwt
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		Expires:  exp,
		MaxAge:   int(time.Until(exp).Seconds()),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	// prepare response object
	var res = response.NewResponse(200, "User was authenticated", nil)
	res.Send(w)
}

// delete all cookies and delete refresh token from db
func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {

}
