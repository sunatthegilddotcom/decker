package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var packCommand = &cobra.Command{
	Use:   "pack",
	Short: "Pack a package folder into a Tarball file",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := "."

		if len(args) > 0 {
			inputPath = args[0]
		}

		fileName, err := core.PackPackage(inputPath, ".")

		if err == nil {
			fmt.Println("Package has been created: " + fileName)
		} else {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(packCommand)
}
