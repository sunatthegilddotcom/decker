package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
)

var rootCmd = &cobra.Command{
	Use:          "decker",
	Short:        "Decker script manager powered by https://godecker.io",
	SilenceUsage: true,
}

//Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	if err := core.InitConfig(); err != nil {
		panic(err)
	}
}
