package intervalrtree

import (
	"math"

	"github.com/spatial-go/geoos/index"
)

// Node define Node interface
type Node interface {
	Max() float64
	Min() float64
	Query(queryMin, queryMax float64, visitor index.ItemVisitor)
}

// IRNode define interval R-tree node.
type IRNode struct {
	min float64
	max float64
}

// Max return r-tree max.
func (in *IRNode) Max() float64 {
	return in.max
}

// Min return r-tree max.
func (in *IRNode) Min() float64 {
	return in.min
}

// Query  Search for intervals in the index which intersect the given closed interval and apply the visitor to them.
func (in *IRNode) Query(queryMin, queryMax float64, visitor index.ItemVisitor) {

}

// IsIntersects ...
func (in *IRNode) IsIntersects(queryMin, queryMax float64) bool {
	if in.Min() > queryMax || in.Max() < queryMin {
		return false
	}
	return true
}

func (in *IRNode) buildExtent(n1, n2 Node) {
	in.min = math.Min(n1.Min(), n2.Min())
	in.max = math.Max(n1.Max(), n2.Max())
}

// BranchNode define interval R-tree  branch node.
type BranchNode struct {
	*IRNode
	node1 Node
	node2 Node
}

// NewBranchNode ...
func NewBranchNode(n1, n2 Node) *BranchNode {
	it := &BranchNode{node1: n1, node2: n2, IRNode: &IRNode{}}
	it.buildExtent(n1, n2)
	return it
}

// Query  Search for intervals in the index which intersect the given closed interval and apply the visitor to them.
func (in *BranchNode) Query(queryMin, queryMax float64, visitor index.ItemVisitor) {
	if !in.IsIntersects(queryMin, queryMax) {
		return
	}
	if in.node1 != nil {
		in.node1.Query(queryMin, queryMax, visitor)
	}
	if in.node2 != nil {
		in.node2.Query(queryMin, queryMax, visitor)
	}
}

// LeafNode define interval R-tree  leaf node.
type LeafNode struct {
	*IRNode
	item interface{}
}

// Query  Search for intervals in the index which intersect the given closed interval and apply the visitor to them.
func (in *LeafNode) Query(queryMin, queryMax float64, visitor index.ItemVisitor) {
	if !in.IsIntersects(queryMin, queryMax) {
		return
	}
	visitor.VisitItem(in.item)
}
