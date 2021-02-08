package measure

import (
	"math"

	"github.com/spatial-go/geoos"
)

// Distance Calculate distance, return unit: meter
func Distance(fromPoint, toPoint *geoos.Point) float64 {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat0 := fromPoint[1] * rad
	lng0 := fromPoint[0] * rad
	lat1 := toPoint[1] * rad
	lng1 := toPoint[0] * rad
	theta := lng1 - lng0
	dist := math.Acos(math.Sin(lat0)*math.Sin(lat1) + math.Cos(lat0)*math.Cos(lat1)*math.Cos(theta))
	return dist * radius
}
