package geoos

// Ring represents a set of ring on the earth.
type Ring LineString

// GeoJSONType returns the GeoJSON type for the ring.
func (r Ring) GeoJSONType() string {
	return TypePolygon
}

// Dimensions returns 2 because a Ring is a 2d object.
func (r Ring) Dimensions() int {
	return 2
}

// Nums num of linstrings
func (r Ring) Nums() int {
	return 1
}

// Bound returns a rect around the ring. Uses rectangular coordinates.
func (r Ring) Bound() Bound {
	return MultiPoint(r).Bound()
}

// EqualRing compares two rings. Returns true if lengths are the same
// and all points are Equal.
func (r Ring) EqualRing(ring Ring) bool {
	return MultiPoint(r).Equal(MultiPoint(ring))
}

// Equal checks if the Ring represents the same Geometry or vector.
func (r Ring) Equal(g Geometry) bool {
	if g.GeoJSONType() != r.GeoJSONType() {
		return false
	}
	return r.EqualRing(g.(Ring))
}

// Area returns the area of a polygonal geometry. The area of a ring is 0.
func (r Ring) Area() (float64, error) {
	return 0.0, nil
}
