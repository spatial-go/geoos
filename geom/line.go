package geom

import "github.com/spatial-go/geos/geo"

// NewLinearRing returns a new geometry of type LinearRing, initialized with the
// given coordinates. The number of coordinates must either be zero (none
// given), in which case it's an empty geometry (IsEmpty() == true), or >= 4.
func NewLinearRing(coords ...Coord) (geo.GEOSGeometry, error) {
	return nil,nil
}

// NewLineString returns a new geometry of type LineString, initialized with the
// given coordinates. If no coordinates are given, it's an empty geometry
// (IsEmpty() == true).
func NewLineString(coords ...Coord) (geo.GEOSGeometry, error) {
	return nil, nil
}
