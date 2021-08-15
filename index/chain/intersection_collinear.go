package chain

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// Finds interior intersections between line segments , and adds them.
type IntersectionCollinear struct {
	Intersections relate.IntersectionPointLine
	result        matrix.Collection
	line          matrix.LineMatrix
}

// This method is called by clients  to process intersections for two segments being intersected.
// Note that some clients (such as <code>MonotoneChain</code>s) may optimize away
// this call for segment pairs which they have determined do not intersect
func (ii *IntersectionCollinear) ProcessIntersections(
	e0 matrix.LineMatrix, segIndex0 int,
	e1 matrix.LineMatrix, segIndex1 int) {
	// don't bother intersecting a segment with itself
	if e0.Equals(e1) && segIndex0 == segIndex1 {
		ii.result = append(ii.result, e0)
		return
	}
	if segIndex0 > len(e0)-1 || segIndex1 > len(e1)-1 {
		return
	}
	mark, ips := relate.Intersection(e0[segIndex0], e0[segIndex0+1], e1[segIndex1], e1[segIndex1+1])

	if mark {
		if tes, _ := (matrix.Matrix(e0[segIndex0])).Compare(matrix.Matrix(e0[segIndex0+1])); tes > 0 {
			sort.Sort(ips)
		} else {
			sort.Sort(sort.Reverse(ips))
		}
		if len(ips) > 1 {
			for _, v := range ips {
				if !v.Matrix.Equals(matrix.Matrix(ii.line[len(ii.line)-1])) {
					ii.line = append(ii.line, v.Matrix)
				}
			}

		}
		if ii.line != nil && len(ii.line) > 1 {
			ii.result = append(ii.result, ii.line)
			ii.line = matrix.LineMatrix{}
		} else {
			ii.result = append(ii.result, ips[0].Matrix)
		}
	}
}

// Always process all intersections
func (ii *IntersectionCollinear) IsDone() bool {
	return false
}

// Always process all intersections
func (ii *IntersectionCollinear) Result() interface{} {
	return ii.result
}
