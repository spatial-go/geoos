package graph

import (
	"log"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// GenerateGraph create graph with matrix.
func GenerateGraph(m matrix.Steric) (Graph, error) {
	g := &MatrixGraph{}
	switch mType := m.(type) {
	case matrix.Matrix:
		g.AddNode(&Node{Value: mType, NodeType: PNode})
	case matrix.LineMatrix:
		return lineCreateGraph(mType)
	case matrix.PolygonMatrix:
		return polyCreateGraph(mType)
	case matrix.Collection:
		return collCreateGraph(mType)
	default:
		return nil, algorithm.ErrNotMatchType

	}
	return g, nil
}

// GenerateGraphCollection create graph with collection.
func GenerateGraphCollection(m matrix.Steric) ([]Graph, error) {

	switch mType := m.(type) {
	case matrix.Collection:
		gc := make([]Graph, len(mType))
		for i, v := range mType {
			if g, err := GenerateGraph(v); err == nil {
				gc[i] = g
			}
		}
	default:
		return nil, algorithm.ErrNotGraph

	}
	return nil, algorithm.ErrNotMatchType
}

// lineCreateGraph create graph with line.
func lineCreateGraph(m matrix.LineMatrix) (Graph, error) {
	g := &MatrixGraph{}

	node := &Node{Value: m, NodeType: LNode}
	g.AddNode(node)

	return g, nil
}

// polyCreateGraph create graph with polygon.
func polyCreateGraph(m matrix.PolygonMatrix) (Graph, error) {
	g := &MatrixGraph{}

	g.nodes = make([]*Node, 0, len(m)/2)
	g.edges = make([]map[int]int, 0, len(m)/2)

	node := &Node{Value: m, NodeType: ANode}
	g.AddNode(node)

	return g, nil
}

// collCreateGraph create graph with polygon.
func collCreateGraph(coll matrix.Collection) (Graph, error) {
	g := &MatrixGraph{}
	for _, m := range coll {
		var nodeType int
		switch m.(type) {
		case matrix.Matrix:
			nodeType = PNode
		case matrix.LineMatrix:
			nodeType = CNode
		case matrix.PolygonMatrix:
			nodeType = ANode
		default:
			nodeType = ANode
		}
		node := &Node{Value: m, NodeType: nodeType}
		g.AddNode(node)

	}
	return g, nil
}

// IntersectionHandle handle graph with m1 and m2.
func IntersectionHandle(m1, m2 matrix.Steric, g1, g2 Graph) error {
	switch m1Type := m1.(type) {
	case matrix.Matrix:
		return matrixIntersectionHandle(m1Type, m2, g1, g2)
	case matrix.LineMatrix:
		return lineIntersectionHandle(m1Type, m2, g1, g2)
	case matrix.PolygonMatrix:
		return polyIntersectionHandle(m1Type, m2, g1, g2)
	case matrix.Collection:
		return collIntersectionHandle(m1Type, m2, g1, g2)
	default:
		return algorithm.ErrNotMatchType

	}
}

// polyIntersectionHandle handle graph with m1 polygon and m2.
func polyIntersectionHandle(m1 matrix.PolygonMatrix, m2 matrix.Steric, g1, g2 Graph) error {

	switch m2Type := m2.(type) {
	case matrix.Matrix:
		for _, l := range m1 {
			if err := matrixAndLineHandle(m2Type, l, g2, g1); err != nil {
				log.Println(err)
			}
		}
		return nil
	case matrix.LineMatrix:
		return lineAndPolygonHandle(m2Type, m1, g2, g1)
	case matrix.PolygonMatrix:
		return polygonAndPolygonHandle(m1, m2Type, g1, g2)
	case matrix.Collection:
		return collIntersectionHandle(m2Type, m1, g2, g1)
	default:
		return algorithm.ErrNotMatchType
	}
}

// collIntersectionHandle handle graph with m1 polygon and m2.
func collIntersectionHandle(m1 matrix.Collection, m2 matrix.Steric, g1, g2 Graph) error {
	for _, m := range m1 {
		switch mType := m.(type) {
		case matrix.Matrix:
			_ = matrixIntersectionHandle(mType, m2, g1, g2)
		case matrix.LineMatrix:
			_ = lineIntersectionHandle(mType, m2, g1, g2)
		case matrix.PolygonMatrix:
			_ = polyIntersectionHandle(mType, m2, g1, g2)
		}
	}
	return nil
}

// lineIntersectionHandle handle graph with m1 line and m2.
func lineIntersectionHandle(m1 matrix.LineMatrix, m2 matrix.Steric, g1, g2 Graph) error {

	switch m2Type := m2.(type) {
	case matrix.Matrix:
		return matrixAndLineHandle(m2Type, m1, g2, g1)
	case matrix.LineMatrix:
		return lineAndLineHandle(m1, m2Type, g1, g2)
	case matrix.PolygonMatrix:
		return lineAndPolygonHandle(m1, m2Type, g1, g2)
	case matrix.Collection:
		return collIntersectionHandle(m2Type, m1, g2, g1)
	default:
		return algorithm.ErrNotMatchType
	}
}

// matrixIntersectionHandle handle graph with m1 matrix and m2.
func matrixIntersectionHandle(m1 matrix.Matrix, m2 matrix.Steric, g1, g2 Graph) error {

	switch m2Type := m2.(type) {
	case matrix.Matrix:
		return nil
	case matrix.LineMatrix:
		return matrixAndLineHandle(m1, m2Type, g1, g2)
	case matrix.PolygonMatrix:
		for _, l := range m2Type {
			if err := matrixAndLineHandle(m1, l, g1, g2); err != nil {
				log.Println(err)
			}
		}
		return nil
	case matrix.Collection:
		return collIntersectionHandle(m2Type, m1, g2, g1)
	default:
		return algorithm.ErrNotMatchType
	}
}

// matrixIntersectionHandle handle graph with m1 and m2.
func matrixAndLineHandle(m1 matrix.Matrix, m2 matrix.LineMatrix, _, g2 Graph) error {

	if m1.Equals(matrix.Matrix(m2[0])) || m1.Equals(matrix.Matrix(m2[len(m2)-1])) {
		node := &Node{Value: m1, NodeType: PNode}
		g2.AddNode(node)
		g2.AddEdge(g2.Nodes()[0], node)
		return nil
	}
	lines := m2.ToLineArray()
	line := matrix.LineMatrix{}
	inLine := false
	var nodeIntersect *Node
	for _, ls := range lines {

		line = append(line, ls.P0)
		if m1.Equals(ls.P0) {
			nodeIntersect = &Node{Value: m1, NodeType: PNode}
			g2.AddNode(nodeIntersect)
			node := &Node{Value: line, NodeType: LNode}
			g2.AddNode(node)
			inLine = true
			g2.AddEdge(nodeIntersect, node)
			line = matrix.LineMatrix{}
			line = append(line, m1)
		} else if m1.Equals(ls.P1) {
			continue
		} else {
			if inLine, _ = relate.InLine(m1, ls.P0, ls.P1); inLine {
				line = append(line, m1)
				node := &Node{Value: line, NodeType: LNode}
				g2.AddNode(node)
				nodeIntersect = &Node{Value: m1, NodeType: PNode}
				g2.AddNode(nodeIntersect)

				g2.AddEdge(nodeIntersect, node)
				line = matrix.LineMatrix{}
				line = append(line, m1)
			}
		}
	}
	if inLine {
		line = append(line, m2[len(m2)-1])
		node := &Node{Value: line, NodeType: LNode}
		g2.AddNode(node)
		g2.AddEdge(nodeIntersect, node)
	}
	return nil
}

// lineAndLineHandle handle graph with m1 and m2.
func lineAndLineHandle(m1, m2 matrix.LineMatrix, g1, g2 Graph) error {
	g := []Graph{g1, g2}
	corrNodes := IntersectLine(m1, m2)
	if len(corrNodes[0]) > 0 || len(corrNodes[1]) > 0 {
		g1.Nodes()[0].Stat = false
		g2.Nodes()[0].Stat = false
	}
	for i, corrs := range corrNodes {
		for _, corr := range corrs {
			node := &Node{Value: corr.InterNode, NodeType: PNode}

			nodeLine := &Node{Value: corr.CorrelationNode, NodeType: LNode}

			g[i].AddNode(node)
			g[i].AddNode(nodeLine)
			g[i].AddEdge(node, nodeLine)
		}
	}
	return nil
}

// lineAndPolygonHandle handle graph with m1 and m2.
func lineAndPolygonHandle(m1 matrix.LineMatrix, m2 matrix.PolygonMatrix, g1, g2 Graph) error {

	corrNodes := IntersectLinePolygon(m1, m2)
	isHandle := map[Graph]bool{}
	for i, corrs := range corrNodes {
		gNum := g1
		if i > 0 {
			gNum = g2
		}
		if len(corrs) > 0 && !isHandle[gNum] {
			gNum.Nodes()[0].Stat = false
			isHandle[gNum] = true
		}
		startNode, endNode := &Node{}, &Node{}
		startNodeLine, endNodeLine := &Node{}, &Node{}
		for j, corr := range corrs {

			node := &Node{Value: corr.InterNode, NodeType: PNode}
			nodeLine := &Node{Value: corr.CorrelationNode, NodeType: LNode}

			if j == 0 {
				startNode = node
				startNodeLine = nodeLine
			}
			if j == len(corrs)-1 {
				endNode = node
				endNodeLine = nodeLine
			}

			gNum.AddNode(node)
			gNum.AddNode(nodeLine)
			gNum.AddEdge(node, nodeLine)
		}
		if i == 0 && m1.IsClosed() {
			if startNode.Value != nil && startNode.Value.EqualsExact(matrix.Matrix(m1[0]), calc.DefaultTolerance) {
				gNum.AddEdge(startNode, endNodeLine)
			}
			if endNode.Value != nil && endNode.Value.EqualsExact(matrix.Matrix(m1[len(m1)-1]), calc.DefaultTolerance) {
				gNum.AddEdge(endNode, startNodeLine)
			}
		}
		if i > 0 {
			if startNode.Value != nil && startNode.Value.EqualsExact(matrix.Matrix(m2[i-1][0]), calc.DefaultTolerance) {
				gNum.AddEdge(startNode, endNodeLine)
			}
			if endNode.Value != nil && endNode.Value.EqualsExact(matrix.Matrix(m2[i-1][len(m2[i-1])-1]), calc.DefaultTolerance) {
				gNum.AddEdge(endNode, startNodeLine)
			}
		}
	}

	if m1.IsClosed() {
		RingNodeHandle(matrix.PolygonMatrix{m1}, m2, g1, g2)
	}

	return nil
}

// polygonAndPolygonHandle handle graph with m1 and m2.
func polygonAndPolygonHandle(m1, m2 matrix.PolygonMatrix, g1, g2 Graph) error {
	g := []Graph{g1, g2}
	twoCorrNodes := IntersectPolygons(m1, m2)
	for i, corrNodes := range twoCorrNodes {
		if len(corrNodes) > 0 && len(corrNodes[0]) > 0 {
			g[i].Nodes()[0].Stat = false
		}
		for k, corrs := range corrNodes {
			startNode, endNode := &Node{}, &Node{}
			startNodeLine, endNodeLine := &Node{}, &Node{}
			for j, corr := range corrs {
				node := &Node{Value: corr.InterNode, NodeType: PNode}

				nodeLine := &Node{Value: corr.CorrelationNode, NodeType: LNode}
				if j == 0 {
					startNode = node
					startNodeLine = nodeLine
				}
				if j == len(corrs)-1 {
					endNode = node
					endNodeLine = nodeLine
				}

				g[i].AddNode(node)
				if nodeLine.Value.Nums() != 0 {
					g[i].AddNode(nodeLine)
					g[i].AddEdge(node, nodeLine)
				}
			}
			var m matrix.PolygonMatrix
			if i == 0 {
				m = m1
			} else {
				m = m2
			}
			if startNode.Value != nil && startNode.Value.EqualsExact(matrix.Matrix(m[k][0]), calc.DefaultTolerance) {
				g[i].AddEdge(startNode, endNodeLine)
			}
			if endNode.Value != nil && endNode.Value.EqualsExact(matrix.Matrix(m[k][len(m[k])-1]), calc.DefaultTolerance) {
				g[i].AddEdge(endNode, startNodeLine)
			}

		}
	}

	RingNodeHandle(m1, m2, g1, g2)
	return nil
}
