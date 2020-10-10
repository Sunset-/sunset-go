package casts

import "strconv"

func Float32ToStr(f float32, prec int) string {
	return strconv.FormatFloat(float64(f), 'f', prec, 64)
}

func Float64ToStr(f float32, prec int) string {
	return strconv.FormatFloat(float64(f), 'f', prec, 64)
}
