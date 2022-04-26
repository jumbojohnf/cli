package module

import (
	"testing"

	"github.com/funcgql/cli/cliio"
	"github.com/funcgql/cli/shell"
	"github.com/funcgql/cli/shell/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_Dependencies(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	fixtureContent, err := cliio.FileOf("fixtures/deps_stdout.txt").StringContent()
	require.NoError(t, err)

	shellAPI := mocks.NewMockAPI(mockCtrl)
	shellAPI.EXPECT().
		ExecuteIn("MODULE_PATH", "go", "list", "-m", "all").
		Return(shell.Output{
			Stdout: fixtureContent,
		}, nil)

	m := module{
		absPath: "MODULE_PATH",
	}

	results, err := m.Dependencies(shellAPI)

	require.NoError(t, err)
	require.Equal(t, map[string]Dependency{
		"github.com/funcgql/cli": {
			ImportPath: "github.com/funcgql/cli",
		},
		"github.com/cpuguy83/go-md2man/v2": {
			ImportPath: "github.com/cpuguy83/go-md2man/v2",
			Version:    "v2.0.1",
		},
		"github.com/davecgh/go-spew": {
			ImportPath: "github.com/davecgh/go-spew",
			Version:    "v1.1.0",
		},
		"github.com/golang/mock": {
			ImportPath: "github.com/golang/mock",
			Version:    "v1.6.0",
		},
	}, results)
}
