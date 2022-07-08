package angle

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Quadrant four const
const (
	NE = iota + 1
	NW
	SW
	SE
)

// QuadrantFloat Returns the quadrant of a directed line segment (specified as x and y displacements, which cannot both be 0).
func QuadrantFloat(dx, dy float64) (int, error) {
	if dx == 0.0 && dy == 0.0 {
		return 0, algorithm.ErrQuadrant
	}
	if dx >= 0.0 {
		if dy >= 0.0 {
			return NE, nil
		}

		return SE, nil
	}
	if dy >= 0.0 {
		return NW, nil
	}
	return SW, nil
}

// Quadrant Returns the quadrant of a directed line segment from p0 to p1.
func Quadrant(p0, p1 matrix.Matrix) (int, error) {
	if p1[0] == p0[0] && p1[1] == p0[1] {
		return 0, algorithm.ErrQuadrant
	}
	return QuadrantFloat(p1[0]-p0[0], p1[1]-p0[1])
}
