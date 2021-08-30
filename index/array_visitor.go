package index

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

// ArrayVisitor Builds an array of all visited items.
type ArrayVisitor struct {
	ItemsArray []interface{}
}

// VisitItem Visits an item.
func (a *ArrayVisitor) VisitItem(item interface{}) {
	a.ItemsArray = append(a.ItemsArray, item)
}

// Items returns items.
func (a *ArrayVisitor) Items() interface{} {
	return a.ItemsArray
}

// LineSegmentVisitor ItemVisitor subclass to reduce volume of query results.
type LineSegmentVisitor struct {

	// LineSegmentVisitor ItemVisitor subclass to reduce volume of query results.
	QuerySeg          *matrix.LineSegment
	ItemsArrayLineSeg []*matrix.LineSegment
}

// VisitItem ...
func (l *LineSegmentVisitor) VisitItem(item interface{}) {
	seg := item.(*matrix.LineSegment)
	if envelope.IsIntersectsTwo(seg.P0, seg.P1, l.QuerySeg.P0, l.QuerySeg.P1) {
		l.ItemsArrayLineSeg = append(l.ItemsArrayLineSeg, seg)
	}
}

// Items returns items.
func (l *LineSegmentVisitor) Items() interface{} {
	return l.ItemsArrayLineSeg
}
