package gateway

import (
	"fmt"

	"github.com/funcgql/cli/cmd/flag"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/gateway"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Apollo gateway for the deploy target(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		functionTypes := flag.TargetFunctionTypes()
		if len(functionTypes) <= 0 {
			return errors.New("at least one cloud function type flag must be specified")
		}

		cfg, err := config.LoadFromRepoRoot()
		if err != nil {
			return err
		}

		for _, functionType := range functionTypes {
			fmt.Println("🌉 Generating", functionType, "Apollo gateway")
			newGateway := gateway.New(functionType, cfg)
			if err := newGateway.ExportIndexFile(cfg); err != nil {
				return err
			}

			fmt.Println("📦 Installing NPM packages")
			if err := newGateway.InstallPackages(); err != nil {
				return err
			}

			fmt.Println("✅ Added new", functionType, "Apollo gateway")
		}

		return nil
	},
}
