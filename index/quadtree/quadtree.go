// Package quadtree A Quadtree is a spatial index structure for efficient range querying
//  of items bounded by 2D rectangles.
package quadtree

import (
	"log"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

// Quadtree A Quadtree is a spatial index structure for efficient range querying
//  of items bounded by 2D rectangles.
//  Geometries can be indexed by using their Envelopes.
//  Any type of Object can also be indexed as
//  long as it has an extent that can be represented by an  Envelope.
type Quadtree struct {
	Root      *Root
	MinExtent float64
}

// EnsureExtent Ensure that the envelope for the inserted item has non-zero extents.
//  Use the current minExtent to pad the envelope, if necessary
func EnsureExtent(itemEnv *envelope.Envelope, minExtent float64) *envelope.Envelope {

	minx := itemEnv.MinX
	maxx := itemEnv.MaxX
	miny := itemEnv.MinY
	maxy := itemEnv.MaxY
	// has a non-zero extent
	if minx != maxx && miny != maxy {
		return itemEnv
	}
	// pad one or both extents
	if minx == maxx {
		minx = minx - minExtent/2.0
		maxx = maxx + minExtent/2.0
	}
	if miny == maxy {
		miny = miny - minExtent/2.0
		maxy = maxy + minExtent/2.0
	}
	return envelope.FourFloat(minx, maxx, miny, maxy)
}

// NewQuadtree  Constructs a Quadtree with zero items.
func NewQuadtree() *Quadtree {
	qt := &Quadtree{}
	qt.Root = &Root{Node: &Node{}, origin: matrix.Matrix{0, 0}}
	return qt
}

// Depth Returns the number of levels in the tree.
func (q *Quadtree) Depth() int {
	//I don't think it's possible for root to be null. Perhaps we should
	//remove the check. [Jon Aquino]
	//Or make an assertion [Jon Aquino 10/29/2003]
	if q.Root != nil {
		return q.Root.Depth()
	}
	return 0
}

// IsEmpty Tests whether the index contains any items.
func (q *Quadtree) IsEmpty() bool {
	if q.Root == nil {
		return true
	}
	return q.Root.IsEmpty()
}

// Size Returns the number of items in the tree.
func (q *Quadtree) Size() int {
	if q.Root != nil {
		return q.Root.Size()
	}
	return 0
}

// Insert insert a single item to the tree.
func (q *Quadtree) Insert(itemEnv *envelope.Envelope, item interface{}) error {
	q.CollectStats(itemEnv)
	insertEnv := EnsureExtent(itemEnv, q.MinExtent)
	q.Root.Insert(insertEnv, item)
	return nil
}

// Remove Removes a single item from the tree.
func (q *Quadtree) Remove(itemEnv *envelope.Envelope, item interface{}) bool {
	posEnv := EnsureExtent(itemEnv, q.MinExtent)
	return q.Root.Remove(posEnv, item)
}

// Query Queries the tree and returns items which may lie in the given search envelope.
func (q *Quadtree) Query(searchEnv *envelope.Envelope) interface{} {
	visitor := &index.ArrayVisitor{}
	if err := q.QueryVisitor(searchEnv, visitor); err != nil {
		log.Println(err)
	}
	return visitor.Items()
}

// QueryVisitor Queries the tree and visits items which may lie in the given search envelope.
func (q *Quadtree) QueryVisitor(searchEnv *envelope.Envelope, visitor index.ItemVisitor) error {
	q.Root.Visit(searchEnv, visitor)
	return nil
}

// CollectStats ...
func (q *Quadtree) CollectStats(itemEnv *envelope.Envelope) {
	delX := itemEnv.Width()
	if delX < q.MinExtent && delX > 0.0 {
		q.MinExtent = delX
	}
	delY := itemEnv.Height()
	if delY < q.MinExtent && delY > 0.0 {
		q.MinExtent = delY
	}
}

var (
	_ index.SpatialIndex = &Quadtree{}
)
