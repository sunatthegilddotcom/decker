package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var checkCommand = &cobra.Command{
	Use:   "check",
	Short: "Check whether a package folder contains the required files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a path for a package")
		}

		return core.CheckPackage(args[0])
	},
}

func init() {
	rootCmd.AddCommand(checkCommand)
}
