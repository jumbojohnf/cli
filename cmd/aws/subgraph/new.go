package subgraph

import (
	"path/filepath"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/go/tools"
	"github.com/funcgql/cli/go/work"
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
		moduleAbsPath := filepath.Join(cfg.RootAbsPath, module.DirNameFromModuleName(moduleName))
		targetModule := module.WithNameAndAbsPath(moduleName, moduleAbsPath)
		newModuleDir := cliio.DirOf(targetModule.AbsPath())
		if err := newModuleDir.Make(); err != nil {
			return errors.Wrapf(err, "failed to create new module directory %s", newModuleDir.AbsPath())
		}

		toolsGoTemplate := tools.NewTemplate(targetModule.Name())
		if _, err := toolsGoTemplate.Export(cfg.RootAbsPath); err != nil {
			return errors.Wrap(err, "failed to create new module tools.go file")
		}

		if _, err := targetModule.Export(cfg.RootAbsPath); err != nil {
			return errors.Wrap(err, "failed to create new go.mod file")
		}

		// Export go.work file last since it needs to include the newly created module.
		goWorkTemplate := work.NewTemplate()
		if _, err := goWorkTemplate.Export(cfg.RootAbsPath); err != nil {
			return errors.Wrapf(err, "failed to create new module go.work file in %s", cfg.RootAbsPath)
		}

		if err := targetModule.Tidy(); err != nil {
			return errors.Wrapf(err, "failed to tidy %s", targetModule.Name())
		}
		// TODO: Run gqlgen init
		// TODO: Remove server.go
		// TODO: Generate main.go
		return nil
	},
}
