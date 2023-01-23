package valid

func ReviewStatus(role int) bool {
	roles := map[int]string{
		0:    "AUDIT_APPROVED",
		2100: "WRONG_VIDEO_INFO",
		2200: "WRONG_VIDEO_CONTENT",
	}

	if _, ok := roles[role]; !ok {
		return false
	}

	return true
}
