package api

func HandleHealth() string {
	return "ok"
}

func HandleWidget(id string) string {
	return "widget:" + id
}

func HandleSearch(term string) string {
	if term == "" {
		return "empty"
	}
	return "search:" + term
}
