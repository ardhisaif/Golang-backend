package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"

)

func New(rt *mux.Router, db *gorm.DB) {
	router := rt.PathPrefix("/user").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)


	router.HandleFunc("/", ctrl.GetAll).Methods("GET")
	router.HandleFunc("/", ctrl.Register).Methods("POST")
	router.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	router.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}
