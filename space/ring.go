package space

import (
	"errors"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// Ring represents a set of ring on the earth.
type Ring LineString

// GeoJSONType returns the GeoJSON type for the ring.
func (r Ring) GeoJSONType() string {
	return TypeLineString
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
	return LineString(r).Bound()
}

// EqualRing compares two rings. Returns true if lengths are the same
// and all points are Equal.
func (r Ring) EqualRing(ring Ring) bool {
	return LineString(r).Equal(LineString(ring))
}

// Equal checks if the Ring represents the same Geometry or vector.
func (r Ring) Equal(g Geometry) bool {
	if g.GeoJSONType() != r.GeoJSONType() {
		return false
	}
	return r.EqualRing(g.(Ring))
}

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (r Ring) EqualsExact(g Geometry, tolerance float64) bool {
	if r.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	return LineString(r).Equal(LineString(g.(Ring)))
}

// Area returns the area of a polygonal geometry. The area of a ring is 0.
func (r Ring) Area() (float64, error) {
	return measure.Area(r.ToMatrix()), nil
}

// ToMatrix returns the LineMatrix of a Ring geometry.
func (r Ring) ToMatrix() matrix.LineMatrix {
	return matrix.LineMatrix(r)
}

// IsEmpty returns true if the Geometry is empty.
func (r Ring) IsEmpty() bool {
	return r == nil || len(r) == 0
}

// Distance returns distance Between the two Geometry.
func (r Ring) Distance(g Geometry) (float64, error) {
	return LineString(r).Distance(g)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (r Ring) SpheroidDistance(g Geometry) (float64, error) {
	return LineString(r).SpheroidDistance(g)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
// The boundary of a lineal geometry is always a zero-dimensional geometry (which may be empty).
func (r Ring) Boundary() (Geometry, error) {
	return nil, errors.New("ring's boundary should be nil")
}

// Length Returns the length of this LineString
func (r Ring) Length() float64 {
	return LineString(r).Length()
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (r Ring) IsSimple() bool {
	elem := ElementValid{LineString(r)}
	return elem.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func (r Ring) Centroid() Point {
	return Centroid(LineString(r))
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (r Ring) UniquePoints() MultiPoint {
	mp := MultiPoint{}
	for i, v := range r {
		if i == len(r)-1 {
			break
		}
		mp = append(mp, v)
	}
	return mp
}
