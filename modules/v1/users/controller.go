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
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
	}

	helpers.New(response, http.StatusOK, false).Send(w)
}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := c.repo.Register(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	helpers.New(response, http.StatusOK, false).Send(w)
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
		return
	}

	response, err := c.repo.Update(&data, id)
	if err != nil {
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
		return
	}

	helpers.New(response, http.StatusOK, false).Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.User
	_, err := c.repo.Delete(&data, id)
	if err != nil {
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
		return
	}

	helpers.New("delete successfully", http.StatusOK, false).Send(w)
}