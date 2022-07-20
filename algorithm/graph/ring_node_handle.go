package graph

import "github.com/spatial-go/geoos/algorithm/matrix"

// RingNodeHandle handle graph with m1 and m2.
func RingNodeHandle(m1, m2 matrix.PolygonMatrix, g1, g2 Graph) {

	if gIntersection, err := g1.Intersection(g2); err != nil {
		lNum := 0
		for _, n := range gIntersection.Nodes() {
			if n.NodeType == LNode {
				lNum++
			}
		}
		if lNum > 1 {
			startPointHandle(m1, g1)
			startPointHandle(m2, g2)
		}
	}
}

func startPointHandle(m matrix.PolygonMatrix, g Graph) {
	for _, line := range m {
		start := &Node{Value: matrix.Matrix(line[0])}
		if startNode, ok := g.Node(start); ok {
			lastLine := matrix.LineMatrix{}
			for _, v := range g.Nodes() {
				if l, ok := v.Value.(matrix.LineMatrix); ok {
					if len(l) > 0 {
						if matrix.Matrix(line[0]).Equals(matrix.Matrix(l[len(l)-1])) &&
							!matrix.Matrix(lastLine[len(lastLine)-1]).Equals(matrix.Matrix(l[0])) {
							g.AddEdge(startNode, v)
						}
						lastLine = l
					}
				}
			}

			maxPointNode := 0
			lineNodeIndexs := []int{}
			for k, v := range g.Edges()[startNode.Index] {
				if v == PointLine {
					maxPointNode++
					lineNodeIndexs = append(lineNodeIndexs, k)
				}
			}
			if maxPointNode == 2 {
				g.DeleteNode(startNode)

				lineNodes := []*Node{g.Nodes()[lineNodeIndexs[0]],
					g.Nodes()[lineNodeIndexs[1]]}

				newLineMatrix := matrix.LineMatrix{}
				if matrix.Matrix(lineNodes[0].Value.(matrix.LineMatrix)[0]).Equals(matrix.Matrix(line[0])) {
					newLineMatrix = append(newLineMatrix, lineNodes[1].Value.(matrix.LineMatrix)...)
					newLineMatrix = append(newLineMatrix, lineNodes[0].Value.(matrix.LineMatrix)[1:]...)
				} else {
					newLineMatrix = append(newLineMatrix, lineNodes[0].Value.(matrix.LineMatrix)...)
					newLineMatrix = append(newLineMatrix, lineNodes[1].Value.(matrix.LineMatrix)[1:]...)
				}
				node := &Node{Value: newLineMatrix, NodeType: LNode}
				g.AddNode(node)

				for i, index := range lineNodeIndexs {
					for k := range g.Edges()[index] {
						nodeLink := g.Nodes()[k]
						g.DeleteEdge(lineNodes[i], nodeLink)
						g.AddEdge(nodeLink, node)
					}
				}

			}
		}
	}
}
