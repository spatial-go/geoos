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

// Equal compares two multi line strings. Returns true if lengths are the same
// and all points are Equal.
func (mls MultiLineString) Equal(multiLineString MultiLineString) bool {
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

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool {
	return b.Min[0] > b.Max[0] || b.Min[1] > b.Max[1]
}

// Top returns the top of the bound.
func (b Bound) Top() float64 {
	return b.Max[1]
}

// Bottom returns the bottom of the bound.
func (b Bound) Bottom() float64 {
	return b.Min[1]
}

// Right returns the right of the bound.
func (b Bound) Right() float64 {
	return b.Max[0]
}

// Left returns the left of the bound.
func (b Bound) Left() float64 {
	return b.Min[0]
}

// LeftTop returns the upper left point of the bound.
func (b Bound) LeftTop() Point {
	return Point{b.Left(), b.Top()}
}

// RightBottom return the lower right point of the bound.
func (b Bound) RightBottom() Point {
	return Point{b.Right(), b.Bottom()}
}

// Area returns the area of a polygonal geometry. The area of a MultiLineString is 0.
func (mls MultiLineString) Area() (float64, error) {
	return 0.0, nil
}
