// Package coordtransform is for transform coord.
package coordtransform

import "math"

// LLToMercator ...
func LLToMercator(lng, lat float64) (x, y float64) {
	x = lng * 20037508.34 / 180
	y = math.Log(math.Tan((90+lat)*math.Pi/360)) / (math.Pi / 180)
	y = y * 20037508.34 / 180
	return x, y
}

// MercatorToLL ...
func MercatorToLL(x, y float64) (lng, lat float64) {
	lng = x / 20037508.34 * 180
	lat = y / 20037508.34 * 180
	lat = 180 / math.Pi * (2*math.Atan(math.Exp(lat*math.Pi/180)) - math.Pi/2)
	return lng, lat
}
