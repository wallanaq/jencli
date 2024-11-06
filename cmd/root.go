package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/jencli/cmd/auth"
)

var Verbose bool
var Version string = "v0.1.0"

var RootCmd = &cobra.Command{
	Use:     "jencli",
	Short:   "A CLI tool for automating tasks and managing jobs in Jenkins",
	Version: Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetOutput(os.Stdout)
		} else {
			log.SetOutput(io.Discard)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(auth.LoginCmd)
}
