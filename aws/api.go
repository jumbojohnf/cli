package aws

import (
	"github.com/funcgql/cli/config"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type API interface {
	CreateLambdaRole() error

	execute(args ...string) (shell.Output, error)
}

func NewAPI() (API, error) {
	cfg, err := config.LoadFromRepoRoot()
	if err != nil {
		return nil, err
	} else if cfg.AWS == nil {
		return nil, errors.New("missing AWS configuration")
	}

	const binaryPath = "aws"
	return &api{
		binaryPath: binaryPath,
		cfg:        *cfg.AWS,
	}, nil
}

type api struct {
	binaryPath string
	cfg        config.AWSConfig
}

func (a *api) execute(args ...string) (shell.Output, error) {
	return shell.Execute(a.binaryPath, args...)
}
