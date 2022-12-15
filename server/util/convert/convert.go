package convert

import "strconv"

func UintToString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

func StringToInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return res
}

func StringToUint(v string) uint {
	return uint(StringToInt(v))
}
