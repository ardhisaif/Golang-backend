package command

import (
	"github.com/ardhisaif/golang_backend/database/orm"
	"github.com/ardhisaif/golang_backend/helpers/serve"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short:"Simple backend golang",
}

func init(){
	initCommand.AddCommand(serve.ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}