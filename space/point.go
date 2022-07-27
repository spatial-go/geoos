// Package space A representation of a linear vector geometry.
// include point line ring polygon multipoint multiline multrpolygon collection bound.
package space

import (
	"math/rand"
	"reflect"

	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/buffer/simplify"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space/spaceerr"
)

// Point describes a geographic point
type Point matrix.Matrix

// GeoJSONType returns GeoJSON type for the point
func (p Point) GeoJSONType() string {
	return TypePoint
}

// Dimensions returns 0 because a point is a 0d object.
func (p Point) Dimensions() int {
	return 0
}

// Bound returns a single point bound of the point.
func (p Point) Bound() Bound {
	return Bound{p, p}
}

// Nums num of points
func (p Point) Nums() int {
	return 1
}

// IsCollection returns true if the Geometry is  collection.
func (p Point) IsCollection() bool {
	return false
}

// Lat returns the vertical, latitude coordinate of the point.
func (p Point) Lat() float64 {
	return p[1]
}

// Lon returns the horizontal, longitude coordinate of the point.
func (p Point) Lon() float64 {
	return p[0]
}

// Y returns the vertical coordinate of the point.
func (p Point) Y() float64 {
	return p[1]
}

// X returns the horizontal coordinate of the point.
func (p Point) X() float64 {
	return p[0]
}

// EqualsPoint checks if the point represents the same point or vector.
func (p Point) EqualsPoint(point Point) bool {
	return matrix.Matrix(p).Equals(matrix.Matrix(point))
}

// Equals checks if the point represents the same Geometry or vector.
func (p Point) Equals(g Geometry) bool {
	if g.GeoJSONType() != p.GeoJSONType() {
		return false
	}
	return p.EqualsPoint(g.(Point))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (p Point) EqualsExact(g Geometry, tolerance float64) bool {
	if g.GeoJSONType() != p.GeoJSONType() {
		return false
	}
	if p.IsEmpty() && g.IsEmpty() {
		return true
	}
	if p.IsEmpty() != g.IsEmpty() {
		return false
	}
	if tolerance == 0 {
		return p.Equals(g)
	}
	return measure.PlanarDistance(matrix.Matrix(p), matrix.Matrix(g.(Point))) <= tolerance
}

// Generate implements the Generator interface for Points
func (p Point) Generate(r *rand.Rand, _ int) reflect.Value {
	for i := range p {
		p[i] = r.Float64()
	}
	return reflect.ValueOf(p)
}

// ToMatrix returns the Steric of a  geometry.
func (p Point) ToMatrix() matrix.Steric {
	return matrix.Matrix(p)
}

// Area returns the area of a polygonal geometry. The area of a point is 0.
func (p Point) Area() (float64, error) {
	return 0.0, nil
}

// IsEmpty returns true if the Geometry is empty.
func (p Point) IsEmpty() bool {
	return len(p) == 0
}

// Distance returns distance Between the two Geometry.
func (p Point) Distance(g Geometry) (float64, error) {
	return Distance(p, g, measure.PlanarDistance)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (p Point) SpheroidDistance(g Geometry) (float64, error) {
	return Distance(p, g, measure.SpheroidDistance)
}

// Boundary returns the closure of the combinatorial boundary of this Geometry.
func (p Point) Boundary() (Geometry, error) {
	return nil, spaceerr.ErrBoundBeNil
}

// IsSimple returns true if this Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (p Point) IsSimple() bool {
	return true
}

// Length Returns the length of this geometry
func (p Point) Length() float64 {
	return 0.0
}

// Centroid Computes the centroid point of a geometry.
func (p Point) Centroid() Point {
	return p
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func (p Point) UniquePoints() MultiPoint {
	return MultiPoint{p}
}

// Simplify returns a "simplified" version of the given geometry using the Douglas-Peucker algorithm,
// May not preserve topology
func (p Point) Simplify(tolerance float64) Geometry {
	result := simplify.Simplify(p.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// SimplifyP returns a geometry simplified by amount given by tolerance.
// Unlike Simplify, SimplifyP guarantees it will preserve topology.
func (p Point) SimplifyP(tolerance float64) Geometry {
	tls := &simplify.TopologyPreservingSimplifier{}
	result := tls.Simplify(p.ToMatrix(), tolerance)
	return TransGeometry(result)
}

// Buffer Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (p Point) Buffer(width float64, quadsegs int) Geometry {
	buff := buffer.Buffer(p.ToMatrix(), width, quadsegs)
	switch b := buff.(type) {
	case matrix.LineMatrix:
		return LineString(b)
	case matrix.PolygonMatrix:
		return Polygon(b)
	}
	return nil
}

// BufferInMeter Returns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (p Point) BufferInMeter(width float64, quadsegs int) Geometry {
	return BufferInMeter(p, width, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (p Point) Envelope() Geometry {
	return p
}

// ConvexHull computes the convex hull of a geometry. The convex hull is the smallest convex geometry
// that encloses all geometries in the input.
// In the general case the convex hull is a Polygon.
// The convex hull of two or more collinear points is a two-point LineString.
// The convex hull of one or more identical points is a Point.
func (p Point) ConvexHull() Geometry {
	result := buffer.ConvexHullWithGeom(p.ToMatrix()).ConvexHull()
	return TransGeometry(result)
}

// PointOnSurface Returns a POINT guaranteed to intersect a surface.
func (p Point) PointOnSurface() Geometry {
	m := buffer.InteriorPoint(p.ToMatrix())
	return Point(m)
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (p Point) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the ring property.
func (p Point) IsRing() bool {
	return p.IsClosed() && p.IsSimple()
}

// IsValid returns true if the  geometry is valid.
func (p Point) IsValid() bool {
	return p.IsCorrect()
}

// IsCorrect returns true if the geometry struct is Correct.
func (p Point) IsCorrect() bool {
	return len(p) >= 2
}

// CoordinateSystem return Coordinate System.
func (p Point) CoordinateSystem() int {
	return defaultCoordinateSystem()
}

// Filter Performs an operation with the provided .
func (p Point) Filter(f matrix.Filter) Geometry {
	return p
}

// Geom return Geometry without Coordinate System.
func (p Point) Geom() Geometry {
	return p
}
