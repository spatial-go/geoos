package dbscan

// This code is heavily based on https://godoc.org/code.google.com/p/eaburns/kdtree
//
// Original code is under New BSD License.
// Author: Ethan Burns <burns.ethan@gmail.com>

import (
	"sort"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/clusters"
)

// KDTree is implementation of K-D Tree, with Points separated from
// nodes.
//
// Nodes (T) hold only indices into Points slice
type KDTree struct {
	Points clusters.PointList
	Root   *T
}

// A T is a the node of a K-D tree.  A *T is the root of a K-D tree,
// and nil is an empty K-D tree.
type T struct {
	// Point is the K-dimensional point associated with the
	// data of this node.
	PointID  int
	EqualIDs []int

	split       int
	left, right *T
}

// Insert returns a new K-D tree with the given node inserted.
// Inserting a node that is already a member of a K-D tree
// invalidates that tree.
func (tree *KDTree) Insert(point geoos.Point) {
	tree.Points = append(tree.Points, point)
	tree.Root = tree.insert(tree.Root, 0, &T{PointID: len(tree.Points) - 1})
}

func (tree *KDTree) insert(t *T, depth int, n *T) *T {
	if t == nil {
		n.split = depth % 2
		n.left, n.right = nil, nil
		return n
	}
	if tree.Points[n.PointID][t.split] < tree.Points[t.PointID][t.split] {
		t.left = tree.insert(t.left, depth+1, n)
	} else {
		t.right = tree.insert(t.right, depth+1, n)
	}
	return t
}

// InRange appends all nodes in the K-D tree that are within a given
// distance from the given point to the given slice, which may be nil.
// To  avoid allocation, the slice can be pre-allocated with a larger
// capacity and re-used across multiple calls to InRange.
func (tree *KDTree) InRange(pt geoos.Point, dist float64, nodes []int) []int {
	if dist < 0 {
		return nodes
	}
	return tree.inRange(tree.Root, pt, dist, nodes)
}

func (tree *KDTree) inRange(t *T, pt geoos.Point, r float64, nodes []int) []int {
	if t == nil {
		return nodes
	}

	diff := pt[t.split] - tree.Points[t.PointID][t.split]

	thisSide, otherSide := t.right, t.left
	if diff < 0 {
		thisSide, otherSide = t.left, t.right
	}

	p1 := geoos.Point{}
	p1[1-t.split] = (pt[1-t.split] + tree.Points[t.PointID][1-t.split]) / 2
	p1[t.split] = pt[t.split]

	p2 := geoos.Point{}
	p2[1-t.split] = (pt[1-t.split] + tree.Points[t.PointID][1-t.split]) / 2
	p2[t.split] = tree.Points[t.PointID][t.split]

	// dis := p1.sqDist(&p2)
	dis := DistanceSphericalFast(p1, p2)
	nodes = tree.inRange(thisSide, pt, r, nodes)
	if dis <= r*r {
		// if tree.Points[t.PointID].sqDist(pt) < r*r {
		if DistanceSphericalFast(tree.Points[t.PointID], pt) < r*r {
			nodes = append(nodes, t.PointID)
			nodes = append(nodes, t.EqualIDs...)
		}
		nodes = tree.inRange(otherSide, pt, r, nodes)
	}

	return nodes
}

// Height returns the height of the K-D tree.
func (tree *KDTree) Height() int {
	return tree.Root.height()
}

func (t *T) height() int {
	if t == nil {
		return 0
	}
	ht := t.left.height()
	if rht := t.right.height(); rht > ht {
		ht = rht
	}
	return ht + 1
}

// NewKDTree returns a new K-D tree built using the given nodes.
func NewKDTree(points clusters.PointList) *KDTree {
	result := &KDTree{
		Points: points,
	}

	if len(points) > 0 {
		result.Root = buildTree(0, preSort(result.Points))
	}

	return result
}

// buildTree does iteration of node building: it finds median
// point, and builds tree node with median (and all points equal to
// median), calling itself recursively for left and right subtrees
func buildTree(depth int, nodes *preSorted) *T {
	split := depth % 2
	switch len(nodes.cur[split]) {
	case 0:
		return nil
	case 1:
		return &T{
			PointID: nodes.cur[split][0],
			split:   split,
		}
	}
	med, equal, left, right := nodes.splitMed(split)
	return &T{
		PointID:  med,
		EqualIDs: equal,
		split:    split,
		left:     buildTree(depth+1, &left),
		right:    buildTree(depth+1, &right),
	}
}

// preSorted holds the nodes pre-sorted on each dimension.
type preSorted struct {
	points clusters.PointList

	// cur is the currently sorted set of point IDs by dimension
	cur [2][]int
}

// PreSort returns the nodes pre-sorted on each dimension.
func preSort(points clusters.PointList) *preSorted {
	p := new(preSorted)
	p.points = points
	for i := range p.cur {
		p.cur[i] = make([]int, len(points))
		for j := range p.cur[i] {
			p.cur[i][j] = j
		}
		sort.Sort(&nodeSorter{i, p.cur[i], points})
	}
	return p
}

// SplitMed returns the median node on the split dimension and two
// preSorted structs that contain the nodes (still sorted on each
// dimension) that are less than and greater than or equal to the
// median node value on the given splitting dimension.
func (p *preSorted) splitMed(dim int) (med int, equal []int, left, right preSorted) {
	m := len(p.cur[dim]) / 2
	for m > 0 && p.points[p.cur[dim][m-1]][dim] == p.points[p.cur[dim][m]][dim] {
		m--
	}
	mh := m
	for mh < len(p.cur[dim])-1 && p.points[p.cur[dim][mh+1]].Equal(p.points[p.cur[dim][m]]) {
		mh++
	}
	med = p.cur[dim][m]
	equal = p.cur[dim][m+1 : mh+1]
	pivot := p.points[med][dim]

	left.points = p.points
	left.cur[dim] = p.cur[dim][:m]

	right.points = p.points
	right.cur[dim] = p.cur[dim][mh+1:]

	for d := range p.cur {
		if d == dim {
			continue
		}

		left.cur[d] = make([]int, 0, len(p.cur))
		right.cur[d] = make([]int, 0, len(p.cur))

		for _, n := range p.cur[d] {
			if n == med {
				continue
			}
			skip := false
			for _, x := range equal {
				if n == x {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			if p.points[n][dim] < pivot {
				left.cur[d] = append(left.cur[d], n)
			} else {
				right.cur[d] = append(right.cur[d], n)
			}
		}
	}

	return
}

// A nodeSorter implements sort.Interface, sorting the nodes
// in ascending order of their point values on the split dimension.
type nodeSorter struct {
	split  int
	nodes  []int
	points clusters.PointList
}

func (n *nodeSorter) Len() int {
	return len(n.nodes)
}

func (n *nodeSorter) Swap(i, j int) {
	n.nodes[i], n.nodes[j] = n.nodes[j], n.nodes[i]
}

func (n *nodeSorter) Less(i, j int) bool {
	a, b := n.points[n.nodes[i]][n.split], n.points[n.nodes[j]][n.split]
	if a == b {
		return n.points[n.nodes[i]][1-n.split] < n.points[n.nodes[j]][1-n.split]
	}
	return a < b
}
