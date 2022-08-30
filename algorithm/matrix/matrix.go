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

	// Boundary returns the closure of the combinatorial boundary of this Steric.
	// The boundary of a lineal geometry is always a zero-dimensional geometry (which may be empty).
	Boundary() (Steric, error)

	// Num of geometries
	Nums() int

	// Equals returns true if the Steric represents the same Geometry or vector.
	Equals(s Steric) bool

	// Proximity returns true if the Steric represents the Proximity Geometry or vector.
	Proximity(s Steric) bool

	// EqualsExact Returns true if the two Steric are exactly equal,
	// up to a specified distance tolerance.
	// Two Steric are exactly equal within a distance tolerance
	EqualsExact(g Steric, tolerance float64) bool

	// IsEmpty returns true if the Matrix is empty.
	IsEmpty() bool

	Bound() Bound

	// Filter Performs an operation with the provided .
	Filter(f Filter) Steric
}

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
// MultiPoint{p1, p2}.Bound()
type Bound []Matrix

// Equals checks if the Bound represents the same Geometry or vector.
func (b Bound) Equals(g Bound) bool {
	return b[0].Equals(g[0]) && b[1].Equals(g[1])
}

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool {
	if b == nil || len(b) < 2 || b[0] == nil || b[1] == nil {
		return true
	}
	return b[0][0] > b[1][0] || b[0][1] > b[1][1]
}

// ToRing converts the bound into a loop defined
// by the boundary of the box.
func (b Bound) ToRing() LineMatrix {
	return LineMatrix{
		b[0],
		{b[1][0], b[0][1]},
		b[1],
		{b[0][0], b[1][1]},
		b[0],
	}
}

// ToPolygon converts the bound into a Polygon object.
func (b Bound) ToPolygon() PolygonMatrix {
	return PolygonMatrix{b.ToRing()}
}

// Contains determines if the point is within the bound.
// Points on the boundary are considered within.
func (b Bound) Contains(m Matrix) bool {
	if m[1] < b[0][1] || b[1][1] < m[1] {
		return false
	}

	if m[0] < b[0][0] || b[1][0] < m[0] {
		return false
	}

	return true
}

// ContainsBound determines if the bound is within the bound.
func (b Bound) ContainsBound(bound Bound) bool {
	if b.IsEmpty() || bound.IsEmpty() {
		return false
	}
	return bound[0][0] >= b[0][0] &&
		bound[1][0] <= b[1][0] &&
		bound[0][1] >= b[0][1] &&
		bound[1][1] <= b[1][1]
}

// IntersectsBound Tests if the region defined by other
// intersects the region of this Envelope.
func (b Bound) IntersectsBound(other Bound) bool {
	if b.IsEmpty() || other.IsEmpty() {
		return false
	}
	return !(other[0][0] > b[1][0] ||
		other[1][0] < b[0][0] ||
		other[0][1] > b[1][1] ||
		other[1][1] < b[0][1])
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

// Boundary returns the closure of the combinatorial boundary of this Matrix.
func (m Matrix) Boundary() (Steric, error) {
	return nil, algorithm.ErrBoundBeNil
}

// Nums num of matrix
func (m Matrix) Nums() int {
	return 1
}

// IsEmpty returns true if the Matrix is empty.
func (m Matrix) IsEmpty() bool {
	return len(m) == 0
}

// Bound returns a single point bound of the point.
func (m Matrix) Bound() Bound {
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

// Proximity returns true if the Steric represents the Proximity Geometry or vector.
func (m Matrix) Proximity(ms Steric) bool {
	return m.EqualsExact(ms, calc.DefaultTolerance)
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
			return m.Equals(mm)
		}

		if m[0]-mm[0] == 0 && m[1]-mm[1] == 0 {
			return true
		}

		if math.Abs(m[0]-mm[0]) > tolerance || math.Abs(m[1]-mm[1]) > tolerance {
			return false
		}
		return math.Hypot((m[0]-mm[0]), (m[1]-mm[1])) <= tolerance
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

// compile time checks
var (
	_ Steric = Matrix{}
	_ Steric = LineMatrix{}
	_ Steric = PolygonMatrix{}
	_ Steric = MultiPolygonMatrix{}

	_ Steric = Collection{}
)
