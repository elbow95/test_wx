package util

import "strconv"

func Str2Int64(s string) int64 {
	id, _ := strconv.ParseInt(s, 0, 64)
	return id
}

func Int642Str(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StrSliceToInt64(ss []string) []int64 {
	ids := make([]int64, 0, len(ss))
	for _, s := range ss {
		ids = append(ids, Str2Int64(s))
	}
	return ids
}
