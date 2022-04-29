package aws

import (
	"fmt"

	"github.com/funcgql/cli/aws"
	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/repopath"
	"github.com/funcgql/cli/rover"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the local environment for AWS development",
	RunE: func(cmd *cobra.Command, args []string) error {
		shellAPI := shell.NewAPI()
		repoPathAPI := repopath.NewAPI(shellAPI)

		cfg, err := config.LoadFromRepoRoot(repoPathAPI)
		if err != nil {
			return err
		}

		if cfg.AWS == nil {
			return errors.Errorf("missing AWS configuration %s", config.ConfigFilename)
		}

		graphModulesDir := cliio.DirOf(cfg.GraphModulesAbsPath)

		if graphModuleDirExists, err := graphModulesDir.Exists(); err != nil {
			return errors.Wrapf(err, "could not find config file in: ", graphModulesDir.AbsPath())
		} else if !graphModuleDirExists {
			return errors.Errorf("config file does not exist in repo's directory🤬\nCurrent repo directory: ", graphModulesDir)
		}

		fmt.Println("🌳 Setting up AWS development environment in", graphModulesDir)

		roverAPI, err := rover.NewAPI(shellAPI)
		if err != nil {
			return err
		}
		if hasRover, err := roverAPI.HasCLI(); err != nil {
			return errors.Wrap(err, "failed to determine if Apollo Rover CLI is already installed")
		} else if hasRover {
			fmt.Println("✅  Apollo Rover CLI already installed")
		} else {
			if err := roverAPI.InstallCLI(); err != nil {
				return errors.Wrap(err, "failed to install Apollo Rover CLI")
			}
		}

		awsAPI, err := aws.NewAPI(shellAPI, repoPathAPI, cfg)
		if err != nil {
			return errors.Wrap(err, "failed to instantiate AWS API")
		}
		if err := awsAPI.InstallCLI(); err != nil {
			return errors.Wrap(err, "failed to install AWS CLI")
		}

		return nil
	},
}
