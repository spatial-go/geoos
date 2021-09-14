package overlay

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// UnaryUnion returns a Geometry containing the union.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func UnaryUnion(matrix4 matrix.Steric) matrix.Steric {
	if c, ok := matrix4.(matrix.Collection); ok {
		return UnaryUnionByHalf(c, 0, len(c))
	}
	return nil
}

// UnaryUnionByHalf returns Unions a section of a list using a recursive binary union on each half of the section.
func UnaryUnionByHalf(matrix4 matrix.Collection, start, end int) matrix.Steric {
	if matrix4 == nil {
		return nil
	}
	if end-start <= 1 {
		return Union(matrix4[start], nil)
	} else if end-start == 2 {
		return Union(matrix4[start], matrix4[start+1])
	} else {
		mid := (end + start) / 2
		g0 := UnaryUnionByHalf(matrix4, start, mid)
		g1 := UnaryUnionByHalf(matrix4, mid, end)
		return Union(g0, g1)
	}
}

// Union  Computes the Union of two geometries,either or both of which may be null.
func Union(m0, m1 matrix.Steric) (result matrix.Steric) {
	switch m := m0.(type) {
	case matrix.Matrix:
		over := &PointOverlay{Subject: m, Clipping: m1}
		result, _ = over.Union()
	case matrix.LineMatrix:
		over := &LineOverlay{PointOverlay: &PointOverlay{Subject: m0, Clipping: m1}}
		result, _ = over.Union()
	case matrix.PolygonMatrix:
		polyOver := &PolygonOverlay{PointOverlay: &PointOverlay{Subject: m0, Clipping: m1}}
		result, _ = polyOver.Union()
	case matrix.Collection:
		result = Union(UnaryUnion(m), m1)
	}
	return result
}

// UnionPolygon  Computes the UnionPolygon of two geometries,either or both of which may be null.
func UnionPolygon(m0, m1 matrix.PolygonMatrix) matrix.Steric {

	polyOver := &PolygonOverlay{PointOverlay: &PointOverlay{Subject: m0, Clipping: m1}}

	result, _ := polyOver.Union()
	return result
}

// UnionLine  Computes the Union of two geometries,either or both of which may be null.
func UnionLine(m0, m1 matrix.LineMatrix) matrix.Steric {
	result := matrix.Collection{}
	ils := IntersectLine(m0, m1)
	for _, il := range ils {
		result = append(result, matrix.LineMatrix{il.Ips[0].Matrix, il.Ips[1].Matrix})
	}
	if sd, err := SymDifference(m0, m1); err == nil {
		result = append(result, sd.(matrix.Collection)...)
	}
	return result
}
