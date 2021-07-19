package space

import (
	"github.com/spatial-go/geoos/algorithm/measure"
)

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

// EqualsExact Returns true if the two Geometrys are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (mls MultiLineString) EqualsExact(g Geometry, tolerance float64) bool {
	if mls.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	for i, v := range mls {
		if v.EqualsExact((g.(MultiLineString)[i]), tolerance) {
			return false
		}
	}
	return true
}

// Area returns the area of a polygonal geometry. The area of a MultiLineString is 0.
func (mls MultiLineString) Area() (float64, error) {
	return 0.0, nil
}

// IsEmpty returns true if the Geometry is empty.
func (mls MultiLineString) IsEmpty() bool {
	return mls == nil || len(mls) == 0
}

// Distance returns distance Between the two Geometry.
func (mls MultiLineString) Distance(g Geometry) (float64, error) {
	elem := &ElementDistance{mls}
	return elem.distanceWithFunc(g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (mls MultiLineString) SpheroidDistance(g Geometry) (float64, error) {
	elem := &ElementDistance{mls}
	return elem.distanceWithFunc(g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (mls MultiLineString) Boundary() (Geometry, error) {
	bdyPts := []Point{}
	for _, v := range mls {
		if len(v) == 0 {
			continue
		}
		bdyPts = append(bdyPts, v[0], v[len(v)-1])
	}
	// return Point or MultiPoint
	if len(bdyPts) == 1 {
		return bdyPts[0], nil
	}
	// this handles 0 points case as well
	return MultiPoint(bdyPts), nil
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (mls MultiLineString) IsClosed() bool {
	if mls.IsEmpty() {
		return false
	}
	for _, v := range mls {
		if !v.IsClosed() {
			return false
		}
	}
	return true
}

// Length Returns the length of this MultiLineString
func (mls MultiLineString) Length() float64 {
	length := 0.0
	for _, v := range mls {
		length += v.Length()
	}
	return length
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (mls MultiLineString) IsSimple() bool {
	elem := ElementValid{mls}
	return elem.IsSimple()
}

// Centroid Computes the centroid point of a geometry.
func (mls MultiLineString) Centroid() Point {
	return Centroid(mls)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (mls MultiLineString) UniquePoints() MultiPoint {
	mp := MultiPoint{}
	for _, v := range mls {
		mp = append(mp, v.UniquePoints()...)
	}
	return mp
}
