package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Get and Set global options",
}

var delConfigCommand = &cobra.Command{
	Use:   "del <key>",
	Short: "Delete a global option",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a key")
		}

		if !core.SetConfigValue(args[0], "") {
			return errors.New(args[0] + " is a invalid key")
		}

		return core.SaveConfig()
	},
}

var getConfigCommand = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a global option",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a key")
		}

		value, found := core.GetConfigValue(args[0])

		if !found {
			return errors.New(args[0] + " is a invalid key")
		}

		fmt.Println(value)

		return nil
	},
}

var setConfigCommand = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a global option",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("You must specify a key and value")
		}

		if !core.SetConfigValue(args[0], args[1]) {
			return errors.New(args[0] + " is a invalid key")
		}

		return core.SaveConfig()
	},
}

func init() {
	configCommand.AddCommand(delConfigCommand)
	configCommand.AddCommand(getConfigCommand)
	configCommand.AddCommand(setConfigCommand)
	rootCmd.AddCommand(configCommand)
}
