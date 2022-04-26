package module

import (
	"testing"

	"github.com/funcgql/cli/cliio"
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

func Test_InstallAllTools_noVersion(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := module{
		absPath: "fixtures/module1",
	}

	shellAPI := mocks.NewMockAPI(mockCtrl)

	depsFixtureContent, err := cliio.FileOf("fixtures/deps_stdout.txt").StringContent()
	require.NoError(t, err)
	shellAPI.EXPECT().
		ExecuteIn("fixtures/module1", "go", "list", "-m", "all").
		Return(shell.Output{
			Stdout: depsFixtureContent,
		}, nil)

	shellAPI.EXPECT().
		ExecuteIn("fixtures/module1", "go", "install", "github.com/golang/mock/mockgen@latest").
		Return(shell.Output{}, nil)

	err = m.InstallAllTools(shellAPI)

	require.NoError(t, err)
}

func Test_InstallAllTools_hasVersion(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := module{
		absPath: "fixtures/module1",
	}

	shellAPI := mocks.NewMockAPI(mockCtrl)

	depsFixtureContent, err := cliio.FileOf("fixtures/deps_stdout_mockgen.txt").StringContent()
	require.NoError(t, err)
	shellAPI.EXPECT().
		ExecuteIn("fixtures/module1", "go", "list", "-m", "all").
		Return(shell.Output{
			Stdout: depsFixtureContent,
		}, nil)

	shellAPI.EXPECT().
		ExecuteIn("fixtures/module1", "go", "install", "github.com/golang/mock/mockgen@v1.6.0").
		Return(shell.Output{}, nil)

	err = m.InstallAllTools(shellAPI)

	require.NoError(t, err)
}
