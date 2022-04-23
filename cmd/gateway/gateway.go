package gateway

import (
	"github.com/spf13/cobra"
)

var GatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Perform operations related to the Apollo gateway",
}

func init() {
	GatewayCmd.AddCommand(newCmd)
}
