package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

const (
	// ConfigFileName ...
	ConfigFileName = "config.json"
	// DefaultServer contains the public decker registry
	DefaultServer = "http://registry.godecker.io/"
)

var (
	filename   string
	properties map[string]string
)

// Delete ...
func Delete(key string) {
	delete(properties, key)
}

// Get ...
func Get(key string) string {
	return properties[key]
}

// GetServer ...
func GetServer() string {
	if server := Get("server"); server != "" {
		return server
	}
	return DefaultServer
}

// GetToken ...
func GetToken() string {
	return Get("token")
}

// List ...
func List() map[string]string {
	props := make(map[string]string)

	for k, v := range properties {
		props[k] = v
	}

	return props
}

// Set ...
func Set(key, value string) {
	properties[key] = value
}

// Save ...
func Save() error {
	b, _ := json.MarshalIndent(properties, "", "  ")
	return ioutil.WriteFile(filename, b, 0777)
}

// Init ...
func Init(baseDir string) error {
	if err := os.MkdirAll(baseDir, 0777); err != nil {
		return err
	}

	filename = path.Join(baseDir, ConfigFileName)

	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&properties)
}

func init() {
	properties = make(map[string]string)
}
