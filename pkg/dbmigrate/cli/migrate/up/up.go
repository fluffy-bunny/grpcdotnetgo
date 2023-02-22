package up

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/dbmigrate/cli/migrate/utils"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:               "up",
	Short:             "migrates the db up",
	PersistentPreRunE: utils.UpPersistentPreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.Migrate(utils.MigrateUp)
	},
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(command)
}
