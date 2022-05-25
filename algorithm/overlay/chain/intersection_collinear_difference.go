package chain

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// IntersectionCollinearDifference Finds interior intersections between line segments , and adds them.
type IntersectionCollinearDifference struct {
	mono   []*MonotoneChain
	Edge   matrix.LineMatrix
	result matrix.Collection
}

// ProcessIntersections This method is called by clients  to process intersections for two segments being intersected.
// Note that some clients (such as <codeProcessIntersections>MonotoneChain</code>s) may optimize away
// this call for segment pairs which they have determined do not intersect
func (ii *IntersectionCollinearDifference) ProcessIntersections(
	e0 matrix.LineMatrix, segIndex0 int,
	e1 matrix.LineMatrix, segIndex1 int) {
	// don't bother intersecting a segment with itself
	if e0.Equals(e1) && segIndex0 == segIndex1 {
		return
	}
	if segIndex0 > len(e0)-1 || segIndex1 > len(e1)-1 {
		return
	}
	mark, ips := relate.Intersection(e0[segIndex0], e0[segIndex0+1], e1[segIndex1], e1[segIndex1+1])
	if mark {
		var monos = []*MonotoneChain{}
		var mono *MonotoneChain
		var edge matrix.LineMatrix
		var line matrix.LineMatrix
		var startPos int
		monoNum := 0
		for _, v := range ii.mono {
			if segIndex0 >= v.Start && segIndex0 < v.End && v.Context.(bool) {
				monoNum++
				mono = v
				monos = append(monos, v)
			}
		}
		if monoNum > 1 {
			for _, v := range monos {
				if markInter := relate.IsIntersectionEdge(v.Edge, matrix.LineMatrix{e1[segIndex1], e1[segIndex1+1]}); markInter {
					mono = v
					edge = mono.Edge
					mono.Context = false
				}
			}
		} else if monoNum == 1 {
			edge = mono.Edge
			mono.Context = false
		} else {
			edge = ii.Edge
			mono = &MonotoneChain{Edge: edge, Start: 0, End: len(edge) - 1, Context: false}
			ii.mono = append(ii.mono, mono)
		}

		if ips[0].IsCollinear {
			sort.Sort(ips)
		}
		if matrix.Matrix(edge[0]).Equals(ips[0].Matrix) {
			line = append(line, edge[:segIndex0-mono.Start]...)
		} else {
			line = append(line, edge[:segIndex0+1-mono.Start]...)
			line = append(line, ips[0].Matrix)
		}
		if len(line) > 1 {
			ii.mono = append(ii.mono, &MonotoneChain{Edge: line, Start: mono.Start, End: segIndex0 + 1, Context: true})
		}
		if matrix.Matrix(edge[len(edge)-1]).Equals(ips[len(ips)-1].Matrix) {
			startPos = segIndex0 + 2
		} else {
			startPos = segIndex0 + 1
		}
		line = matrix.LineMatrix{}
		line = append(line, ips[len(ips)-1].Matrix)
		line = append(line, edge[startPos-mono.Start:]...)
		if len(line) > 1 {
			ii.mono = append(ii.mono, &MonotoneChain{Edge: line, Start: startPos - 1, End: mono.End, Context: true})
		}

	}
}

// IsDone Always process all intersections
func (ii *IntersectionCollinearDifference) IsDone() bool {
	return false
}

// Result returns result.
func (ii *IntersectionCollinearDifference) Result() interface{} {
	for _, v := range ii.mono {
		if v.Context.(bool) {
			ii.result = append(ii.result, v.Edge)
		}
	}
	if ii.result == nil && len(ii.mono) <= 0 {
		ii.result = append(ii.result, ii.Edge)
	}
	return ii.result
}
