package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	// DefaultService has the public registry url
	DefaultService = "http://registry.godecker.io/"
)

var configFile string
var configDefaults map[string]string
var configOverride map[string]string

// LoadConfig ...
func LoadConfig() error {
	file, err := os.Open(configFile)

	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&configOverride)
}

// SaveConfig ...
func SaveConfig() error {
	b, _ := json.MarshalIndent(configOverride, "", "  ")
	return ioutil.WriteFile(configFile, b, 0777)
}

// GetConfigValue ...
func GetConfigValue(key string) (string, bool) {
	defaultValue, found := configDefaults[key]

	if !found {
		return "", false
	}

	value, _ := configOverride[key]

	if strings.Trim(value, " ") == "" {
		value = defaultValue
	}

	return value, true
}

// SetConfigValue ...
func SetConfigValue(key, value string) bool {
	_, found := configDefaults[key]

	if !found {
		return false
	}

	value = strings.Trim(value, " ")

	if value != "" {
		configOverride[key] = value
	} else {
		delete(configOverride, key)
	}

	return true
}

func init() {
	configFile = path.Join(os.Getenv("HOME"), ".decker")

	configDefaults = make(map[string]string)
	configOverride = make(map[string]string)

	configDefaults["token"] = ""
	configDefaults["service"] = DefaultService
}
