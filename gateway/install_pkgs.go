package gateway

import (
	"github.com/funcgql/cli/npm"
	"github.com/pkg/errors"
)

func (g gateway) InstallPackages() error {
	var packages = []string{
		"@apollo/gateway",
		"apollo-server-lambda",
		"aws-lambda",
		"aws-sdk",
	}

	if err := npm.InstallIn(g.absPath, packages...); err != nil {
		return errors.Wrapf(err, "failed to install NPM packages in %s", g.absPath)
	}
	return nil
}
