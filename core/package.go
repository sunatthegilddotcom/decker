package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/jhoonb/archivex"
	"github.com/kennygrant/sanitize"
)

const (
	// BinFolder ...
	BinFolder = "bin"
	// PackageFile ...
	PackageFile = "package.json"
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

// ReadFromFile ...
func (p *Package) ReadFromFile(input string) error {
	file, err := os.Open(input)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(p)

	if err != nil {
		return err
	}

	return nil
}

// WriteToFile ...
func (p *Package) WriteToFile(output string) error {
	fileName := path.Join(output, PackageFile)

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

// Validate ...
func (p *Package) Validate() error {
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

// PackageManager ...
type PackageManager struct {
}

// Init ...
func (p *PackageManager) Init(output string) error {
	err := os.MkdirAll(output, 0777)

	if err != nil {
		return err
	}

	binPath := path.Join(output, "bin")

	os.Mkdir(binPath, 0777)

	spec := new(Package)
	spec.Name = path.Base(output)
	spec.Version = "0.0.0"
	spec.Description = "A short description of your package"
	spec.Author = GetOSUserFullname()
	spec.Keywords = []string{}
	spec.Repository = "https://github.com/godecker/" + spec.Name
	spec.Runtime = "bash"
	spec.Main = spec.Name + ".sh"
	spec.WriteToFile(output)

	ioutil.WriteFile(path.Join(binPath, spec.Main), []byte("#!/bin/bash\n\necho \"hello\""), 0777)

	return nil
}

// Check ...
func (p *PackageManager) Check(inputPath string) error {
	jsonPath := path.Join(inputPath, PackageFile)

	success, err := IsFile(jsonPath)

	if err != nil {
		return err
	}

	if !success {
		return errors.New("package folder: " + PackageFile + " not found")
	}

	binPath := path.Join(inputPath, BinFolder)

	success, err = IsDirectory(binPath)

	if err != nil {
		return err
	}

	if !success {
		return errors.New("package folder: " + BinFolder + " is not a directory")
	}

	pkg := new(Package)

	if err := pkg.ReadFromFile(jsonPath); err != nil {
		return err
	}

	err = pkg.Validate()

	if err != nil {
		return errors.New(PackageFile + ": " + err.Error())
	}

	success, err = IsFile(path.Join(binPath, pkg.Main))

	if err != nil {
		return err
	}

	if !success {
		return errors.New(PackageFile + " : file " + pkg.Main + " not found in " + BinFolder + " folder")
	}

	return nil
}

// Pack ...
func (p *PackageManager) Pack(inputPath string, outputPath string) (fileName string, err error) {
	err = p.Check(inputPath)

	if err != nil {
		return
	}

	jsonPath := path.Join(inputPath, PackageFile)
	binPath := path.Join(inputPath, BinFolder)

	pkg := new(Package)
	pkg.ReadFromFile(jsonPath)

	fileName = pkg.Name
	if pkg.Version != "" {
		fileName += "-" + pkg.Version
	}
	fileName += ".tar.gz"
	filePath := path.Join(outputPath, fileName)

	tar := new(archivex.TarFile)
	tar.Compressed = true

	if err = tar.Create(filePath); err != nil {
		return
	}

	defer tar.Close()

	if err = tar.AddFile(jsonPath); err != nil {
		return
	}

	if err = tar.AddAll(binPath, true); err != nil {
		return
	}

	return
}
