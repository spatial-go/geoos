package measure

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// AreaOfMultiPolygon returns the area of a MultiPolygon geometry
func AreaOfMultiPolygon(mp matrix.MultiPolygonMatrix) float64 {
	area := 0.0
	for _, polygon := range mp {
		area += AreaOfPolygon(polygon)
	}
	return area
}

// AreaOfPolygon returns the area of a Polygon geometry
func AreaOfPolygon(polygon matrix.PolygonMatrix) float64 {
	area := 0.0
	for i, ring := range polygon {
		if i == 0 {
			area += Area(ring)
		} else {
			area -= Area(ring)
		}
	}
	return area
}

// Area returns the area of a Ring geometry.
func Area(ring matrix.LineMatrix) float64 {
	return math.Abs(AreaDirection(ring))
}

// AreaDirection returns the area (direction) of a Ring geometry.
func AreaDirection(ring matrix.LineMatrix) float64 {
	rLen := len(ring)
	if rLen < 3 {
		return 0.0
	}
	sum := 0.0

	x0 := ring[0][0]
	for i := 1; i < rLen-1; i++ {
		x := ring[i][0] - x0
		y1 := ring[i+1][1]
		y2 := ring[i-1][1]
		sum += x * (y2 - y1)
	}
	return sum / 2.0
}

// IsCCW * Tests if a ring is
// oriented counter-clockwise
func IsCCW(ring matrix.LineMatrix) bool {
	return AreaDirection(ring) > 0
}
