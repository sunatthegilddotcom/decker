package core

import (
	"os"
	"os/user"
	"strings"
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

// CompactServer ...
func CompactServer(server string) string {
	server = strings.ToLower(strings.Trim(server, " "))

	if server == "" {
		return ""
	}

	if !strings.HasSuffix(server, "/") {
		server += "/"
	}

	if server == DefaultServer {
		return ""
	}

	return server
}

// ResolveServer ...
func ResolveServer(server string) string {
	server = strings.Trim(server, " ")

	if server == "" {
		server = Config.DefaultServer
	}

	if server == "" {
		server = DefaultServer
	}

	server = strings.ToLower(server)

	if !strings.HasSuffix(server, "/") {
		server += "/"
	}

	return server
}
