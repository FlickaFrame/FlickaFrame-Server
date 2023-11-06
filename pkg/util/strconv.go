package util

import "strconv"

func MustString2Int64(x string) int64 {
	if len(x) == 0 {
		return 0
	}
	ret, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}
