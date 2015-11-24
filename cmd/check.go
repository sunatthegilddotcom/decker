package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker-cli/core"
)

func init() {
	var checkCommand = &cobra.Command{
		Use:   "check",
		Short: "Check whether a package contains the required files",
		Run: func(cmd *cobra.Command, args []string) {
			inputPath := "."

			if len(args) > 0 {
				inputPath = args[0]
			}

			packageManager := core.PackageManager{}
			err := packageManager.Check(inputPath)

			if err != nil {
				fmt.Println(err)
			}
		},
	}

	RootCmd.AddCommand(checkCommand)
}
