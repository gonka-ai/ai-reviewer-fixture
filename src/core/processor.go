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
	return "normalized:" + label
}

func NormalizeSubnetLabel(label string) string {
	if label == "" {
		return ""
	}
	return "subnet:" + label
}
