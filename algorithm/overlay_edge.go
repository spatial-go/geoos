package algorithm

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// Edge ...
type Edge struct {
	Vertexs     []Vertex
	IsClockwise bool
	NowStatus   int
}

// AreaDirection Returns area  <0 if direction is true, area > 0 else.
func (e *Edge) AreaDirection() float64 {
	var ring matrix.LineMatrix
	for _, v := range e.Vertexs {
		ring = append(ring, v.Matrix)
	}
	return measure.AreaDirection(ring)
}

// SetClockwise with AreaDirection.
func (e *Edge) SetClockwise() {
	if e.AreaDirection() > 0 {
		e.IsClockwise = true
	} else {
		e.IsClockwise = false
	}
}

// Vertex overlay point.
type Vertex struct {
	matrix.Matrix
	IsIntersectionPoint, IsEntering, IsChecked bool
}

// X Returns x  .
func (v *Vertex) X() float64 {
	return v.Matrix[0]
}

// Y Returns y  .
func (v *Vertex) Y() float64 {
	return v.Matrix[1]
}

// Sub Returns p - point  .
func (v *Vertex) Sub(point *Vertex) *Vertex {
	x := v.X() - point.X()
	y := v.Y() - point.Y()
	return &Vertex{Matrix: matrix.Matrix{x, y}}
}

// Equal Returns p == point  .
func (v *Vertex) Equal(point *Vertex) bool {
	return v.X() == point.X() && v.Y() == point.Y()
}

// Line  straight line  .
type Line struct {
	Start, End *Vertex
	IsMain     bool
}
