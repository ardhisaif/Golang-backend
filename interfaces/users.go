package interfaces

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
)

type UserRepoIF interface {
	FindAll() (*model.Users, error)
	FindByUsername(name string) (*model.User, error)
	Register(user *model.User) (*model.User, error)
	Update(user *model.User, id string) (*model.User, error)
	Delete(user *model.User, id string) (*model.User, error)
}

type UserSvcIF interface {
	FindAll(data *model.User) *helpers.Response 
	Register(data *model.User) *helpers.Response 
	Update(data *model.User, id string) *helpers.Response 
	Delete(data *model.User, id string) *helpers.Response
}