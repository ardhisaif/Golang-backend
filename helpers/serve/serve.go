package serve

import (
	"log"
	"net/http"

	"github.com/ardhisaif/golang_backend/router"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Short: "for running api server",
	RunE: serve,
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

	log.Println("app run on port 3001")
	return http.ListenAndServe(":3001", r)
}