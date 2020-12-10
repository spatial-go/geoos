package base

type Point struct {
	X float64
	Y float64
}

func (p Point) GeoJSONType() string {
	return TypePoint
}
// Dimensions returns 0 because a point is a 0d object.
func (p Point) Dimensions() int {
	return 0
}


// Y returns the vertical coordinate of the point.
func (p Point) GetY() float64 {
	return p.Y
}

// X returns the horizontal coordinate of the point.
func (p Point) GetX() float64 {
	return p.X
}

// Lat returns the vertical, latitude coordinate of the point.
func (p Point) GetLat() float64 {
	return p.X
}

// Lon returns the horizontal, longitude coordinate of the point.
func (p Point) GetLon() float64 {
	return p.Y
}


