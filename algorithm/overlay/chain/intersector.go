package chain

import "github.com/spatial-go/geoos/algorithm/matrix"

// Intersector Processes possible intersections detected.
// detects that two SegmentStrings <i>might</i> intersect.
type Intersector interface {
	// This func is called by clients of interface to process
	// intersections for two segments of the lineSegments being intersected.
	ProcessIntersections(
		e0 matrix.LineMatrix, segIndex0 int,
		e1 matrix.LineMatrix, segIndex1 int)

	// Reports whether the client of this class needs to continue testing all intersections in an arrangement.
	IsDone() bool

	// Result return results.
	Result() interface{}
}
