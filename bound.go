package geoos

var emptyBound = Bound{Min: Point{1, 1}, Max: Point{-1, -1}}

// A Bound represents a closed box or rectangle.
// To create a bound with two points you can do something like:
//	orb.MultiPoint{p1, p2}.Bound()
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
		Point{b.Max.X, b.Min.Y},
		b.Max,
		Point{b.Min.X, b.Max.Y},
		b.Min,
	}
}

// Equal returns if two bounds are equal.
func (b Bound) Equal(c Bound) bool {
	return b.Min == c.Min && b.Max == c.Max
}
