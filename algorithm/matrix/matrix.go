package matrix

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/algoerr"
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

	// IsEmpty returns true if the Matrix is empty.
	IsEmpty() bool

	Bound() []Matrix
}

// LineSegment is line.
type LineSegment struct {
	P0, P1 Matrix
}

// Matrix is a one-dimensional matrix.
type Matrix []float64

// LineMatrix is a two-dimensional matrix.
type LineMatrix [][]float64

// PolygonMatrix is a three-dimensional matrix.
type PolygonMatrix [][][]float64

// MultiPolygonMatrix is a four-dimensional matrix.
type MultiPolygonMatrix [][][][]float64

// Dimensions returns 0 because a matrix is a 0d object.
func (m Matrix) Dimensions() int {
	return 0
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (m Matrix) BoundaryDimensions() int {
	return calc.FALSE
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

// Dimensions returns 0 because a line matrix is a 0d object.
func (l LineMatrix) Dimensions() int {
	return 1
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (l LineMatrix) BoundaryDimensions() int {
	if l.IsClosed() {
		return calc.FALSE
	}
	return 0
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
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
	return l == nil || len(l) == 0
}

// Bound returns a rect around the line string. Uses rectangular coordinates.
func (l LineMatrix) Bound() []Matrix {
	if len(l) == 0 {
		return []Matrix{}
	}

	b := []Matrix{{0, 0}, {0, 0}}
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

// Dimensions returns 0 because a polygon matrix is a 0d object.
func (p PolygonMatrix) Dimensions() int {
	return 2
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (p PolygonMatrix) BoundaryDimensions() int {
	return 1
}

// Nums num of polygon matrix
func (p PolygonMatrix) Nums() int {
	return 1
}

// IsEmpty returns true if the Matrix is empty.
func (p PolygonMatrix) IsEmpty() bool {
	return p == nil || len(p) == 0
}

// Bound returns a bound around the polygon.
func (p PolygonMatrix) Bound() []Matrix {
	if len(p) == 0 {
		return []Matrix{}
	}
	return LineMatrix(p[0]).Bound()
}

// Dimensions returns 0 because a 3 Dimensions matrix is a 0d object.
func (m MultiPolygonMatrix) Dimensions() int {
	return 3
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (m MultiPolygonMatrix) BoundaryDimensions() int {
	return 1
}

// Nums num of polygon matrix
func (m MultiPolygonMatrix) Nums() int {
	return len(m)
}

// IsEmpty returns true if the Matrix is empty.
func (m MultiPolygonMatrix) IsEmpty() bool {
	return m == nil || len(m) == 0
}

// Bound returns a bound around the multi-polygon.
func (m MultiPolygonMatrix) Bound() []Matrix {
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

// A Collection is a collection of Sterics that is also a Steric.
type Collection []Steric

// Dimensions returns the max of the dimensions of the collection.
func (c Collection) Dimensions() int {
	max := -1
	for _, g := range c {
		if d := g.Dimensions(); d > max {
			max = d
		}
	}
	return max
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (c Collection) BoundaryDimensions() int {
	dimension := calc.FALSE
	for _, g := range c {
		if g.BoundaryDimensions() > dimension {
			dimension = g.BoundaryDimensions()
		}
	}
	return dimension
}

// Nums ...
func (c Collection) Nums() int {
	return len(c)
}

// IsEmpty returns true if the Matrix is empty.
func (c Collection) IsEmpty() bool {
	return c == nil || len(c) == 0
}

// Bound returns a bound around the Collection.
func (c Collection) Bound() []Matrix {
	if len(c) == 0 {
		return []Matrix{}
	}
	b := c[0].Bound()
	for i := 1; i < len(c); i++ {
		bound := c[1].Bound()
		b[0][0] = math.Min(b[0][0], bound[0][0])
		b[0][1] = math.Min(b[0][1], bound[0][1])
		b[1][0] = math.Min(b[1][0], bound[1][0])
		b[1][1] = math.Min(b[1][1], bound[1][1])
	}

	return b
}

// Equals returns  true if the two MultiPolygonMatrix are equal
func (c Collection) Equals(ms Steric) bool {
	if mm, ok := ms.(Collection); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (c == nil) {
			return false
		}

		if len(mm) != len(c) {
			return false
		}

		for i := range mm {
			if !c[i].Equals(mm[i]) {
				return false
			}
		}
		return true
	}
	return false
}

// Compare returns 0 if m1==m2,1 if positive ,-1 else
// Compares Coordinate for order.
func (m Matrix) Compare(m1 Matrix) (int, error) {
	// If one is nil, the other must also be nil.
	if (m1 == nil) != (m == nil) {
		return -2, algoerr.ErrNilSteric
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

// Equals returns  true if the two PolygonMatrix are equal
func (p PolygonMatrix) Equals(ms Steric) bool {
	if mm, ok := ms.(PolygonMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (p == nil) {
			return false
		}

		if len(mm) != len(p) {
			return false
		}

		for i := range mm {
			if !LineMatrix(p[i]).Equals(LineMatrix(mm[i])) {
				return false
			}
		}
		return true
	}
	return false
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
	}
	return true
}

// TransMatrixs trans steric to array matrixs.
func TransMatrixs(inputGeom Steric) []Matrix {

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
			p := TransMatrixs(v)
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
