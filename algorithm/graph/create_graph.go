// package graph ...
package graph

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// CreateGraph create graph with matrix.
func CreateGraph(m matrix.Steric) (Graph, error) {
	g := &MatrixGraph{}
	switch mType := m.(type) {
	case matrix.Matrix:
		g.AddNode(&Node{0, mType, true, false, PNode})
	case matrix.LineMatrix:
		return lineCreateGraph(mType)
	case matrix.PolygonMatrix:
		return polyCreateGraph(mType)
	case matrix.Collection:
		return nil, algorithm.ErrNotGraphCollection
	default:
		return nil, algorithm.ErrNotMatchType

	}
	return nil, algorithm.ErrNotMatchType
}

// CreateGraphCollection create graph with matrix.
func CreateGraphCollection(m matrix.Steric) ([]Graph, error) {

	switch mType := m.(type) {
	case matrix.Collection:
		gc := []Graph{}
		for _, v := range mType {
			if g, err := CreateGraph(v); err == nil {
				gc = append(gc, g)
			}
		}
	default:
		return nil, algorithm.ErrNotGraph

	}
	return nil, algorithm.ErrNotMatchType
}

// lineCreateGraph create graph with matrix.
func lineCreateGraph(m matrix.LineMatrix) (Graph, error) {
	g := &MatrixGraph{}
	node := &Node{0, m, true, false, CNode}
	g.AddNode(node)

	return g, nil
}

// polyCreateGraph create graph with matrix.
func polyCreateGraph(m matrix.PolygonMatrix) (Graph, error) {
	g := &MatrixGraph{}

	node := &Node{0, m, true, false, ANode}
	g.AddNode(node)

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
		return algorithm.ErrNotGraphCollection
	default:
		return algorithm.ErrNotMatchType

	}
}

// polyIntersectionHandle handle graph with m1 and m2.
func polyIntersectionHandle(m1 matrix.PolygonMatrix, m2 matrix.Steric, g1, g2 Graph) error {
	for _, l := range m1 {
		lineIntersectionHandle(l, m2, g1, g2)
	}
	return nil
}

// lineIntersectionHandle handle graph with m1 and m2.
func lineIntersectionHandle(m1 matrix.LineMatrix, m2 matrix.Steric, g1, g2 Graph) error {

	switch m2Type := m2.(type) {
	case matrix.Matrix:
		return matrixAndLineHandle(m2Type, m1, g2, g1)
	case matrix.LineMatrix:
		return lineAndLineHandle(m1, m2Type, g1, g2)
	case matrix.PolygonMatrix:
		for _, l := range m2Type {
			lineAndLineHandle(m1, l, g1, g2)
		}
		return nil
	case matrix.Collection:
		return algorithm.ErrNotGraphCollection
	default:
		return algorithm.ErrNotMatchType
	}
}

// matrixIntersectionHandle handle graph with m1 and m2.
func matrixIntersectionHandle(m1 matrix.Matrix, m2 matrix.Steric, g1, g2 Graph) error {

	switch m2Type := m2.(type) {
	case matrix.Matrix:
		return nil
	case matrix.LineMatrix:
		return matrixAndLineHandle(m1, m2Type, g1, g2)
	case matrix.PolygonMatrix:
		for _, l := range m2Type {
			matrixAndLineHandle(m1, l, g1, g2)
		}
		return nil
	case matrix.Collection:
		return algorithm.ErrNotGraphCollection
	default:
		return algorithm.ErrNotMatchType
	}
}

// matrixIntersectionHandle handle graph with m1 and m2.
func matrixAndLineHandle(m1 matrix.Matrix, m2 matrix.LineMatrix, g1, g2 Graph) error {
	lines := m2.ToLineArray()
	line := matrix.LineMatrix{}
	inLine := false
	var nodeIntersect *Node
	for _, ls := range lines {
		line = append(line, ls.P0)
		if inLine, _ = relate.InLine(m1, ls.P0, ls.P1); inLine {
			line = append(line, m1)
			node := &Node{0, line, true, false, LNode}
			g2.AddNode(node)
			nodeIntersect = &Node{0, m1, true, false, PNode}
			g2.AddNode(nodeIntersect)

			g2.AddEdge(nodeIntersect, node)
			line = matrix.LineMatrix{}
			line = append(line, m1)
		}
	}
	if inLine {
		line = append(line, m2[len(m2)-1])
		node := &Node{0, line, true, false, LNode}
		g2.AddNode(node)
		g2.AddEdge(nodeIntersect, node)
	}
	return nil
}

// lineAndLineHandle handle graph with m1 and m2.
func lineAndLineHandle(m1, m2 matrix.LineMatrix, g1, g2 Graph) error {
	// mark = false
	// for i := range m1 {
	// 	for j := range m2 {
	// 		if i < len(aLine)-1 && j < len(bLine)-1 {
	// 			markInter, ips := Intersection(matrix.Matrix(aLine[i]),
	// 				matrix.Matrix(aLine[i+1]),
	// 				matrix.Matrix(bLine[j]),
	// 				matrix.Matrix(bLine[j+1]))
	// 			if markInter {
	// 				mark = markInter
	// 				ps = append(ps, ips...)
	// 			}
	// 		}
	// 	}
	// }
	// filt := &UniqueIntersectionEdgeFilter{}
	// for _, v := range ps {
	// 	filt.Filter(v)
	// }
	// ps = filt.Ips
	// return

	lines := m1.ToLineArray()
	line := matrix.LineMatrix{}
	inLine := false
	var nodeIntersect *Node
	for _, ls := range lines {
		line = append(line, ls.P0)
		// todo
		// if inLine, _ = relate.InLine(m1, ls.P0, ls.P1); inLine {
		// 	line = append(line, m1)
		// 	node := &Node{0, line, true, false, LNode}
		// 	g2.AddNode(node)
		// 	nodeIntersect = &Node{0, m1, true, false, PNode}
		// 	g2.AddNode(nodeIntersect)

		// 	g2.AddEdge(nodeIntersect, node)
		// 	line = matrix.LineMatrix{}
		// 	line = append(line, m1)
		// }
	}
	if inLine {
		line = append(line, m2[len(m2)-1])
		node := &Node{0, line, true, false, LNode}
		g2.AddNode(node)
		g2.AddEdge(nodeIntersect, node)
	}
	return nil
}

// IntersectLine returns a Collection  that represents that part of geometry A intersect with geometry B.
func IntersectLine(m1, m2 matrix.LineMatrix) chain.CorrelationNodeResult {
	smi := &chain.SegmentMutualIntersector{SegmentMutual: m1}
	icd := &chain.IntersectionCorrelation{Edge: m1, Edge1: m2}
	smi.Process(m2, icd)
	result := icd.Result()
	return result.(chain.CorrelationNodeResult)
}

// lineAndPolygonHandle handle graph with m1 and m2.
func lineAndPolygonHandle(m1 matrix.LineMatrix, m2 matrix.PolygonMatrix, g1, g2 Graph) error {
	// TODO
	// lines := m2.ToLineArray()
	// line := matrix.LineMatrix{}
	// inLine := false
	// var nodeIntersect *Node
	// for _, ls := range lines {
	// 	line = append(line, ls.P0)
	// 	// todo
	// 	// if inLine, _ = relate.InLine(m1, ls.P0, ls.P1); inLine {
	// 	// 	line = append(line, m1)
	// 	// 	node := &Node{0, line, true, false, LNode}
	// 	// 	g2.AddNode(node)
	// 	// 	nodeIntersect = &Node{0, m1, true, false, PNode}
	// 	// 	g2.AddNode(nodeIntersect)

	// 	// 	g2.AddEdge(nodeIntersect, node)
	// 	// 	line = matrix.LineMatrix{}
	// 	// 	line = append(line, m1)
	// 	// }
	// }
	// if inLine {
	// 	line = append(line, m2[len(m2)-1])
	// 	node := &Node{0, line, true, false, LNode}
	// 	g2.AddNode(node)
	// 	g2.AddEdge(nodeIntersect, node)
	// }
	return nil
}
