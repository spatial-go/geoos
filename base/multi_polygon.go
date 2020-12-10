package base

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

func (mp MultiPolygon)Area() (float64, error){
	s := NormalStrategy()
	return s.Area(mp)
}