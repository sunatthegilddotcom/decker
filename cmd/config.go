package cmd

import (
	"fmt"

	"github.com/viniciuschiele/decker/core"

	"github.com/spf13/cobra"
)

func init() {
	var configCommand = &cobra.Command{
		Use:   "config",
		Short: "get and set options",
	}

	var setCommand = &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set a global option",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				fmt.Println("config set: expected two arguments")
				return
			}

			config, err := core.GetConfig()

			if err != nil {
				panic(err)
			}

			err = config.Set(args[0], args[1])

			if err != nil {
				fmt.Println("config set: " + err.Error())
				return
			}

			err = core.SaveConfig(config)

			if err != nil {
				panic(err)
			}
		},
	}

	configCommand.AddCommand(setCommand)
	RootCmd.AddCommand(configCommand)
}
