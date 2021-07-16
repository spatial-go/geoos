package space

import (
	"errors"
	"math"
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

// ToPolygon converts the bound into a Polygon object.
func (b Bound) ToPolygon() Polygon {
	return Polygon{b.ToRing()}
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

// EqualBound returns if two bounds are equal.
func (b Bound) EqualBound(c Bound) bool {
	return b.Min.EqualPoint(c.Min) && b.Max.EqualPoint(c.Max)
}

// Equal checks if the Bound represents the same Geometry or vector.
func (b Bound) Equal(g Geometry) bool {
	if g.GeoJSONType() != b.GeoJSONType() {
		return false
	}
	return b.EqualBound(g.(Bound))
}

// EqualsExact Returns true if the two Geometrys are exactly equal,
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
		return 0, errors.New("Geometry is nil")
	}
	return b.ToRing().Distance(g)
}

// SpheroidDistance returns  spheroid distance Between the two Geometry.
func (b Bound) SpheroidDistance(g Geometry) (float64, error) {
	if b.IsEmpty() && g.IsEmpty() {
		return 0, nil
	}
	if b.IsEmpty() != g.IsEmpty() {
		return 0, errors.New("Geometry is nil")
	}
	return b.ToRing().SpheroidDistance(g)
}

// Boundary returns the closure of the combinatorial boundary of this space.Geometry.
func (b Bound) Boundary() (Geometry, error) {
	return nil, errors.New("Bound's boundary should be nil")
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
