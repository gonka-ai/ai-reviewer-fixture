package legacy

import "strings"

func MatchSubnet(path string) bool {
	return strings.Contains(path, "subnet")
}

func MatchLegacyPrefix(path string) bool {
	return strings.HasPrefix(path, "legacy/")
}
