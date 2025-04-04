package dto

import (
	"github.com/MicroMolekula/auth-service/internal/models"
)

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func UserModelToDto(user *models.User) *User {
	return &User{
		int(user.ID),
		user.Email,
	}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserRegisterToModel(user *UserRegister) *models.User {
	return &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
