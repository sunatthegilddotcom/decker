package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

func init() {
	var configCommand = &cobra.Command{
		Use:   "config",
		Short: "Get and Set global options",
	}

	var delCommand = &cobra.Command{
		Use:   "del <key>",
		Short: "Delete a global option",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("config del: expected one argument")
				return
			}

			config, err := core.GetConfig()

			if err != nil {
				panic(err)
			}

			err = config.Set(args[0], "")

			if err != nil {
				fmt.Println("config del: " + err.Error())
				return
			}

			err = core.SaveConfig(config)

			if err != nil {
				panic(err)
			}
		},
	}

	var getCommand = &cobra.Command{
		Use:   "get <key>",
		Short: "Get a global option",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("config get: expected one argument")
				return
			}

			config, err := core.GetConfig()

			if err != nil {
				panic(err)
			}

			value, err := config.Get(args[0])

			if err != nil {
				fmt.Println("config get: " + err.Error())
				return
			}

			fmt.Println(value)
		},
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

	configCommand.AddCommand(delCommand)
	configCommand.AddCommand(getCommand)
	configCommand.AddCommand(setCommand)
	rootCmd.AddCommand(configCommand)
}
