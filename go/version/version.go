package version

import (
	"strings"

	"golang.org/x/mod/semver"
)

type Version interface {
	Full() string
	MajorMinor() string
}

func (v version) Full() string {
	return strings.TrimPrefix(v.semverValue, "v")
}

func (v version) MajorMinor() string {
	return strings.TrimPrefix(semver.MajorMinor(v.semverValue), "v")
}

type version struct {
	semverValue string
}
