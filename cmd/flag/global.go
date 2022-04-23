package flag

import (
	"github.com/spf13/cobra"
)

func AddGlobalFlags(rootCmd *cobra.Command) {
	addGlobalTargetFunctionTypeFlags(rootCmd)
}
