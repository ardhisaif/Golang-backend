package reservation

import (
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
)

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo}
}

func (r *service) GetAll() *helpers.Response {
	data, err := r.repo.FindAll()
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) History() *helpers.Response {
	data, err := r.repo.History()
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Search(name string) *helpers.Response {
	data, err := r.repo.Search(name)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Sort(name string) *helpers.Response {
	data, err := r.repo.Sort(name)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Create(data *model.Reservation) *helpers.Response {
	data, err := r.repo.Create(data)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Update(data *model.Reservation, id string) *helpers.Response {
	data, err := r.repo.Update(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string) *helpers.Response {
	_, err := r.repo.Pay(reservation, vehicle, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Payment success")
}

func (r *service) Delete(data *model.Reservation, id string) *helpers.Response {
	_, err := r.repo.Delete(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Payment success")
}
