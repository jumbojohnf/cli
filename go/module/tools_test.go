package module

import (
	"testing"

	"github.com/funcgql/cli/shell"
	"github.com/funcgql/cli/shell/mocks"
	"github.com/golang/mock/gomock"
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

func Test_Install(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	targetTool := Tool{
		ImportPath:    "IMPORT_PATH",
		moduleAbsPath: "MODULE_PATH",
	}

	shellAPI := mocks.NewMockAPI(mockCtrl)
	shellAPI.EXPECT().
		ExecuteIn("MODULE_PATH", "go", "install", "IMPORT_PATH@VERSION").
		Return(shell.Output{}, nil)

	err := targetTool.Install("VERSION", shellAPI)

	require.NoError(t, err)
}
