package orm

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "cmd for database migration",
	RunE: dbMigrate,
}

var migUp bool
var migDown bool

func init(){
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", true, "for running auto migration")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "for running auto reset migration")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	if migDown {
		return db.Migrator().DropTable(&model.Reservation{}, &model.Vehicle{}, &model.User{})
	}

	if migUp {
		return db.AutoMigrate(&model.Reservation{}, &model.Vehicle{}, &model.User{})
	}

	return nil
}