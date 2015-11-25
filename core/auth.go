package core

import "os"

// GetToken ...
func GetToken() string {
	return os.Getenv("Decker.Token")
}

// GetService ...
func GetService() string {
	return os.Getenv("Decker.Service")
}

// IsAuthenticated ...
func IsAuthenticated() bool {
	if GetToken() == "" {
		return false
	}
	return true
}

// Login ...
func Login(username, password, service string) error {
	os.Setenv("Decker.AccessToken", "sda6sda7s6adsa7sd68asd")
	os.Setenv("Decker.Service", service)

	return nil
}
