package simplify

import (
	"log"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
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
	if err := l.index.Insert(envelope.TwoMatrix(seg.P0, seg.P1), seg); err != nil {
		log.Println(err)
	}
}

// Remove remove a TaggedLineSegment.
func (l *LineSegmentIndex) Remove(seg *TaggedLineSegment) {
	l.index.Remove(envelope.TwoMatrix(seg.P0, seg.P1), seg.LineSegment)
}

// Query query LineSegment returns array TaggedLineSegment.
func (l *LineSegmentIndex) Query(querySeg *matrix.LineSegment) []*matrix.LineSegment {
	env := envelope.TwoMatrix(querySeg.P0, querySeg.P1)

	visitor := &index.LineSegmentVisitor{QuerySeg: querySeg}
	if err := l.index.QueryVisitor(env, visitor); err != nil {
		log.Println(err)
		return nil
	} else {
		itemsFound := visitor.Items()
		return itemsFound.([]*matrix.LineSegment)
	}
}
