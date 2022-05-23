// Package matrix Define spatial matrix base.
package matrix

// PolygonMatrix is a three-dimensional matrix.
type PolygonMatrix [][][]float64

// Dimensions returns 2 because a polygon matrix is a 2d object.
func (p PolygonMatrix) Dimensions() int {
	return 2
}

// BoundaryDimensions Compute the IM entry for the intersection of the boundary
// of a geometry with the Exterior.
func (p PolygonMatrix) BoundaryDimensions() int {
	return 1
}

// Boundary returns the closure of the combinatorial boundary of this Polygon.
func (p PolygonMatrix) Boundary() (Steric, error) {
	if p.IsEmpty() {
		return Collection{}, nil
	}
	if len(p) <= 1 {
		return Collection{LineMatrix(p[0])}, nil
	}
	rings := Collection{}
	for _, v := range p {
		rings = append(rings, LineMatrix(v))
	}
	return rings, nil
}

// Nums num of polygon matrix
func (p PolygonMatrix) Nums() int {
	return 1
}

// IsEmpty returns true if the Matrix is empty.
func (p PolygonMatrix) IsEmpty() bool {
	return len(p) == 0
}

// Bound returns a bound around the polygon.
func (p PolygonMatrix) Bound() Bound {
	if len(p) == 0 {
		return []Matrix{}
	}
	return LineMatrix(p[0]).Bound()
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

// Proximity returns true if the Steric represents the Proximity Geometry or vector.
func (p PolygonMatrix) Proximity(ms Steric) bool {
	if mm, ok := ms.(PolygonMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (p == nil) {
			return false
		}

		if len(mm) != len(p) {
			return false
		}

		for i := range mm {
			if !LineMatrix(p[i]).Proximity(LineMatrix(mm[i])) {
				return false
			}
		}
		return true
	}
	return false
}

// EqualsExact returns  true if the two Matrix are equalexact
func (p PolygonMatrix) EqualsExact(ms Steric, tolerance float64) bool {
	if mm, ok := ms.(PolygonMatrix); ok {
		// If one is nil, the other must also be nil.
		if (mm == nil) != (p == nil) {
			return false
		}

		if len(mm) != len(p) {
			return false
		}

		for i := range mm {
			if !LineMatrix(p[i]).EqualsExact(LineMatrix(mm[i]), tolerance) {
				return false
			}
		}
		return true
	}
	return false
}

// Filter Performs an operation with the provided .
func (p PolygonMatrix) Filter(f Filter) Steric {
	if f.IsChanged() {
		poly := PolygonMatrix{}
		for _, v := range p {
			r := LineMatrix(v).Filter(f).(LineMatrix)
			if !Matrix(r[len(r)-1]).Equals(Matrix(r[0])) {
				r = append(r, r[0])
			}
			poly = append(poly, r)
		}
		return poly
	}
	for _, v := range p {
		_ = LineMatrix(v).Filter(f)
	}
	return p
}

// IsRectangle returns true if  the polygon is rectangle.
func (p PolygonMatrix) IsRectangle() bool {

	if p.IsEmpty() || len(p) > 1 {
		return false
	}
	if len(p[0]) != 5 {
		return false
	}
	// check vertices have correct values
	for i := 0; i < 5; i++ {
		x := p[0][i][0]
		if !(x == p.Bound()[0][0] || x == p.Bound()[1][1]) {
			return false
		}
		y := p[0][i][1]
		if !(y == p.Bound()[0][1] || y == p.Bound()[1][1]) {
			return false
		}
	}

	// check vertices are in right order
	for i := 0; i < 4; i++ {
		x0 := p[0][i][0]
		y0 := p[0][i][1]
		x1 := p[0][i+1][0]
		y1 := p[0][i+1][1]
		xChanged := x0 != x1
		yChanged := y0 != y1
		if xChanged == yChanged {
			return false
		}
	}
	return true
}
