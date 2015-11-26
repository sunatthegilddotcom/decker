package pkg

import (
	"encoding/json"
	"errors"
	"os"
	"path"

	"github.com/kennygrant/sanitize"
)

const (
	// PackageFileName ...
	PackageFileName = "package.json"
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

// LoadFile ...
func LoadFile(inputPath string) (*Package, error) {
	file, err := os.Open(inputPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	pkg := new(Package)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(pkg)

	if err != nil {
		return nil, err
	}

	return pkg, nil
}

// Save ...
func (p *Package) Save(output string) error {
	fileName := path.Join(output, PackageFileName)

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

// Check ...
func (p *Package) Check() error {
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}

	if p.Name != sanitize.Name(p.Name) {
		return errors.New("name has invalid characters")
	}

	if p.Version != sanitize.Name(p.Version) {
		return errors.New("version has invalid characters")
	}

	if p.Runtime == "" {
		return errors.New("runtime cannot be empty")
	}

	if p.Runtime != sanitize.Name(p.Runtime) {
		return errors.New("runtime has invalid characters")
	}

	if p.Main == "" {
		return errors.New("main cannot be empty")
	}

	if p.Main != sanitize.Name(p.Main) {
		return errors.New("Main has invalid characters")
	}

	return nil
}
