package casts

import (
	"math"
	"strconv"
)

func Float32ToStr(f float32, prec int) string {
	return strconv.FormatFloat(float64(f), 'f', prec, 64)
}

func Float64ToStr(f float32, prec int) string {
	return strconv.FormatFloat(float64(f), 'f', prec, 64)
}

func FloatToInt(f float64) int {
	return int(math.Floor(f))
}
