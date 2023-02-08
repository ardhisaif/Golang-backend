package interfaces

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
)

type Auth_ServiceIF interface {
	Login(input *model.User) *helpers.Response
}