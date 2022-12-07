package convert

import "strconv"

func UintToString(n uint) string {
	return strconv.FormatUint(uint64(n), 10)
}

// func StringToInt64(n string) int64 {
// 	r, err := strconv.ParseInt(n, 10, 64)
// 	if err != nil {

// 	}
// 	return r
// }
