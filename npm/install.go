package npm

import (
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

func InstallIn(dir string, packages ...string) error {
	args := append([]string{"install"}, packages...)
	if output, err := shell.ExecuteIn(dir, "npm", args...); err != nil {
		return errors.Wrapf(err, "failed to install npm packages %s %s", packages, output.Combined)
	}
	return nil
}
