package version_test

import (
	"testing"

	"github.com/funcgql/cli/go/version"
	"github.com/stretchr/testify/require"
)

func Test_Curernt(t *testing.T) {
	result := version.Current()

	require.NotEmpty(t, result.Full())
	require.NotEmpty(t, result.MajorMinor())
}
