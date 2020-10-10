package casts

import "strconv"

func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func StrToInt32(str string) int32 {
	return int32(StrToInt(str))
}

func StrToInt64(str string) int64 {
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

func StrToFloat32(str string) float32 {
	f, _ := strconv.ParseFloat(str, 32)
	return float32(f)
}

func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
