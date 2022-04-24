package rover

import (
	"strings"

	"github.com/funcgql/cli/cliio"
	"github.com/pkg/errors"
)

func (a *api) HasCLI() (bool, error) {
	output, err := a.execute("--version")
	if err != nil {
		return false, err
	}
	return !strings.Contains(output.Combined, cliVersion), nil
}

func (a *api) InstallCLI() error {
	installer, err := cliio.Download(cliInstallerURL)
	if err != nil {
		return errors.Wrap(err, "failed to download Rover CLI installer script")
	}

	if _, err := a.shellAPI.Execute("/bin/sh", installer.Name()); err != nil {
		return err
	}

	return nil
}

const (
	cliVersion      = "0.4.8"
	cliInstallerURL = "https://rover.apollo.dev/nix/latest"
)
