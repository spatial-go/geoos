package kdtree

import "github.com/spatial-go/geoos/algorithm/matrix"

// KdNode A node of a KdTree, which represents one or more points in the same location.
type KdNode struct {
	matrix.Matrix
	Data interface{}
	Left *KdNode
	//the right node of the tree
	Right *KdNode
	Count int
}

// X Returns the X coordinate of the node
func (k *KdNode) X() float64 {
	return k.Matrix[0]
}

// Y Returns the Y coordinate of the node
func (k *KdNode) Y() float64 {
	return k.Matrix[1]
}

// Increments counts of points at this location
func (k *KdNode) increment() {
	k.Count++
}

// IsRepeated Tests whether more than one point with this value have been inserted (up to the tolerance)
func (k *KdNode) IsRepeated() bool {
	return k.Count > 1
}
