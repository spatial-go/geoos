package geoos

// A Collection is a collection of geometries that is also a Geometry.
type Collection []Geometry

// GeoJSONType returns the geometry collection type.
func (c Collection) GeoJSONType() string {
	return TypeCollection
}

// Dimensions returns the max of the dimensions of the collection.
func (c Collection) Dimensions() int {
	max := -1
	for _, g := range c {
		if d := g.Dimensions(); d > max {
			max = d
		}
	}
	return max
}

func (c Collection) Nums() int {
	return len(c)
}

// Equal compares two collections. Returns true if lengths are the same
// and all the sub geometries are the same and in the same order.
func (c Collection) Equal(collection Collection) bool {
	if len(c) != len(collection) {
		return false
	}
	for i, g := range c {
		if !Equal(g, collection[i]) {
			return false
		}
	}
	return true
}
