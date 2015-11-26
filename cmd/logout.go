package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var logoutCommand = &cobra.Command{
	Use:   "logout",
	Short: "Log out from a Decker registry, if no server is specified \"https://registry.godecker.io/\" is the default.",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := ""

		if len(args) > 0 {
			server = args[0]
		}

		return core.Logout(server)
	},
}

func init() {
	rootCmd.AddCommand(logoutCommand)
}
