package module

import (
	"testing"

	"github.com/funcgql/cli/shell"
	"github.com/funcgql/cli/shell/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Install(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	targetTool := Tool{
		ImportPath:    "IMPORT_PATH",
		moduleAbsPath: "MODULE_PATH",
	}

	shellAPI := mocks.NewMockAPI(mockCtrl)
	shellAPI.EXPECT().
		ExecuteWithIOIn("MODULE_PATH", "go", "install", "IMPORT_PATH@VERSION").
		Return(shell.Output{}, nil)

	err := targetTool.Install("VERSION", shellAPI)

	require.NoError(t, err)
}
