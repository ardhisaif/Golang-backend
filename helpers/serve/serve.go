package serve

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ardhisaif/golang_backend/router"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"github.com/rs/cors"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Short: "for running api server",
	RunE: serve,
}

func corsHandler() *cors.Cors {
	t := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return t
}

func serve(cmd *cobra.Command, args []string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r, err := router.NewApp()
	if err != nil {
		return err
	}
	
	var address string = "0.0.0.0:8080"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "0.0.0.0:" + PORT
	}

	corss := corsHandler()

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
		Handler:      corss.Handler(r),
	}


	log.Println("app run on port", address)
	return srv.ListenAndServe()
}