package overlay

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func SymDifference(m0, m1 matrix.Steric) (matrix.Steric, error) {

	result := matrix.Collection{}
	if res, err := Difference(m0, m1); err == nil {
		if r, ok := res.(matrix.Collection); ok {
			for _, v := range r {
				result = append(result, v)
			}
		} else {
			result = append(result, res)
		}
	}
	if res, err := Difference(m1, m0); err == nil {
		if r, ok := res.(matrix.Collection); ok {
			for _, v := range r {
				result = append(result, v)
			}
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
		newLine := &LineOverlay{PointOverlay: &PointOverlay{Subject: m, Clipping: m1}}
		if result, err := newLine.Difference(); err == nil {
			if len(result.(matrix.Collection)) == 1 {
				return result.(matrix.Collection)[0], nil
			}
			return result, nil
		}
		return nil, err
	case matrix.PolygonMatrix:
		newPoly := &PolygonOverlay{PointOverlay: &PointOverlay{Subject: m, Clipping: m1}}
		return newPoly.Difference()
	default:
		return nil, algoerr.ErrNotSupportCollection

	}
}
