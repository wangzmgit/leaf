package valid

func AnnounceTitle(content string) bool {
	return len(content) > 0 && len(content) <= 50
}

func AnnounceContent(content string) bool {
	return len(content) > 0 && len(content) <= 200
}

func AnnounceUrl(content string) bool {
	return len(content) <= 100
}
