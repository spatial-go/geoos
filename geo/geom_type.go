package geo
/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

// GeometryType ...
type GeometryType C.int

const (
	// POINT is a 0-dimensional geometric object, a single location is geometric
	// space.
	POINT GeometryType = C.GEOS_POINT
	// LINESTRING is a curve with linear interpolation between points.
	LINESTRING GeometryType = C.GEOS_LINESTRING
	// LINEARRING is a linestring that is both closed and simple.
	LINEARRING GeometryType = C.GEOS_LINEARRING
	// POLYGON is a planar surface with 1 exterior boundary and 0 or more
	// interior boundaries.
	POLYGON GeometryType = C.GEOS_POLYGON
	// MULTIPOINT is a 0-dimensional geometry collection, the elements of which
	// are restricted to points.
	MULTIPOINT GeometryType = C.GEOS_MULTIPOINT
	// MULTILINESTRING is a 1-dimensional geometry collection, the elements of
	// which are restricted to linestrings.
	MULTILINESTRING GeometryType = C.GEOS_MULTILINESTRING
	// MULTIPOLYGON is a 2-dimensional geometry collection, the elements of
	// which are restricted to polygons.
	MULTIPOLYGON GeometryType = C.GEOS_MULTIPOLYGON
	// GEOMETRYCOLLECTION is a geometric object that is a collection of some
	// number of geometric objects.
	GEOMETRYCOLLECTION GeometryType = C.GEOS_GEOMETRYCOLLECTION
)

var cGeomTypeIds = map[C.int]GeometryType{
	C.GEOS_POINT:              POINT,
	C.GEOS_LINESTRING:         LINESTRING,
	C.GEOS_LINEARRING:         LINEARRING,
	C.GEOS_POLYGON:            POLYGON,
	C.GEOS_MULTIPOINT:         MULTIPOINT,
	C.GEOS_MULTILINESTRING:    MULTILINESTRING,
	C.GEOS_MULTIPOLYGON:       MULTIPOLYGON,
	C.GEOS_GEOMETRYCOLLECTION: GEOMETRYCOLLECTION,
}

var geometryTypes = map[GeometryType]string{
	POINT:              "Point",
	LINESTRING:         "LineString",
	LINEARRING:         "LinearRing",
	POLYGON:            "Polygon",
	MULTIPOINT:         "MultiPoint",
	MULTILINESTRING:    "MultiLineString",
	MULTIPOLYGON:       "MultiPolygon",
	GEOMETRYCOLLECTION: "GeometryCollection",
}
// Type returns the SFS type of the geometry.
func  Type(g GEOSGeometry) (GeometryType, error) {
	i := C.GEOSGeomTypeId_r(geosContext,g)
	if i == -1 {
		return -1, Error()
	}
	return cGeomTypeIds[i], nil
}