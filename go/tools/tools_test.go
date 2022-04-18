package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_goToolsModuleNamesIn(t *testing.T) {
	results, err := goToolsIn("fixtures")

	require.NoError(t, err)
	require.Equal(t, []goTools{
		{
			modulePath: "fixtures/module1",
			toolNames: []string{
				"github.com/module1/tool1",
				"github.com/module1/tool2",
			},
		},
		{
			modulePath: "fixtures/module2",
			toolNames: []string{
				"github.com/module2/tool1",
				"github.com/module2/tool2",
			},
		}}, results)
}
