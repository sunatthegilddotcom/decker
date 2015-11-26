package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Get and Set global options",
}

var getServerCommand = &cobra.Command{
	Use:   "get-server",
	Short: "get default server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(core.ResolveServer(""))
	},
}

var setServerCommand = &cobra.Command{
	Use:   "set-server <server>",
	Short: "Set default server",
	RunE: func(cmd *cobra.Command, args []string) error {
		core.Config.DefaultServer = ""

		if len(args) > 0 {
			core.Config.DefaultServer = core.CompactServer(args[0])
		}

		return core.Config.Save()
	},
}

func init() {
	configCommand.AddCommand(getServerCommand)
	configCommand.AddCommand(setServerCommand)
	rootCmd.AddCommand(configCommand)
}
