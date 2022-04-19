package aws

import (
	"github.com/funcgql/cli/config"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type API struct {
	binaryPath string
	cfg        config.AWSConfig
}

func NewAPI() (*API, error) {
	cfg, err := config.LoadFromRepoRoot()
	if err != nil {
		return nil, err
	} else if cfg.AWS == nil {
		return nil, errors.New("missing AWS configuration")
	}

	const binaryPath = "aws"
	return &API{
		binaryPath: binaryPath,
		cfg:        *cfg.AWS,
	}, nil
}

func (r *API) execute(args ...string) (shell.Output, error) {
	return shell.Execute(r.binaryPath, args...)
}
