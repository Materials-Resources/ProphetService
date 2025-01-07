package helpers

import (
	"strconv"
)

func GetValueOrDefault[T any](ptr *T, defaultVal T) T {
	if ptr == nil {
		return defaultVal
	}
	return *ptr
}

func StringToInt32(s string, strict bool) int32 {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil && strict {
		panic("invalid int32 format")
	}
	return int32(val)
}
