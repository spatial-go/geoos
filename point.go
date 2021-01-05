package geoos

// Point describes a geographic point
type Point struct {
	X float64
	Y float64
}

// GeoJSONType returns GeoJSON type for the point
func (p Point) GeoJSONType() string {
	return TypePoint
}

// Dimensions returns 0 because a point is a 0d object.
func (p Point) Dimensions() int {
	return 0
}

// Nums num of points
func (p Point) Nums() int {
	return 1
}

// Lat returns the vertical, latitude coordinate of the point.
func (p Point) Lat() float64 {
	return p.X
}

// Lon returns the horizontal, longitude coordinate of the point.
func (p Point) Lon() float64 {
	return p.Y
}

// Equal checks if the point represents the same point or vector.
func (p Point) Equal(point Point) bool {
	return p.X == point.X && p.Y == point.Y
}
