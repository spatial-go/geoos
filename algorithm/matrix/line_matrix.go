// Package matrix Define spatial matrix base.
package matrix

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// LineMatrix is a two-dimensional matrix.
type LineMatrix [][]float64

// Dimensions returns 0 because a line matrix is a 0d object.
func (l LineMatrix) Dimensions() int {
	return 1
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (l LineMatrix) BoundaryDimensions() int {
	if l.IsClosed() {
		return calc.ImFalse
	}
	return 0
}

// Boundary returns the closure of the combinatorial boundary of this LineMatrix.
// The boundary of a lineal geometry is always a zero-dimensional geometry (which may be empty).
func (l LineMatrix) Boundary() (Steric, error) {
	if l.IsClosed() {
		return nil, spaceerr.ErrBoundBeNil
	}
	return Collection{Matrix(l[0]), Matrix(l[len(l)-1])}, nil
}

// IsClosed Returns TRUE if the line's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (l LineMatrix) IsClosed() bool {
	if l.IsEmpty() {
		return false
	}
	return Matrix(l[0]).Equals(Matrix(l[len(l)-1]))
}

// Nums num of line matrix
func (l LineMatrix) Nums() int {
	return 1
}

// IsEmpty returns true if the Matrix is empty.
func (l LineMatrix) IsEmpty() bool {
	return len(l) == 0
}

// Bound returns a rect around the line string. Uses rectangular coordinates.
func (l LineMatrix) Bound() Bound {
	if len(l) == 0 {
		return []Matrix{}
	}

	b := []Matrix{{math.MaxFloat64, math.MaxFloat64}, {0, 0}}
	for _, p := range l {
		b[0][0] = math.Min(b[0][0], p[0])
		b[0][1] = math.Min(b[0][1], p[1])
		b[1][0] = math.Max(b[1][0], p[0])
		b[1][1] = math.Max(b[1][1], p[1])
	}

	return b
}

// ToLineArray returns the LineArray
func (l LineMatrix) ToLineArray() (lines []*LineSegment) {
	return LineArray(l)
}

// Equals returns  true if the two LineMatrix are equal
func (l LineMatrix) Equals(ms Steric) bool {
	if mm, ok := ms.(LineMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (l == nil) {
			return false
		}

		if len(mm) != len(l) {
			return false
		}

		for i := range mm {
			if !Matrix(l[i]).Equals(Matrix(mm[i])) {
				return false
			}
		}
		return true
	}
	return false
}

// Proximity returns true if the Steric represents the Proximity Geometry or vector.
func (l LineMatrix) Proximity(ms Steric) bool {
	if mm, ok := ms.(LineMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (l == nil) {
			return false
		}

		if len(mm) != len(l) {
			return false
		}

		for i := range mm {
			havePoint := false
			for _, v := range l {
				if Matrix(v).Proximity(Matrix(mm[i])) {
					havePoint = true
					break
				}
			}
			if !havePoint {
				for i := range mm {
					havePointReverse := false
					for _, v := range l.Reverse() {
						if Matrix(v).Proximity(Matrix(mm[i])) {
							havePointReverse = true
							break
						}
					}
					if !havePointReverse {
						return false
					}
				}
				return false
			}
		}
		return true
	}
	return false
}

// EqualsExact returns  true if the two Matrix are equalexact
func (l LineMatrix) EqualsExact(ms Steric, tolerance float64) bool {
	if mm, ok := ms.(LineMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (l == nil) {
			return false
		}

		if len(mm) != len(l) {
			return false
		}

		for i := range mm {
			if !Matrix(l[i]).EqualsExact(Matrix(mm[i]), tolerance) {
				return false
			}
		}
		return true
	}
	return false
}

// Filter Performs an operation with the provided .
func (l LineMatrix) Filter(f Filter) Steric {
	f.FilterMatrixes(TransMatrixes(l))
	if f.IsChanged() {
		l = l[:0]
		for _, v := range f.Matrixes() {
			l = append(l, v)
		}
		return l
	}
	return l
}

// Reverse  this LineMatrix.
func (l LineMatrix) Reverse() LineMatrix {
	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return l
}
