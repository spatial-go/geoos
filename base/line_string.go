package base

// LineString represents a set of points to be thought of as a polyline.
type LineString []Point

// GeoJSONType returns the GeoJSON type for the object.
func (ls LineString) GeoJSONType() string {
	return TypeLineString
}
// Dimensions returns 1 because a LineString is a 1d object.
func (ls LineString) Dimensions() int {
	return 1
}

func (ls LineString) Boundary() (*Geometry ,error){
	s := NormalStrategy()
	return s.Boundary(ls)
}