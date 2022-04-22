package cmd

import (
	"os"

	"github.com/funcgql/cli/cmd/aws"
	"github.com/funcgql/cli/cmd/flag"
	"github.com/funcgql/cli/cmd/gateway"
	"github.com/funcgql/cli/cmd/subgraph"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "funcgql",
	Short: "A CLI for writing a federated GQL application on serverless cloud providers",
	Long:  "The CLI can be configured via a funcgql.yaml file within the Git repository",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(
		&flag.IsLambdaTargetFunctionType,
		"lambda",
		false,
		"If AWS lambda is a deploy target",
	)

	rootCmd.AddCommand(aws.AWSCmd)
	rootCmd.AddCommand(gateway.GatewayCmd)
	rootCmd.AddCommand(subgraph.SubgraphCmd)
}
