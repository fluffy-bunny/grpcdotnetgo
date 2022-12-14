package version

import (
	"fmt"

	main_version "github.com/fluffy-bunny/grpcdotnetgo/pkg/cobracore/cmd/version"
	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(main_version.Version)
	},
}
var Verbose bool
var many []string

func Init(rootCmd *cobra.Command) {
	command.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	command.Flags().StringArrayVarP(&many, "env", "e", []string{}, "sa is many")
	rootCmd.AddCommand(command)
}
