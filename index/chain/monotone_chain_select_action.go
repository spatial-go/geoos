package chain

import "github.com/spatial-go/geoos/algorithm/matrix"

type SelectAction interface {
	// This function can be overridden if the original chains are needed.
	Select(mc1 *MonotoneChain, startIndex int)
}

// The action for the internal iterator for performing envelope select queries on a MonotoneChain
type MonotoneChainSelectAction struct {
	// these envelopes are used during the MonotoneChain search process
	//Envelope tempEnv1 = new Envelope();

	selectedSegment *matrix.LineSegment
}

// This method is overridden to process a segment in the context of the parent chain.
func (m *MonotoneChainSelectAction) Select(mc *MonotoneChain, startIndex int) {
	m.selectedSegment = mc.GetLineSegment(startIndex, m.selectedSegment)
}
