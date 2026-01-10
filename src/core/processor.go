package core

func ScoreWidget(name string) int {
	if name == "" {
		return 0
	}
	return len(name) * 2
}

func NormalizeLabel(label string) string {
	if label == "" {
		return ""
	}
	return label
}
