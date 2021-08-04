package overlay

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Overlay  Computes the overlay of two geometries,either or both of which may be nil.
type Overlay interface {
	// Union  Computes the Union of two geometries,either or both of which may be nil.
	Union() (matrix.Steric, error)
	// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
	Intersection() (matrix.Steric, error)
	// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
	// One can think of this as GeometryA - Intersection(A,B).
	// If A is completely contained in B then an empty geometry collection is returned.
	Difference() (matrix.Steric, error)
	// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
	// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
	// One can think of this as Union(geomA,geomB) - Intersection(A,B).
	SymDifference() (matrix.Steric, error)
}

// PointOverlay  Computes the overlay of two geometries,either or both of which may be nil.
type PointOverlay struct {
	subject, clipping matrix.Steric
}

// Union  Computes the Union of two geometries,either or both of which may be nil.
func (p *PointOverlay) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if s, ok := p.subject.(matrix.Matrix); ok {
		if c, ok := p.clipping.(matrix.Matrix); ok {
			if s.Equals(c) {
				return s, nil
			}
			return matrix.Collection{s, c}, nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *PointOverlay) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	if s, ok := p.subject.(matrix.Matrix); ok {
		if c, ok := p.clipping.(matrix.Matrix); ok {
			if s.Equals(c) {
				return s, nil
			}
			return nil, nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *PointOverlay) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}
	if s, ok := p.subject.(matrix.Matrix); ok {
		if c, ok := p.clipping.(matrix.Matrix); ok {
			if s.Equals(c) {
				return nil, nil
			}
			return s, nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *PointOverlay) DifferenceReverse() (matrix.Steric, error) {
	newP := &PointOverlay{subject: p.clipping, clipping: p.subject}
	return newP.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *PointOverlay) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil {
		result = append(result, res)
	}
	if res, err := p.DifferenceReverse(); err == nil {
		result = append(result, res)
	}
	return result, nil
}

// unionCheck  Computes the Union of two geometries,either or both of which may be null.
func (p *PointOverlay) unionCheck() (matrix.Steric, bool) {

	if p.subject == nil && p.clipping == nil {
		return nil, false
	}
	if p.subject == nil {
		return p.clipping, false
	}

	if p.clipping == nil {
		return p.subject, false
	}

	return nil, true
}

// intersectionCheck  Computes the Union of two geometries,either or both of which may be null.
func (p *PointOverlay) intersectionCheck() (matrix.Steric, bool) {

	if p.subject == nil && p.clipping == nil {
		return nil, false
	}
	if p.subject == nil {
		return nil, false
	}

	if p.clipping == nil {
		return nil, false
	}

	return nil, true
}

// differenceCheck  Computes the Union of two geometries,either or both of which may be null.
func (p *PointOverlay) differenceCheck() (matrix.Steric, bool) {

	if p.subject == nil && p.clipping == nil {
		return nil, false
	}
	if p.subject == nil {
		return nil, false
	}

	if p.clipping == nil {
		return p.subject, false
	}

	return nil, true
}
