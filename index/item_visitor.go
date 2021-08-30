package index

// ItemVisitor A visitor for items in a SpatialIndex.
type ItemVisitor interface {
	// VisitItem Visits an item in the index.
	VisitItem(item interface{})

	Items() interface{}
}

// compile time checks
var (
	_ ItemVisitor = &ArrayVisitor{}
	_ ItemVisitor = &LineSegmentVisitor{}
)
