// Package chain Define Monotone Chains,it is a way of partitioning the segments of a linestring to
// allow for fast searching of intersections.
package chain

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

// MonotoneChain Monotone Chains are a way of partitioning the segments of a linestring to allow for fast searching of intersections.
// They have the following properties:
//  the segments within a monotone chain never intersect each other
//  he envelope of any contiguous subset of the segments in a monotone chain
//  is equal to the envelope of the endpoints of the subset.
type MonotoneChain struct {
	Edge       matrix.LineMatrix
	Start, End int
	Env        *envelope.Envelope
	Context    interface{}
	ID         int // useful for optimizing chain comparisons
}

// Envelope Gets the envelope of the chain.
func (m *MonotoneChain) Envelope() *envelope.Envelope {
	return m.EnvelopeExpansion(0.0)
}

// EnvelopeExpansion Gets the envelope for this chain,expanded by a given distance.
func (m *MonotoneChain) EnvelopeExpansion(expansionDistance float64) *envelope.Envelope {
	if m.Env == nil {
		/**
		 * The monotonicity property allows fast envelope determination
		 */
		m.Env = envelope.Bound(m.Edge.Bound())
		if expansionDistance > 0.0 {
			m.Env.ExpandBy(expansionDistance)
		}
	}
	return m.Env
}

// LineSegment Gets the line segment starting at <code>index</code>
func (m *MonotoneChain) LineSegment(index int, ls *matrix.LineSegment) *matrix.LineSegment {
	ls.P0 = m.Edge[index]
	ls.P1 = m.Edge[index+1]
	return ls
}

// Select Determine all the line segments in the chain whose envelopes overlap the searchEnvelope, and process them.
func (m *MonotoneChain) Select(searchEnv *envelope.Envelope, mcs *MonotoneChainSelectAction) {
	m.computeSelect(searchEnv, m.Start, m.End, mcs)
}

func (m *MonotoneChain) computeSelect(
	searchEnv *envelope.Envelope,
	start0, end0 int,
	mcs *MonotoneChainSelectAction) {
	p0 := m.Edge[start0]
	p1 := m.Edge[end0]

	// terminating condition for the recursion
	if end0-start0 == 1 {
		mcs.Select(m, start0)
		return
	}
	// nothing to do if the envelopes don't overlap
	if !searchEnv.IsIntersects(envelope.TwoMatrix(p0, p1)) {
		return
	}
	// the chains overlap, so split each in half and iterate  (binary search)
	mid := (start0 + end0) / 2

	// check terminating conditions before recursing
	if start0 < mid {
		m.computeSelect(searchEnv, start0, mid, mcs)
	}
	if mid < end0 {
		m.computeSelect(searchEnv, mid, end0, mcs)
	}
}

// ComputeOverlaps Determines the line segments in two chains which may overlap,
// and passes them to an overlap action.
func (m *MonotoneChain) ComputeOverlaps(mc *MonotoneChain, mco overlapAction) {
	m.computeOverlapsTwo(m.Start, m.End, mc, mc.Start, mc.End, 0.0, mco)
}

// ComputeOverlapsTolerance Determines the line segments in two chains which may overlap,
//  using an overlap distance tolerance, and passes them to an overlap action.
func (m *MonotoneChain) ComputeOverlapsTolerance(mc *MonotoneChain, overlapTolerance float64, mco overlapAction) {
	m.computeOverlapsTwo(m.Start, m.End, mc, mc.Start, mc.End, overlapTolerance, mco)
}

// Uses an efficient mutual binary search strategy to determine which pairs of chain segments
// may overlap, and calls the given overlap action on them.
func (m *MonotoneChain) computeOverlapsTwo(
	start0, end0 int,
	mc *MonotoneChain,
	start1, end1 int,
	overlapTolerance float64,
	mco overlapAction) {

	// terminating condition for the recursion
	if end0-start0 == 1 && end1-start1 == 1 {
		mco.Overlap(m, start0, mc, start1)
		return
	}
	// nothing to do if the envelopes of these subchains don't overlap
	if !m.OverlapsMonotoneChain(start0, end0, mc, start1, end1, overlapTolerance) {
		return
	}
	// the chains overlap, so split each in half and iterate  (binary search)
	mid0 := (start0 + end0) / 2
	mid1 := (start1 + end1) / 2

	if start0 < mid0 {
		if start1 < mid1 {
			m.computeOverlapsTwo(start0, mid0, mc, start1, mid1, overlapTolerance, mco)
		}
		if mid1 < end1 {
			m.computeOverlapsTwo(start0, mid0, mc, mid1, end1, overlapTolerance, mco)
		}
	}
	if mid0 < end0 {
		if start1 < mid1 {
			m.computeOverlapsTwo(mid0, end0, mc, start1, mid1, overlapTolerance, mco)
		}
		if mid1 < end1 {
			m.computeOverlapsTwo(mid0, end0, mc, mid1, end1, overlapTolerance, mco)
		}
	}
}

// OverlapsMonotoneChain Tests whether the envelope of a section of the chain overlaps (intersects) the envelope of a section of another target chain.
// This test is efficient due to the monotonicity property of the sections (i.e. the envelopes can be are determined
// from the section endpoints rather than a full scan).
func (m *MonotoneChain) OverlapsMonotoneChain(
	start0, end0 int,
	mc *MonotoneChain,
	start1, end1 int,
	overlapTolerance float64) bool {
	if overlapTolerance > 0.0 {
		return Overlaps(m.Edge[start0], m.Edge[end0], mc.Edge[start1], mc.Edge[end1], overlapTolerance)
	}
	return envelope.IsIntersectsTwo(m.Edge[start0], m.Edge[end0], mc.Edge[start1], mc.Edge[end1])
}

// Overlaps Tests whether the envelope of a section of the chain overlaps (intersects) the envelope of a section of another target chain.
// This test is efficient due to the monotonicity property of the sections (i.e. the envelopes can be are determined
// from the section endpoints rather than a full scan).
func Overlaps(p1, p2, q1, q2 matrix.Matrix, overlapTolerance float64) bool {
	minQ := math.Min(q1[0], q2[0])
	maxQ := math.Max(q1[0], q2[0])
	minP := math.Min(p1[0], p2[0])
	maxP := math.Max(p1[0], p2[0])

	if minP > maxQ+overlapTolerance {
		return false
	}
	if maxP < minQ-overlapTolerance {
		return false
	}
	minQ = math.Min(q1[1], q2[1])
	maxQ = math.Max(q1[1], q2[1])
	minP = math.Min(p1[1], p2[1])
	maxP = math.Max(p1[1], p2[1])

	if minP > maxQ+overlapTolerance {
		return false
	}
	if maxP < minQ-overlapTolerance {
		return false
	}
	return true
}
