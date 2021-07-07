package overlay

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Edge ...
type Edge struct {
	points      []Point
	isClockwise bool
	nowStatus   int
}

// AreaDirection Returns area  <0 if direction is true, area > 0 else.
func (c *Edge) AreaDirection() float64 {
	var ring matrix.LineMatrix
	for _, v := range c.points {
		ring = append(ring, v.matrix)
	}
	return algorithm.AreaDirection(ring)
}

// SetClockwise with AreaDirection.
func (c *Edge) SetClockwise() {
	if c.AreaDirection() > 0 {
		c.isClockwise = true
	} else {
		c.isClockwise = false
	}
}

// Point overlay point.
type Point struct {
	matrix                                     matrix.Matrix
	isIntersectionPoint, isEntering, isChecked bool
}

// X Returns x  .
func (p *Point) X() float64 {
	return p.matrix[0]
}

// Y Returns y  .
func (p *Point) Y() float64 {
	return p.matrix[1]
}

// Sub Returns p - point  .
func (p *Point) Sub(point *Point) *Point {
	x := p.X() - point.X()
	y := p.Y() - point.Y()
	return &Point{matrix: matrix.Matrix{x, y}}
}

// Equal Returns p == point  .
func (p *Point) Equal(point *Point) bool {
	return p.X() == point.X() && p.Y() == point.Y()
}

// Line  straight line  .
type Line struct {
	start, end *Point
	isMain     bool
}
