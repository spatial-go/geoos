package overlay

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// LineOverlay  Computes the overlay of two geometries,either or both of which may be nil.
type LineOverlay struct {
	*PointOverlay
}

// Union  Computes the Union of two geometries,either or both of which may be nil.
func (p *LineOverlay) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if s, ok := p.subject.(matrix.LineMatrix); ok {
		if c, ok := p.clipping.(matrix.LineMatrix); ok {
			return LineMerge(matrix.Collection{s, c}), nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *LineOverlay) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	if s, ok := p.subject.(matrix.LineMatrix); ok {
		if c, ok := p.clipping.(matrix.LineMatrix); ok {
			result := matrix.Collection{}
			for _, il := range IntersectLine(s, c) {
				result = append(result, matrix.LineMatrix{il.line.P0, il.line.P1})
			}
			return LineMerge(result), nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *LineOverlay) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}
	if s, ok := p.subject.(matrix.LineMatrix); ok {
		if c, ok := p.clipping.(matrix.LineMatrix); ok {
			var err error
			if result, err := differenceLine(s, c); err == nil {
				if len(result.(matrix.Collection)) == 1 {
					return result.(matrix.Collection)[0], nil
				}
				return result, nil
			}
			return nil, err
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *LineOverlay) DifferenceReverse() (matrix.Steric, error) {
	newPoly := &LineOverlay{PointOverlay: &PointOverlay{subject: p.clipping, clipping: p.subject}}
	return newPoly.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *LineOverlay) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil {
		result = append(result, res)
	}
	if res, err := p.DifferenceReverse(); err == nil {
		result = append(result, res)
	}
	return result, nil
}
