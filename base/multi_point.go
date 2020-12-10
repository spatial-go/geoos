package base

// A MultiPoint represents a set of points in the 2D Eucledian or Cartesian plane.
type MultiPoint []Point

// GeoJSONType returns the GeoJSON type for the object.
func (mp MultiPoint) GeoJSONType() string {
	return TypeMultiPoint
}
// Dimensions returns 0 because a MultiPoint is a 0d object.
func (mp MultiPoint) Dimensions() int {
	return 0
}
