package reservation

import (
	"encoding/json"
	"fmt"
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

	result := c.service.
}

func (c *controller) History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	response, err := c.repo.History()
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response).Send(w)
}

func (c *controller) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	name := r.FormValue("name")
	response, err := c.repo.Search(name)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response).Send(w)
}

func (c *controller) Sort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	name := r.FormValue("name")
	fmt.Println(name)
	fmt.Println("ma...........")
	response, err := c.repo.Sort(name)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data successfully retrieved/transmitted!", response).Send(w)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var data model.Reservation
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := c.repo.Create(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.New(http.StatusOK, "Data successfully created!", response).Send(w)
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.Reservation
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

func (c *controller) Payment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var reservation model.Reservation
	var vehicle model.Vehicle

	_ , err := c.repo.Pay(&reservation, &vehicle, id)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Payment success").Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.Reservation
	_, err := c.repo.Delete(&data, id)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	helpers.New(http.StatusOK, "Data succesfulfully deleted!").Send(w)
}
