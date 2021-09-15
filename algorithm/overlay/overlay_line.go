package overlay

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
	"github.com/spatial-go/geoos/algorithm/relate"
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
	if ps, ok := p.Subject.(matrix.LineMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix:
			return Union(pc, ps), nil
		case matrix.LineMatrix:
			ins, _ := p.Intersection()
			sds, _ := p.SymDifference()
			ic := ins.(matrix.Collection)
			if len(ic) == 1 {
				if _, ok := ic[0].(matrix.Matrix); ok {
					return sds.(matrix.Collection), nil
				}
			}
			return append(ins.(matrix.Collection), sds.(matrix.Collection)...), nil
		case matrix.PolygonMatrix:
			return matrix.Collection{ps, pc}, nil
		case matrix.Collection:
			return append(pc, ps), nil
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *LineOverlay) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	var line matrix.LineMatrix
	if l, ok := p.Subject.(matrix.LineMatrix); ok {
		line = l
	} else {
		return nil, algorithm.ErrNotMatchType
	}
	switch c := p.Clipping.(type) {
	case matrix.Matrix:
		if mark := relate.InLineMatrix(c, line); mark {
			return c, nil
		}
		return nil, nil
	case matrix.LineMatrix:
		result := intersectLine(line, c)
		return LineMerge(result), nil
	case matrix.PolygonMatrix:
		result := matrix.Collection{}
		for _, ring := range c {
			res := intersectLine(line, ring[:len(ring)-1])
			result = append(result, res...)
		}
		filt := &matrix.UniqueArrayFilter{}
		result = result.Filter(filt).(matrix.Collection)
		return LineMerge(result), nil
	}
	return nil, algorithm.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *LineOverlay) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}
	if s, ok := p.Subject.(matrix.LineMatrix); ok {
		if c, ok := p.Clipping.(matrix.LineMatrix); ok {
			var err error
			if result, err := differenceLine(s, c); err == nil {
				return result, nil
			}
			return nil, err
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *LineOverlay) DifferenceReverse() (matrix.Steric, error) {
	newPoly := &LineOverlay{PointOverlay: &PointOverlay{Subject: p.Clipping, Clipping: p.Subject}}
	return newPoly.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *LineOverlay) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil {
		result = append(result, res.(matrix.Collection)...)
	}
	if res, err := p.DifferenceReverse(); err == nil {
		result = append(result, res.(matrix.Collection)...)
	}
	return result, nil
}

// intersectLine returns a array  that represents that part of geometry A intersect with geometry B.
func intersectLine(m, m1 matrix.LineMatrix) matrix.Collection {
	smi := &chain.SegmentMutualIntersector{SegmentMutual: m}
	icd := &chain.IntersectionCollinear{Edge: m}
	smi.Process(m1, icd)
	result := icd.Result()
	return result.(matrix.Collection)
}

// IntersectLine returns a array  that represents that part of geometry A intersect with geometry B.
func IntersectLine(m, m1 matrix.LineMatrix) []relate.IntersectionResult {

	mark, ips := relate.IntersectionEdge(m, m1)
	if !mark || len(ips) < 1 {
		return nil
	}
	ils := []relate.IntersectionResult{}
	il := relate.IntersectionResult{Ips: relate.IntersectionPointLine{}}
	for i, line := range m.ToLineArray() {
		for _, ip := range ips {
			if in, _ := relate.InLine(ip.Matrix, line.P0, line.P1); in {
				il.Pos = i
				il.Line = *line
				il.Start = i
				il.End = i + 1
				il.Ips = append(il.Ips, ip)
			}
		}
		if tes, _ := line.P0.Compare(line.P1); tes > 0 {
			sort.Sort(il.Ips)
		} else {
			sort.Sort(sort.Reverse(il.Ips))
		}
		if len(il.Ips) > 0 {
			ils = append(ils, il)
		}
		il = relate.IntersectionResult{Ips: relate.IntersectionPointLine{}}
	}

	return ils
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
