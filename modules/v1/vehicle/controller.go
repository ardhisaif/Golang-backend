package vehicle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

func UploadImage(w http.ResponseWriter, r *http.Request) string {
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		helpers.New(http.StatusInternalServerError, err.Error()).Send(w)
		return ""
	}
	fmt.Println(file)
	err = r.ParseForm()
	if err != nil {
		helpers.New(http.StatusInternalServerError, err.Error()).Send(w)
		return ""
	}

	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	err = os.MkdirAll("./public", os.ModePerm)
	if err != nil {
		helpers.New(http.StatusInternalServerError, err.Error()).Send(w)
		return ""
	}

	imageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename))

	pathFile := filepath.Join("./public/" + imageName)

	dst, err := os.Create(pathFile)
	if err != nil {
		helpers.New(http.StatusInternalServerError, err.Error()).Send(w)
		return ""
	}
	defer dst.Close()

	return imageName
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") // set header to json

	var decoder = schema.NewDecoder()
	var data model.Vehicle
	// imageName:= UploadImage(w, r)
	err := r.ParseForm()
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	filename := handler.Filename
	fileLocation := filepath.Join("./public/", filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = decoder.Decode(&data, r.PostForm)
	if err != nil {
		helpers.New(http.StatusBadRequest, err.Error()).Send(w)
		return
	}
	fmt.Println(fileLocation)
	
	data.Image = fileLocation
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
