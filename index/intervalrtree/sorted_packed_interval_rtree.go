// Package intervalrtree Contains structs to implement an R-tree index for one-dimensional intervals.
package intervalrtree

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

// SortedPackedIntervalRTree  A static index on a set of 1-dimensional intervals,
//  using an R-Tree packed based on the order of the interval midpoints.
//  It supports range searching,
//  where the range is an interval of the real line (which may be a single point).
//  A common use is to index 1-dimensional intervals which
//  are the projection of 2-D objects onto an axis of the coordinate system.
type SortedPackedIntervalRTree struct {
	leaves LeafNodes
	root   Node
}

// Insert Adds an item to the index which is associated with the given interval
func (s *SortedPackedIntervalRTree) Insert(queryEnv *envelope.Envelope, item interface{}) error {
	min := queryEnv.MinX
	max := queryEnv.MaxX
	if s.root != nil {
		return index.ErrRTreeQueried
	}
	s.leaves = append(s.leaves, &LeafNode{&IRNode{min, max}, item})
	return nil
}

func (s *SortedPackedIntervalRTree) init() {
	// already built
	if s.root != nil {
		return
	}
	/**
	 * if leaves is empty then nothing has been inserted.
	 * In this case it is safe to leave the tree in an open state
	 */
	if len(s.leaves) == 0 {
		return
	}
	s.buildRoot()
}

func (s *SortedPackedIntervalRTree) buildRoot() {
	if s.root != nil {
		return
	}
	s.root = s.buildTree()
}

func (s *SortedPackedIntervalRTree) buildTree() Node {

	// sort the leaf nodes
	sort.Sort(s.leaves)

	// now group nodes into blocks of two and build tree up recursively
	src := s.leaves
	var dest, temp LeafNodes = make(LeafNodes, 1), make(LeafNodes, 1)

	for true {
		s.buildLevel(src, dest)
		if len(dest) == 1 {
			return dest[0]
		}
		temp = src
		src = dest
		dest = temp
	}
	return nil
}

func (s *SortedPackedIntervalRTree) buildLevel(src, dest LeafNodes) {
	//level++;
	for i := 0; i < len(src); i += 2 {
		if i == 0 {
			dest[i] = src[i]
		} else {
			dest = append(dest, src[i])
		}
		if i+1 < len(src) {
			node := NewBranchNode(src[i], src[i+1])
			dest = append(dest, node)

		}
	}
}

// Query Search for intervals in the index which intersect the given closed interval and apply the visitor to them.
func (s *SortedPackedIntervalRTree) Query(queryEnv *envelope.Envelope) interface{} {
	visitor := &index.ArrayVisitor{}
	s.QueryVisitor(queryEnv, visitor)
	return visitor.Items()
}

// QueryVisitor Search for intervals in the index which intersect the given closed interval and apply the visitor to them.
func (s *SortedPackedIntervalRTree) QueryVisitor(queryEnv *envelope.Envelope, visitor index.ItemVisitor) error {
	min := queryEnv.MinX
	max := queryEnv.MaxX
	s.init()

	// if root is null tree must be empty
	if s.root == nil {
		return index.ErrTreeIsNil
	}
	s.root.Query(min, max, visitor)
	return nil
}

// Remove Removes a single item from the tree.
func (s *SortedPackedIntervalRTree) Remove(itemEnv *envelope.Envelope, item interface{}) bool {
	// TODO Auto-generated method stub
	return false
}

// LeafNodes ...
type LeafNodes []Node

// Len ...
func (l LeafNodes) Len() int {
	return len(l)
}

// Less ...
func (l LeafNodes) Less(i, j int) bool {

	mid1 := (l[i].Min() + l[i].Max()) / 2
	mid2 := (l[j].Min() + l[j].Max()) / 2
	if mid1 < mid2 {
		return true
	}
	return false
}

// Swap ...
func (l LeafNodes) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

var (
	_ index.SpatialIndex = &SortedPackedIntervalRTree{}
)
