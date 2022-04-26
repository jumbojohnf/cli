package subgraph

import (
	"fmt"

	"github.com/funcgql/cli/cmd/flag"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/gqlgen"
	"github.com/funcgql/cli/repopath"
	"github.com/funcgql/cli/shell"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [name of subgraph module | subgraph module in the current directory]",
	Short: "Update the source code of a subgraph module",
	Long: "Update the source code of the specified subgraph module or the subgraph module in the current working " +
		"directory based on the schema file",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return updateNamed(args[0])
		} else {
			return updateCurrentDir()
		}
	},
}

func init() {
	updateCmd.SetHelpFunc(func(cmd *cobra.Command, strings []string) {
		for _, targetFunctionTypeFlag := range flag.AllTargetFunctionTypeFlags {
			cmd.Flags().MarkHidden(string(targetFunctionTypeFlag))
		}

		cmd.Parent().HelpFunc()(cmd, strings)
	})
}

func updateNamed(moduleName string) error {
	shellAPI := shell.NewAPI()
	repoPathAPI := repopath.NewAPI(shellAPI)

	cfg, err := config.LoadFromRepoRoot(repoPathAPI)
	if err != nil {
		return err
	}
	targetModule, exists, err := module.WithName(moduleName, cfg)
	if err != nil {
		return err
	} else if !exists {
		return errors.Errorf("module %s does not exist", moduleName)
	}

	return updateModule(targetModule)
}

func updateCurrentDir() error {
	shellAPI := shell.NewAPI()

	currentDirModule, exists, err := module.CurrentDir(shellAPI)
	if err != nil {
		return err
	} else if !exists {
		return errors.Wrap(err, "current directory does not contain a subgraph go module")
	}

	return updateModule(currentDirModule)
}

func updateModule(targetModule module.Module) error {
	shellAPI := shell.NewAPI()

	fmt.Println("🐭 Updating module", targetModule.Name(), "tools")
	if err := targetModule.InstallAllTools(shellAPI); err != nil {
		return err
	}

	fmt.Println("🏗  Updating subgraph source code of", targetModule.Name())
	if err := gqlgen.NewAPI(shellAPI).Generate(targetModule); err != nil {
		return err
	}

	fmt.Println("✅ Successfully updated module", targetModule.Name())
	return nil
}
