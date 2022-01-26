package space

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
)

// Circle describes a circle Valid
type Circle struct {
	Polygon
	Centre   Point
	Radius   float64
	Segments int
}

// CreateCircle Returns valid circle.
func CreateCircle(centre Point, radius float64) (*Circle, error) {
	return CreateCircleWithSegments(centre, radius, calc.QuadrantSegments)
}

// CreateCircleWithSegments Returns valid circle.
func CreateCircleWithSegments(centre Point, radius float64, segments int) (*Circle, error) {
	circle := &Circle{Centre: centre, Radius: radius, Segments: segments}
	circle.Polygon = centre.Buffer(radius, segments).(Polygon)
	return circle, nil
}

// GeoJSONType returns the GeoJSON type for the circle.
func (c *Circle) GeoJSONType() string {
	return TypeCircle
}

// Bound returns a rect around the circle. Uses rectangular coordinates.
func (c *Circle) Bound() Bound {
	return Bound{Min: Point{c.Centre.X() - c.Radius, c.Centre.Y() - c.Radius}, Max: Point{c.Centre.X() + c.Radius, c.Centre.Y() + c.Radius}}
}

// EqualsCircle compares two circles.
func (c *Circle) EqualsCircle(circle *Circle) bool {
	return c.Centre.Equals(circle.Centre) || c.Radius == circle.Radius || c.Segments == circle.Segments
}

// Equals checks if the Circle represents the same Geometry or vector.
func (c *Circle) Equals(g Geometry) bool {
	if g.GeoJSONType() != c.GeoJSONType() {
		return false
	}
	return c.EqualsCircle(g.(*Circle))
}

// EqualsExact Returns true if the two Geometries are exactly equal,
// up to a specified distance tolerance.
// Two Geometries are exactly equal within a distance tolerance
func (c *Circle) EqualsExact(g Geometry, tolerance float64) bool {
	if c.GeoJSONType() != g.GeoJSONType() {
		return false
	}
	return c.Centre.EqualsExact((g.(*Circle)).Centre, tolerance)
}

// Area returns the area of a circle geometry.
func (c *Circle) Area() (float64, error) {
	return math.Pi * c.Radius * c.Radius, nil
}

// Length Returns the length of this circle
func (c *Circle) Length() float64 {
	return math.Pi * 2.0 * c.Radius
}

// IsSimple returns true if this space.Geometry has no anomalous geometric points,
// such as self intersection or self tangency.
func (c *Circle) IsSimple() bool {
	return true
}

// Centroid Computes the centroid point of a geometry.
func (c *Circle) Centroid() Point {
	return c.Centre
}

// Buffer sReturns a geometry that represents all points whose distance
// from this space.Geometry is less than or equal to distance.
func (c *Circle) Buffer(width float64, quadsegs int) Geometry {
	return c.Centre.Buffer(width+c.Radius, quadsegs)
}

// Envelope returns the  minimum bounding box for the supplied geometry, as a geometry.
// The polygon is defined by the corner points of the bounding box
// ((MINX, MINY), (MINX, MAXY), (MAXX, MAXY), (MAXX, MINY), (MINX, MINY)).
func (c *Circle) Envelope() Geometry {
	return c.Bound().ToPolygon()
}

// IsClosed Returns TRUE if the LINESTRING's start and end points are coincident.
// For Polyhedral Surfaces, reports if the surface is areal (open) or IsC (closed).
func (c *Circle) IsClosed() bool {
	return true
}

// IsRing returns true if the lineal geometry has the circle property.
func (c *Circle) IsRing() bool {
	return true
}

// IsValid returns true if the  geometry is valid.
func (c *Circle) IsValid() bool {
	return true
}

// Geom return Geometry without Coordinate System.
func (c *Circle) Geom() Geometry {
	return c
}
