package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Version(t *testing.T) {
	v := version{
		semverValue: "v1.18.1",
	}

	require.Equal(t, "1.18.1", v.Full())
	require.Equal(t, "1.18", v.MajorMinor())
}
