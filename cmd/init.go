package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

func init() {
	var initCreate = &cobra.Command{
		Use:   "init",
		Short: "Generates code scaffolding for a package",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("You must specify a path to initialize a new package")
				return
			}

			packageManager := core.PackageManager{}
			err := packageManager.Init(args[0])

			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	RootCmd.AddCommand(initCreate)
}
