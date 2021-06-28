package geoos

// MultiLineString is a set of polylines.
type MultiLineString []LineString

// GeoJSONType returns the GeoJSON type for the object.
func (mls MultiLineString) GeoJSONType() string {
	return TypeMultiLineString
}

// Dimensions returns 1 because a MultiLineString is a 2d object.
func (mls MultiLineString) Dimensions() int {
	return 1
}

// Nums num of multiLinstrings
func (mls MultiLineString) Nums() int {
	return len(mls)
}

// Bound returns a bound around all the line strings.
func (mls MultiLineString) Bound() Bound {
	if len(mls) == 0 {
		return emptyBound
	}

	bound := mls[0].Bound()
	for i := 1; i < len(mls); i++ {
		bound = bound.Union(mls[i].Bound())
	}

	return bound
}

// Union extends this bound to contain the union of this and the given bound.
func (b Bound) Union(other Bound) Bound {
	if other.IsEmpty() {
		return b
	}

	b = b.Extend(other.Min)
	b = b.Extend(other.Max)
	b = b.Extend(other.LeftTop())
	b = b.Extend(other.RightBottom())

	return b
}

// EqualMultiLineString compares two multi line strings. Returns true if lengths are the same
// and all points are Equal.
func (mls MultiLineString) EqualMultiLineString(multiLineString MultiLineString) bool {
	if len(mls) != len(multiLineString) {
		return false
	}
	for i, ls := range mls {
		if !ls.Equal(multiLineString[i]) {
			return false
		}
	}

	return true
}

// Equal checks if the MultiLineString represents the same Geometry or vector.
func (mls MultiLineString) Equal(g Geometry) bool {
	if g.GeoJSONType() != mls.GeoJSONType() {
		return false
	}
	return mls.EqualMultiLineString(g.(MultiLineString))
}

// Area returns the area of a polygonal geometry. The area of a MultiLineString is 0.
func (mls MultiLineString) Area() (float64, error) {
	return 0.0, nil
}
