package version

import (
	"regexp"
	"runtime"
)

func Current() string {
	versionStr := runtime.Version()
	regex := regexp.MustCompile("go(.+)")
	matches := regex.FindAllStringSubmatch(versionStr, -1)
	for _, match := range matches {
		if len(match) == 2 {
			return match[1]
		}
	}

	return ""
}
