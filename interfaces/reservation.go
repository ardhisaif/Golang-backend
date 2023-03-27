package interfaces

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
)

type ReservationRepoIF interface {
	FindAll(userID string) (*model.Reservations, error)
	FindByUserID(id string, userID string) (*model.Reservations, error)
	FindByID(id string, userID string) (*model.Reservation, error)
	History(userID string) (*model.Reservations, error)
	Search(name string, userID string) (*model.Reservations, error)
	Sort(name string, userID string) (*model.Reservations, error)
	Create(user *model.Reservation) (*model.Reservation, error)
	Update(user *model.Reservation, id string) (*model.Reservation, error)
	Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string) (*model.Reservation, error)
	Delete(user *model.Reservation, id string) (*model.Reservation, error)
}

type ReservationSvcIF interface {
	GetAll(userID string) *helpers.Response
	GetByID(id string, userID string) *helpers.Response 
	History(userID string) *helpers.Response
	Search(name string, userID string) *helpers.Response
	Sort(name string, userID string) *helpers.Response
	Create(data *model.Reservation, userID string) *helpers.Response
	Update(data *model.Reservation, id string, userID string) *helpers.Response
	Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string, userID string) *helpers.Response
	Delete(data *model.Reservation, id string) *helpers.Response
}
