package chain

import "github.com/spatial-go/geoos/algorithm/matrix"

// SelectAction The action for the internal iterator for performing envelope select queries on a MonotoneChain
type SelectAction interface {
	// This function can be overridden if the original chains are needed.
	Select(mc1 *MonotoneChain, startIndex int)
}

// MonotoneChainSelectAction The action for the internal iterator for performing envelope select queries on a MonotoneChain
type MonotoneChainSelectAction struct {
	// these envelopes are used during the MonotoneChain search process
	//Envelope tempEnv1 = new Envelope();

	selectedSegment *matrix.LineSegment
}

// Select This method is overridden to process a segment in the context of the parent chain.
func (m *MonotoneChainSelectAction) Select(mc *MonotoneChain, startIndex int) {
	m.selectedSegment = mc.LineSegment(startIndex, m.selectedSegment)
}
