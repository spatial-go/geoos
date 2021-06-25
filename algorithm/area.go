package algorithm

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func AreaOfMultiPolygon(mp matrix.MatrixMultiPolygon) float64 {
	area := 0.0
	for _, polygon := range mp {
		area += AreaOfPolygon(polygon)
	}
	return area
}

func AreaOfPolygon(polygon matrix.MatrixPolygon) float64 {
	area := 0.0
	for i, ring := range polygon {
		if i == 0 {
			area += ofRing(ring)
		} else {
			area -= ofRing(ring)
		}
	}
	return area
}

// OfRing returns the area of a Ring geometry.
func ofRing(ring matrix.MatrixLine) float64 {
	rlen := len(ring)
	if rlen < 3 {
		return 0.0
	}
	sum := 0.0

	x0 := ring[0][0]
	for i := 1; i < rlen-1; i++ {
		x := ring[i][0] - x0
		y1 := ring[i+1][1]
		y2 := ring[i-1][1]
		sum += x * (y2 - y1)
	}
	return math.Abs(sum / 2.0)
}
