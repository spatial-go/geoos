package clipping

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/graph/de9im"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
)

// LineClipping  Computes the overlay of two geometries,either or both of which may be nil.
type LineClipping struct {
	*PointClipping
}

// Union  Computes the Union of two geometries,either or both of which may be nil.
func (p *LineClipping) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.LineMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix:
			return Union(pc, ps), nil
		case matrix.LineMatrix:
			gu, _ := ClipHandle(ps, pc).Union()
			return graph.GenerateSteric(gu)
		case matrix.PolygonMatrix:
			switch im := de9im.IM(ps, pc); {
			case im.IsDisjoint():
				return matrix.Collection{ps, pc}, nil
			case im.IsCoveredBy():
				return pc, nil
			}

			gu, _ := ClipHandle(ps, pc).Union()
			return graph.GenerateSteric(gu)
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *LineClipping) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.LineMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix:
			return Intersection(pc, ps)
		case matrix.LineMatrix:
			gi, _ := ClipHandle(ps, pc).Intersection()
			return graph.GenerateSteric(gi)
		case matrix.PolygonMatrix:
			switch im := de9im.IM(ps, pc); {
			case im.IsDisjoint():
				return matrix.Collection{}, nil
			case im.IsCoveredBy():
				return ps, nil
			}
			gu, _ := ClipHandle(ps, pc).Union()
			for _, v := range gu.Nodes() {
				switch im := de9im.IM(v.Value, pc); {
				case !im.IsWithin() || v.Value.Equals(ps):
					gu.DeleteNode(v)
				}
			}
			return graph.GenerateSteric(gu)
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *LineClipping) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}

	if ps, ok := p.Subject.(matrix.LineMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix:
			return ps, nil
		case matrix.LineMatrix:
			gd, _ := ClipHandle(ps, pc).Difference()
			return graph.GenerateSteric(gd)

			// var err error
			// if result, err := differenceLine(ps, pc); err == nil {
			// 	return result, nil
			// }
			// return nil, err
		case matrix.PolygonMatrix:
			var result matrix.Collection
			for _, v := range pc {
				if res, err := differenceLine(ps, v); err == nil {
					result = append(result, res.(matrix.Collection)...)
				}
			}
			return result, nil
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *LineClipping) DifferenceReverse() (matrix.Steric, error) {
	newPoly := &LineClipping{PointClipping: &PointClipping{Subject: p.Clipping, Clipping: p.Subject}}
	return newPoly.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *LineClipping) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil {
		result = append(result, res.(matrix.Collection)...)
	}
	if res, err := p.DifferenceReverse(); err == nil {
		result = append(result, res.(matrix.Collection)...)
	}
	return result, nil
}

func differenceLine(m, m1 matrix.LineMatrix) (matrix.Steric, error) {
	smi := &chain.SegmentMutualIntersector{SegmentMutual: m}
	icd := &chain.IntersectionCollinearDifference{Edge: m}
	smi.Process(m1, icd)
	result := icd.Result()
	if m, ok := result.(matrix.Collection); ok && m != nil {
		return m, nil
	}
	return matrix.Collection{}, nil
}
