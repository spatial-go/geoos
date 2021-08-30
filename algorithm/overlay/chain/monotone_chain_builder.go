package chain

import (
	"github.com/spatial-go/geoos/algorithm/calc/angle"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Chains Computes a list of the MonotoneChains for a list of coordinates.
func Chains(edge matrix.LineMatrix) []*MonotoneChain {
	return ChainsContext(edge, nil)
}

// ChainsContext Computes a list of the MonotoneChains for a list of coordinates,attaching a context data object to each.
func ChainsContext(edge matrix.LineMatrix, context interface{}) []*MonotoneChain {
	mcList := []*MonotoneChain{}
	chainStart := 0
	for chainStart < len(edge)-1 {
		chainEnd := findChainEnd(edge, chainStart)
		mc := &MonotoneChain{Edge: edge, Start: chainStart, End: chainEnd, Context: context}
		mcList = append(mcList, mc)
		chainStart = chainEnd
	}
	return mcList
}

// Finds the index of the last point in a monotone chain starting at a given point.
// Repeated points (0-length segments) are included in the monotone chain returned.
func findChainEnd(pts matrix.LineMatrix, start int) int {
	safeStart := start
	// skip any zero-length segments at the start of the sequence
	// (since they cannot be used to establish a quadrant)
	for safeStart < len(pts)-1 && matrix.Matrix(pts[safeStart]).Equals(matrix.Matrix(pts[safeStart+1])) {
		safeStart++
	}
	// check if there are NO non-zero-length segments
	if safeStart >= len(pts)-1 {
		return len(pts) - 1
	}
	// determine overall quadrant for chain (which is the starting quadrant)
	chainQuad, _ := angle.Quadrant(pts[safeStart], pts[safeStart+1])
	last := start + 1
	for last < len(pts) {
		// skip zero-length segments, but include them in the chain
		if !matrix.Matrix(pts[last-1]).Equals(matrix.Matrix(pts[last])) {
			// compute quadrant for next possible segment in chain
			quad, _ := angle.Quadrant(pts[last-1], pts[last])
			if quad != chainQuad {
				break
			}
		}
		last++
	}
	return last - 1
}
