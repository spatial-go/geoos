package calc

import (
	"fmt"
	"strconv"
)

// DecimalFloat10 Returns float64 , DecimalPlaces  10.
func DecimalFloat10(x float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf(DecimalPlaces, x), 64)
	return value
}
