package clipping

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/operation"
)

// PointClipping  Computes the overlay of two geometries.
type PointClipping struct {
	Subject, Clipping matrix.Steric
}

// Union  Computes the Union of two geometries, if one is encountered.
func (p *PointClipping) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.Matrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix:
			if ps.Equals(pc) {
				return ps, nil
			}
			return matrix.Collection{ps, pc}, nil
		case matrix.LineMatrix:
			if pc.IsClosed() {
				if operation.InLineMatrix(ps, pc) {
					return matrix.PolygonMatrix{pc}, nil
				}
				if operation.IsPnPolygon(ps, pc) {
					return matrix.PolygonMatrix{pc}, nil
				}
			}
			if operation.InLineMatrix(ps, pc) {
				return pc, nil
			}

			return matrix.Collection{ps, pc}, nil
		case matrix.PolygonMatrix:
			switch pointInPolygon, _ := operation.InPolygon(ps, pc); pointInPolygon {
			case operation.OnlyInLine, operation.OnlyInPolygon:
				return pc, nil
			case operation.OnlyOutPolygon:
				return matrix.Collection{ps, pc}, nil
			}
		case matrix.Collection:
			var result matrix.Collection
			for _, v := range pc {
				over := &PointClipping{Subject: ps, Clipping: v}
				res, _ := over.Union()
				if _, ok = res.(matrix.Collection); ok {
					result = append(result, res.(matrix.Collection)...)
				} else {
					result = append(result, res)
				}
			}
			return result, nil
		}

	}
	return nil, algorithm.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *PointClipping) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.Matrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix:
			if ps.Equals(pc) {
				return ps, nil
			}
			return nil, nil
		case matrix.LineMatrix:
			if mark := operation.InLineMatrix(ps, pc); mark {
				return ps, nil
			}
			return nil, nil
		case matrix.PolygonMatrix:

			switch pointInPolygon, _ := operation.InPolygon(ps, pc); pointInPolygon {
			case operation.OnlyInLine, operation.OnlyInPolygon:
				return ps, nil
			case operation.OnlyOutPolygon:
				return nil, nil
			}
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *PointClipping) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.Matrix); ok {
		if pc, ok := p.Clipping.(matrix.Matrix); ok {
			if ps.Equals(pc) {
				return nil, nil
			}
			return ps, nil
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *PointClipping) DifferenceReverse() (matrix.Steric, error) {
	newP := &PointClipping{Subject: p.Clipping, Clipping: p.Subject}
	return newP.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *PointClipping) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil {
		result = append(result, res)
	}
	if res, err := p.DifferenceReverse(); err == nil {
		result = append(result, res)
	}
	return result, nil
}

// unionCheck  check  two geometries for union.
func (p *PointClipping) unionCheck() (matrix.Steric, bool) {

	if p.Subject == nil && p.Clipping == nil {
		return nil, false
	}
	if p.Subject == nil {
		return p.Clipping, false
	}

	if p.Clipping == nil {
		return p.Subject, false
	}

	return nil, true
}

// intersectionCheck  check  two geometries for intersection.
func (p *PointClipping) intersectionCheck() (matrix.Steric, bool) {

	if p.Subject == nil && p.Clipping == nil {
		return nil, false
	}
	if p.Subject == nil {
		return nil, false
	}

	if p.Clipping == nil {
		return nil, false
	}

	return nil, true
}

// differenceCheck check  two geometries for difference.
func (p *PointClipping) differenceCheck() (matrix.Steric, bool) {

	if p.Subject == nil && p.Clipping == nil {
		return nil, false
	}
	if p.Subject == nil {
		return nil, false
	}

	if p.Clipping == nil {
		return p.Subject, false
	}

	return nil, true
}
