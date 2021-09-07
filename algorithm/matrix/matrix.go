// Package matrix Define spatial matrix base.
package matrix

import (
	"math"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/calc"
)

// Steric is the interface implemented by other Steric objects
type Steric interface {
	// e.g. 0d, 1d, 2d
	Dimensions() int

	BoundaryDimensions() int

	// Num of geometries
	Nums() int

	// Equals returns true if the Geometry represents the same Geometry or vector.
	Equals(s Steric) bool

	// EqualsExact Returns true if the two Geometries are exactly equal,
	// up to a specified distance tolerance.
	// Two Geometries are exactly equal within a distance tolerance
	EqualsExact(g Steric, tolerance float64) bool

	// IsEmpty returns true if the Matrix is empty.
	IsEmpty() bool

	Bound() []Matrix

	// Filter Performs an operation with the provided .
	Filter(f Filter) Steric
}

// Matrix is a one-dimensional matrix.
type Matrix []float64

// Dimensions returns 0 because a matrix is a 0d object.
func (m Matrix) Dimensions() int {
	return 0
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (m Matrix) BoundaryDimensions() int {
	return calc.ImFalse
}

// Nums num of matrix
func (m Matrix) Nums() int {
	return 1
}

// IsEmpty returns true if the Matrix is empty.
func (m Matrix) IsEmpty() bool {
	return m == nil || len(m) == 0
}

// Bound returns a single point bound of the point.
func (m Matrix) Bound() []Matrix {
	return []Matrix{m, m}
}

// Compare returns 0 if m1==m2,1 if positive ,-1 else
// Compares Coordinate for order.
func (m Matrix) Compare(m1 Matrix) (int, error) {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m == nil) {
		return -2, algorithm.ErrNilSteric
	}

	if m1[0] < m[0] {
		return -1, nil
	}
	if m1[0] > m[0] {
		return 1, nil
	}
	if m1[1] < m[1] {
		return -1, nil
	}
	if m1[1] > m[1] {
		return 1, nil
	}
	return 0, nil
}

// Equals returns  true if the two Matrix are equal
func (m Matrix) Equals(ms Steric) bool {
	if mm, ok := ms.(Matrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (m == nil) {
			return false
		}

		if len(mm) != len(m) {
			return false
		}

		for i := range mm {
			if mm[i] != m[i] {
				return false
			}
		}
		return true
	}
	return false
}

// EqualsExact returns  true if the two Matrix are equalexact
func (m Matrix) EqualsExact(ms Steric, tolerance float64) bool {
	if mm, ok := ms.(Matrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (m == nil) {
			return false
		}

		if len(mm) != len(m) {
			return false
		}

		if tolerance == 0 {
			return m.Equals(ms)
		}

		return math.Sqrt((m[0]-mm[0])*(m[0]-mm[0])+(m[1]-mm[1])*(m[1]-mm[1])) <= tolerance
	}
	return false
}

// Filter Performs an operation with the provided .
func (m Matrix) Filter(f Filter) Steric {
	return m
}

// TransMatrixes trans steric to array matrixes.
func TransMatrixes(inputGeom Steric) []Matrix {

	switch m := inputGeom.(type) {
	case Matrix:
		return []Matrix{m}
	case LineMatrix:
		tm := []Matrix{}
		for _, v := range m {
			tm = append(tm, v)
		}
		return tm
	case PolygonMatrix:
		tm := []Matrix{}
		for _, v := range m {
			for _, p := range v {
				tm = append(tm, p)
			}
		}
		return tm
	case Collection:
		tm := []Matrix{}
		for _, v := range m {
			p := TransMatrixes(v)
			tm = append(tm, p...)
		}
		return tm
	default:
		return []Matrix{}
	}
}

// LineArray returns the LineArray
func LineArray(l LineMatrix) (lines []*LineSegment) {
	for i := 0; i < len(l)-1; i++ {
		lines = append(lines, &LineSegment{Matrix(l[i]), Matrix(l[i+1])})
	}
	return
}

// compile time checks
var (
	_ Steric = Matrix{}
	_ Steric = LineMatrix{}
	_ Steric = PolygonMatrix{}
	_ Steric = MultiPolygonMatrix{}

	_ Steric = Collection{}
)
