package quadtree

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Root QuadRoot is the root of a single Quadtree.  It is centred at the origin,
// and does not have a defined extent.
type Root struct {
	*Node
	origin matrix.Matrix
}

// Insert an item into the quadtree this is the root of.
func (r *Root) Insert(itemEnv *matrix.Envelope, item interface{}) {

	index := SubnodeIndex(itemEnv, 0.0, 0.0)
	// if index is -1, itemEnv must cross the X or Y axis.
	if index == -1 {
		r.Add(item)
		return
	}

	node := r.Subnode[index]

	if node == nil || !node.Env.Contains(itemEnv) {
		largerNode := CreateExpanded(node, itemEnv)
		r.Subnode[index] = largerNode
	}

	r.InsertContained(r.Subnode[index], itemEnv, item)
}

// InsertContained insert an item which is known to be contained in the tree rooted at
//  the given QuadNode root.  Lower levels of the tree will be created if necessary to hold the item.
func (r *Root) InsertContained(tree *Node, itemEnv *matrix.Envelope, item interface{}) {
	isZeroX := IsZeroWidth(itemEnv.MinX, itemEnv.MaxX)
	isZeroY := IsZeroWidth(itemEnv.MinY, itemEnv.MaxY)
	var node *Node
	if isZeroX || isZeroY {
		node = tree.Find(itemEnv)
	} else {
		node = tree.GetNode(itemEnv)
	}
	node.Add(item)
}

// IsSearchMatch ...
func (r *Root) IsSearchMatch(searchEnv *matrix.Envelope) bool {
	return true
}
