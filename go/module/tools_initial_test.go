package module

import (
	"testing"

	"github.com/funcgql/cli/shell"
	"github.com/funcgql/cli/shell/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_InstallInitialTools_hasVersion(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := module{
		absPath: "fixtures/module1",
	}

	shellAPI := mocks.NewMockAPI(mockCtrl)

	shellAPI.EXPECT().
		ExecuteWithIOIn("fixtures/module1", "go", "install", "github.com/golang/mock/mockgen@v1.6.0").
		Return(shell.Output{}, nil)

	err := m.InstallInitialTools(shellAPI)

	require.NoError(t, err)
}
