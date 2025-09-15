package mapper

import (
	"github.com/aaronlyy/go-api-example/internal/models"
	"github.com/aaronlyy/go-api-example/internal/dto"
)

func UserToDTO(userModel models.User) dto.UserCreated {
	return dto.UserCreated{
		UUID: userModel.UUID,
		Username: userModel.Username,
		Email: userModel.Email,
	}
}