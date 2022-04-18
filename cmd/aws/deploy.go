package aws

import (
	"github.com/funcgql/cli/aws"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the given module to AWS Lambda",
	RunE: func(cmd *cobra.Command, args []string) error {
		awsAPI, err := aws.NewAPI()
		if err != nil {
			return err
		}
		return awsAPI.CreateLambdaRole()
	},
}
