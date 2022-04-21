package aws

import (
	"fmt"

	"github.com/funcgql/cli/apollo"
	"github.com/funcgql/cli/aws"
	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/tools"
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

		if _, err := dir.Exists(); err != nil {
			return errors.Errorf("could not find config file", err)
		}

		if bool, err := dir.Exists(); !bool && err == nil {
			return errors.Errorf("config file does not exist in working directory🤬\nCurrent working directory: ", dir)
		}

		if bool, err := dir.Exists(); bool && err == nil {
			fmt.Println("🌳 Setting up AWS development environment in", dir)
		}

		if err := tools.InstallAllIn(cfg.GraphModulesAbsPath); err != nil {
			return errors.Wrap(err, "failed to install go tools")
		}

		if err := apollo.InstallRoverCLI(); err != nil {
			return errors.Wrap(err, "failed to install Apollo Rover CLI")
		}

		if err := aws.InstallCLI(); err != nil {
			return errors.Wrap(err, "failed to install AWS CLI")
		}

		return nil
	},
}
