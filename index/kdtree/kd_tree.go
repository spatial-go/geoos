// Package kdtree A tree is a k-D tree index over 2-D point data.
package kdtree

import (
	"log"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/index"
	"github.com/spatial-go/geoos/index/quadtree"
)

// KdTree An implementation of a 2-D KD-Tree. KD-trees provide fast range searching and fast lookup for point data.
// This implementation supports detecting and snapping points which are closer
// than a given distance tolerance.
// If the same point (up to tolerance) is inserted
// more than once, it is snapped to the existing node.
// In other words, if a point is inserted which lies within the tolerance of a node already in the index,
// it is snapped to that node.
// When a point is snapped to a node then a new node is not created but the count of the existing node
// is incremented.
// If more than one node in the tree is within tolerance of an inserted point,
// the closest and then lowest node is snapped to.
type KdTree struct {
	root          *KdNode
	numberOfNodes int64
	tolerance     float64
}

// ToMatrixesNotIncludeRepeated Converts a collection of KdNodes to an array of matrixes.
func (k *KdTree) ToMatrixesNotIncludeRepeated(kdnodes []*KdNode) []matrix.Matrix {
	return k.ToMatrixes(kdnodes, false)
}

// ToMatrixes Converts a collection of KdNodes to an array of matrixes.
// specifying whether repeated nodes should be represented by multiple coordinates.
func (k *KdTree) ToMatrixes(kdnodes []*KdNode, includeRepeated bool) []matrix.Matrix {
	ms := []matrix.Matrix{}
	for _, node := range kdnodes {
		count := 1
		if includeRepeated {
			count = node.Count
		}
		for i := 0; i < count; i++ {
			ms = append(ms, node.Matrix)
		}
	}
	return ms
}

// IsEmpty Tests whether the index contains any items.
func (k *KdTree) IsEmpty() bool {
	return k.root == nil
}

// InsertNoData Inserts a new point in the kd-tree, with no data.
func (k *KdTree) InsertNoData(p matrix.Matrix) *KdNode {
	return k.InsertMatrix(p, nil)
}

//InsertMatrix Inserts a new point into the kd-tree.
func (k *KdTree) InsertMatrix(p matrix.Matrix, data interface{}) *KdNode {
	if k.root == nil {
		k.root = &KdNode{Matrix: p, Data: data}
		return k.root
	}

	/**
	 * Check if the point is already in the tree, up to tolerance.
	 * If tolerance is zero, this phase of the insertion can be skipped.
	 */
	if k.tolerance > 0 {
		matchNode := k.FindBestMatchNode(p)
		if matchNode != nil {
			// point already in index - increment counter
			matchNode.increment()
			return matchNode
		}
	}

	return k.insertExact(p, data)
}

//Insert Inserts a new point into the kd-tree.
func (k *KdTree) Insert(env *envelope.Envelope, data interface{}) error {
	if m, ok := data.(matrix.Matrix); ok {
		k.InsertNoData(m)
		return nil
	}
	return index.ErrNotMatchType
}

// FindBestMatchNode Finds the node in the tree which is the best match for a point
//  being inserted.
//  The match is made deterministic by returning the lowest of any nodes which
//  lie the same distance from the point.
//  There may be no match if the point is not within the distance tolerance of any
//  existing node.
func (k *KdTree) FindBestMatchNode(p matrix.Matrix) *KdNode {
	visitor := &BestMatchVisitor{Matrix: p, tolerance: k.tolerance}
	if err := k.QueryVisitor(visitor.QueryEnvelope(), visitor); err != nil {
		log.Println(err)
	}
	return visitor.MatchNode
}

// insertExact Inserts a point known to be beyond the distance tolerance of any existing node.
//  The point is inserted at the bottom of the exact splitting path,
//  so that tree shape is deterministic.
func (k *KdTree) insertExact(p matrix.Matrix, data interface{}) *KdNode {
	currentNode := k.root
	leafNode := k.root
	isOddLevel := true
	isLessThan := true

	/**
	 * Traverse the tree, first cutting the plane left-right (by X ordinate)
	 * then top-bottom (by Y ordinate)
	 */
	for currentNode != nil {
		isInTolerance := measure.PlanarDistance(p, currentNode.Matrix) <= k.tolerance

		// check if point is already in tree (up to tolerance) and if so simply
		// return existing node
		if isInTolerance {
			currentNode.increment()
			return currentNode
		}

		if isOddLevel {
			isLessThan = p[0] < currentNode.X()
		} else {
			isLessThan = p[1] < currentNode.Y()
		}
		leafNode = currentNode
		if isLessThan {
			//System.out.print("L");
			currentNode = currentNode.Left
		} else {
			//System.out.print("R");
			currentNode = currentNode.Right
		}

		isOddLevel = !isOddLevel
	}
	//System.out.println("<<");
	// no node found, add new leaf node to tree
	k.numberOfNodes++
	node := &KdNode{Matrix: p, Data: data}
	if isLessThan {
		leafNode.Left = node
	} else {
		leafNode.Right = node
	}
	return node
}

// QueryNode Performs a range search of the points in the index and visits all nodes found.
func (k *KdTree) QueryNode(currentNode *KdNode,
	queryEnv *envelope.Envelope, odd bool, visitor index.ItemVisitor) error {
	if currentNode == nil {
		return index.ErrTreeIsNil
	}
	var min, max, discriminant float64
	if odd {
		min = queryEnv.MinX
		max = queryEnv.MaxX
		discriminant = currentNode.X()
	} else {
		min = queryEnv.MinY
		max = queryEnv.MaxY
		discriminant = currentNode.Y()
	}
	searchLeft := min < discriminant
	searchRight := discriminant <= max

	// search is computed via in-order traversal
	if searchLeft {
		if err := k.QueryNode(currentNode.Left, queryEnv, !odd, visitor); err != nil {
			log.Println(err)
		}
	}
	if queryEnv.Contains(envelope.Matrix(currentNode.Matrix)) {
		visitor.VisitItem(currentNode)
	}
	if searchRight {
		if err := k.QueryNode(currentNode.Right, queryEnv, !odd, visitor); err != nil {
			log.Println(err)
		}
	}
	return nil
}

// QueryNodePoint Performs a range search of the points in the index and visits all nodes found.
func (k *KdTree) QueryNodePoint(currentNode *KdNode,
	queryPt matrix.Matrix, odd bool) *KdNode {
	if currentNode == nil {
		return nil
	}
	if currentNode.Matrix.Equals(queryPt) {
		return currentNode
	}
	var ord, discriminant float64
	if odd {
		ord = queryPt[0]
		discriminant = currentNode.X()
	} else {
		ord = queryPt[1]
		discriminant = currentNode.Y()
	}
	searchLeft := ord < discriminant

	if searchLeft {
		return k.QueryNodePoint(currentNode.Left, queryPt, !odd)
	}
	return k.QueryNodePoint(currentNode.Right, queryPt, !odd)
}

// QueryVisitor Performs a range search of the points in the index and visits all nodes found.
func (k *KdTree) QueryVisitor(queryEnv *envelope.Envelope, visitor index.ItemVisitor) error {
	return k.QueryNode(k.root, queryEnv, true, visitor)
}

// Query  Performs a range search of the points in the index.
func (k *KdTree) Query(qEnv *envelope.Envelope) interface{} {
	bmv := &BestMatchVisitor{}
	if err := k.QueryVisitor(qEnv, bmv); err != nil {
		log.Println(err)
	}

	return bmv.MatchNode
}

// QueryMatrix Searches for a given point in the index and returns its node if found.
func (k *KdTree) QueryMatrix(queryPt matrix.Matrix) *KdNode {
	return k.QueryNodePoint(k.root, queryPt, true)
}

// Depth Computes the Depth of the tree.
func (k *KdTree) Depth() int {
	return k.DepthNode(k.root)
}

// DepthNode Computes the Depth of the tree.
func (k *KdTree) DepthNode(currentNode *KdNode) int {
	if currentNode == nil {
		return 0
	}
	dL := k.DepthNode(currentNode.Left)
	dR := k.DepthNode(currentNode.Right)
	if dL > dR {
		return 1 + dL
	}
	return 1 + dR
}

// Size Computes the Size (number of items) in the tree.
func (k *KdTree) Size() int {
	return k.SizeNode(k.root)
}

// SizeNode Computes the Size (number of items) in the tree.
func (k *KdTree) SizeNode(currentNode *KdNode) int {
	if currentNode == nil {
		return 0
	}
	sizeL := k.SizeNode(currentNode.Left)
	sizeR := k.SizeNode(currentNode.Right)
	return 1 + sizeL + sizeR
}

// Remove Removes a single item from the tree.
func (k *KdTree) Remove(itemEnv *envelope.Envelope, item interface{}) bool {
	return false
}

// BestMatchVisitor A visitor for items in a SpatialIndex.
type BestMatchVisitor struct {
	tolerance float64
	MatchNode *KdNode
	matchDist float64
	matrix.Matrix
}

// QueryEnvelope ...
func (b *BestMatchVisitor) QueryEnvelope() *envelope.Envelope {
	queryEnv := envelope.Matrix(b.Matrix)
	queryEnv.ExpandBy(b.tolerance)
	return queryEnv
}

// VisitItem Visits an item in the index.
func (b *BestMatchVisitor) VisitItem(item interface{}) {
	node := item.(*KdNode)
	dist := measure.PlanarDistance(b.Matrix, node.Matrix)
	isInTolerance := dist <= b.tolerance
	if !isInTolerance {
		return
	}
	update := false

	if b.MatchNode == nil || dist < b.matchDist {
		update = true
	}
	// if distances are the same, record the lesser coordinate
	if b.MatchNode != nil {
		comp, _ := node.Matrix.Compare(b.MatchNode.Matrix)
		if dist == b.matchDist && comp < 1 {
			update = true
		}
	}

	if update {
		b.MatchNode = node
		b.matchDist = dist
	}
}

// Items returns items.
func (b *BestMatchVisitor) Items() interface{} {
	return b.MatchNode
}

var (
	_ index.ItemVisitor  = &BestMatchVisitor{}
	_ index.SpatialIndex = &quadtree.Quadtree{}
	_ index.SpatialIndex = &KdTree{}
)
