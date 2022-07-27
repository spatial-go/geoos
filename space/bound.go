package space

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/space/spaceerr"
)

var emptyBound = Bound{Min: Point{1, 1}, Max: Point{-1, -1}}

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
// MultiPoint{p1, p2}.Bound()
type Bound struct {
	Min, Max Point
}

// GeoJSONType returns the GeoJSON type for the object.
func (b Bound) GeoJSONType() string {
	return TypeBound
}

// Dimensions returns 2 because a Bound is a 2d object.
func (b Bound) Dimensions() int {
	return 2
}

// Nums ...
func (b Bound) Nums() int {
	return 2
}

// IsCollection returns true if the Geometry is  collection.
func (b Bound) IsCollection() bool {
	return false
}

// ToPolygon converts the bound into a Polygon object.
func (b Bound) ToPolygon() Polygon {
	return Polygon{b.ToRing()}
}

// ToMatrix returns the Steric of a  geometry.
func (b Bound) ToMatrix() matrix.Steric {
	return b.ToPolygon().ToMatrix()
}

// ToRing converts the bound into a loop defined
// by the boundary of the box.
func (b Bound) ToRing() Ring {
	return Ring{
		b.Min,
		{b.Max.X(), b.Min.Y()},
		b.Max,
		{b.Min.X(), b.Max.Y()},
		b.Min,
	}
}

// Extend grows the bound to include the new point.
func (b Bound) Extend(point Point) Bound {
	// already included, no big deal
	if b.Contains(point) {
		return b
	}

	return Bound{
		Min: Point{
			math.Min(b.Min[0], point[0]),
			math.Min(b.Min[1], point[1]),
		},
		Max: Point{
			math.Max(b.Max[0], point[0]),
			math.Max(b.Max[1], point[1]),
		},
	}
}

// Contains determines if the point is within the bound.
// Points on the boundary are considered within.
func (b Bound) Contains(point Point) bool {
	if point[1] < b.Min[1] || b.Max[1] < point[1] {
		return false
	}

	if point[0] < b.Min[0] || b.Max[0] < point[0] {
		return false
	}

	return true
}

// ContainsBound determines if the bound is within the bound.
func (b Bound) ContainsBound(bound Bound) bool {
	if b.IsEmpty() || bound.IsEmpty() {
		return false
	}
	return bound.Min.X() >= b.Min.X() &&
		bound.Max.X() <= b.Max.X() &&
		bound.Min.Y() >= b.Min.Y() &&
		bound.Max.Y() <= b.Max.Y()
}

// Bound returns the the same bound.
func (b Bound) Bound() Bound {
	return b
}

// EqualsBound returns if two bounds are equal.
func (b Bound) EqualsBound(c Bound) bool {
	return b.Min.EqualsPoint(c.Min) && b.Max.EqualsPoint(c.Max)
}

// Equals checks if the Bound represents the same Geometry or vector.
func (b Bound) Equals(g Geometry) bool {
	if g.GeoJSONType() != b.GeoJSONType() {
		return false
	}
	return b.EqualsBound(g.(Bound))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (b Bound) EqualsExact(g Geometry, tolerance float64) bool {
	if b.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	if b.IsEmpty() && g.IsEmpty() {
		return true
	}
	if b.IsEmpty() != g.IsEmpty() {
		return false
	}

	return b.Max.EqualsExact(g.(Bound).Max, tolerance) && b.Min.EqualsExact(g.(Bound).Min, tolerance)
}

// Area returns the area of a polygonal geometry. The area of a bound is 0.
func (b Bound) Area() (float64, error) {
	return b.ToPolygon().Area()
}

// IsEmpty returns true if it contains zero area or if
// it's in some malformed negative state where the left point is larger than the right.
// This can be caused by padding too much negative.
func (b Bound) IsEmpty() bool {
	if b.Max == nil || b.Min == nil {
		return true
	}
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

// Distance returns distance Between the two Geometry.
func (b Bound) Distance(g Geometry) (float64, error) {
	if b.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if b.IsEmpty() != g.IsEmpty() {
		return 0, spaceerr.ErrNilGeometry
	}
	return b.ToRing().Distance(g)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (b Bound) SpheroidDistance(g Geometry) (float64, error) {
	if b.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if b.IsEmpty() != g.IsEmpty() {
		return 0, spaceerr.ErrNilGeometry
	}
	return b.ToRing().SpheroidDistance(g)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (b Bound) Boundary() (Geometry, error) {
	return nil, spaceerr.ErrNotSupportBound
}

// Length Returns the length of this LineString
func (b Bound) Length() float64 {
	return b.ToRing().Length()
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (b Bound) IsSimple() bool {
	return true
}

// Centroid Computes the centroid point of a geometry.
func (b Bound) Centroid() Point {
	return Centroid(b.ToRing())
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (b Bound) UniquePoints() MultiPoint {
	return MultiPoint{b.Max, b.Min}
}

// IntersectsBound Tests if the region defined by other
// intersects the region of this Envelope.
func (b Bound) IntersectsBound(other Bound) bool {
	if b.IsEmpty() || other.IsEmpty() {
		return false
	}
	return !(other.Min.X() > b.Max.X() ||
		other.Max.X() < b.Min.X() ||
		other.Min.Y() > b.Max.Y() ||
		other.Max.Y() < b.Min.Y())
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (b Bound) Simplify(tolerance float64) Geometry {
	return b.ToPolygon()
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (b Bound) SimplifyP(tolerance float64) Geometry {
	return b.ToPolygon()
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (b Bound) Buffer(width float64, quadsegs int) Geometry {
	return b.ToPolygon().Buffer(width, quadsegs)
}

// BufferInMeter Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (b Bound) BufferInMeter(width float64, quadsegs int) Geometry {
	return b.ToPolygon().BufferInMeter(width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (b Bound) Envelope() Geometry {
	return b.ToPolygon()
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (b Bound) ConvexHull() Geometry {
	return b.ToPolygon().ConvexHull()
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (b Bound) PointOnSurface() Geometry {
	return b.ToPolygon().PointOnSurface()
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (b Bound) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the ring property.
func (b Bound) IsRing() bool {
	return b.IsClosed() && b.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (b Bound) IsValid() bool {
	return b.Min.IsValid() && b.Max.IsValid()
}

// IsCorrect returns true if the geometry struct is Correct.
func (b Bound) IsCorrect() bool {
	return b.Min.IsCorrect() && b.Max.IsCorrect()
}

// CoordinateSystem return Coordinate System.
func (b Bound) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (b Bound) Filter(f matrix.Filter) Geometry {
	return b
}

// Geom return Geometry without Coordinate System.
func (b Bound) Geom() Geometry {
	return b
}
