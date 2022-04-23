package subgraph

import (
	"fmt"

	"github.com/funcgql/cli/cmd/flag"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/module"
	goworktemplate "github.com/funcgql/cli/go/work/template"
	"github.com/funcgql/cli/gqlgen"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [name of new module]",
	Short: "Create a new subgraph function go module",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		functionTypes := flag.TargetFunctionTypes()
		if len(functionTypes) <= 0 {
			return errors.New("at least one cloud function type flag must be specified")
		}

		cfg, err := config.LoadFromRepoRoot()
		if err != nil {
			return err
		}

		moduleName := args[0]
		fmt.Println("🐭 Creating go module", moduleName)
		newModule, err := module.New(moduleName, cfg)
		if err != nil {
			errors.Wrapf(err, "failed to create new go module %s", moduleName)
		}
		if err := newModule.InstallTools(); err != nil {
			return err
		}

		// Export go.work file after module is created since it needs to include the newly created module.
		fmt.Println("🐭 Updating go.work file")
		goWorkTemplate := goworktemplate.New()
		if err := goWorkTemplate.Export(cfg.GraphModulesAbsPath); err != nil {
			return errors.Wrapf(err, "failed to update go.work file in %s", cfg.GraphModulesAbsPath)
		}

		fmt.Println("🚧 Generating subgraph initial code")
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
		fmt.Println("🏗  Updating subgraph initial code")
		if err := gqlgenAPI.Generate(newModule.AbsPath()); err != nil {
			return errors.Wrapf(err, "failed to update generated GQL code in %s", newModule.AbsPath())
		}

		fmt.Println("✅ Added new", moduleName, "subgraph")
		return nil
	},
}
