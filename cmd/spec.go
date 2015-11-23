package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

var cmdSpec = &cobra.Command{
	Use:   "spec",
	Short: "Generates a package.json file for a new package",
	Run: func(cmd *cobra.Command, args []string) {
		generateSpec()
	},
}

func init() {
	RootCmd.AddCommand(cmdSpec)
}

func generateSpec() {
	json.Marshal(nil)
}
