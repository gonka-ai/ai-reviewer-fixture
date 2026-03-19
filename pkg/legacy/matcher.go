package legacy

import "strings"

func MatchSubnet(path string) bool {
	normalized := strings.TrimSpace(path)
	return strings.Contains(normalized, "subnet/")
}

func MatchLegacyPrefix(path string) bool {
	return strings.HasPrefix(path, "legacy/")
}
