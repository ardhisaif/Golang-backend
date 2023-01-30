package reservation

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

func (c *controller) History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	response, err := c.repo.FindAll()
	if err != nil {
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
	}

	helpers.New(response, http.StatusOK, false).Send(w)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var data model.Reservation
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := c.repo.Create(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	helpers.New(response, http.StatusOK, false).Send(w)
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.Reservation
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

func (c *controller) Payment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var reservation model.Reservation
	var vehicle model.Vehicle
	// err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	// if err != nil {
	// 	helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
	// 	return
	// }

	_ , err := c.repo.Pay(&reservation, &vehicle, id)
	if err != nil {
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
		return
	}

	helpers.New("payment success", http.StatusOK, false).Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.Reservation
	_, err := c.repo.Delete(&data, id)
	if err != nil {
		helpers.New(err.Error(), http.StatusBadRequest, true).Send(w)
		return
	}

	helpers.New("delete successfully", http.StatusOK, false).Send(w)
}
