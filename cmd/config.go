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

		config, err := core.GetConfig()

		if err != nil {
			return err
		}

		err = config.Set(args[0], "")

		if err != nil {
			return err
		}

		return core.SaveConfig(config)
	},
}

var getConfigCommand = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a global option",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a key")
		}

		config, err := core.GetConfig()

		if err != nil {
			return err
		}

		value, err := config.Get(args[0])

		if err != nil {
			return err
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

		config, err := core.GetConfig()

		if err != nil {
			return err
		}

		err = config.Set(args[0], args[1])

		if err != nil {
			return err
		}

		return core.SaveConfig(config)
	},
}

func init() {
	configCommand.AddCommand(delConfigCommand)
	configCommand.AddCommand(getConfigCommand)
	configCommand.AddCommand(setConfigCommand)
	rootCmd.AddCommand(configCommand)
}
