package rover

import (
	"os"
	"path/filepath"

	"github.com/funcgql/cli/shell"
)

type API interface {
	HasCLI() (bool, error)
	InstallCLI() error
}

func NewAPI(shellAPI shell.API) (API, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	const binaryRelativePath = ".rover/bin/rover"
	binaryPath := filepath.Join(userHomeDir, binaryRelativePath)
	return &api{
		binaryPath: binaryPath,
		shellAPI:   shellAPI,
	}, nil
}

type api struct {
	binaryPath string
	shellAPI   shell.API
}
