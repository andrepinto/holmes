package version

import (
	"github.com/blang/semver"
)

func CheckVersion(version string, baseVersion string) bool{
	v, _ := semver.Parse(version)
	valid, _ := semver.ParseRange(baseVersion)
	result := valid(v)
	return result
}
