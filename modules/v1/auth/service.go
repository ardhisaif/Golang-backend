package auth

import (
	"fmt"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
)

type service struct {
	repo interfaces.UserRepoIF
}

type tokenData struct {
	Token  string `json:"token"`
	UserID string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func NewService(repo interfaces.UserRepoIF) *service {
	return &service{repo}
}

func (s *service) Login(input *model.User) *helpers.Response {
	user, err := s.repo.FindByUsername(input.Name)
	if err != nil {
		return helpers.New(http.StatusUnauthorized, "Incorrect username or password")
	}

	if helpers.CheckPass(input.Password, user.Password) {
		return helpers.New(http.StatusUnauthorized, "Incorrect username or password")
	}

	jwt := helpers.NewToken(user.UserID, user.Role)
	token, err := jwt.Create()
	if err != nil {
		fmt.Println(err)
		return helpers.New(http.StatusInternalServerError, err)
	}

	return helpers.New(http.StatusOK, "Login success", tokenData{token, user.UserID, user.Name, user.Email, user.Role})
}
