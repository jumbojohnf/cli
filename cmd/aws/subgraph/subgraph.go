package subgraph

import (
	"github.com/spf13/cobra"
)

var SubgraphCmd = &cobra.Command{
	Use:   "subgraph",
	Short: "Perform operations related to a single subgraph",
}

func init() {
	SubgraphCmd.AddCommand(newCmd)
}
