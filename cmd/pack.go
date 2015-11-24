package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

func init() {
	var packCommand = &cobra.Command{
		Use:   "pack",
		Short: "Pack a package folder into a Tarball file",
		Run: func(cmd *cobra.Command, args []string) {
			inputPath := "."

			if len(args) > 0 {
				inputPath = args[0]
			}

			packageManager := core.PackageManager{}
			fileName, err := packageManager.Pack(inputPath, ".")

			if err == nil {
				fmt.Println("Package has been created: " + fileName)
			} else {
				fmt.Println(err.Error())
			}
		},
	}

	RootCmd.AddCommand(packCommand)
}
