package reservation

import (
	"github.com/ardhisaif/golang_backend/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	router := rt.PathPrefix("/reservation").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	router.HandleFunc("/", middleware.Handle(ctrl.GetAll, middleware.AuthMiddle("user"))).Methods("GET")
	// router.HandleFunc("/{id}", ctrl.GetByID).Methods("GET")
	router.HandleFunc("/{id}", middleware.Handle(ctrl.GetByID, middleware.AuthMiddle("user"))).Methods("GET")
	// router.HandleFunc("/", ctrl.Create).Methods("POST")
	router.HandleFunc("/", middleware.Handle(ctrl.Create, middleware.AuthMiddle("user"))).Methods("POST")
	// router.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	router.HandleFunc("/{id}", middleware.Handle(ctrl.Update, middleware.AuthMiddle("user"))).Methods("PUT")
	// router.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
	router.HandleFunc("/{id}", middleware.Handle(ctrl.Delete, middleware.AuthMiddle("user"))).Methods("DELETE")

	// router.HandleFunc("/payment/{id}", ctrl.Payment).Methods("PUT")
	router.HandleFunc("/payment/{id}", middleware.Handle(ctrl.Payment, middleware.AuthMiddle("user"))).Methods("PUT")
	// router.HandleFunc("/history", ctrl.History).Methods("GET")
	router.HandleFunc("/history", middleware.Handle(ctrl.History, middleware.AuthMiddle("user"))).Methods("GET")
	// router.HandleFunc("/history/search/", ctrl.Search).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/history/search/", middleware.Handle(ctrl.Search, middleware.AuthMiddle("user"))).Queries("name", "{name}").Methods("GET")
	// router.HandleFunc("/sort/", ctrl.Sort).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/sort/", middleware.Handle(ctrl.Sort, middleware.AuthMiddle("user"))).Queries("name", "{name}").Methods("GET")
	
}
