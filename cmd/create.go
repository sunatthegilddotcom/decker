package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker-cli/core"
)

func init() {
	var cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Generates code scaffolding for a package",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("You must specify a path")
				return
			}

			packageManager := core.PackageManager{}
			err := packageManager.Create(args[0])

			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	RootCmd.AddCommand(cmdCreate)
}
