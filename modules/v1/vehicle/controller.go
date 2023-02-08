package vehicle

import (
	"encoding/json"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/ardhisaif/golang_backend/helpers"
	"github.com/ardhisaif/golang_backend/interfaces"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type controller struct {
	service interfaces.VehicleSvcIF
}

func NewCtrl(service interfaces.VehicleSvcIF) *controller {
	return &controller{service}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	c.service.FindAll().Send(w)
}

func (c *controller) Popular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	c.service.Popular().Send(w)
}

func (c *controller) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	name := r.FormValue("name")
	c.service.Search(name).Send(w)
}

func (c *controller) Sort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	name := r.FormValue("name")
	c.service.Sort(name).Send(w)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var decoder = schema.NewDecoder()
	var data model.Vehicle
	err := r.ParseForm()
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	fileUrl, err := helpers.CloudUpload(file)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}
	defer file.Close()

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}
	data.Image = fileUrl
	c.service.Create(&data).Send(w)
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&data) // implement req json
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	c.service.Update(&data, id).Send(w)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	id := mux.Vars(r)["id"]
	var data *model.Vehicle

	c.service.Delete(data, id).Send(w)
}
