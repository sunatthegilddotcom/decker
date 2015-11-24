package core

import (
	"encoding/json"
	"os"
	"path"
)

// Package represets a package.json file
type Package struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Keywords    []string `json:"keywords"`
	Repository  string   `json:"repository"`
	Runtime     string   `json:"runtime"`
	Main        string   `json:"main"`
}

// WriteToFile ...
func (p *Package) WriteToFile(output string) error {
	fileName := path.Join(output, "package.json")

	if stat, _ := os.Stat(fileName); stat != nil {
		return nil
	}

	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	b, _ := json.MarshalIndent(p, "", "  ")

	file.Write(b)
	file.Close()

	return nil
}

// PackageManager ...
type PackageManager struct {
}

// Create ...
func (p *PackageManager) Create(output string) error {
	err := os.MkdirAll(output, 0777)

	if err != nil {
		return err
	}

	os.Mkdir(path.Join(output, "bin"), 0777)

	spec := new(Package)
	spec.Name = path.Base(output)
	spec.Version = "0.0.0"
	spec.Description = "A short description of your package"
	spec.Author = GetOSUserFullname()
	spec.Keywords = []string{}
	spec.Repository = "https://github.com/godecker/" + spec.Name
	spec.Runtime = "bash"
	spec.Main = spec.Name
	spec.WriteToFile(output)

	return nil
}
