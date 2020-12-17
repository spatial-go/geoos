package geos

// Polygon is a closed area. The first LineString is the outer ring.
// The others are the holes. Each LineString is expected to be closed
// ie. the first point matches the last.
type Polygon []Ring

// GeoJSONType returns the GeoJSON type for the object.
func (p Polygon) GeoJSONType() string {
	return TypePolygon
}

// Dimensions returns 2 because a Polygon is a 2d object.
func (p Polygon) Dimensions() int {
	return 2
}

// Nums num of polygons
func (p Polygon) Nums() int {
	return 1
}

func (p Polygon) Area() (float64, error) {
	s := NormalStrategy()
	return s.Area(p)
}

func (p Polygon) Boundary() (Geometry, error) {
	s := NormalStrategy()
	return s.Boundary(p)
}
