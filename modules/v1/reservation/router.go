package reservation

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	router := rt.PathPrefix("/reservation").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	router.HandleFunc("/", ctrl.GetAll).Methods("GET")
	router.HandleFunc("/{id}", ctrl.GetByID).Methods("GET")
	router.HandleFunc("/", ctrl.Create).Methods("POST")
	router.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	router.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")

	router.HandleFunc("/payment/{id}", ctrl.Payment).Methods("PUT")
	router.HandleFunc("/history", ctrl.History).Methods("GET")
	router.HandleFunc("/history/search/", ctrl.Search).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/sort/", ctrl.Sort).Queries("name", "{name}").Methods("GET")
	
}
