package aws

import (
	"github.com/funcgql/cli/config"

	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

type AWS struct {
	binaryPath string
	cfg        config.AWSConfig
}

func NewAPI() (*AWS, error) {
	cfg, err := config.LoadFromRepoRoot()
	if err != nil {
		return nil, err
	} else if cfg.AWS == nil {
		return nil, errors.New("missing AWS configuration")
	}

	const binaryPath = "aws"
	return &AWS{
		binaryPath: binaryPath,
		cfg:        *cfg.AWS,
	}, nil
}

func (r *AWS) execute(args ...string) (shell.Output, error) {
	return shell.Execute(r.binaryPath, args...)
}
