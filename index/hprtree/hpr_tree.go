package hprtree

import (
	"math"
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

// HPRTree const parameter.
const (
	EnvSize             = 4
	HilbertLevel        = 12
	DefaultNodeCapacity = 16
)

// HPRTree A Hilbert-Packed R-tree.  This is a static R-tree which is packed by
// using the Hilbert ordering of the tree items.
// The tree is constructed by sorting the items by the Hilbert code of the midpoint of their envelope.
type HPRTree struct {
	nodeCapacity int
	totalExtent  *envelope.Envelope

	Items           []interface{}
	layerStartIndex []int
	nodeBounds      []float64
	isBuilt         bool
}

// HPRTree return default HPRTree.
func (h *HPRTree) HPRTree() *HPRTree {
	h = &HPRTree{}
	h.nodeCapacity = DefaultNodeCapacity
	h.totalExtent = &envelope.Envelope{}
	return h
}

// Size Gets the number of items in the index.
func (h *HPRTree) Size() int {
	return len(h.Items)
}

// Insert Adds a spatial item with an extent specified by the given Envelope to the index
func (h *HPRTree) Insert(itemEnv *envelope.Envelope, item interface{}) {
	if h.isBuilt {
		return
	}
	h.Items = append(h.Items, &Item{itemEnv, item})
	h.totalExtent.ExpandToIncludeEnv(itemEnv)
}

// Query Queries the index for all items whose extents intersect the given search  Envelope
// Note that some kinds of indexes may also return objects which do not in fact
//  intersect the query envelope.
func (h *HPRTree) Query(searchEnv *envelope.Envelope) interface{} {
	h.build()

	if !h.totalExtent.IsIntersects(searchEnv) {
		return h.Items
	}

	visitor := &index.ArrayVisitor{}
	h.QueryVisitor(searchEnv, visitor)
	return visitor.Items
}

// QueryVisitor Queries the index for all items whose extents intersect the given search Envelope,
// and applies an  ItemVisitor to them.
// Note that some kinds of indexes may also return objects which do not in fact
// intersect the query envelope.
func (h *HPRTree) QueryVisitor(searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	h.build()
	if !h.totalExtent.IsIntersects(searchEnv) {
		return
	}
	if h.layerStartIndex == nil {
		h.queryItems(0, searchEnv, visitor)
	} else {
		h.queryTopLayer(searchEnv, visitor)
	}
}
func (h *HPRTree) queryTopLayer(searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	layerIndex := len(h.layerStartIndex) - 2
	layerSize := h.layerSize(layerIndex)
	// query each node in layer
	for i := 0; i < layerSize; i += EnvSize {
		h.queryNode(layerIndex, i, searchEnv, visitor)
	}
}

func (h *HPRTree) queryNode(layerIndex, nodeOffset int, searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	layerStart := h.layerStartIndex[layerIndex]
	nodeIndex := layerStart + nodeOffset
	if !h.isIntersects(nodeIndex, searchEnv) {
		return
	}
	if layerIndex == 0 {
		childNodesOffset := nodeOffset / EnvSize * h.nodeCapacity
		h.queryItems(childNodesOffset, searchEnv, visitor)
	} else {
		childNodesOffset := nodeOffset * h.nodeCapacity
		h.queryNodeChildren(layerIndex-1, childNodesOffset, searchEnv, visitor)
	}
}

func (h *HPRTree) isIntersects(nodeIndex int, env *envelope.Envelope) bool {
	//nodeIntersectsCount++;
	isBeyond := (env.MaxX < h.nodeBounds[nodeIndex]) || (env.MaxY < h.nodeBounds[nodeIndex+1]) ||
		(env.MinX > h.nodeBounds[nodeIndex+2]) || (env.MinY > h.nodeBounds[nodeIndex+3])
	return !isBeyond
}

func (h *HPRTree) queryNodeChildren(layerIndex, blockOffset int, searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	layerStart := h.layerStartIndex[layerIndex]
	layerEnd := h.layerStartIndex[layerIndex+1]
	for i := 0; i < h.nodeCapacity; i++ {
		nodeOffset := blockOffset + EnvSize*i
		// don't query past layer end
		if layerStart+nodeOffset >= layerEnd {
			break
		}
		h.queryNode(layerIndex, nodeOffset, searchEnv, visitor)
	}
}

func (h *HPRTree) queryItems(blockStart int, searchEnv *envelope.Envelope, visitor index.ItemVisitor) {
	for i := 0; i < h.nodeCapacity; i++ {
		itemIndex := blockStart + i
		// don't query past end of items
		if itemIndex >= h.Size() {
			break
		}
		// visit the item if its envelope intersects search env
		item := h.Items[itemIndex].(*Item)
		//nodeIntersectsCount++;
		if h.isIntersectsEnv(item.Env, searchEnv) {
			//if (item.getEnvelope().intersects(searchEnv)) {
			visitor.VisitItem(item.Item)
		}
	}
}

// isIntersectsEnv Tests whether two envelopes intersect.
// Avoids the null check in {@link Envelope#intersects(Envelope)}.
func (h *HPRTree) isIntersectsEnv(env1, env2 *envelope.Envelope) bool {
	return !(env2.MinX > env1.MaxX ||
		env2.MaxX < env1.MinX ||
		env2.MinY > env1.MaxY ||
		env2.MaxY < env1.MinY)
}

func (h *HPRTree) layerSize(layerIndex int) int {
	layerStart := h.layerStartIndex[layerIndex]
	layerEnd := h.layerStartIndex[layerIndex+1]
	return layerEnd - layerStart
}

// Remove Removes a single item from the tree.
func (h *HPRTree) Remove(itemEnv *envelope.Envelope, item interface{}) bool {
	// TODO Auto-generated method stub
	return false
}

// build Builds the index, if not already built.
func (h *HPRTree) build() {
	// skip if already built
	if h.isBuilt {
		return
	}
	h.isBuilt = true
	// don't need to build an empty or very small tree
	if h.Size() <= h.nodeCapacity {
		return
	}
	h.sortItems()
	//dumpItems(items);

	h.layerStartIndex = h.computeLayerIndices(h.Size(), h.nodeCapacity)
	// allocate storage
	nodeCount := h.layerStartIndex[len(h.layerStartIndex)-1] / 4
	h.nodeBounds = h.createBoundsArray(nodeCount)

	// compute tree nodes
	h.computeLeafNodes(h.layerStartIndex[1])
	for i := 1; i < len(h.layerStartIndex)-1; i++ {
		h.computeLayerNodes(i)
	}
	//dumpNodes();
}

func (h *HPRTree) createBoundsArray(size int) []float64 {
	a := make([]float64, 4*size)
	for i := 0; i < size; i++ {
		index := 4 * i
		a[index] = math.MaxFloat64
		a[index+1] = math.MaxFloat64
		a[index+2] = -math.MaxFloat64
		a[index+3] = -math.MaxFloat64
	}
	return a
}

func (h *HPRTree) computeLayerNodes(layerIndex int) {
	layerStart := h.layerStartIndex[layerIndex]
	childLayerStart := h.layerStartIndex[layerIndex-1]
	layerSize := h.layerSize(layerIndex)
	childLayerEnd := layerStart
	for i := 0; i < layerSize; i += EnvSize {
		childStart := childLayerStart + h.nodeCapacity*i
		h.computeNodeBounds(layerStart+i, childStart, childLayerEnd)
	}
}

func (h *HPRTree) computeNodeBounds(nodeIndex, blockStart, nodeMaxIndex int) {
	for i := 0; i <= h.nodeCapacity; i++ {
		index := blockStart + 4*i
		if index >= nodeMaxIndex {
			break
		}
		h.updateNodeBounds(nodeIndex, h.nodeBounds[index], h.nodeBounds[index+1], h.nodeBounds[index+2], h.nodeBounds[index+3])
	}
}

func (h *HPRTree) computeLeafNodes(layerSize int) {
	for i := 0; i < layerSize; i += EnvSize {
		h.computeLeafNodeBounds(i, h.nodeCapacity*i/4)
	}
}

func (h *HPRTree) computeLeafNodeBounds(nodeIndex, blockStart int) {
	for i := 0; i <= h.nodeCapacity; i++ {
		itemIndex := blockStart + i
		if itemIndex >= h.Size() {
			break
		}
		env := h.Items[itemIndex].(Item).Env
		h.updateNodeBounds(nodeIndex, env.MinX, env.MinY, env.MaxX, env.MaxY)
	}
}

func (h *HPRTree) updateNodeBounds(nodeIndex int, minX, minY, maxX, maxY float64) {
	if minX < h.nodeBounds[nodeIndex] {
		h.nodeBounds[nodeIndex] = minX
	}
	if minY < h.nodeBounds[nodeIndex+1] {
		h.nodeBounds[nodeIndex+1] = minY
	}
	if maxX > h.nodeBounds[nodeIndex+2] {
		h.nodeBounds[nodeIndex+2] = maxX
	}
	if maxY > h.nodeBounds[nodeIndex+3] {
		h.nodeBounds[nodeIndex+3] = maxY
	}
}

func (h *HPRTree) getNodeEnvelope(i int) *envelope.Envelope {
	return envelope.FourFloat(h.nodeBounds[i], h.nodeBounds[i+1], h.nodeBounds[i+2], h.nodeBounds[i+3])
}

func (h *HPRTree) computeLayerIndices(itemSize, nodeCapacity int) []int {
	layerIndexList := []int{}
	layerSize := itemSize
	index := 0
	for layerSize > 1 {
		layerIndexList = append(layerIndexList, index)
		layerSize := h.numNodesToCover(layerSize, nodeCapacity)
		index += EnvSize * layerSize
	}
	return layerIndexList
}

/**
 * Computes the number of blocks (nodes) required to
 * cover a given number of children.
 *
 * @param nChild
 * @param nodeCapacity
 * @return the number of nodes needed to cover the children
 */
func (h *HPRTree) numNodesToCover(nChild, nodeCapacity int) int {
	mult := nChild / nodeCapacity
	total := mult * nodeCapacity
	if total == nChild {
		return mult

	}
	return mult + 1
}

/**
 * Gets the extents of the internal index nodes
 *
 * @return a list of the internal node extents
 */
func (h *HPRTree) getBounds() []*envelope.Envelope {
	numNodes := len(h.nodeBounds) / 4
	bounds := make([]*envelope.Envelope, numNodes)
	// create from largest to smallest
	for i := numNodes - 1; i >= 0; i-- {
		boundIndex := 4 * i
		bounds[i] = envelope.FourFloat(h.nodeBounds[boundIndex], h.nodeBounds[boundIndex+2],
			h.nodeBounds[boundIndex+1], h.nodeBounds[boundIndex+3])
	}
	return bounds
}

func (h *HPRTree) sortItems() {
	comp := &ItemComparator{items: h.Items, encoder: (&HilbertEncoder{}).HilbertEncoder(HilbertLevel, h.totalExtent)}
	sort.Sort(comp)
	h.Items = comp.items
}

// ItemComparator sort items by HilbertEncoder
type ItemComparator struct {
	items   []interface{}
	encoder *HilbertEncoder
}

// Len ...
func (it *ItemComparator) Len() int {
	return len(it.items)
}

// Less ...
func (it *ItemComparator) Less(i, j int) bool {

	hCode1 := it.encoder.encode(it.items[i].(Item).Env)
	hCode2 := it.encoder.encode(it.items[j].(Item).Env)
	return hCode1 < hCode2
}

// Swap ...
func (it ItemComparator) Swap(i, j int) {
	it.items[i], it.items[j] = it.items[j], it.items[i]
}

var (
	_ index.SpatialIndex = &HPRTree{}
)
