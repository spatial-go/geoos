package simplify

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/index"
	"github.com/spatial-go/geoos/index/quadtree"
)

// LineSegmentIndex An spatial index on a set of  LineSegments.
// Supports adding and removing items.
type LineSegmentIndex struct {
	index *quadtree.Quadtree
}

// Add  add a LineMatrix.
func (l *LineSegmentIndex) Add(line matrix.LineMatrix) {
	segs := line.ToLineArray()
	for i := 0; i < len(segs); i++ {
		seg := segs[i]
		l.AddSegment(seg)
	}
}

// AddSegment add a LineSegment.
func (l *LineSegmentIndex) AddSegment(seg *matrix.LineSegment) {
	l.index.Insert(matrix.EnvelopeTwoMatrix(seg.P0, seg.P1), seg)
}

// Remove remove a TaggedLineSegment.
func (l *LineSegmentIndex) Remove(seg *TaggedLineSegment) {
	l.index.Remove(matrix.EnvelopeTwoMatrix(seg.P0, seg.P1), seg.LineSegment)
}

// Query query LineSegment returns array TaggedLineSegment.
func (l *LineSegmentIndex) Query(querySeg *matrix.LineSegment) []*TaggedLineSegment {
	env := matrix.EnvelopeTwoMatrix(querySeg.P0, querySeg.P1)

	visitor := &LineSegmentVisitor{QuerySeg: querySeg}
	l.index.QueryVisitor(env, visitor.ItemVisitor)
	itemsFound := visitor.Items

	return itemsFound
}

// LineSegmentVisitor ItemVisitor subclass to reduce volume of query results.
type LineSegmentVisitor struct {
	index.ItemVisitor
	QuerySeg *matrix.LineSegment
	Items    []*TaggedLineSegment
}

// VisitItem ...
func (l *LineSegmentVisitor) VisitItem(item *TaggedLineSegment) {
	seg := item
	if matrix.IsIntersectsTwo(seg.P0, seg.P1, l.QuerySeg.P0, l.QuerySeg.P1) {
		l.Items = append(l.Items, item)
	}
}
