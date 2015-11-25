package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var packCommand = &cobra.Command{
	Use:   "pack",
	Short: "Pack a package folder into a Tarball file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a path for a package")
		}

		fileName, err := core.PackPackage(args[0], ".")

		if err != nil {
			return err
		}

		fmt.Println("Package has been created: " + fileName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(packCommand)
}
