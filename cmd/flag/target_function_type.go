package flag

import "github.com/funcgql/cli/functype"

var IsLambdaTargetFunctionType bool

func TargetFunctionTypes() []functype.FunctionType {
	var results []functype.FunctionType
	if IsLambdaTargetFunctionType {
		results = append(results, functype.Lambda)
	}
	return results
}
