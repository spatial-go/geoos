package geo
/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

// GeometryType represents the various geometry types supported by GEOS, and
// correspond to OGC Simple Features geometry types.
type GEOSGeometryType C.int

const (
	// POINT is a 0-dimensional geometric object, a single location is geometric
	// space.
	POINT GEOSGeometryType = C.GEOS_POINT
	// LINESTRING is a curve with linear interpolation between points.
	LINESTRING GEOSGeometryType = C.GEOS_LINESTRING
	// LINEARRING is a linestring that is both closed and simple.
	LINEARRING GEOSGeometryType = C.GEOS_LINEARRING
	// POLYGON is a planar surface with 1 exterior boundary and 0 or more
	// interior boundaries.
	POLYGON GEOSGeometryType = C.GEOS_POLYGON
	// MULTIPOINT is a 0-dimensional geometry collection, the elements of which
	// are restricted to points.
	MULTIPOINT GEOSGeometryType = C.GEOS_MULTIPOINT
	// MULTILINESTRING is a 1-dimensional geometry collection, the elements of
	// which are restricted to linestrings.
	MULTILINESTRING GEOSGeometryType = C.GEOS_MULTILINESTRING
	// MULTIPOLYGON is a 2-dimensional geometry collection, the elements of
	// which are restricted to polygons.
	MULTIPOLYGON GEOSGeometryType = C.GEOS_MULTIPOLYGON
	// GEOMETRYCOLLECTION is a geometric object that is a collection of some
	// number of geometric objects.
	GEOMETRYCOLLECTION GEOSGeometryType = C.GEOS_GEOMETRYCOLLECTION
)
var cGeomTypeIds = map[C.int]GEOSGeometryType{
	C.GEOS_POINT:              POINT,
	C.GEOS_LINESTRING:         LINESTRING,
	C.GEOS_LINEARRING:         LINEARRING,
	C.GEOS_POLYGON:            POLYGON,
	C.GEOS_MULTIPOINT:         MULTIPOINT,
	C.GEOS_MULTILINESTRING:    MULTILINESTRING,
	C.GEOS_MULTIPOLYGON:       MULTIPOLYGON,
	C.GEOS_GEOMETRYCOLLECTION: GEOMETRYCOLLECTION,
}

var geometryTypes = map[GEOSGeometryType]string{
	POINT:              "Point",
	LINESTRING:         "LineString",
	LINEARRING:         "LinearRing",
	POLYGON:            "Polygon",
	MULTIPOINT:         "MultiPoint",
	MULTILINESTRING:    "MultiLineString",
	MULTIPOLYGON:       "MultiPolygon",
	GEOMETRYCOLLECTION: "GeometryCollection",
}
