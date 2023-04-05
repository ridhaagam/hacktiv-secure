package migrate

import (
	"fmt"
	"secure/challenge-2/infrastructure/repository/postgres"

	"github.com/spf13/cobra"
)

var (
	Postgres bool
)

// PostgresCmd represents the postgres command
var PostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Migrate PostgreSQL database",
	Long:  `The postgres command is used to migrate the PostgreSQL database to its latest schema version`,
	Run: func(cmd *cobra.Command, args []string) {
		if Postgres {
			err := postgres.MigratePostgre(PostgresDB)
			if err != nil {
				_ = fmt.Errorf("fatal error in migrating postgres: %s", err)
				panic(err)
			}
			return
		}

		cmd.Help()
	},
}

func init() {
	// migrating flag
	PostgresCmd.PersistentFlags().BoolVarP(&Postgres, "migrate", "m", false, "perform database migration")
}
