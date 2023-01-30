package vehicle

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	router := rt.PathPrefix("/vehicle").Subrouter()
	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	router.HandleFunc("/", ctrl.GetAll).Methods("GET")
	router.HandleFunc("/popular", ctrl.Popular).Methods("GET")
	router.HandleFunc("/search/", ctrl.Search).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/sort/", ctrl.Sort).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/", ctrl.Create).Methods("POST")
	router.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	router.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}
