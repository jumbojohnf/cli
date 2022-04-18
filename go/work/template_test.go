package work

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_topLevelModuleDirNames(t *testing.T) {
	results, err := goWorkTemplate{}.topLevelModuleDirNames("fixtures")

	require.NoError(t, err)
	require.Equal(t, []string{"mod_1", "mod_2"}, results)
}

func Test_Export(t *testing.T) {
	content, err := NewTemplate().Render("fixtures")
	require.NoError(t, err)
	fmt.Println(content)
	require.Contains(t, content, "go ")
	require.Contains(t, content, `use (
  ./mod_1
  ./mod_2
)`)
}
