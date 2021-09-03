// Package filter Define matrix data filter function.
package filter

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Filter  An interface  which use the values of the coordinates in a  Geometry.
//  Coordinate filters can be used to implement centroid and
//  envelope computation, and many other functions.
type Filter interface {

	// FilterSteric  Performs an operation with the provided .
	FilterSteric(matrix matrix.Steric)

	// Filter  Performs an operation with the provided .
	Filter(matrix matrix.Matrix)

	// Matrixes ...
	Matrixes() []matrix.Matrix
}

// UniqueArrayFilter  A Filter that extracts a unique array.
type UniqueArrayFilter struct {
	matrixes []matrix.Matrix
}

// Matrixes  Returns the gathered Coordinates.
func (u *UniqueArrayFilter) Matrixes() []matrix.Matrix {
	return u.matrixes
}

// Filter Performs an operation with the provided .
func (u *UniqueArrayFilter) Filter(matrix matrix.Matrix) {
	u.add(matrix)
}

// FilterSteric Performs an operation with the provided .
func (u *UniqueArrayFilter) FilterSteric(matr matrix.Steric) {
	//set empty
	u.matrixes = u.matrixes[:0]
	switch m := matr.(type) {
	case matrix.Matrix:
		u.Filter(m)
	case matrix.LineMatrix:
		for _, v := range m {
			u.Filter(v)
		}
	case matrix.PolygonMatrix:
		for _, v := range m {
			u.FilterSteric(matrix.LineMatrix(v))
		}
	case matrix.Collection:
		for _, v := range m {
			u.FilterSteric(v)
		}
	}
}

func (u *UniqueArrayFilter) add(matrix matrix.Matrix) {
	hasMatrix := false
	for _, v := range u.matrixes {
		if v.Equals(matrix) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.matrixes = append(u.matrixes, matrix)
	}
}
