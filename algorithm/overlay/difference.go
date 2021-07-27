package overlay

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
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
		if result, err := differenceLine(m, m1.(matrix.LineMatrix)); err == nil {
			if len(result.(matrix.Collection)) == 1 {
				return result.(matrix.Collection)[0], nil
			}
			return result, nil
		}
		return nil, err
	case matrix.PolygonMatrix:
		//poly := Intersection(m0.(matrix.PolygonMatrix), m1.(matrix.PolygonMatrix))
		//TODO
		return nil, nil
	default:
		return nil, algoerr.ErrNotSupportCollection

	}
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
				il.ips = append(il.ips, ip)
			}
		}
		sort.Sort(il.ips)
		ils = append(ils, il)
		il = IntersectionLineSegment{}
	}
	result := matrix.Collection{}
	line := matrix.LineMatrix{}
	startPos := 0
	for _, v := range ils {
		if matrix.Matrix(m[v.pos]).Equals(v.ips[0].Matrix) {
			line = append(line, m[startPos:v.pos]...)
		} else {
			line = append(line, m[startPos:v.pos+1]...)
			line = append(line, v.ips[0].Matrix)
		}
		if len(line) > 1 {
			result = append(result, line)
		}
		if v.pos < len(m)-1 && matrix.Matrix(m[v.pos+1]).Equals(v.ips[len(ips)-1].Matrix) {
			startPos = v.pos + 2
		} else {
			startPos = v.pos + 1
		}
		line = matrix.LineMatrix{}
		line = append(line, v.ips[len(v.ips)-1].Matrix)
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
	ips  relate.IntersectionPointLine
}
