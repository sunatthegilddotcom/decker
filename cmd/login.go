package cmd

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/viniciuschiele/decker/core"
	"golang.org/x/crypto/ssh/terminal"
)

var loginCommand = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Decker registry server, if no server is specified \"https://registry.godecker.io/\" is the default.",
	RunE: func(cmd *cobra.Command, args []string) error {
		server := ""

		if len(args) > 0 {
			server = args[0]
		}

		username := cmd.Flag("username").Value.String()
		password := cmd.Flag("password").Value.String()

		if username == "" {
			password = ""

			fmt.Printf("Username: ")
			scanner := bufio.NewScanner(os.Stdin)
			if !scanner.Scan() {
				return nil
			}
			username = scanner.Text()
		}

		if password == "" {
			fmt.Printf("Password: ")
			passwordInBytes, err := terminal.ReadPassword(int(syscall.Stdin))
			fmt.Println()

			if err != nil {
				return err
			}
			password = string(passwordInBytes)
		}

		err := core.Login(username, password, server)

		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("Welcome!")
		fmt.Println()

		return nil
	},
}

func init() {
	loginCommand.Flags().StringP("username", "u", "", "Username")
	loginCommand.Flags().StringP("password", "p", "", "Password")

	rootCmd.AddCommand(loginCommand)
}
