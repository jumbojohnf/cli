package config_test

import (
	"testing"

	"github.com/funcgql/cli/config"
	"github.com/stretchr/testify/require"
)

func Test_LoadIn_aws_lambda(t *testing.T) {
	result, err := config.LoadFrom("fixtures/aws_lambda")

	require.NoError(t, err)
	require.Equal(t, &config.Config{
		GraphModulesRelPath: "GQL_ROOT_PATH",
		GraphModulesAbsPath: "fixtures/aws_lambda/GQL_ROOT_PATH",
		AWS: &config.AWSConfig{
			Lambda: &config.LambdaConfig{
				RoleName: "LAMBDA_ROLE_NAME",
			},
		},
	}, result)
}

func Test_LoadIn_base(t *testing.T) {
	result, err := config.LoadFrom("fixtures/base")

	require.NoError(t, err)
	require.Equal(t, &config.Config{
		GraphModulesRelPath: "GQL_ROOT_PATH",
		GraphModulesAbsPath: "fixtures/base/GQL_ROOT_PATH",
		AWS:                 nil,
	}, result)
}
