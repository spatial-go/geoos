package overlay

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
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

// IntersectLine returns a array  that represents that part of geometry A intersect with geometry B.
func IntersectLine(m, m1 matrix.LineMatrix) []IntersectionLineSegment {
	mark, ips := relate.IntersectionEdge(m, m1)
	if !mark || len(ips) <= 1 {
		return nil
	}
	ils := []IntersectionLineSegment{}
	il := IntersectionLineSegment{Ips: relate.IntersectionPointLine{}}
	for i, line := range m.ToLineArray() {
		for _, ip := range ips {
			if relate.InLine(ip.Matrix, line.P0, line.P1) {
				il.pos = i
				il.line = *line
				il.Ips = append(il.Ips, ip)
			}
		}
		if tes, _ := line.P0.Compare(line.P1); tes > 0 {
			sort.Sort(il.Ips)
		} else {
			sort.Sort(sort.Reverse(il.Ips))
		}
		if len(il.Ips) > 1 {
			ils = append(ils, il)
		}
		il = IntersectionLineSegment{Ips: relate.IntersectionPointLine{}}
	}

	return ils
}

func differenceLine(m, m1 matrix.LineMatrix) (matrix.Steric, error) {
	mark, ips := relate.IntersectionEdge(m, m1)
	if !mark || len(ips) <= 1 {
		return m, nil
	}
	ils := []IntersectionLineSegment{}
	il := IntersectionLineSegment{}
	for i, line := range m.ToLineArray() {
		for _, ip := range ips {
			if relate.InLine(ip.Matrix, line.P0, line.P1) {
				il.pos = i
				il.line = *line
				il.Ips = append(il.Ips, ip)
			}
		}
		sort.Sort(il.Ips)
		ils = append(ils, il)
		il = IntersectionLineSegment{}
	}
	result := matrix.Collection{}
	line := matrix.LineMatrix{}
	startPos := 0
	for _, v := range ils {
		if matrix.Matrix(m[v.pos]).Equals(v.Ips[0].Matrix) {
			line = append(line, m[startPos:v.pos]...)
		} else {
			line = append(line, m[startPos:v.pos+1]...)
			line = append(line, v.Ips[0].Matrix)
		}
		if len(line) > 1 {
			result = append(result, line)
		}
		if v.pos < len(m)-1 && matrix.Matrix(m[v.pos+1]).Equals(v.Ips[len(ips)-1].Matrix) {
			startPos = v.pos + 2
		} else {
			startPos = v.pos + 1
		}
		line = matrix.LineMatrix{}
		line = append(line, v.Ips[len(v.Ips)-1].Matrix)
	}
	line = append(line, m[startPos:]...)
	if len(line) > 1 {
		result = append(result, line)
	}
	return result, nil
}

// IntersectionLineSegment ...
type IntersectionLineSegment struct {
	pos  int
	line matrix.LineSegment
	Ips  relate.IntersectionPointLine
}
