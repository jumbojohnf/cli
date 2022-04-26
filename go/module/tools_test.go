package module

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Tools(t *testing.T) {
	m := module{
		absPath: "fixtures/module1",
	}
	results, err := m.Tools()

	require.NoError(t, err)
	require.Equal(t, []Tool{
		{
			ImportPath:    "github.com/golang/mock/mockgen",
			moduleAbsPath: m.absPath,
		},
	}, results)
}
