// Package clipping the spatial geometric operation and reconstruction between entities is realized.
// a method for spatial geometry operation by update geometric relation graph.
package clipping

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/graph"
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

// Union  Computes the Union of two geometries, if one is encountered.
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

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
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

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func SymDifference(m0, m1 matrix.Steric) (matrix.Steric, error) {

	result := matrix.Collection{}
	if res, err := Difference(m0, m1); err == nil {
		if r, ok := res.(matrix.Collection); ok {
			result = append(result, r...)
		} else {
			result = append(result, res)
		}
	}
	if res, err := Difference(m1, m0); err == nil {
		if r, ok := res.(matrix.Collection); ok {
			result = append(result, r...)
		} else {
			result = append(result, res)
		}
	}
	return result, nil
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func Difference(m0, m1 matrix.Steric) (matrix.Steric, error) {
	switch m := m0.(type) {
	case matrix.Matrix:
		return m0, nil
	case matrix.LineMatrix:
		var err error
		newLine := &LineClipping{PointClipping: &PointClipping{Subject: m, Clipping: m1}}
		if result, err := newLine.Difference(); err == nil {
			return result, nil
		}
		return nil, err
	case matrix.PolygonMatrix:
		newPoly := &PolygonClipping{PointClipping: &PointClipping{Subject: m, Clipping: m1}}
		return newPoly.Difference()
	default:
		return nil, algorithm.ErrNotSupportCollection

	}
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

// LineMerge returns a Geometry containing the LineMerges.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func LineMerge(ml matrix.Collection) ([]matrix.LineMatrix, error) {
	lines := make([]matrix.LineMatrix, len(ml))
	for i, v := range ml {
		if line, ok := v.(matrix.LineMatrix); ok {
			lines[i] = line
		} else {
			return nil, algorithm.ErrUnknownType(ml)
		}
	}

	return mergeByHalf(lines, 0, len(lines))
}

// mergeByHalf returns Unions a section of a list using a recursive binary union on each half of the section.
func mergeByHalf(matrix4 []matrix.LineMatrix, start, end int) (result []matrix.LineMatrix, err error) {
	if matrix4 == nil {
		return nil, nil
	}
	if end-start <= 1 {
		result, err = merge([]matrix.LineMatrix{matrix4[start]}, nil)
	} else if end-start == 2 {
		result, err = merge([]matrix.LineMatrix{matrix4[start]}, []matrix.LineMatrix{matrix4[start+1]})
	} else {
		mid := (end + start) / 2
		g0, _ := mergeByHalf(matrix4, start, mid)
		g1, _ := mergeByHalf(matrix4, mid, end)
		result, err = merge(g0, g1)
	}
	return
}

func merge(ps, pc []matrix.LineMatrix) (result []matrix.LineMatrix, err error) {
	cs := matrix.CollectionFromMultiLineMatrix(ps)
	cc := matrix.CollectionFromMultiLineMatrix(pc)
	clip := graph.MergeHandle(cs, cc)
	gu, _ := clip.Union()
	gi, _ := clip.Intersection()
	return linkmerge(gu, gi)
}
