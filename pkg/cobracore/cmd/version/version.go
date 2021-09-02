package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

func SetVersion(version string) {
	Version = version
}

var command = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
var Verbose bool
var many []string

func Init(rootCmd *cobra.Command) {
	command.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	command.Flags().StringArrayVarP(&many, "env", "e", []string{}, "sa is many")
	rootCmd.AddCommand(command)
}
