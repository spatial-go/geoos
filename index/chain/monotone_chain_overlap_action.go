package chain

import "github.com/spatial-go/geoos/algorithm/matrix"

type overlapAction interface {
	// This function can be overridden if the original chains are needed.
	Overlap(mc1 *MonotoneChain, start1 int, mc2 *MonotoneChain, start2 int)
}

// The action for the internal iterator for performing overlap queries on a MonotoneChain
type MonotoneChainOverlapAction struct {
	overlapSeg1 *matrix.LineSegment
	overlapSeg2 *matrix.LineSegment
}

// This function can be overridden if the original chains are needed.
func (m *MonotoneChainOverlapAction) Overlap(mc1 *MonotoneChain, start1 int, mc2 *MonotoneChain, start2 int) {
	if m.overlapSeg1 == nil {
		m.overlapSeg1 = &matrix.LineSegment{}
	}
	if m.overlapSeg2 == nil {
		m.overlapSeg2 = &matrix.LineSegment{}
	}
	m.overlapSeg1 = mc1.GetLineSegment(start1, m.overlapSeg1)
	m.overlapSeg2 = mc2.GetLineSegment(start2, m.overlapSeg2)
}
