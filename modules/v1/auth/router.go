package auth

import (
	"github.com/ardhisaif/golang_backend/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	router := rt.PathPrefix("/auth").Subrouter()
	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	router.HandleFunc("/login", ctrl.Login).Methods("POST")
}
