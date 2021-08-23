package simplify

import "github.com/spatial-go/geoos/algorithm/matrix"

// TaggedLineString Represents a LineString which can be modified to a simplified shape.
//  This class provides an attribute which specifies the minimum allowable length for the modified result.
type TaggedLineString struct {
	ParentLine  matrix.LineMatrix
	Segs        []*TaggedLineSegment
	resultSegs  []*matrix.LineSegment
	MinimumSize int
}

// GetParentMatrixes ...
func (t *TaggedLineString) GetParentMatrixes() []matrix.Matrix {
	return matrix.TransMatrixes(t.ParentLine)
}

// GetResultMatrixes ...
func (t *TaggedLineString) GetResultMatrixes() []matrix.Matrix {
	return extractCoordinates(t.resultSegs)
}

// GetResultSize ...
func (t *TaggedLineString) GetResultSize() int {
	resultSegsSize := len(t.resultSegs)
	if resultSegsSize == 0 {
		return 0
	}
	return resultSegsSize + 1
}

// GetSegment ...
func (t *TaggedLineString) GetSegment(i int) *TaggedLineSegment {
	if t == nil || t.Segs == nil {
		return nil
	}
	return t.Segs[i]
}

func (t *TaggedLineString) initTaggedLine() {
	pts := matrix.TransMatrixes(t.ParentLine)
	t.Segs = make([]*TaggedLineSegment, len(pts)-1)
	for i := range pts[:len(pts)-1] {
		seg := TaggedLineSegmentFour(pts[i], pts[i+1], t.ParentLine, i)
		t.Segs[i] = seg
	}
}

// AddToResult ...
func (t *TaggedLineString) AddToResult(seg *matrix.LineSegment) {
	t.resultSegs = append(t.resultSegs, seg)
}

// AsLine ...
func (t *TaggedLineString) AsLine() matrix.LineMatrix {
	line := matrix.LineMatrix{}
	for _, v := range extractCoordinates(t.resultSegs) {
		line = append(line, v)
	}
	return line
}

func extractCoordinates(segs []*matrix.LineSegment) []matrix.Matrix {
	pts := make([]matrix.Matrix, len(segs)+1)
	seg := &TaggedLineSegment{}
	for i, v := range segs {
		seg.LineSegment = v
		pts[i] = v.P0
	}
	// add last point
	pts[len(pts)-1] = seg.P1
	return pts
}
