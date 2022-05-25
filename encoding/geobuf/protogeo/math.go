package protogeo

import (
	"math"
)

var (
	// MaxPrecision ...
	MaxPrecision = uint(math.Pow10(9))
)

// GetPrecision ...
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

// IntWithPrecision ///
func IntWithPrecision(point float64, precision uint) int64 {
	return int64(math.Round(point * float64(precision)))
}

// FloatWithPrecision ...
func FloatWithPrecision(point int64, precision uint32) float64 {
	return float64(point) / float64(precision)
}

// EncodePrecision ...
func EncodePrecision(precision uint) uint32 {
	return uint32(math.Ceil(math.Log(float64(precision)) / math.Ln10))
}

// DecodePrecision ...
func DecodePrecision(precision uint32) float64 {
	return math.Pow10(int(precision))
}
