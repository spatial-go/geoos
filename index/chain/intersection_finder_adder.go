package chain

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// IntersectionFinderAdder Finds intersections between line segments , and adds them.
type IntersectionFinderAdder struct {
	Intersections relate.IntersectionPointLine
}

// ProcessIntersections This method is called by clients  to process intersections for two segments being intersected.
// Note that some clients (such as <code>MonotoneChain</code>s) may optimize away
// this call for segment pairs which they have determined do not intersect
func (ii *IntersectionFinderAdder) ProcessIntersections(
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
		ii.Intersections = append(ii.Intersections, ips...)
	}
}

// IsDone Always process all intersections
func (ii *IntersectionFinderAdder) IsDone() bool {
	return false
}

// Result returns result.
func (ii *IntersectionFinderAdder) Result() interface{} {
	return ii.Intersections
}
