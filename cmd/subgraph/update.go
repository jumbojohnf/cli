package subgraph

import (
	"fmt"
	"os"

	"github.com/funcgql/cli/cmd/flag"
	"github.com/funcgql/cli/config"
	"github.com/funcgql/cli/go/module"
	"github.com/funcgql/cli/gqlgen"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [name of subgraph module | subgraph module in current directory]",
	Short: "Update the source code of a subgraph module",
	Long: "Update the source code of the specified subgraph module or the subgraph module in the current working " +
		"directory based on the schema file",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return updateModule(args[0])
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

func updateModule(moduleName string) error {
	cfg, err := config.LoadFromRepoRoot()
	if err != nil {
		return err
	}
	targetModule, exists, err := module.WithName(moduleName, cfg)
	if err != nil {
		return err
	} else if !exists {
		return errors.Errorf("module %s does not exist", moduleName)
	}

	fmt.Println("🏗  Updating subgraph source code of", moduleName)
	if err := gqlgen.NewAPI().Generate(targetModule.AbsPath()); err != nil {
		return err
	}

	fmt.Println("✅ Successfully updated module", moduleName)
	return nil
}

func updateCurrentDir() error {
	workingDirPath, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to determine current working directory path")
	}

	fmt.Println("🏗  Updating subgraph source code in current directory")
	if err := gqlgen.NewAPI().Generate(workingDirPath); err != nil {
		return err
	}

	fmt.Println("✅ Successfully updated subgraph module in", workingDirPath)
	return nil
}
