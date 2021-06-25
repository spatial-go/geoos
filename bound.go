package geoos

import (
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
	return TypePolygon
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
		Point{b.Max.X(), b.Min.Y()},
		b.Max,
		Point{b.Min.X(), b.Max.Y()},
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

// Bound returns the the same bound.
func (b Bound) Bound() Bound {
	return b
}

// Equal returns if two bounds are equal.
func (b Bound) Equal(c Bound) bool {
	return b.Min == c.Min && b.Max == c.Max
}

// Area returns the area of a polygonal geometry. The area of a bound is 0.
func (b Bound) Area() (float64, error) {
	return 0.0, nil
}
