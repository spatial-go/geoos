package index

// ArrayVisitor Builds an array of all visited items.
type ArrayVisitor struct {
	Items []interface{}
}

// VisitItem Visits an item.
func (a *ArrayVisitor) VisitItem(item interface{}) {
	a.Items = append(a.Items, item)
}
