package simplify

import "github.com/spatial-go/geoos/algorithm/matrix"

// TaggedLineSegment A LineSegment which is tagged with its location in a parent  Geometry.
//  Used to index the segments in a geometry and recover the segment locations
//  from the index.
type TaggedLineSegment struct {
	*matrix.LineSegment
	Parent matrix.Steric
	Index  int
}

// TaggedLineSegmentFour four parameters create TaggedLineSegmentFour.
func TaggedLineSegmentFour(p0, p1 matrix.Matrix, parent matrix.Steric, index int) *TaggedLineSegment {
	tls := &TaggedLineSegment{LineSegment: &matrix.LineSegment{P0: p0, P1: p1}, Parent: parent, Index: index}
	return tls
}

// TaggedLineSegmentTwo  two parameters create TaggedLineSegmentFour.
func TaggedLineSegmentTwo(p0, p1 matrix.Matrix) *TaggedLineSegment {
	return TaggedLineSegmentFour(p0, p1, nil, -1)
}
