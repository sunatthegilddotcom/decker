package cmd

import (
	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/config"
)

var logoutCommand = &cobra.Command{
	Use:   "logout",
	Short: "Log out from a Decker registry",
	RunE: func(cmd *cobra.Command, args []string) error {
		config.Delete("token")
		return config.Save()
	},
}

func init() {
	rootCmd.AddCommand(logoutCommand)
}
