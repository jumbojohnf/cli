package apollo

import (
	"fmt"
	"strings"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
)

func InstallRoverCLI() error {
	if needsInstall, _ := needsInstall(); !needsInstall {
		fmt.Println("✅  Rover CLI already installed")
		return nil
	}

	installer, err := cliio.Download(cliInstallerURL)
	if err != nil {
		return errors.Wrap(err, "failed to download Rover CLI installer script")
	}

	fmt.Println("🌝 Installing Apollo Rover CLI")
	if _, err := shell.Execute("/bin/sh", installer.Name()); err != nil {
		return err
	}

	return nil
}

const (
	cliVersion      = "0.4.8"
	cliInstallerURL = "https://rover.apollo.dev/nix/latest"
)

func needsInstall() (bool, error) {
	api, err := NewRoverAPI()
	if err != nil {
		return false, err
	}

	output, err := api.execute("--version")
	if err != nil {
		return true, err
	}
	return !strings.Contains(output.Combined, cliVersion), nil
}
