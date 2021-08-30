package quadtree

import (
	"reflect"

	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

// Node Represents a node of a  Quadtree.  Nodes contain
//  items which have a spatial extent corresponding to the node's position in the quadtree.
type Node struct {
	Items            []interface{}
	Subnode          [4]*Node
	Env              *envelope.Envelope
	Centrex, Centrey float64
	Level            int
}

// SubnodeIndex Gets the index of the subquad that wholly contains the given envelope.
// If none does, returns -1.
func SubnodeIndex(env *envelope.Envelope, centrex, centrey float64) int {

	subnodeIndex := -1
	if env.MinX >= centrex {
		if env.MinY >= centrey {
			subnodeIndex = 3
		}
		if env.MaxY <= centrey {
			subnodeIndex = 1
		}
	}
	if env.MaxX <= centrex {
		if env.MinY >= centrey {
			subnodeIndex = 2
		}
		if env.MaxY <= centrey {
			subnodeIndex = 0
		}
	}
	return subnodeIndex
}

// NewNode ...
func NewNode(env *envelope.Envelope) *Node {
	key := NewKeyEnv(env)
	node := NewNodeEnv(key.Env, key.Level)
	return node
}

// NewNodeExpanded ...
func NewNodeExpanded(node *Node, addEnv *envelope.Envelope) *Node {
	expandEnv := envelope.Env(addEnv)
	if node != nil {
		expandEnv.ExpandToIncludeEnv(node.Env)
	}
	largerNode := NewNode(expandEnv)
	if node != nil {
		largerNode.InsertNode(node)
	}
	return largerNode
}

// NewNodeEnv ...
func NewNodeEnv(env *envelope.Envelope, level int) *Node {
	n := &Node{}
	//this.parent = parent;
	n.Env = env
	n.Level = level
	n.Centrex = (env.MinX + env.MaxX) / 2
	n.Centrey = (env.MinY + env.MaxY) / 2
	return n
}

// HasItems ...
func (n *Node) HasItems() bool {
	if n == nil || n.Items == nil || len(n.Items) == 0 {
		return false
	}
	return true
}

// Add ...
func (n *Node) Add(item interface{}) {
	n.Items = append(n.Items, item)
	//DEBUG itemCount++;
	//DEBUG System.out.print(itemCount);
}

// Remove Removes a single item from this subtree.
func (n *Node) Remove(itemEnv *envelope.Envelope, item interface{}) bool {
	// use envelope to restrict nodes scanned
	if !n.IsSearchMatch(itemEnv) {
		return false
	}
	found := false
	for i := 0; i < 4; i++ {
		if n.Subnode[i] != nil && n.Subnode[i].HasItems() {
			found = n.Subnode[i].Remove(itemEnv, item)
			if found {
				// trim subtree if empty
				if n.Subnode[i].IsPrunable() {
					n.Subnode[i] = &Node{}
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

	found = n.remove(item)
	return found
}

// remove Removes a single item from this subtree.
func (n *Node) remove(item interface{}) bool {
	for i, v := range n.Items {
		if reflect.DeepEqual(v, item) {
			n.Items = append(n.Items[:i], n.Items[i+1:]...)
			return true
		}
	}
	return false
}

// IsPrunable ...
func (n *Node) IsPrunable() bool {
	return !(n.HasChildren() || n.HasItems())
}

// HasChildren ...
func (n *Node) HasChildren() bool {
	for i := 0; i < 4; i++ {
		if !n.Subnode[i].HasItems() {
			return true
		}
	}
	return false
}

// Visit ...
func (n *Node) Visit(searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	if !n.IsSearchMatch(searchEnv) {
		return
	}
	// this node may have items as well as subnodes (since items may not
	// be wholely contained in any single subnode
	n.VisitItems(searchEnv, n.Env, visitor)

	for i := 0; i < 4; i++ {
		if !n.Subnode[i].IsEmpty() {
			n.Subnode[i].Visit(searchEnv, visitor)
		}
	}
}

// VisitItems ...
func (n *Node) VisitItems(searchEnv, nodeEnv *envelope.Envelope, visitor index.ItemVisitor) {
	// would be nice to filter items based on search envelope, but can't until they contain an envelope
	for _, v := range n.Items {
		if searchEnv.IsIntersects(nodeEnv) {
			visitor.VisitItem(v)
		}
	}
}

// Depth ...
func (n *Node) Depth() int {
	maxSubDepth := 0
	for i := 0; i < 4; i++ {
		if !n.Subnode[i].IsEmpty() {
			sqd := n.Subnode[i].Depth()
			if sqd > maxSubDepth {
				maxSubDepth = sqd
			}
		}
	}
	return maxSubDepth + 1
}

// Size ...
func (n *Node) Size() int {
	subSize := 0
	for i := 0; i < 4; i++ {
		if !n.Subnode[i].IsEmpty() {
			subSize += n.Subnode[i].Size()
		}
	}
	return subSize + len(n.Items)
}

// NodeCount ...
func (n *Node) NodeCount() int {
	subSize := 0
	for i := 0; i < 4; i++ {
		if !n.Subnode[i].IsEmpty() {
			subSize += n.Subnode[i].Size()
		}
	}
	return subSize + 1
}

// IsEmpty ...
func (n *Node) IsEmpty() bool {
	if n == nil {
		return true
	}
	isEmpty := true
	if n.Items != nil && len(n.Items) > 0 {
		isEmpty = false
	} else {
		for i := 0; i < 4; i++ {
			if n.Subnode[i] == nil {
				continue
			}
			if !n.Subnode[i].HasItems() {
				if !n.Subnode[i].IsEmpty() {
					isEmpty = false
					break
				}
			}
		}
	}
	return isEmpty
}

// IsSearchMatch ...
func (n *Node) IsSearchMatch(searchEnv *envelope.Envelope) bool {
	if searchEnv == nil {
		return false
	}
	return n.Env.IsIntersects(searchEnv)
}

// GetNode Returns the subquad containing the envelope searchEnv.
// Creates the subquad if it does not already exist.
func (n *Node) GetNode(searchEnv *envelope.Envelope) *Node {
	subnodeIndex := SubnodeIndex(searchEnv, n.Centrex, n.Centrey)
	// if subquadIndex is -1 searchEnv is not contained in a subquad
	if subnodeIndex != -1 {
		// create the quad if it does not exist
		node := n.GetSubnode(subnodeIndex)
		// recursively search the found/created quad
		return node.GetNode(searchEnv)
	}
	return n
}

// Find Returns the smallest existing  node containing the envelope.
func (n *Node) Find(searchEnv *envelope.Envelope) *Node {
	subnodeIndex := SubnodeIndex(searchEnv, n.Centrex, n.Centrey)
	if subnodeIndex == -1 {
		return n
	}
	if n.Subnode[subnodeIndex] != nil {
		// query lies in subquad, so search it
		node := n.Subnode[subnodeIndex]
		return node.Find(searchEnv)
	}
	// no existing subquad, so return this one anyway
	return n
}

// InsertNode ...
func (n *Node) InsertNode(node *Node) {
	//System.out.println(env);
	//System.out.println(quad.env);
	index := SubnodeIndex(node.Env, n.Centrex, n.Centrey)
	//System.out.println(index);
	if node.Level == n.Level-1 {
		n.Subnode[index] = node
		//System.out.println("inserted");
	} else {
		// the quad is not a direct child, so make a new child quad to contain it
		// and recursively insert the quad
		childNode := n.CreateSubnode(index)
		childNode.InsertNode(node)
		n.Subnode[index] = childNode
	}
}

// GetSubnode get the subquad for the index. If it doesn't exist, create it
func (n *Node) GetSubnode(index int) *Node {
	if n.Subnode[index] == nil {
		n.Subnode[index] = n.CreateSubnode(index)
	}
	return n.Subnode[index]
}

// CreateSubnode ...
func (n *Node) CreateSubnode(index int) *Node {
	// create a new subquad in the appropriate quadrant

	minx := 0.0
	maxx := 0.0
	miny := 0.0
	maxy := 0.0

	switch index {
	case 0:
		minx = n.Env.MinX
		maxx = n.Centrex
		miny = n.Env.MinY
		maxy = n.Centrey
	case 1:
		minx = n.Centrex
		maxx = n.Env.MaxX
		miny = n.Env.MinY
		maxy = n.Centrey
	case 2:
		minx = n.Env.MinX
		maxx = n.Centrex
		miny = n.Centrey
		maxy = n.Env.MaxY
	case 3:
		minx = n.Centrex
		maxx = n.Env.MaxX
		miny = n.Centrey
		maxy = n.Env.MaxY
	}
	sqEnv := envelope.FourFloat(minx, maxx, miny, maxy)
	node := NewNodeEnv(sqEnv, n.Level-1)
	return node
}
