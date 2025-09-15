package controller

import (
	"fmt"
	"net/http"

	"github.com/aaronlyy/go-api-example/internal/dto"
	"github.com/aaronlyy/go-api-example/internal/mapper"
	"github.com/aaronlyy/go-api-example/internal/repository"
	"github.com/aaronlyy/go-api-example/internal/response"
	"github.com/aaronlyy/go-api-example/internal/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

// create a new controller struct
type UserController struct {
	DB *pgxpool.Pool
}

func NewUserController(db *pgxpool.Pool) UserController {
	return UserController{DB: db}
}

// attach method to controller, needs ResponseWriter and Request
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {

	// create new struct for parsing request body
	var rb dto.CreateUserRequest

	// parse request body to struct
	if err := util.ParseBody(r.Body, &rb); err != nil {
		response.NewResponse(500, "error parsing body or missing data", nil).Send(w)
		return
	}

	fmt.Printf("New user created: %s\n", rb.Username)

	response.NewResponse(201, "New user created", nil).Send(w)
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	userRepo := repository.NewUsersRepository(c.DB)
	users, err := userRepo.ListAll(r.Context())

	if err != nil {
		response.NewResponse(400, "error getting users", nil).Send(w)
		return
	}

	fmt.Printf("users: %v", users)

	response.NewResponse(200, "got all users", nil).Send(w)

}

func (c *UserController) GetOneByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	userRepo := repository.NewUsersRepository(c.DB)
	user, err := userRepo.GetOneByUsername(r.Context(), username)

	if err != nil {
		response.NewResponse(400, "could not get user", nil).Send(w)
		return
	}

	// model to dto
	userDTO := mapper.UserToDTO(user)

	response.NewResponse(200, "got user", userDTO).Send(w)
}

// attach method to controller, needs ResponseWriter and Request
func (c *UserController) Deactivate(w http.ResponseWriter, r *http.Request) {
	// prepare response object
	response.NewResponse(200, "User deactivated", nil).Send(w)
}
