package geos

// MultiPolygon is a set of polygons.
type MultiPolygon []Polygon

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPolygon) GeoJSONType() string {
	return TypeMultiPolygon
}

// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPolygon) Dimensions() int {
	return 2
}

// Nums num of polygons
func (mp MultiPolygon) Nums() int {
	return len(mp)
}

// Area Returns the area of this polygonal geometry
func (mp MultiPolygon) Area() (float64, error) {
	s := NormalStrategy()
	return s.Area(mp)
}
