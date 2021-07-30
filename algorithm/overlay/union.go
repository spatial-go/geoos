package overlay

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/calc"
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
func Union(m0, m1 matrix.PolygonMatrix) matrix.PolygonMatrix {

	if m0 == nil && m1 == nil {
		return nil
	}
	if m0 == nil {
		return m1
	}

	if m1 == nil {
		return m0
	}

	return unionActual(m0, m1, Merge)
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

// unionActual the actual unioning of two polygonal geometries.
func unionActual(m0, m1 matrix.PolygonMatrix, ath Atherton) matrix.PolygonMatrix {

	subject := &algorithm.Plane{}
	for _, v2 := range m0 {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				subject.AddPoint(&algorithm.Vertex{Matrix: matrix.Matrix(v1)})
			}
		}
		subject.CloseRing()
		subject.Rank = calc.MAIN
	}
	clipping := &algorithm.Plane{}
	for _, v2 := range m1 {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				clipping.AddPoint(&algorithm.Vertex{Matrix: matrix.Matrix(v1)})
			}
		}
		clipping.CloseRing()
		clipping.Rank = calc.CUT
	}
	poly := Weiler(subject, clipping, ath)
	var result matrix.PolygonMatrix
	for _, v2 := range poly.Rings {
		var edge matrix.LineMatrix
		for _, v1 := range v2.Vertexs {
			edge = append(edge, v1.Matrix)
		}
		if !matrix.Matrix(edge[len(edge)-1]).Equals(matrix.Matrix(edge[0])) {
			edge = append(edge, edge[0])
		}
		result = append(result, edge)
	}

	return result
}
