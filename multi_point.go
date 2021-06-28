package geoos

import "github.com/spatial-go/geoos/algorithm/matrix"

// A MultiPoint represents a set of points in the 2D Eucledian or Cartesian plane.
type MultiPoint matrix.LineMatrix

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPoint) GeoJSONType() string {
	return TypeMultiPoint
}

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPoint) Dimensions() int {
	return 0
}

// Nums num of multiPoint.
func (mp MultiPoint) Nums() int {
	return len(mp)
}

// Bound returns a bound around the points. Uses rectangular coordinates.
func (mp MultiPoint) Bound() Bound {
	if len(mp) == 0 {
		return emptyBound
	}

	b := Bound{mp[0], mp[0]}
	for _, p := range mp {
		b = b.Extend(p)
	}

	return b
}

// EqualMultiPoint compares two MultiPoint objects. Returns true if lengths are the same
// and all points are Equal, and in the same order.
func (mp MultiPoint) EqualMultiPoint(multiPoint MultiPoint) bool {
	if len(mp) != len(multiPoint) {
		return false
	}
	for i, v := range mp.ToPointArray() {
		if !v.Equal(Point(multiPoint[i])) {
			return false
		}
	}
	return true
}

// Equal checks if the MultiPoint represents the same Geometry or vector.
func (mp MultiPoint) Equal(g Geometry) bool {
	if g.GeoJSONType() != mp.GeoJSONType() {
		return false
	}
	return mp.EqualMultiPoint(g.(MultiPoint))
}

// Area returns the area of a polygonal geometry. The area of a multipoint is 0.
func (mp MultiPoint) Area() (float64, error) {
	return 0.0, nil
}

// ToPointArray returns the PointArray
func (mp MultiPoint) ToPointArray() (pa []Point) {
	for _, v := range mp {
		pa = append(pa, v)
	}
	return
}
