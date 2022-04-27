package aws

import (
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/repopath"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type API interface {
	CreateLambdaRole() error
	HasCLI() (bool, error)
	InstallCLI() error
}

func NewAPI(shellAPI shell.API, repoPathAPI repopath.API, cfg *config.Config) (API, error) {
	if cfg.AWS == nil {
		return nil, errors.New("missing AWS configuration")
	}

	const binaryPath = "aws"
	return &api{
		binaryPath: binaryPath,
		cfg:        *cfg.AWS,
		shellAPI:   shellAPI,
	}, nil
}

type api struct {
	binaryPath string
	cfg        config.AWSConfig
	shellAPI   shell.API
}
