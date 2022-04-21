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
		if err := newModule.InstallTools(); err != nil {
			return err
		}

		fmt.Println("🚧 Generating subgraph initial code", moduleName)
		gqlgenAPI := gqlgen.NewAPI()
		if err := gqlgenAPI.Init(newModule.AbsPath(), newModule, functionTypes); err != nil {
			return errors.Wrapf(err, "failed to run initialize GQL in %s", newModule.AbsPath())
		}

		// Run tidy last after all the generated code is in place.
		fmt.Println("🧹 Tidying", moduleName)
		if err := newModule.Tidy(); err != nil {
			return errors.Wrapf(err, "failed to tidy %s", newModule.Name())
		}

		// Run generate again to update to templated schema after module has been updated.
		fmt.Println("🏗  Updating subgraph initial code", moduleName)
		if err := gqlgenAPI.Generate(newModule.AbsPath()); err != nil {
			return errors.Wrapf(err, "failed to update generated GQL code in %s", newModule.AbsPath())
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
