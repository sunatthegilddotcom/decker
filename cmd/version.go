package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Show the Decker verion information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client Version: 0.1.0")
	},
}

func init() {
	RootCmd.AddCommand(cmdVersion)
}
