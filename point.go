package geos

// Point Describes a geographic point
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
