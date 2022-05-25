package chain

import (
	"fmt"
	"sort"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// IntersectionCorrelation Finds intersections between line segments , and adds them.
type IntersectionCorrelation struct {
	Edge, Edge1 matrix.LineMatrix
	result0     IntersectionNodeOfLine
	result1     IntersectionNodeOfLine
	// 0 not check, 1 equals,-1 not equals
	isEquals int
}

// ProcessIntersections This method is called by clients  to process intersections for two segments being intersected.
// Note that some clients (such as <code>MonotoneChain</code>s) may optimize away
// this call for segment pairs which they have determined do not intersect
func (ii *IntersectionCorrelation) ProcessIntersections(
	e0 matrix.LineMatrix, segIndex0 int,
	e1 matrix.LineMatrix, segIndex1 int) {
	// don't bother intersecting a segment with itself

	if ii.isEquals == 1 {
		return
	}
	if ii.isEquals == 0 {
		if e0.EqualsExact(e1, calc.DefaultTolerance) {
			ii.isEquals = 1
			// if  edge is closed, return
			if e0.IsClosed() {
				return
			}
			inr0 := &IntersectionNodeResult{0, 1, relate.IntersectionPoint{Matrix: matrix.Matrix(ii.Edge[0]), IsCollinear: true},
				matrix.LineSegment{P0: ii.Edge[0], P1: ii.Edge[1]}, matrix.LineSegment{P0: ii.Edge[0], P1: ii.Edge[1]}, ii.Edge1}
			ii.result0 = append(ii.result0, inr0)
			ii.result0 = append(ii.result0, inr0)
			if !matrix.Matrix(ii.Edge[len(ii.Edge)-1]).Equals(matrix.Matrix(ii.Edge[0])) {
				inr0 = &IntersectionNodeResult{len(ii.Edge) - 2, len(ii.Edge) - 1, relate.IntersectionPoint{Matrix: matrix.Matrix(ii.Edge[len(ii.Edge)-1]), IsCollinear: true},
					matrix.LineSegment{P0: ii.Edge[0], P1: ii.Edge[len(ii.Edge)-1]}, matrix.LineSegment{P0: ii.Edge[0], P1: ii.Edge[len(ii.Edge)-1]}, ii.Edge1}
				ii.result0 = append(ii.result0, inr0)
				ii.result0 = append(ii.result0, inr0)
			}

			inr0 = &IntersectionNodeResult{0, 1, relate.IntersectionPoint{Matrix: matrix.Matrix(ii.Edge1[0]), IsCollinear: true},
				matrix.LineSegment{P0: ii.Edge1[0], P1: ii.Edge1[1]}, matrix.LineSegment{P0: ii.Edge1[0], P1: ii.Edge1[1]}, ii.Edge}
			ii.result1 = append(ii.result1, inr0)
			ii.result1 = append(ii.result1, inr0)
			if !matrix.Matrix(ii.Edge[len(ii.Edge)-1]).Equals(matrix.Matrix(ii.Edge[0])) {
				inr0 = &IntersectionNodeResult{len(ii.Edge1) - 2, len(ii.Edge1) - 1, relate.IntersectionPoint{Matrix: matrix.Matrix(ii.Edge1[len(ii.Edge1)-1]), IsCollinear: true},
					matrix.LineSegment{P0: ii.Edge1[0], P1: ii.Edge1[len(ii.Edge1)-1]}, matrix.LineSegment{P0: ii.Edge1[0], P1: ii.Edge1[len(ii.Edge1)-1]}, ii.Edge}
				ii.result1 = append(ii.result1, inr0)
				ii.result1 = append(ii.result1, inr0)
			}
			return
		}
		ii.isEquals = -1

	}
	if segIndex0 > len(e0)-1 || segIndex1 > len(e1)-1 {
		return
	}

	mark, ips := relate.Intersection(e0[segIndex0], e0[segIndex0+1], e1[segIndex1], e1[segIndex1+1])

	if mark {
		if tes, _ := (matrix.Matrix(e0[segIndex0])).Compare(matrix.Matrix(e0[segIndex0+1])); tes > 0 {
			sort.Sort(ips)
		} else {
			sort.Sort(sort.Reverse(ips))
		}

		for _, ip := range ips {
			inr0 := &IntersectionNodeResult{segIndex0, segIndex0 + 1, ip,
				matrix.LineSegment{P0: e0[segIndex0], P1: e0[segIndex0+1]},
				matrix.LineSegment{P0: e1[segIndex1], P1: e1[segIndex1+1]},
				ii.Edge1,
			}
			inr1 := &IntersectionNodeResult{segIndex1, segIndex1 + 1, ip,
				matrix.LineSegment{P0: e1[segIndex1], P1: e1[segIndex1+1]},
				matrix.LineSegment{P0: e0[segIndex0], P1: e0[segIndex0+1]},
				ii.Edge,
			}

			ii.result0 = append(ii.result0, inr0)
			ii.result1 = append(ii.result1, inr1)
		}
	}
}

// IsDone Always process all intersections
func (ii *IntersectionCorrelation) IsDone() bool {
	return ii.isEquals == 1
}

// Result returns result.
func (ii *IntersectionCorrelation) Result() interface{} {
	sort.Sort(ii.result0)
	sort.Sort(ii.result1)

	return []IntersectionNodeOfLine{ii.result0, ii.result1}
}

// CorrelationNodeResult ...
type CorrelationNodeResult [][]*IntersectionCorrelationNode

// IntersectionCorrelationNode ...
type IntersectionCorrelationNode struct {
	//Pos             int
	InterNode       matrix.Matrix
	CorrelationNode matrix.LineMatrix
}

func (ic *IntersectionCorrelationNode) String() string {
	return fmt.Sprintf("CorrelationNode{InterNode:%v,CorrelationNode:%v}\n", ic.InterNode, ic.CorrelationNode)

}

// IntersectionNodeResult ...
type IntersectionNodeResult struct {
	Pos       int
	End       int
	InterNode relate.IntersectionPoint
	Line      matrix.LineSegment
	OtherLine matrix.LineSegment
	OtherEdge matrix.LineMatrix
}

// IntersectionNodeOfLine overlay point array.
type IntersectionNodeOfLine []*IntersectionNodeResult

// Len ...
func (ipl IntersectionNodeOfLine) Len() int {
	return len(ipl)
}

// Less ...
func (ipl IntersectionNodeOfLine) Less(i, j int) bool {
	if ipl[i].Pos == ipl[j].Pos {
		line := ipl[i].Line
		if line.P0 == nil {
			return false
		}
		if tes, err := line.P0.Compare(line.P1); err == nil {
			return ipl[i].InterNode.Compare(&ipl[j].InterNode, tes)
		}
	}
	return ipl[i].Pos < ipl[j].Pos
}

// Swap ...
func (ipl IntersectionNodeOfLine) Swap(i, j int) {
	ipl[i], ipl[j] = ipl[j], ipl[i]
}
