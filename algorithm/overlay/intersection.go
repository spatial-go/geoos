package overlay

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

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

// Intersection  Computes the Intersection of two geometries,either or both of which may be null.
func Intersection(m0, m1 matrix.PolygonMatrix) matrix.PolygonMatrix {

	if m0 == nil && m1 == nil {
		return nil
	}
	if m0 == nil {
		return m1
	}

	if m1 == nil {
		return m0
	}

	return unionActual(m0, m1, Clip)
}
