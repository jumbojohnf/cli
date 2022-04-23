package rover

import (
	"os"
	"path/filepath"

	"github.com/funcgql/cli/shell"
)

type API interface {
	execute(args ...string) (shell.Output, error)
}

func NewAPI() (API, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	const binaryRelativePath = ".rover/bin/rover"
	binaryPath := filepath.Join(userHomeDir, binaryRelativePath)
	return &api{binaryPath: binaryPath}, nil
}

type api struct {
	binaryPath string
}

func (a *api) execute(args ...string) (shell.Output, error) {
	return shell.Execute(a.binaryPath, args...)
}
