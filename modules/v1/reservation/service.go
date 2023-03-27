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

func (r *service) GetAll(userID string) *helpers.Response {
	data, err := r.repo.FindAll(userID)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) GetByID(id string, userID string) *helpers.Response {
	data, err := r.repo.FindByID(id, userID)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) History(userID string) *helpers.Response {
	data, err := r.repo.History(userID )
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Search(name string, userID string) *helpers.Response {
	data, err := r.repo.Search(name, userID )
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Sort(name string, userID string) *helpers.Response {
	data, err := r.repo.Sort(name, userID)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Create(data *model.Reservation, userID string ) *helpers.Response {
	data.UserID = userID
	data, err := r.repo.Create(data)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Update(data *model.Reservation, id string, userID string) *helpers.Response {
	data.UserID = userID
	data, err := r.repo.Update(data, id)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", data)
}

func (r *service) Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string, userID string) *helpers.Response {
	data, err := r.repo.FindByID(id, userID)
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())
	}

	_, err = r.repo.Pay(data, vehicle, id )
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}
	return helpers.New(http.StatusOK, "Payment success", data)
}

func (r *service) Delete(data *model.Reservation, id string) *helpers.Response {
	_, err := r.repo.Delete(data, id )
	if err != nil {
		return helpers.New(http.StatusBadRequest, err.Error())

	}

	return helpers.New(http.StatusOK, "Delete success")
}
