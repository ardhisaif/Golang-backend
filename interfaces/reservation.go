package interfaces

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
)

type ReservationRepoIF interface {
	FindAll() (*model.Reservations, error)
	FindByUserID(id string) (*model.Reservations, error)
	FindByID(id string) (*model.Reservation, error)
	History() (*model.Reservations, error)
	Search(name string) (*model.Reservations, error)
	Sort(name string) (*model.Reservations, error)
	Create(user *model.Reservation) (*model.Reservation, error)
	Update(user *model.Reservation, id string) (*model.Reservation, error)
	Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string) (*model.Reservation, error)
	Delete(user *model.Reservation, id string) (*model.Reservation, error)
}

type ReservationSvcIF interface {
	GetAll() *helpers.Response
	GetByID(id string) *helpers.Response 
	History() *helpers.Response
	Search(name string) *helpers.Response
	Sort(name string) *helpers.Response
	Create(data *model.Reservation) *helpers.Response
	Update(data *model.Reservation, id string) *helpers.Response
	Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string) *helpers.Response
	Delete(data *model.Reservation, id string) *helpers.Response
}
