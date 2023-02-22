package migrate

import (
	"fmt"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/dbmigrate/cli/migrate/down"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/dbmigrate/cli/migrate/force"
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/dbmigrate/cli/migrate/up"

	// justified
	_ "github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/snowflake"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/golang-migrate/migrate/v4/source/google_cloud_storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validArgs = []string{
	"up",
	"down",
}

// ExactArgsAndOnlyValidArgs returns an error if there is not at least N args.
func ExactArgsAndOnlyValidArgs(n int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return cobra.OnlyValidArgs(cmd, args)
	}
}

var command = &cobra.Command{
	Use:       "migrate",
	Short:     "migrates db",
	ValidArgs: validArgs,
	Args:      ExactArgsAndOnlyValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("in migrate, and we should NEVER see this")
		return nil
	},
}
var verbose bool
var connectionString string
var databaseName string
var source string
var step uint
var sleep int

// Init ...
func Init(rootCmd *cobra.Command) {
	// verbose
	command.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose logging")
	viper.BindPFlag("verbose", command.PersistentFlags().Lookup("verbose"))

	// database
	command.PersistentFlags().StringVarP(&connectionString, "database", "d", "", "Run migrations against this database (driver://url)")
	viper.BindPFlag("database", command.PersistentFlags().Lookup("database"))
	viper.BindEnv("database", "DATABASE")

	// database-name
	command.PersistentFlags().StringVarP(&databaseName, "database-name", "b", "", "Replace database-name in the --database connection string")
	viper.BindPFlag("database-name", command.PersistentFlags().Lookup("database-name"))
	viper.BindEnv("database-name", "DATABASE_NAME")

	// step
	command.PersistentFlags().UintVarP(&step, "step", "n", 0, "[optional] number or migrate steps.  If not present, then migrate is full up or down ")
	viper.BindPFlag("step", command.PersistentFlags().Lookup("step"))

	// source, any source
	command.PersistentFlags().StringVarP(&source, "source", "s", "", "github source, i.e. github://mattes:personal-access-token@mattes/migrate_test")
	viper.BindPFlag("source", command.PersistentFlags().Lookup("source"))
	viper.BindEnv("source", "SOURCE")

	// sleep time after successful migration
	command.PersistentFlags().IntVar(&sleep, "sleep", 0, "[optional] number of minutes to sleep (negative number == infinite).  ")
	viper.BindPFlag("sleep", command.PersistentFlags().Lookup("sleep"))

	rootCmd.AddCommand(command)
	up.Init(command)
	down.Init(command)
	force.Init(command)
}
