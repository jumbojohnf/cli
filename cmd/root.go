package cmd

import (
	"os"

	"github.com/funcgql/cli/cmd/aws"
	"github.com/funcgql/cli/cmd/subgraph"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "funcgql",
	Short: "A CLI for writing a federated GQL application on serverless cloud providers",
	Long:  "The CLI can be configured via a funcgql.json file within the Git repository",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(aws.AWSCmd)
	rootCmd.AddCommand(subgraph.SubgraphCmd)
}
