package json_converter

import (
	"github.com/spf13/cobra"
)

var excelCmd = &cobra.Command{
	Use:     "excel",
	Short:   "Convert json to excel file",
	Example: "json-converter excel /Users/kecci/Downloads/message_transaction_2022_nov.json /Users/kecci/Downloads/message_transaction_2022_nov.xlsx",
	Aliases: []string{"xlsx"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		JsonToExcel(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(excelCmd)
}
