package aws

import (
	"fmt"
	"os"
	"strings"

	"github.com/funcgql/cli/shell"
	"github.com/neakor/cloud-functions-gql/util"
	"github.com/pkg/errors"
)

func InstallCLI() error {
	if needsInstall, _ := needsInstall(); !needsInstall {
		fmt.Println("✅  AWS CLI already installed")
		return nil
	}

	installer, err := util.Download(cliInstallerURL)
	if err != nil {
		return errors.Wrap(err, "failed to download AWS CLI installer")
	}
	defer os.Remove(installer.Name())

	fmt.Println("🌤  Installing AWS CLI")
	if output, err := shell.Execute("sudo", "installer", "-pkg", installer.Name(), "-target", "/"); err != nil {
		return errors.Wrapf(err, "failed to run AWS CLI installer %s %s", installer.Name(), output.Combined)
	}

	fmt.Println("✅  Please run 'aws configure' to finish configuring the AWS CLI")

	return nil
}

const (
	cliVersion      = "aws-cli/2.5.2"
	cliInstallerURL = "https://awscli.amazonaws.com/AWSCLIV2.pkg"
)

func needsInstall() (bool, error) {
	api, err := NewAPI()
	if err != nil {
		return false, err
	}

	output, err := api.execute("--version")
	if err != nil {
		return true, err
	}

	return !strings.Contains(output.Combined, cliVersion), nil
}
