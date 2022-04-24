package aws

import (
	"os"
	"strings"

	"github.com/neakor/cloud-functions-gql/util"
	"github.com/pkg/errors"
)

func (a *api) HasCLI() (bool, error) {
	output, err := a.execute("--version")
	if err != nil {
		return true, err
	}

	return !strings.Contains(output.Combined, cliVersion), nil
}

func (a *api) InstallCLI() error {
	installer, err := util.Download(cliInstallerURL)
	if err != nil {
		return errors.Wrap(err, "failed to download AWS CLI installer")
	}
	defer os.Remove(installer.Name())

	if output, err := a.shellAPI.Execute("sudo", "installer", "-pkg", installer.Name(), "-target", "/"); err != nil {
		return errors.Wrapf(err, "failed to run AWS CLI installer %s %s", installer.Name(), output.Combined)
	}

	return nil
}

const (
	cliVersion      = "aws-cli/2.5.2"
	cliInstallerURL = "https://awscli.amazonaws.com/AWSCLIV2.pkg"
)
