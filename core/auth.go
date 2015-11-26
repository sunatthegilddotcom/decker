package core

import "os"

// GetAuth ...
func GetAuth() (server, token string) {
	server = os.Getenv("decker.server")
	token = os.Getenv("decker.token")

	if token == "" {
		server = Config.DefaultServer
		token = Config.Auths[server]
	}

	return server, token
}

// SetAuth ...
func SetAuth(server, token string) {
	os.Setenv("decker.server", server)
	os.Setenv("decker.token", token)
}

// IsAuthenticated ...
func IsAuthenticated() bool {
	_, token := GetAuth()
	return token != ""
}

// Login ...
func Login(username, password, server string) error {
	SetAuth(server, "asd546asd5a6d4a5s")

	server = ResolveServer(server)

	// registry

	server = CompactServer(server)

	Config.Auths[server] = "sas6as6a6s7a7s"
	return Config.Save()
}

// Logout ...
func Logout(server string) error {
	SetAuth("", "")

	server = CompactServer(server)

	delete(Config.Auths, server)
	return Config.Save()
}
