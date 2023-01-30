package main

import (
	"log"
	"net/http"

	"github.com/ardhisaif/golang_backend/router"
)

func main() {
	r, err := router.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("app run on port 3001")
	err = http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(err)
	}
}
