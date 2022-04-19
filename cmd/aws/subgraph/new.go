package subgraph

import (
	"fmt"

	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/gqlgen"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [name of new module]",
	Short: "Create a new subgraph lambda function go module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadFromRepoRoot()
		if err != nil {
			return err
		}

		moduleName := args[0]
		fmt.Println("🐭 Creating go module", moduleName)
		newModule, err := module.New(moduleName, cfg.RootAbsPath)
		if err != nil {
			errors.Wrapf(err, "failed to create new go module %s", moduleName)
		}

		fmt.Println("🚧 Generating subgraph initial code", moduleName)
		if err := gqlgen.NewAPI().Init(newModule.AbsPath(), moduleName); err != nil {
			return errors.Wrapf(err, "failed to run gqlgen init in %s", cfg.RootAbsPath)
		}

		// Run tidy last after all the generated code is in place.
		fmt.Println("🧹 Tidying", moduleName)
		if err := newModule.Tidy(); err != nil {
			return errors.Wrapf(err, "failed to tidy %s", newModule.Name())
		}
		return nil
	},
}
