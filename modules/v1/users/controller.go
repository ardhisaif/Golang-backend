package users

import (
	"encoding/json"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
)

type controller struct {
	service interfaces.UserSvcIF
}

func NewCtrl(service interfaces.UserSvcIF) *controller {
	return &controller{service}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json
	
	var data model.User
	c.service.FindAll(&data).Send(w)
}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&data)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	c.service.Register(&data).Send(w)
	
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(&data)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	c.service.Update(&data, id).Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.User

	c.service.Update(&data, id).Send(w)
}