package graph

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
)

// Merge ...
type Merge struct {
	Arg      []matrix.Steric // the arg(s) of the operation
	ArgGraph []Graph
	IsDesc   bool
}

// MergeHandle handle graph with m1 and m2,returns graph of intersection , union, difference and sym  difference.
func MergeHandle(m0, m1 matrix.Steric) *Merge {
	arg := []matrix.Steric{m0, m1}

	merge := &Merge{Arg: arg, ArgGraph: make([]Graph, 2)}

	for i, v := range merge.Arg {
		merge.ArgGraph[i], _ = GenerateGraph(v)
	}

	if err := mergeHandle(merge.Arg[0], merge.Arg[1], merge.ArgGraph[0], merge.ArgGraph[1]); err != nil {
		return nil
	}

	return merge
}

// Intersection  Computes the Intersection of two Graph.
func (c *Merge) Intersection() (Graph, error) {
	return c.ArgGraph[0].Intersection(c.ArgGraph[1])
}

// Union  Computes the Union of two Graph.
func (c *Merge) Union() (Graph, error) {
	return c.ArgGraph[0].Union(c.ArgGraph[1])
}

// Difference returns a Graph that represents that part of Graph A that does not intersect with Graph B.
// One can think of this as GraphA - Intersection(A,B).
func (c *Merge) Difference() (Graph, error) {
	return c.ArgGraph[0].Difference(c.ArgGraph[1])
}

// SymDifference returns a Graph that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
//
// One can think of this as Union(A,B) - Intersection(A,B).
func (c *Merge) SymDifference() (Graph, error) {
	return c.ArgGraph[0].SymDifference(c.ArgGraph[1])
}

// mergeHandle handle graph with m1 and m2.
func mergeHandle(m1, m2 matrix.Steric, g1, g2 Graph) error {
	switch mType := m1.(type) {
	case matrix.LineMatrix:
		_ = linemergeHandle(mType, m2, g1, g2)
	case matrix.Collection:
		_ = collMergeHandle(mType, m2, g1, g2)
	default:
		for _, g := range g1.Nodes() {
			g1.DeleteNode(g)
		}
	}

	return nil

}

// linemergeHandle handle graph with m1 and m2.
func linemergeHandle(m1 matrix.LineMatrix, m2 matrix.Steric, g1, g2 Graph) error {
	switch mType := m2.(type) {
	case matrix.LineMatrix:
		_ = lineMergeLine(m1, mType, g1, g2)
	case matrix.Collection:
		_ = collMergeHandle(mType, m1, g2, g1)
	default:
		for _, g := range g2.Nodes() {
			g1.DeleteNode(g)
		}
	}
	return nil
}

// collMergeHandle handle graph with m1 polygon and m2.
func collMergeHandle(m1 matrix.Collection, m2 matrix.Steric, g1, g2 Graph) error {
	for _, m := range m1 {
		switch mType := m.(type) {
		case matrix.LineMatrix:
			_ = linemergeHandle(mType, m2, g1, g2)
		}
	}
	return nil
}

// lineMergeLine handle graph with m1 and m2.
func lineMergeLine(m1, m2 matrix.LineMatrix, g1, g2 Graph) error {
	g := []Graph{g1, g2}
	corrNodes := mergeLine(m1, m2)
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

// mergeLine returns a Collection  that represents that part of geometry A intersect with geometry B.
func mergeLine(m1, m2 matrix.LineMatrix) chain.CorrelationNodeResult {
	smi := &chain.SegmentMutualIntersector{SegmentMutual: m1}
	icd := &chain.MergeCorrelation{Edge: m1, Edge1: m2}
	smi.Process(m2, icd)
	inols := icd.Result()

	m := []matrix.LineMatrix{m1, m2}
	if inolsType, ok := inols.([]chain.IntersectionNodeOfLine); ok {
		result := make(chain.CorrelationNodeResult, len(inolsType))
		for i, inol := range inolsType {
			inol = sortIntersectionNode(inol)
			resultCorr := createCorrelationNode(inol, m[i])
			result[i] = resultCorr
		}
		return result
	}
	return nil
}
