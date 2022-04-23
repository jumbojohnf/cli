package flag

import (
	"fmt"

	"github.com/funcgql/cli/functype"
	"github.com/spf13/cobra"
)

var AllTargetFunctionTypeFlags = functype.All

func TargetFunctionTypes() []functype.FunctionType {
	var results []functype.FunctionType
	if targetLambdaFunctionType {
		results = append(results, functype.Lambda)
	}
	return results
}

var (
	targetLambdaFunctionType bool

	targetFunctionFlagRefs = map[functype.FunctionType]*bool{
		functype.Lambda: &targetLambdaFunctionType,
	}
)

func addGlobalTargetFunctionTypeFlags(rootCmd *cobra.Command) {
	for _, functionType := range functype.All {
		rootCmd.PersistentFlags().BoolVar(
			targetFunctionFlagRefs[functionType],
			string(functionType),
			false,
			fmt.Sprintf("If %s is a deploy target", functionType),
		)
	}
}
