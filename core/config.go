package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

const (
	// ConfigFileName contains the config file name
	ConfigFileName = ".decker"
	// DefaultServer contains the public decker registry
	DefaultServer = "http://registry.godecker.io/"
)

var (
	// Config is a singleton instance of ConfigFile
	Config *ConfigFile
	// ConfigDir contains the config folder
	ConfigDir string
)

// ConfigFile ~/.decker/config.json file info
type ConfigFile struct {
	Auths         map[string]string `json:"auths"`
	DefaultServer string            `json:"defaultServer,omitempty"`
	filename      string            // Internal use
}

// InitConfig ...
func InitConfig() error {
	Config.filename = path.Join(ConfigDir, ConfigFileName)

	file, err := os.Open(Config.filename)

	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(Config)
}

// Save ...
func (p *ConfigFile) Save() error {
	b, _ := json.MarshalIndent(Config, "", "  ")
	return ioutil.WriteFile(p.filename, b, 0777)
}

func init() {
	Config = &ConfigFile{Auths: make(map[string]string, 0)}
	ConfigDir = os.Getenv("HOME")
}
