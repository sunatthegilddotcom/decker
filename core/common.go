package core

import "os/user"

// GetOSUserFullname ...
func GetOSUserFullname() string {
	currentUser, _ := user.Current()

	if currentUser == nil {
		return ""
	}

	return currentUser.Name
}
