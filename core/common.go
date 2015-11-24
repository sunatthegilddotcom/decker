package core

import (
	"os"
	"os/user"
)

// GetOSUserFullname ...
func GetOSUserFullname() string {
	currentUser, _ := user.Current()

	if currentUser == nil {
		return ""
	}

	return currentUser.Name
}

// IsDirectory ...
func IsDirectory(path string) (bool, error) {
	stat, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if !stat.IsDir() {
		return false, nil
	}

	return true, nil
}

// IsFile ...
func IsFile(path string) (bool, error) {
	stat, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return false, nil
	}

	return true, nil
}
