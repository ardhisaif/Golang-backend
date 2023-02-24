package reservation

import (
	"encoding/json"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
	"github.com/gorilla/mux"
)

type controller struct {
	service interfaces.ReservationSvcIF
}

func NewCtrl(service interfaces.ReservationSvcIF) *controller {
	return &controller{service}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	response := c.service.GetAll()
	response.Send(w)
}

func (c *controller) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	response := c.service.GetByID(id)
	response.Send(w)
}

func (c *controller) History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	response := c.service.History()
	response.Send(w)
}

func (c *controller) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	name := r.FormValue("name")
	response := c.service.Search(name)
	response.Send(w)
}

func (c *controller) Sort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	name := r.FormValue("name")
	c.service.Sort(name).Send(w)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json
	var data model.Reservation
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	c.service.Create(&data).Send(w)
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

	response := c.service.Update(&data, id)
	response.Send(w)
}

func (c *controller) Payment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var reservation model.Reservation
	var vehicle model.Vehicle

	response := c.service.Pay(&reservation, &vehicle, id)
	response.Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.Reservation
	response := c.service.Delete(&data, id)
	response.Send(w)
}
