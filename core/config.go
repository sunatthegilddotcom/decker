package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	// DeckerFile ...
	DeckerFile = ".decker"
)

// Config ...
type Config struct {
	Token   string `json:"token,omitempty"`
	Service string `json:"service,omitempty"`
}

// Get ...
func (c *Config) Get(name string) (string, error) {
	switch strings.ToLower(name) {
	case "token":
		return c.Token, nil
	case "service":
		return c.Service, nil
	default:
		return "", errors.New(name + " is an invalid name")
	}
}

// Set ...
func (c *Config) Set(name, value string) error {

	switch strings.ToLower(name) {
	case "token":
		c.Token = value
	case "service":
		c.Service = value
	default:
		return errors.New(name + " is an invalid name")
	}

	return nil
}

// GetConfig ...
func GetConfig() (*Config, error) {
	config := new(Config)

	file, err := os.Open(path.Join(os.Getenv("HOME"), DeckerFile))

	if os.IsNotExist(err) {
		return config, nil
	}

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

// SaveConfig ...
func SaveConfig(config *Config) error {
	b, _ := json.MarshalIndent(config, "", "  ")

	configPath := path.Join(os.Getenv("HOME"), DeckerFile)
	return ioutil.WriteFile(configPath, b, 0777)
}
