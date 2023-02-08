package users

import (
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
)

type service struct {
	repo interfaces.UserRepoIF
}

func NewService(repo interfaces.UserRepoIF) *service {
	return &service{repo}
}

func (c *service) FindAll(data *model.User) *helpers.Response {
	response, err := c.repo.FindAll()
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}
	
	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) Register(data *model.User) *helpers.Response {
	hash, err := helpers.HashPass(data.Password)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}
	data.Password = hash

	response, err := c.repo.Register(data)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "User successfully registered!", response)
}

func (c *service) Update(data *model.User, id string) *helpers.Response {
	response, err := c.repo.Update(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "User successfully updated!", response)
}

func (c *service) Delete(data *model.User, id string) *helpers.Response {
	response, err := c.repo.Delete(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "User successfully deleted!", response.UserID)
}