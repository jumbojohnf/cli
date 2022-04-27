package aws

import (
	"github.com/funcgql/cli/aws"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/repopath"
	"github.com/funcgql/cli/shell"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the given module to AWS Lambda",
	RunE: func(cmd *cobra.Command, args []string) error {
		shellAPI := shell.NewAPI()
		repoPathAPI := repopath.NewAPI(shellAPI)
		cfg, err := config.LoadFromRepoRoot(repoPathAPI)
		if err != nil {
			return err
		}

		awsAPI, err := aws.NewAPI(shellAPI, repoPathAPI, cfg)
		if err != nil {
			return err
		}
		return awsAPI.CreateLambdaRole()
	},
}
