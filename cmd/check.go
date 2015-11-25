package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "Check whether a package folder contains the required files",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := "."

		if len(args) > 0 {
			inputPath = args[0]
		}

		return core.CheckPackage(inputPath)
	},
}

func init() {
	rootCmd.AddCommand(checkCommand)
}
