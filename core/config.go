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
	configPath := path.Join(os.Getenv("HOME"), DeckerFile)

	file, err := os.Open(configPath)

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
	configPath := path.Join(os.Getenv("HOME"), DeckerFile)

	b, _ := json.MarshalIndent(config, "", "  ")
	return ioutil.WriteFile(configPath, b, 0777)
}
