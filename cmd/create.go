package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/pkg"
)

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Creates an empty package",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a path to create a new package")
		}

		return pkg.Create(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCommand)
}
