package version

import (
	"regexp"
	"runtime"
)

func Current() Version {
	versionStr := runtime.Version()
	regex := regexp.MustCompile("go(.+)")
	matches := regex.FindAllStringSubmatch(versionStr, -1)
	for _, match := range matches {
		if len(match) == 2 {
			return version{
				// Prefix with 'v' to make it a proper semver.
				semverValue: "v" + match[1],
			}
		}
	}

	return nil
}
