package interfaces

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
)

type VehicleRepoIF interface {
	FindAll() (*model.Vehicles, error)
	FindByID(id string) (*model.Vehicle, error)
	Popular() (*model.Vehicles, error)
	Search(name string) (*model.Vehicles, error)
	FindByType(name string) (*model.Vehicles, error)
	Sort(name string) (*model.Vehicles, error)
	Create(user *model.Vehicle) (*model.Vehicle, error)
	Update(user *model.Vehicle, id string) (*model.Vehicle, error)
	Delete(user *model.Vehicle, id string) (*model.Vehicle, error)
}

type VehicleSvcIF interface {
	FindAll() *helpers.Response
	FindByID(id string) *helpers.Response 
	Popular() *helpers.Response
	Search(name string) *helpers.Response
	FindByType(name string) *helpers.Response
	Sort(name string) *helpers.Response
	Create(data *model.Vehicle) *helpers.Response
	Update(data *model.Vehicle, id string) *helpers.Response
	Delete(data *model.Vehicle, id string) *helpers.Response
}