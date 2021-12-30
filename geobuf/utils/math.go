package utils

import (
	"math"
)

var (
	MaxPrecision = uint(math.Pow10(9))
)

func GetPrecision(point float64) uint {
	var e uint = 1
	for {
		base := math.Round(float64(point * float64(e)))
		if (base/float64(e)) != point && e < MaxPrecision {
			e = e * 10
		} else {
			break
		}
	}
	return e
}

func IntWithPrecision(point float64, precision uint) int64 {
	return int64(math.Round(point * float64(precision)))
}

func FloatWithPrecision(point int64, precision uint32) float64 {
	return float64(point) / float64(precision)
}

func EncodePrecision(precision uint) uint32 {
	return uint32(math.Ceil(math.Log(float64(precision)) / math.Ln10))
}

func DecodePrecision(precision uint32) float64 {
	return math.Pow10(int(precision))
}
