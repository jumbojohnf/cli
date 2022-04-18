package config_test

import (
	"testing"

	"github.com/funcgql/cli/config"
	"github.com/stretchr/testify/require"
)

func Test_LoadIn(t *testing.T) {
	result, err := config.LoadFrom("fixtures")

	require.NoError(t, err)
	require.Equal(t, &config.Config{
		RootPath:    "GQL_ROOT_PATH",
		RootAbsPath: "fixtures/GQL_ROOT_PATH",
		AWS: &config.AWSConfig{
			Lambda: config.LambdaConfig{
				RoleName: "LAMBDA_ROLE_NAME",
			},
		},
	}, result)
}
