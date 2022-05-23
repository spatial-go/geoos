// Package matrix Define spatial matrix base.
package matrix

import (
	"math"
)

// MultiPolygonMatrix is a four-dimensional matrix.
type MultiPolygonMatrix [][][][]float64

// Dimensions returns 3 because a 3 Dimensions matrix is a multi 2D object.
func (m MultiPolygonMatrix) Dimensions() int {
	return 3
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (m MultiPolygonMatrix) BoundaryDimensions() int {
	return 1
}

// Boundary returns the closure of the combinatorial boundary of this Polygon.
func (m MultiPolygonMatrix) Boundary() (Steric, error) {
	if m.IsEmpty() {
		return Collection{}, nil
	}
	rings := Collection{}
	for _, v := range m {
		if r, err := PolygonMatrix(v).Boundary(); err == nil {
			rings = append(rings, r.(Collection)...)
		}
	}
	return rings, nil
}

// Nums num of polygon matrix
func (m MultiPolygonMatrix) Nums() int {
	return len(m)
}

// IsEmpty returns true if the Matrix is empty.
func (m MultiPolygonMatrix) IsEmpty() bool {
	return len(m) == 0
}

// Bound returns a bound around the multi-polygon.
func (m MultiPolygonMatrix) Bound() Bound {
	if len(m) == 0 {
		return []Matrix{}
	}
	b := PolygonMatrix(m[0]).Bound()
	for i := 1; i < len(m); i++ {
		bound := PolygonMatrix(m[i]).Bound()
		b[0][0] = math.Min(b[0][0], bound[0][0])
		b[0][1] = math.Min(b[0][1], bound[0][1])
		b[1][0] = math.Min(b[1][0], bound[1][0])
		b[1][1] = math.Min(b[1][1], bound[1][1])
	}

	return b
}

// Equals returns  true if the two MultiPolygonMatrix are equal
func (m MultiPolygonMatrix) Equals(ms Steric) bool {
	if mm, ok := ms.(MultiPolygonMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (m == nil) {
			return false
		}

		if len(mm) != len(m) {
			return false
		}

		for i := range mm {
			if !PolygonMatrix(m[i]).Equals(PolygonMatrix(mm[i])) {
				return false
			}
		}
		return true
	}
	return false

}

// Proximity returns true if the Steric represents the Proximity Geometry or vector.
func (m MultiPolygonMatrix) Proximity(ms Steric) bool {
	if mm, ok := ms.(MultiPolygonMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (m == nil) {
			return false
		}

		if len(mm) != len(m) {
			return false
		}

		for i := range mm {
			if !PolygonMatrix(m[i]).Proximity(PolygonMatrix(mm[i])) {
				return false
			}
		}
		return true
	}
	return false
}

// EqualsExact returns  true if the two Matrix are equalexact
func (m MultiPolygonMatrix) EqualsExact(ms Steric, tolerance float64) bool {
	if mm, ok := ms.(MultiPolygonMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (m == nil) {
			return false
		}

		if len(mm) != len(m) {
			return false
		}

		for i := range mm {
			if !PolygonMatrix(m[i]).EqualsExact(PolygonMatrix(mm[i]), tolerance) {
				return false
			}
		}
	}
	return true
}

// Filter Performs an operation with the provided .
func (m MultiPolygonMatrix) Filter(f Filter) Steric {
	if f.IsChanged() {
		mPoly := m[:0]
		for _, v := range m {
			p := PolygonMatrix(v).Filter(f)
			mPoly = append(mPoly, p.(PolygonMatrix))
		}
		return mPoly
	}
	for _, v := range m {
		_ = PolygonMatrix(v).Filter(f)
	}
	return m
}
