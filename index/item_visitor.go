package index

// ItemVisitor A visitor for items in a {@link SpatialIndex}.
type ItemVisitor interface {
	// VisitItem Visits an item in the index.
	VisitItem(item interface{})
}
