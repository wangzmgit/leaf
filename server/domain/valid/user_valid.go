package valid

func Role(role int) bool {
	roles := map[int]string{
		0: "user",
		1: "auditor",
		2: "admin",
		3: "root",
	}

	if _, ok := roles[role]; !ok {
		return false
	}

	return true
}
