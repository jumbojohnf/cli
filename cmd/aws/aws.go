package aws

import (
	"github.com/funcgql/cli/cmd/aws/subgraph"
	"github.com/spf13/cobra"
)

var AWSCmd = &cobra.Command{
	Use:   "aws",
	Short: "Perform operations targeting the AWS cloud provider",
}

func init() {
	AWSCmd.AddCommand(deployCmd)
	AWSCmd.AddCommand(initCmd)
	AWSCmd.AddCommand(subgraph.SubgraphCmd)
}
