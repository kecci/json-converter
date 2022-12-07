package json_converter

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:     "",
	Version: version,
	Short:   "json-converter - a simple CLI to conversion json and excel files.",
	Long:    `json-converter - a simple CLI to conversion json and excel files.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: true,
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
