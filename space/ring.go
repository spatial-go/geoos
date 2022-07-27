package space

import (
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space/spaceerr"
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

// IsCollection returns true if the Geometry is  collection.
func (r Ring) IsCollection() bool {
	return false
}

// Bound returns a rect around the ring. Uses rectangular coordinates.
func (r Ring) Bound() Bound {
	return LineString(r).Bound()
}

// EqualsRing compares two rings. Returns true if lengths are the same
// and all points are Equal.
func (r Ring) EqualsRing(ring Ring) bool {
	return LineString(r).Equals(LineString(ring))
}

// Equals checks if the Ring represents the same Geometry or vector.
func (r Ring) Equals(g Geometry) bool {
	if g.GeoJSONType() != r.GeoJSONType() {
		return false
	}
	return r.EqualsRing(g.(Ring))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (r Ring) EqualsExact(g Geometry, tolerance float64) bool {
	if r.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	return LineString(r).Equals(LineString(g.(Ring)))
}

// Area returns the area of a polygonal geometry.
func (r Ring) Area() (float64, error) {
	return measure.Area(r.ToMatrix().(matrix.LineMatrix)), nil
}

// ToMatrix returns the LineMatrix of a Ring geometry.
func (r Ring) ToMatrix() matrix.Steric {
	return LineString(r).ToMatrix()
}

// IsEmpty returns true if the Geometry is empty.
func (r Ring) IsEmpty() bool {
	return len(r) == 0
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
	return nil, spaceerr.ErrBoundBeNil
}

// Length Returns the length of this LineString
func (r Ring) Length() float64 {
	return LineString(r).Length()
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (r Ring) IsSimple() bool {
	return LineString(r).IsSimple()
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

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (r Ring) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(r.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (r Ring) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(r.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (r Ring) Buffer(width float64, quadsegs int) Geometry {
	return LineString(r).Buffer(width, quadsegs)
}

// BufferInMeter Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (r Ring) BufferInMeter(width float64, quadsegs int) Geometry {
	return LineString(r).BufferInMeter(width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (r Ring) Envelope() Geometry {
	return LineString(r).Envelope()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (r Ring) ConvexHull() Geometry {
	return LineString(r).ConvexHull()
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (r Ring) PointOnSurface() Geometry {
	return LineString(r).PointOnSurface()
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (r Ring) IsClosed() bool {
	return LineString(r).IsClosed()
}

// IsRing returns true if the lineal geometry has the ring property.
func (r Ring) IsRing() bool {
	return LineString(r).IsRing()
}

// IsValid returns true if the geometry is valid.
func (r Ring) IsValid() bool {
	return LineString(r).IsValid() && (!r.IsEmpty()) && r.IsRing()
}

// IsCorrect returns true if the geometry struct is Correct.
func (r Ring) IsCorrect() bool {
	return LineString(r).IsCorrect() && (!r.IsEmpty())
}

// CoordinateSystem return Coordinate System.
func (r Ring) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (r Ring) Filter(f matrix.Filter) Geometry {
	line := LineString(r).Filter(f)
	if f.IsChanged() {
		return append(Ring(line.(LineString)), line.(LineString)[0])
	}
	return r
}

// Geom return Geometry without Coordinate System.
func (r Ring) Geom() Geometry {
	return r
}
