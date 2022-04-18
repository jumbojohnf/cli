package apollo

import (
	"os"
	"path/filepath"

	"github.com/funcgql/cli/shell"
)

type Rover struct {
	binaryPath string
}

func NewRoverAPI() (*Rover, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	const binaryRelativePath = ".rover/bin/rover"
	binaryPath := filepath.Join(userHomeDir, binaryRelativePath)
	return &Rover{binaryPath: binaryPath}, nil
}

func (r *Rover) execute(args ...string) (shell.Output, error) {
	return shell.Execute(r.binaryPath, args...)
}
