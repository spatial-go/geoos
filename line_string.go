package geoos

import "github.com/spatial-go/geoos/algorithm/matrix"

// LineString represents a set of points to be thought of as a polyline.
type LineString matrix.LineMatrix

// GeoJSONType returns the GeoJSON type for the linestring.
func (ls LineString) GeoJSONType() string {
	return TypeLineString
}

// Dimensions returns 1 because a LineString is a 1d object.
func (ls LineString) Dimensions() int {
	return 1
}

// Nums num of linestrings
func (ls LineString) Nums() int {
	return 1
}

// // Boundary returns the closure of the combinatorial boundary of this linestring.
// func (ls LineString) Boundary() (Geometry, error) {
// 	s := NormalStrategy()
// 	return s.Boundary(ls)
// }

// Bound returns a rect around the line string. Uses rectangular coordinates.
func (ls LineString) Bound() Bound {
	if len(ls) == 0 {
		return emptyBound
	}

	b := Bound{ls[0], ls[0]}
	for _, p := range ls {
		b = b.Extend(p)
	}

	return b
}

// EqualLineString compares two line strings. Returns true if lengths are the same
// and all points are Equal.
func (ls LineString) EqualLineString(lineString LineString) bool {
	if len(ls) != len(lineString) {
		return false
	}
	for i, v := range ls.ToPointArray() {
		if !v.Equal(Point(lineString[i])) {
			return false
		}
	}
	return true
}

// Equal checks if the LineString represents the same Geometry or vector.
func (ls LineString) Equal(g Geometry) bool {
	if g.GeoJSONType() != ls.GeoJSONType() {
		return false
	}
	return ls.EqualLineString(g.(LineString))
}

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (ls LineString) EqualsExact(g Geometry, tolerance float64) bool {
	if ls.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	line := g.(LineString)
	if ls.IsEmpty() && g.IsEmpty() {
		return true
	}
	if ls.IsEmpty() != g.IsEmpty() {
		return false
	}
	if len(ls) != len(line) {
		return false
	}

	for i, v := range ls {
		if Point(v).EqualsExact(Point(line[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry. The area of a LineString is 0.
func (ls LineString) Area() (float64, error) {
	return 0.0, nil
}

// ToPointArray returns the PointArray
func (ls LineString) ToPointArray() (la []Point) {
	for _, v := range ls {
		la = append(la, v)
	}
	return
}

// IsEmpty returns true if the Geometry is empty.
func (ls LineString) IsEmpty() bool {
	return ls == nil || len(ls) == 0
}
