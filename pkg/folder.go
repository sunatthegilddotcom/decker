package pkg

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/jhoonb/archivex"
)

const (
	// BinFolder ...
	BinFolder = "bin"
)

// Create ...
func Create(output string) error {
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
	spec.Save(output)

	ioutil.WriteFile(path.Join(binPath, spec.Main), []byte("#!/bin/bash\n\necho \"hello\""), 0777)

	return nil
}

// Check ...
func Check(inputPath string) error {
	jsonPath := path.Join(inputPath, PackageFileName)

	success, err := IsFile(jsonPath)

	if err != nil {
		return err
	}

	if !success {
		return errors.New(PackageFileName + " is missing")
	}

	binPath := path.Join(inputPath, BinFolder)

	success, err = IsDirectory(binPath)

	if err != nil {
		return err
	}

	if !success {
		return errors.New(BinFolder + " is not a directory")
	}

	pkg, err := LoadFile(jsonPath)

	if err != nil {
		return err
	}

	err = pkg.Check()

	if err != nil {
		return err
	}

	success, err = IsFile(path.Join(binPath, pkg.Main))

	if err != nil {
		return err
	}

	if !success {
		return errors.New("file " + pkg.Main + " not found in " + BinFolder + " folder")
	}

	return nil
}

// Pack ...
func Pack(inputPath string, outputPath string) (fileName string, err error) {
	err = Check(inputPath)

	if err != nil {
		return
	}

	packageFile := path.Join(inputPath, PackageFileName)
	binFolder := path.Join(inputPath, BinFolder)

	pkg, err := LoadFile(packageFile)

	if err != nil {
		return
	}

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

	if err = tar.AddFile(packageFile); err != nil {
		return
	}

	if err = tar.AddAll(binFolder, true); err != nil {
		return
	}

	return
}
