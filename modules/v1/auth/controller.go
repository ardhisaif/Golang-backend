package auth

import (
	"encoding/json"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
)

type controller struct {
	service interfaces.Auth_ServiceIF
}

func NewCtrl(service interfaces.Auth_ServiceIF) *controller {
	return &controller{service}
}

func (c *controller) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json
	
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.New(http.StatusInternalServerError, err.Error())
		return
	}

	c.service.Login(&user).Send(w)
}
