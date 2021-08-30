package quadtree

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

// Root QuadRoot is the root of a single Quadtree.  It is centred at the origin,
// and does not have a defined extent.
type Root struct {
	*Node
	origin matrix.Matrix
}

// Insert an item into the quadtree this is the root of.
func (r *Root) Insert(itemEnv *envelope.Envelope, item interface{}) {

	index := SubnodeIndex(itemEnv, 0.0, 0.0)
	// if index is -1, itemEnv must cross the X or Y axis.
	if index == -1 {
		r.Add(item)
		return
	}

	node := r.Subnode[index]

	if node == nil || !node.Env.Contains(itemEnv) {
		largerNode := NewNodeExpanded(node, itemEnv)
		r.Subnode[index] = largerNode
	}

	r.InsertContained(r.Subnode[index], itemEnv, item)
}

// InsertContained insert an item which is known to be contained in the tree rooted at
//  the given QuadNode root.  Lower levels of the tree will be created if necessary to hold the item.
func (r *Root) InsertContained(tree *Node, itemEnv *envelope.Envelope, item interface{}) {
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
func (r *Root) IsSearchMatch(searchEnv *envelope.Envelope) bool {
	return true
}

// Remove Removes a single item from this subtree.
func (r *Root) Remove(itemEnv *envelope.Envelope, item interface{}) bool {
	// use envelope to restrict nodes scanned
	if !r.IsSearchMatch(itemEnv) {
		return false
	}
	found := false
	for i := 0; i < 4; i++ {
		if r.Subnode[i] != nil && r.Subnode[i].HasItems() {
			found = r.Subnode[i].Remove(itemEnv, item)
			if found {
				// trim subtree if empty
				if r.Subnode[i].IsPrunable() {
					r.Subnode[i] = &Node{}
				}
				break
			}
		}
	}
	// if item was found lower down, don't need to search for it here
	if found {
		return found
	}
	// otherwise, try and remove the item from the list of items in this node
	found = r.remove(item)
	return found
}

// Visit ...
func (r *Root) Visit(searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	if !r.IsSearchMatch(searchEnv) {
		return
	}
	// this node may have items as well as subnodes (since items may not
	// be wholely contained in any single subnode
	r.VisitItems(searchEnv, r.Env, visitor)

	for i := 0; i < 4; i++ {
		if !r.Subnode[i].IsEmpty() {
			r.Subnode[i].Visit(searchEnv, visitor)
		}
	}
}
