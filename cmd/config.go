package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/config"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Configuration management",
}

var deleteConfigCommand = &cobra.Command{
	Use:   "delete <key>",
	Short: "Delete a property",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a key")
		}

		config.Delete(args[0])
		return config.Save()
	},
}

var getConfigCommand = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a property",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("You must specify a key")
		}

		value := config.Get(args[0])

		fmt.Println(value)

		return nil
	},
}

var listConfigCommand = &cobra.Command{
	Use:   "list",
	Short: "List all the properties",
	Run: func(cmd *cobra.Command, args []string) {
		properties := config.List()

		for k, v := range properties {
			fmt.Println(k + " = " + v)
		}
	},
}

var setConfigCommand = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a property",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("You must specify a key and value")
		}

		config.Set(args[0], args[1])
		return config.Save()
	},
}

func init() {
	configCommand.AddCommand(deleteConfigCommand)
	configCommand.AddCommand(getConfigCommand)
	configCommand.AddCommand(listConfigCommand)
	configCommand.AddCommand(setConfigCommand)
	rootCmd.AddCommand(configCommand)
}
