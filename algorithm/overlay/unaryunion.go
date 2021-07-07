package overlay

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// UnaryUnion returns a Geometry containing the union.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func UnaryUnion(matrix4 matrix.MultiPolygonMatrix) matrix.PolygonMatrix {
	return UnaryUnionByHalf(matrix4, 0, len(matrix4))
}

// UnaryUnionByHalf returns Unions a section of a list using a recursive binary union on each half of the section.
func UnaryUnionByHalf(matrix4 matrix.MultiPolygonMatrix, start, end int) matrix.PolygonMatrix {
	if matrix4 == nil {
		return nil
	}
	if end-start <= 1 {
		return union(matrix4[start], nil)
	} else if end-start == 2 {
		return union(matrix4[start], matrix4[start+1])
	} else {
		mid := (end + start) / 2
		g0 := UnaryUnionByHalf(matrix4, start, mid)
		g1 := UnaryUnionByHalf(matrix4, mid, end)
		return union(g0, g1)
	}
}

// Computes the union of two geometries,either or both of which may be null.
func union(m0, m1 matrix.PolygonMatrix) matrix.PolygonMatrix {

	if m0 == nil && m1 == nil {
		return nil
	}
	if m0 == nil {
		return m1
	}

	if m1 == nil {
		return m0
	}

	return unionActual(m0, m1)
}

// unionActual the actual unioning of two polygonal geometries.
func unionActual(m0, m1 matrix.PolygonMatrix) matrix.PolygonMatrix {

	subject := &Polygon{}
	for _, v2 := range m0 {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				subject.AddPoint(&Point{matrix: matrix.Matrix(v1)})
			}
		}
		subject.CloseRing()
		subject.rank = MAIN
	}
	clipping := &Polygon{}
	for _, v2 := range m1 {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				clipping.AddPoint(&Point{matrix: matrix.Matrix(v1)})
			}
		}
		clipping.CloseRing()
		clipping.rank = CUT
	}
	poly := Weiler(subject, clipping, Merge)
	var result matrix.PolygonMatrix
	for _, v2 := range poly.rings {
		var edge matrix.LineMatrix
		for _, v1 := range v2.points {
			edge = append(edge, v1.matrix)
		}
		if !matrix.Equal(edge[len(edge)-1], edge[0]) {
			edge = append(edge, edge[0])
		}
		result = append(result, edge)
	}

	return result
}
