package orm

// import (
// 	"github.com/spf13/cobra"
// )

// var SeedCmd = &cobra.Command{
// 	Use:   "seed",
// 	Short: "cmd for database seed",
// 	RunE:  dbSeed,
// }

// var seedUp bool
// var seedDown bool

// func init() {
// 	SeedCmd.Flags().BoolVarP(&seedUp, "up", "u", true, "for running auto seed")
// 	SeedCmd.Flags().BoolVarP(&seedDown, "down", "d", false, "for running auto reset seed")
// }

// func dbSeed(cmd *cobra.Command, args []string) error {
// 	db, err := ConnectDB()
// 	if err != nil {
// 		return err
// 	}

// 	if seedDown {
// 		err := seedDown(db)
// 		return err
// 	}

// 	if seedUp {
// 		err := seedDown(db)
// 		return err
// 	}

// 	return err
// }
