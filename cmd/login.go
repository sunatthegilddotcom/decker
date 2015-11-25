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

func init() {
	var username string
	var password string

	var loginCommand = &cobra.Command{
		Use:   "login",
		Short: "Log in to a Decker registry server, if no server is specified \"https://registry.godecker.io/v1/\" is the default.",
		Run: func(cmd *cobra.Command, args []string) {
			service := "https://registry.godecker.io/v1/"

			if len(args) > 0 {
				service = args[0]
			}

			if username == "" {
				password = ""

				fmt.Printf("Username: ")
				scanner := bufio.NewScanner(os.Stdin)
				if !scanner.Scan() {
					return
				}
				username = scanner.Text()
			}

			if password == "" {
				fmt.Printf("Password: ")
				passwordInBytes, err := terminal.ReadPassword(int(syscall.Stdin))
				fmt.Println()

				if err != nil {
					return
				}
				password = string(passwordInBytes)
			}

			authManager := core.AuthManager{}
			err := authManager.Login(username, password, service)

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Welcome!")
		},
	}

	loginCommand.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCommand.Flags().StringVarP(&password, "password", "p", "", "Password")

	RootCmd.AddCommand(loginCommand)
}
