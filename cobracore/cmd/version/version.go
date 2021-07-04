package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var command = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Verbose mode:%v\n", Verbose)
		fmt.Printf("many :%v:%v, \n", len(many), many)
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
var Verbose bool
var many []string

func Init(rootCmd *cobra.Command) {
	command.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	command.Flags().StringArrayVarP(&many, "env", "e", []string{}, "sa is many")
	rootCmd.AddCommand(command)
}
