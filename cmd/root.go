package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/config"
)

var rootCmd = &cobra.Command{
	Use:          "decker",
	Short:        "Decker script manager powered by https://godecker.io",
	SilenceUsage: true,
}

//Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	baseDir := path.Join(os.Getenv("HOME"), ".decker")

	if err := config.Init(baseDir); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
