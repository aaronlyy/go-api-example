package dto

// from request body
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// model to response body
type UserCreated struct {
	UUID string `json:"uuid"`
	Username string `json:"username"`
	Email string `json:"email"`
}
