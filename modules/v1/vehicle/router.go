package vehicle

import (
	"github.com/ardhisaif/golang_backend/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	router := rt.PathPrefix("/vehicle").Subrouter()
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	router.HandleFunc("/", middleware.Handle(ctrl.GetAll, middleware.AuthMiddle("user"))).Methods("GET")
	// router.HandleFunc("/popular", middleware.Handle(ctrl.Popular, middleware.AuthMiddle("user"))).Methods("GET")
	router.HandleFunc("/search", middleware.Handle(ctrl.Search, middleware.AuthMiddle("user"))).Queries("name", "{name}").Methods("GET")
	// router.HandleFunc("/type", middleware.Handle(ctrl.FindByType, middleware.AuthMiddle("user"))).Queries("name", "{name}").Methods("GET")
	// router.HandleFunc("/sort", middleware.Handle(ctrl.Sort, middleware.AuthMiddle("user"))).Queries("name", "{name}").Methods("GET")
	// router.HandleFunc("/", middleware.Handle(ctrl.Create, middleware.AuthMiddle("admin"))).Methods("POST")
	// router.HandleFunc("/{id}", middleware.Handle(ctrl.Update, middleware.AuthMiddle("admin"))).Methods("PUT")
	router.HandleFunc("/{id}", middleware.Handle(ctrl.Delete, middleware.AuthMiddle("admin"))).Methods("DELETE")

	// router.HandleFunc("/", ctrl.GetAll).Methods("GET")
	router.HandleFunc("/popular", ctrl.Popular).Methods("GET")
	// router.HandleFunc("/search/", ctrl.Search).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/type", ctrl.FindByType).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/sort", ctrl.Sort).Queries("name", "{name}").Methods("GET")
	router.HandleFunc("/", ctrl.Create).Methods("POST")
	router.HandleFunc("/{id}", ctrl.Update).Methods("PUT")
	router.HandleFunc("/{id}", ctrl.GetByID).Methods("GET")
	// router.HandleFunc("/{id}", ctrl.Delete).Methods("DELETE")
}
