package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "Check whether a package folder contains the required files",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := "."

		if len(args) > 0 {
			inputPath = args[0]
		}

		err := core.CheckPackage(inputPath)

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCommand)
}
