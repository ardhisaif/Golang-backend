package router

import (
	"fmt"
	"net/http"

	"github.com/ardhisaif/golang_backend/database/orm"
	"github.com/ardhisaif/golang_backend/modules/v1/auth"
	"github.com/ardhisaif/golang_backend/modules/v1/reservation"
	"github.com/ardhisaif/golang_backend/modules/v1/users"
	"github.com/ardhisaif/golang_backend/modules/v1/vehicle"
	"github.com/gorilla/mux"
)

func NewApp() (*mux.Router, error) {
	r := mux.NewRouter()
	db, err := orm.ConnectDB()
	if err != nil {
		return nil, err
	}

	auth.New(r, db)
	users.New(r, db)
	vehicle.New(r, db)
	reservation.New(r, db)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Masuk</h1>")
	})

	return r, nil
}
