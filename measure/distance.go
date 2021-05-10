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

// MercatorDistance scale factor is changed along the meridians as a function of latitude
// https://gis.stackexchange.com/questions/110730/mercator-scale-factor-is-changed-along-the-meridians-as-a-function-of-latitude
// https://gis.stackexchange.com/questions/93332/calculating-distance-scale-factor-by-latitude-for-mercator
func MercatorDistance(d float64, lat float64) float64 {
	e := 0.006694379990141317
	lat = lat * math.Pi / 180
	factor := math.Sqrt(1-math.Pow(e, 2)*math.Pow(math.Sin(lat), 2)) * (1 / math.Cos(lat))
	d = d * factor
	return d
}
