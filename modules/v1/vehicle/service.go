package vehicle

import (
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
)

type service struct {
	repo interfaces.VehicleRepoIF
}

func NewService(repo interfaces.VehicleRepoIF) *service {
	return &service{repo}
}

func (c *service) FindAll() *helpers.Response {
	response, err := c.repo.FindAll()
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) FindByID(id string) *helpers.Response {
	response, err := c.repo.FindByID(id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) Popular() *helpers.Response {
	response, err := c.repo.Popular()
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) Search(name string) *helpers.Response {
	response, err := c.repo.Search(name)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) FindByType(name string) *helpers.Response {
	response, err := c.repo.FindByType(name)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) Sort(name string) *helpers.Response {
	response, err := c.repo.Sort(name)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response)
}

func (c *service) Create(data *model.Vehicle) *helpers.Response {
	data, err := c.repo.Create(data)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}
	return helpers.New(http.StatusOK, "Vehicle successfully created!", data)
}

func (c *service) Update(data *model.Vehicle, id string) *helpers.Response {
	response, err := c.repo.Update(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Vehicle successfully updated!", response)
}

func (c *service) Delete(data *model.Vehicle,id string) *helpers.Response {
	response, err := c.repo.Delete(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Vehicle successfully deleted!", response.VehicleID)
}