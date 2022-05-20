package clipping

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// UnaryUnion returns a Geometry containing the union.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func UnaryUnion(matrix4 matrix.Steric) (result matrix.Steric, err error) {
	if c, ok := matrix4.(matrix.Collection); ok {
		return unaryUnionByHalf(c, 0, len(c))
	}
	return nil, algorithm.ErrUnknownType(matrix4)
}

// unaryUnionByHalf returns Unions a section of a list using a recursive binary union on each half of the section.
func unaryUnionByHalf(matrix4 matrix.Collection, start, end int) (result matrix.Steric, err error) {
	if matrix4 == nil {
		return nil, nil
	}
	if end-start <= 1 {
		result, err = Union(matrix4[start], nil)
	} else if end-start == 2 {
		result, err = Union(matrix4[start], matrix4[start+1])
	} else {
		mid := (end + start) / 2
		g0, _ := unaryUnionByHalf(matrix4, start, mid)
		g1, _ := unaryUnionByHalf(matrix4, mid, end)
		result, err = Union(g0, g1)
	}
	return
}

// Union  Computes the Union of two geometries,either or both of which may be null.
func Union(m0, m1 matrix.Steric) (result matrix.Steric, err error) {
	switch m := m0.(type) {
	case matrix.Matrix:
		over := &PointClipping{Subject: m, Clipping: m1}
		result, err = over.Union()
	case matrix.LineMatrix:
		over := &LineClipping{PointClipping: &PointClipping{Subject: m0, Clipping: m1}}
		result, err = over.Union()
	case matrix.PolygonMatrix:
		polyOver := &PolygonClipping{PointClipping: &PointClipping{Subject: m0, Clipping: m1}}
		result, err = polyOver.Union()
	case matrix.Collection:
		resultColl := matrix.Collection{}
		IsUnion := false
		for _, v := range m {
			un, err := Union(v, m1)
			if _, ok := un.(matrix.Collection); ok || err != nil {
				resultColl = append(resultColl, v)
			} else {
				IsUnion = true
				resultColl = append(resultColl, un)
			}
		}
		if !IsUnion {
			switch coll := m1.(type) {
			case matrix.Collection:
				resultColl = append(resultColl, coll...)
			default:
				resultColl = append(resultColl, coll)
			}
		}
		result = resultColl
	}
	return result, err
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be null.
func Intersection(m0, m1 matrix.Steric) (matrix.Steric, error) {
	switch m := m0.(type) {
	case matrix.Matrix:
		over := &PointClipping{Subject: m, Clipping: m1}
		return over.Intersection()
	case matrix.LineMatrix:
		var err error
		newLine := &LineClipping{PointClipping: &PointClipping{Subject: m, Clipping: m1}}
		if result, err := newLine.Intersection(); err == nil {
			return result, nil
		}
		return nil, err
	case matrix.PolygonMatrix:
		polyOver := &PolygonClipping{PointClipping: &PointClipping{Subject: m, Clipping: m1}}
		return polyOver.Intersection()
	default:
		return nil, algorithm.ErrNotSupportCollection

	}
}
