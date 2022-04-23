package aws

import (
	"fmt"

	"github.com/funcgql/cli/aws"
	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/tools"
	"github.com/funcgql/cli/rover"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the local environment for AWS development",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadFromRepoRoot()
		if err != nil {
			return err
		}

		if cfg.AWS == nil {
			return errors.Errorf("missing AWS configuration %s", config.ConfigFilename)
		}

		dir := cliio.DirOf(cfg.GraphModulesAbsPath)

		if bool, err := dir.Exists(); !bool && err == nil {
			return errors.Errorf("config file does not exist in the repo's directory🤬\nCurrent repo directory: ", cfg.GraphModulesAbsPath)
		} else if _, err := dir.Exists(); err != nil {
			return errors.Wrap(err, "could not find config file")
		}

		if bool, err := dir.Exists(); bool && err == nil {
			fmt.Println("🌳 Setting up AWS development environment in", dir)
		}

		if err := tools.InstallAllIn(cfg.GraphModulesAbsPath); err != nil {
			return errors.Wrap(err, "failed to install go tools")
		}

		if err := rover.InstallCLI(); err != nil {
			return errors.Wrap(err, "failed to install Apollo Rover CLI")
		}

		if err := aws.InstallCLI(); err != nil {
			return errors.Wrap(err, "failed to install AWS CLI")
		}

		return nil
	},
}
