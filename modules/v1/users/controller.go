package users

import (
	"encoding/json"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/gorilla/mux"
)

type controller struct {
	repo *repository
}

func NewCtrl(repo *repository) *controller {
	return &controller{repo}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	response, err := c.repo.FindAll()
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response).Send(w)
}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	hash, err := helpers.HashPass(data.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	data.Password = hash

	response, err := c.repo.Register(&data)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "User successfully registered!", response).Send(w)
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

	response, err := c.repo.Update(&data, id)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data successfully updated!", response).Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.User
	_, err := c.repo.Delete(&data, id)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data successfully deleted!", id).Send(w)
}