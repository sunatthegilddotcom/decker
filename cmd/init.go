package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Creates an empty package",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a path to initialize a new package")
		}

		return core.InitPackage(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCommand)
}
