package util

import "strconv"

func StringToIntUnsafe(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func StringToFloat64Unsafe(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)
	return res
}
