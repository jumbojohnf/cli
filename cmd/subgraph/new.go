package subgraph

import (
	"fmt"

	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/functype"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/gqlgen"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [name of new module]",
	Short: "Create a new subgraph function go module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		functionTypes := parseFunctionTypes()
		if len(functionTypes) <= 0 {
			return errors.New("at least one cloud function type flag must be specified")
		}

		cfg, err := config.LoadFromRepoRoot()
		if err != nil {
			return err
		}

		moduleName := args[0]
		fmt.Println("🐭 Creating go module", moduleName)
		newModule, err := module.New(moduleName, cfg.GraphModulesAbsPath)
		if err != nil {
			errors.Wrapf(err, "failed to create new go module %s", moduleName)
		}

		fmt.Println("🚧 Generating subgraph initial code", moduleName)
		if err := gqlgen.NewAPI().Init(newModule.AbsPath(), moduleName, functionTypes); err != nil {
			return errors.Wrapf(err, "failed to run gqlgen init in %s", cfg.GraphModulesAbsPath)
		}

		// Run tidy last after all the generated code is in place.
		fmt.Println("🧹 Tidying", moduleName)
		if err := newModule.Tidy(); err != nil {
			return errors.Wrapf(err, "failed to tidy %s", newModule.Name())
		}
		return nil
	},
}

var (
	lambda bool
)

func init() {
	newCmd.Flags().BoolVar(&lambda, "lambda", false, "If AWS lambda should be generated as a deploy target")
}

func parseFunctionTypes() []functype.FunctionType {
	var results []functype.FunctionType
	if lambda {
		results = append(results, functype.Lambda)
	}
	return results
}
