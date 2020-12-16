package geos

// Ring represents a set of ring on the earth.
type Ring LineString

// GeoJSONType returns the GeoJSON type for the object.
func (r Ring) GeoJSONType() string {
	return TypePolygon
}

// Dimensions returns 2 because a Ring is a 2d object.
func (r Ring) Dimensions() int {
	return 2
}

// Nums num of linstrings
func (r Ring) Nums() int {
	return 1
}
