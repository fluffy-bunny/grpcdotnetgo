package force

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/dbmigrate/cli/migrate/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var command = &cobra.Command{
	Use:               "force",
	Short:             "Force sets a migration version. It does not check any currently active version in database. It resets the dirty state to false.",
	PersistentPreRunE: utils.UpPersistentPreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.Migrate(utils.MigrateForce)
	},
}
var version int

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(command)
	// a migration version
	command.Flags().IntVar(&version, "version", -1, "[required] a migration version")
	command.MarkFlagRequired("version")
	viper.BindPFlag("version", command.PersistentFlags().Lookup("version"))
	viper.BindEnv("version", "FORCE_VERSION")
}
