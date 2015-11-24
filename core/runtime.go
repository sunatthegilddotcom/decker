package core

// Runtime ...
type Runtime struct {
	Name           string
	FileExtensions []string
}

// Runtimes ...
var Runtimes [2]*Runtime

func init() {
	Runtimes[0] = &Runtime{Name: "exec", FileExtensions: []string{"", ".exe"}}
	Runtimes[1] = &Runtime{Name: "bash", FileExtensions: []string{".sh"}}
}

// FindRuntimeByExtension ...
func FindRuntimeByExtension(fileExtension string) *Runtime {
	for _, runtime := range Runtimes {
		for _, ext := range runtime.FileExtensions {
			if fileExtension == ext {
				return runtime
			}
		}
	}

	return nil
}
