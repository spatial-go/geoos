package dbscan

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/spatial-go/geoos/clusters"
	"github.com/spatial-go/geoos/space"
)

// A pointSlice is a slice of points that implements the quick.Generator
// interface, generating a random set of points on the unit square.
type pointSlice clusters.PointList

func (pointSlice) Generate(r *rand.Rand, size int) reflect.Value {
	ps := make([]space.Point, size)
	for i := range ps {
		ps[i] = append(ps[i], r.Float64(), r.Float64())
	}
	return reflect.ValueOf(ps)
}

// TestInsert tests the insert function, ensuring that random points
// inserted into an empty tree maintain the K-D tree invariant.
func TestInsert(t *testing.T) {
	if err := quick.Check(func(pts pointSlice) bool {
		var tree = NewKDTree(nil)
		for _, p := range pts {
			if p == nil || p.IsEmpty() {
				//tree.Insert(space.Point{0, 0})
			} else {
				tree.Insert(p)
			}
		}
		_, ok := tree.invariantHolds(tree.Root)
		return ok
	}, nil); err != nil {
		t.Error(err)
	}
}

// TestMake tests the Make function, ensuring that a tree built
// using random points respects the K-D tree invariant.
func TestMake(t *testing.T) {
	if err := quick.Check(func(pts pointSlice) bool {
		tree := NewKDTree(clusters.PointList(pts))
		_, ok := tree.invariantHolds(tree.Root)
		return ok
	}, nil); err != nil {
		t.Error(err)
	}
}

// TestInRange tests the InRange function, ensuring that all points
// in the range are reported, and all points reported are indeed in
// the range.
func TestInRange(t *testing.T) {
	if err := quick.Check(func(pts pointSlice, pt space.Point, r float64) bool {
		if pt.IsEmpty() {
			pt = space.Point{0, 0}
		}
		r = math.Abs(r)
		tree := NewKDTree(clusters.PointList(pts))

		in := make(map[int]bool, len(pts))
		for _, n := range tree.InRange(pt, r, nil) {
			in[n] = true
		}

		num := 0
		for i, p := range pts {
			// dis:=pt.sqDist(&p)
			dis := DistanceSphericalFast(pt, p)
			if dis <= r*r {
				num++
				if !in[i] {
					return false
				}
			}
		}
		return num == len(in)
	}, nil); err != nil {
		t.Error(err)
	}
}

// InvariantHolds returns the points in this subtree, and a bool
// that is true if the K-D tree invariant holds.  The K-D tree invariant
// states that all points in the left subtree have values less than that
// of the current node on the splitting dimension, and the points
// in the right subtree have values greater than or equal to that of
// the current node.
func (tree *KDTree) invariantHolds(t *T) ([]space.Point, bool) {
	if t == nil {
		return nil, true
	}

	ok := true

	for _, i := range t.EqualIDs {
		if !tree.Points[i].Equal(tree.Points[t.PointID]) {
			ok = false
			break
		}
	}

	left, leftOk := tree.invariantHolds(t.left)
	right, rightOk := tree.invariantHolds(t.right)

	ok = ok && leftOk && rightOk

	if ok {
		for _, l := range left {
			if l[t.split] >= tree.Points[t.PointID][t.split] {
				ok = false
				break
			}
		}
	}
	if ok {
		for _, r := range right {
			if r[t.split] < tree.Points[t.PointID][t.split] {
				ok = false
				break
			}
		}
	}

	return append(append(left, tree.Points[t.PointID]), right...), ok
}

func TestPreSort(t *testing.T) {
	if err := quick.Check(func(pts pointSlice) bool {
		p := preSort(clusters.PointList(pts))
		for i := range p.cur {
			if !isSortedOnDim(i, p.cur[i], pts) || len(p.cur[i]) != len(pts) {
				return false
			}
		}
		return true
	}, nil); err != nil {
		t.Error(err)
	}
}

func TestPreSort_SplitMed(t *testing.T) {
	if err := quick.Check(func(pts pointSlice, dim int) bool {
		if len(pts) == 0 {
			return true
		}
		if dim < 0 {
			dim = -dim
		}
		dim %= 2

		sorted := preSort(clusters.PointList(pts))
		med, equal, left, right := sorted.splitMed(dim)
		for _, p := range equal {
			if pts[p].Equal(pts[med]) {
				return false
			}
		}

		if len(left.cur[dim])+len(right.cur[dim])+1+len(equal) != len(pts) {
			return false
		}

		for i, p := range [2]*preSorted{&left, &right} {
			for d, ns := range p.cur {
				if len(ns) != len(p.cur[0]) {
					return false
				}
				if !isSortedOnDim(d, ns, pts) {
					return false
				}
				for _, n := range ns {
					if i == 0 && pts[n][dim] >= pts[med][dim] {
						return false
					} else if i == 1 && pts[n][dim] < pts[med][dim] {
						return false
					}
				}
			}
		}
		return true
	}, nil); err != nil {
		t.Error(err)
	}
}

// IsSortedOnDim returns true if the given slice is in sorted order
// on the given dimension.
func isSortedOnDim(dim int, nodes []int, pts pointSlice) bool {
	if len(nodes) == 0 {
		return true
	}
	prev := pts[nodes[0]][dim]
	for _, n := range nodes {
		if pts[n][dim] < prev {
			return false
		}
		prev = pts[n][dim]
	}
	return true
}
