package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var packCommand = &cobra.Command{
	Use:   "pack",
	Short: "Pack a package folder into a Tarball file",
	RunE: func(cmd *cobra.Command, args []string) error {
		inputPath := "."

		if len(args) > 0 {
			inputPath = args[0]
		}

		fileName, err := core.PackPackage(inputPath, ".")

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
