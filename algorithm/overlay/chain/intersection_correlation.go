package chain

import (
	"fmt"
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// IntersectionCorrelation Finds intersections between line segments , and adds them.
type IntersectionCorrelation struct {
	Edge, Edge1 matrix.LineMatrix
	result0     IntersectionNodeOfLine
	result1     IntersectionNodeOfLine
}

// ProcessIntersections This method is called by clients  to process intersections for two segments being intersected.
// Note that some clients (such as <code>MonotoneChain</code>s) may optimize away
// this call for segment pairs which they have determined do not intersect
func (ii *IntersectionCorrelation) ProcessIntersections(
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
		if tes, _ := (matrix.Matrix(e0[segIndex0])).Compare(matrix.Matrix(e0[segIndex0+1])); tes > 0 {
			sort.Sort(ips)
		} else {
			sort.Sort(sort.Reverse(ips))
		}

		for _, ip := range ips {
			inr0 := &IntersectionNodeResult{segIndex0, segIndex0 + 1, ip.Matrix}
			inr1 := &IntersectionNodeResult{segIndex1, segIndex1 + 1, ip.Matrix}

			ii.result0 = append(ii.result0, inr0)
			ii.result1 = append(ii.result1, inr1)
		}
	}
}

// IsDone Always process all intersections
func (ii *IntersectionCorrelation) IsDone() bool {
	if ii.Edge.Equals(ii.Edge1) {
		inr0 := &IntersectionNodeResult{0, 1, ii.Edge[0]}
		ii.result0 = append(ii.result0, inr0)
		inr0 = &IntersectionNodeResult{len(ii.Edge) - 2, len(ii.Edge) - 1, ii.Edge[len(ii.Edge)-1]}
		ii.result0 = append(ii.result0, inr0)

		inr0 = &IntersectionNodeResult{0, 1, ii.Edge1[0]}
		ii.result1 = append(ii.result1, inr0)
		inr0 = &IntersectionNodeResult{len(ii.Edge1) - 2, len(ii.Edge1) - 1, ii.Edge1[len(ii.Edge1)-1]}
		ii.result1 = append(ii.result1, inr0)

		return true

	}
	return false
}

// Result returns result.
func (ii *IntersectionCorrelation) Result() interface{} {
	sort.Sort(ii.result0)
	sort.Sort(ii.result1)
	resultCorr0 := []*IntersectionCorrelationNode{}
	resultCorr1 := []*IntersectionCorrelationNode{}
	lines0 := []matrix.LineMatrix{}
	lines1 := []matrix.LineMatrix{}
	totalIps := 0
	ips := []IntersectionNodeOfLine{{}, {}}
	pos := [2]int{0, 0}
	pos1 := [2]int{0, 0}

	correlationNode0 := matrix.LineMatrix{}
	correlationNode1 := matrix.LineMatrix{}

	for i := 0; i < len(ii.result0); i++ {
		r0 := ii.result0[i]
		r1 := ii.result1[i]

		if i == 0 {
			correlationNode := matrix.LineMatrix{}
			for j := 0; j <= r0.Pos; j++ {
				correlationNode = append(correlationNode, ii.Edge[j])
			}
			correlationNode = append(correlationNode, r0.InterNode)

			lines0 = append(lines0, correlationNode)

			correlationNode = matrix.LineMatrix{}

			for j := 0; j <= r1.Pos; j++ {
				correlationNode = append(correlationNode, ii.Edge1[j])
			}
			correlationNode = append(correlationNode, r1.InterNode)

			lines1 = append(lines1, correlationNode)

			correlationNode = matrix.LineMatrix{}
			totalIps++
			ips[0] = append(ips[0], r0)
			ips[1] = append(ips[1], r1)
		}

		if i == len(ii.result0)-1 {
			correlationNode := matrix.LineMatrix{}
			correlationNode = append(correlationNode, r0.InterNode)
			for j := r0.End; j < len(ii.Edge); j++ {
				correlationNode = append(correlationNode, ii.Edge[j])
			}
			lines0 = append(lines0, correlationNode)
			correlationNode = matrix.LineMatrix{}
			correlationNode = append(correlationNode, r1.InterNode)
			for j := r1.End; j < len(ii.Edge1); j++ {
				correlationNode = append(correlationNode, ii.Edge1[j])
			}

			lines1 = append(lines1, correlationNode)
			//totalIps++
			ips[0] = append(ips[0], r0)
			ips[1] = append(ips[1], r1)

		}

		if i < len(ii.result0)-2 {
			r01 := ii.result0[i+1]
			r11 := ii.result1[i+1]
			r02 := ii.result0[i+2]
			//r12 := ii.result1[i+2]
			if r02.InterNode.Equals(r01.InterNode) {
				if pos[0] == 0 {
					pos[0] = r0.End
					pos[1] = r1.End
					correlationNode0 = append(correlationNode0, r0.InterNode)
					correlationNode1 = append(correlationNode1, r1.InterNode)
				}
				if r0.Pos == r01.Pos || r1.Pos == r11.Pos {
					if i < len(ii.result0)-3 {
						i++
					}
				}
				continue
			}
		}
		if pos[0] == 0 {
			pos[0] = r0.End
			pos[1] = r1.End
			correlationNode0 = append(correlationNode0, r0.InterNode)
			correlationNode1 = append(correlationNode1, r1.InterNode)
		}
		endNode0 := r0.InterNode
		endNode1 := r1.InterNode
		pos1[0] = r0.End
		pos1[1] = r1.End
		if i < len(ii.result0)-2 {
			r01 := ii.result0[i+1]
			r11 := ii.result1[i+1]
			pos1[0] = r01.End
			pos1[1] = r11.End
			endNode0 = r01.InterNode
			endNode1 = r11.InterNode
		}

		for j := pos[0]; j < pos1[0]; j++ {
			correlationNode0 = append(correlationNode0, ii.Edge[j])
		}
		correlationNode0 = append(correlationNode0, endNode0)
		lines0 = append(lines0, correlationNode0)

		for j := pos[1]; j < pos1[1]; j++ {
			correlationNode1 = append(correlationNode1, ii.Edge1[j])
		}
		correlationNode1 = append(correlationNode1, endNode1)

		lines1 = append(lines1, correlationNode1)
		totalIps++
		ips[0] = append(ips[0], r0)
		ips[1] = append(ips[1], r1)

		correlationNode0 = matrix.LineMatrix{}
		correlationNode1 = matrix.LineMatrix{}

	}

	for i := 0; i < totalIps; i++ {

		if !matrix.Matrix(lines0[i][0]).Equals(matrix.Matrix(lines0[i][len(lines0[i])-1])) {
			resultCorr0 = append(resultCorr0,
				&IntersectionCorrelationNode{ips[0][i].Pos, ips[0][i].InterNode, lines0[i]})
		}
		if !matrix.Matrix(lines0[i+1][0]).Equals(matrix.Matrix(lines0[i+1][len(lines0[i+1])-1])) {
			resultCorr0 = append(resultCorr0,
				&IntersectionCorrelationNode{ips[0][i].Pos, ips[0][i].InterNode, lines0[i+1]})
		}

		if !matrix.Matrix(lines1[i][0]).Equals(matrix.Matrix(lines1[i][len(lines1[i])-1])) {
			resultCorr1 = append(resultCorr1,
				&IntersectionCorrelationNode{ips[1][i].Pos, ips[1][i].InterNode, lines1[i]})
		}
		if !matrix.Matrix(lines1[i+1][0]).Equals(matrix.Matrix(lines1[i+1][len(lines1[i+1])-1])) {
			resultCorr1 = append(resultCorr1,
				&IntersectionCorrelationNode{ips[1][i].Pos, ips[1][i].InterNode, lines1[i+1]})
		}
	}

	return CorrelationNodeResult{resultCorr0, resultCorr1}
}

// CorrelationNodeResult ...
type CorrelationNodeResult [2][]*IntersectionCorrelationNode

// IntersectionCorrelationNode ...
type IntersectionCorrelationNode struct {
	Pos             int
	InterNode       matrix.Matrix
	CorrelationNode matrix.LineMatrix
}

func (ic *IntersectionCorrelationNode) String() string {
	return fmt.Sprintf("CorrelationNode{Pos:%v,InterNode:%v,CorrelationNode:%v}\n", ic.Pos, ic.InterNode, ic.CorrelationNode)

}

// IntersectionNodeResult ...
type IntersectionNodeResult struct {
	Pos       int
	End       int
	InterNode matrix.Matrix
}

// IntersectionNodeOfLine overlay point array.
type IntersectionNodeOfLine []*IntersectionNodeResult

// Len ...
func (ipl IntersectionNodeOfLine) Len() int {
	return len(ipl)
}

// Less ...
func (ipl IntersectionNodeOfLine) Less(i, j int) bool {
	return ipl[i].Pos < ipl[j].Pos
}

// Swap ...
func (ipl IntersectionNodeOfLine) Swap(i, j int) {
	ipl[i], ipl[j] = ipl[j], ipl[i]
}
