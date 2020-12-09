package geom

import "C"
import (
	"github.com/spatial-go/geos/geo"
)

// EmptyPolygon returns a new geometry of type Polygon that's empty (i.e.,
// IsEmpty() == true).
func EmptyPolygon() (geo.GEOSGeometry, error) {
	return nil, nil
}

// NewPolygon returns a new geometry of type Polygon, initialized with the given
// shell (exterior ring) and slice of holes (interior rings). The shell and holes
// slice are themselves slices of coordinates. A shell is required, and a
// variadic number of holes (therefore are optional).
//
// To create a new polygon from existing linear ring Geometry objects, use
// PolygonFromGeom.
func NewPolygon(shell []Coord, holes ...[]Coord) (geo.GEOSGeometry, error) {
	return nil,nil
	
}

// PolygonFromGeom returns a new geometry of type Polygon, initialized with the
// given shell (exterior ring) and slice of holes (interior rings). The shell
// and slice of holes are geometry objects, and expected to be LinearRings.
func PolygonFromGeom(shell geo.GEOSGeometry, holes ...geo.GEOSGeometry) (geo.GEOSGeometry, error) {
	
	return nil,nil
}

