package reservation

import (
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
)

type service struct {
	repo interfaces.ReservationRepoIF
}

func NewService(repo interfaces.ReservationRepoIF) *service {
	return &service{repo}
}

func (r *service) GetAll() *helpers.Response {
	data, err := r.repo.FindAll()
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) GetByID(id string) *helpers.Response {
	data, err := r.repo.FindByID(id)
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
	data, err := r.repo.FindByID(id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	_, err = r.repo.Pay(data, vehicle, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}
	return helpers.New(http.StatusOK, "Payment success", data)
}

func (r *service) Delete(data *model.Reservation, id string) *helpers.Response {
	_, err := r.repo.Delete(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Payment success")
}
