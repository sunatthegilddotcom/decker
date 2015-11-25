package core

import "os"

// AuthManager ...
type AuthManager struct {
}

// GetToken ...
func (auth *AuthManager) GetToken() string {
	return os.Getenv("Decker.Token")
}

// GetService ...
func (auth *AuthManager) GetService() string {
	return os.Getenv("Decker.Service")
}

// IsAuthenticated ...
func (auth *AuthManager) IsAuthenticated() bool {
	if auth.GetToken() == "" {
		return false
	}
	return true
}

// Login ...
func (auth *AuthManager) Login(username, password, service string) error {
	os.Setenv("Decker.AccessToken", "sda6sda7s6adsa7sd68asd")
	os.Setenv("Decker.Service", service)

	return nil
}
