package npm

import (
	"github.com/pkg/errors"
)

func (a *api) InstallIn(dir string, packages ...string) error {
	args := append([]string{"install"}, packages...)
	if output, err := a.shellAPI.ExecuteWithIOIn(dir, "npm", args...); err != nil {
		return errors.Wrapf(err, "failed to install npm packages %s %s", packages, output.Combined)
	}
	return nil
}
