package index

import "github.com/spatial-go/geoos/algorithm/matrix"

// ArrayVisitor Builds an array of all visited items.
type ArrayVisitor struct {
	Items []interface{}
}

// VisitItem Visits an item.
func (a *ArrayVisitor) VisitItem(item interface{}) {
	a.Items = append(a.Items, item)
}

// LineSegmentVisitor ItemVisitor subclass to reduce volume of query results.
type LineSegmentVisitor struct {

	// LineSegmentVisitor ItemVisitor subclass to reduce volume of query results.
	QuerySeg *matrix.LineSegment
	Items    []*matrix.LineSegment
}

// VisitItem ...
func (l *LineSegmentVisitor) VisitItem(item interface{}) {
	seg := item.(*matrix.LineSegment)
	if matrix.IsIntersectsTwo(seg.P0, seg.P1, l.QuerySeg.P0, l.QuerySeg.P1) {
		l.Items = append(l.Items, seg)
	}
}
