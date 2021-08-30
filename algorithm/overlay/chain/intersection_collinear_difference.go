package chain

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// IntersectionCollinearDifference Finds interior intersections between line segments , and adds them.
type IntersectionCollinearDifference struct {
	Intersections relate.IntersectionPointLine
	startPos      int
	result        matrix.Collection
	line          matrix.LineMatrix
}

// ProcessIntersections This method is called by clients  to process intersections for two segments being intersected.
// Note that some clients (such as <codeProcessIntersections>MonotoneChain</code>s) may optimize away
// this call for segment pairs which they have determined do not intersect
func (ii *IntersectionCollinearDifference) ProcessIntersections(
	e0 matrix.LineMatrix, segIndex0 int,
	e1 matrix.LineMatrix, segIndex1 int) {
	// don't bother intersecting a segment with itself
	if e0.Equals(e1) && segIndex0 == segIndex1 {
		return
	}
	if segIndex0 > len(e0)-1 || segIndex1 > len(e1)-1 {
		return
	}
	mark, ips := relate.Intersection(e0[segIndex0], e0[segIndex0+1], e1[segIndex1], e1[segIndex1+1])

	if mark {
		if ips[0].IsCollinear {
			sort.Sort(ips)
			if matrix.Matrix(e0[segIndex0]).Equals(ips[0].Matrix) {
				ii.line = append(ii.line, e0[ii.startPos:segIndex0]...)
			} else {
				ii.line = append(ii.line, e0[ii.startPos:segIndex0+1]...)
				ii.line = append(ii.line, ips[0].Matrix)
			}
			if len(ii.line) > 1 {
				ii.result = append(ii.result, ii.line)
			}
			if segIndex0 < len(e0)-1 && matrix.Matrix(e0[segIndex0+1]).Equals(ips[len(ips)-1].Matrix) {
				ii.startPos = segIndex0 + 2
			} else {
				ii.startPos = segIndex0 + 1
			}
			ii.line = matrix.LineMatrix{}
			ii.line = append(ii.line, ips[len(ips)-1].Matrix)
		}
	}
	if segIndex0 == len(e0)-2 {
		ii.line = append(ii.line, e0[ii.startPos:]...)
		if len(ii.line) > 1 {
			ii.result = append(ii.result, ii.line)
		}
	}
}

// IsDone Always process all intersections
func (ii *IntersectionCollinearDifference) IsDone() bool {
	return false
}

// Result returns result.
func (ii *IntersectionCollinearDifference) Result() interface{} {
	return ii.result
}
