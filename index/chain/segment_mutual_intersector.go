package chain

import "github.com/spatial-go/geoos/algorithm/matrix"

// Intersects two sets of SegmentString.
type SegmentMutualIntersector struct {
	SegmentMutual matrix.LineMatrix
}

//the given collection of SegmentStrings and the set of indexed segments.
func (s *SegmentMutualIntersector) Process(segStrings []*matrix.LineSegment, segInt Intersector) {
	monoChains, testChains := []*MonotoneChain{}, []*MonotoneChain{}
	for _, v := range s.SegmentMutual.ToLineArray() {
		monoChains = s.AddToMonoChains(v, monoChains)
	}
	for _, v := range segStrings {
		testChains = s.AddToMonoChains(v, testChains)
	}
	s.IntersectChains(monoChains, testChains, segInt)
}

func (s *SegmentMutualIntersector) AddToMonoChains(segStr *matrix.LineSegment, monoChains []*MonotoneChain) []*MonotoneChain {
	pts := matrix.LineMatrix{segStr.P0, segStr.P1}
	segChains := ChainsContext(pts, segStr)

	monoChains = append(monoChains, segChains...)
	return monoChains
}

func (s *SegmentMutualIntersector) IntersectChains(monoChains []*MonotoneChain, testChains []*MonotoneChain, segInt Intersector) {
	overlapAction := &SegmentOverlapAction{MonotoneChainOverlapAction: &MonotoneChainOverlapAction{}, si: segInt}

	for _, queryChain := range monoChains {
		for _, testChain := range testChains {
			queryChain.ComputeOverlaps(testChain, overlapAction)
			if segInt.IsDone() {
				return
			}
		}
	}
}

type SegmentOverlapAction struct {
	*MonotoneChainOverlapAction
	si Intersector
}

func (s *SegmentOverlapAction) SegmentOverlapAction(si Intersector) {
	s.si = si
}

func (s *SegmentOverlapAction) Overlap(mc1 *MonotoneChain, start1 int, mc2 *MonotoneChain, start2 int) {
	ss1 := mc1.Edge
	ss2 := mc2.Edge
	s.si.ProcessIntersections(ss1, start1, ss2, start2)
}
