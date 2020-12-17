// Package geo provides support for creating and manipulating spatial data.
// At its core, it relies on the GEOS C library for the implementation of
// spatial operations and geometric algorithms.
package geo

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

import (
	"errors"
	"fmt"
)

type GEOSContext C.GEOSContextHandle_t

type GEOSGeometry *C.GEOSGeometry
type GEOSPreparedGeometry *C.GEOSPreparedGeometry
type GEOSCoordSequence C.GEOSCoordSequence
type GEOSBufferParams C.GEOSBufferParams
type GEOSSTRtree C.GEOSSTRtree

type GEOSChar *C.char

var (
	geosContext = InitGeosContext()
)

func InitGeosContext() GEOSContext {
	c := C.geos_initGEOS()
	return GEOSContext(c)
}
func FinishGeosContext(c GEOSContext) {
	C.finishGEOS_r(c)
}

func Version() string {
	return C.GoString(C.GEOSversion())
}

func Error() error {
	return fmt.Errorf("geo: %s", C.GoString(C.geos_get_last_error()))
}

// Returns the area of a polygonal geometry
func Area(wkt string) (float64, error) {
	geoGeom := GeomFromWKTStr(wkt)
	var d C.double
	i := C.GEOSArea_r(geosContext, geoGeom, &d)
	if i == 0 {
		return 0.0, Error()
	}
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return float64(d), nil
}

// Returns the closure of the combinatorial boundary of this Geometry
func Boundary(wkt string) (string, error) {
	geoGeom := GeomFromWKTStr(wkt)
	g := C.GEOSBoundary_r(geosContext, geoGeom)
	s, e := ToWKTStr(g)
	if e != nil {
		return "", e
	}
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return s, nil
}

// Computes the geometric center of a geometry, or equivalently, the center of mass of the geometry as a POINT.
// For [MULTI]POINTs, this is computed as the arithmetic mean of the input coordinates.
// For [MULTI]LINESTRINGs, this is computed as the weighted length of each line segment.
// For [MULTI]POLYGONs, "weight" is thought in terms of area.
// If an empty geometry is supplied, an empty GEOMETRYCOLLECTION is returned.
// If NULL is supplied, NULL is returned.
// If CIRCULARSTRING or COMPOUNDCURVE are supplied, they are converted to linestring wtih CurveToLine first, then same than for LINESTRING
func Centroid(wkt string) (string, error) {
	geoGeom := GeomFromWKTStr(wkt)
	g := C.GEOSGetCentroid_r(geosContext, geoGeom)
	s, e := ToWKTStr(g)
	if e != nil {
		return "", e
	}
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return s, nil
}

// Returns true if this Geometry has no anomalous geometric points, such as self intersection or self tangency
func IsSimple(wkt string) (bool, error) {
	geoGeom := GeomFromWKTStr(wkt)
	defer C.GEOSGeom_destroy_r(geosContext, geoGeom)
	c := C.GEOSisSimple_r(geosContext, geoGeom)
	return boolFromC(c)
}

// returns the 2D Cartesian length of the geometry if it is a LineString, MultiLineString
func Length(wkt string) (float64, error) {
	geoGeom := GeomFromWKTStr(wkt)
	var d C.double
	i := C.GEOSLength_r(geosContext, geoGeom, &d)
	if i == 0 {
		return 0.0, Error()
	}
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return float64(d), nil

}

// returns the minimum 2D Cartesian (planar) distance between two geometries, in projected units (spatial ref units).
func Distance(g1 string, g2 string) (float64, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	var distance C.double
	i := C.GEOSDistance_r(geosContext, geom1, geom2, &distance)
	if i == 0 {
		return 0.0, Error()
	}
	C.GEOSGeom_destroy_r(geosContext, geom1)
	C.GEOSGeom_destroy_r(geosContext, geom2)
	return float64(distance), nil

}

// Returns the Hausdorff distance between two geometries, a measure of how similar or dissimilar 2 geometries are.
// Implements algorithm for computing a distance metric which can be thought of as the "Discrete Hausdorff Distance".
// This is the Hausdorff distance restricted to discrete points for one of the geometries
func HausdorffDistance(g1 string, g2 string) (float64, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)

	var distance C.double
	i := C.GEOSHausdorffDistance_r(geosContext, geom1, geom2, &distance)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(distance), nil
}

// Returns true if this Geometry is an empty geometry.
// If true, then this Geometry represents an empty geometry collection, polygon, point etc.
func IsEmpty(g string) (bool, error) {
	geoGeom := GeomFromWKTStr(g)
	defer C.GEOSGeom_destroy_r(geosContext, geoGeom)
	c := C.GEOSisEmpty_r(geosContext, geoGeom)
	return boolFromC(c)
}

// Envelope ...
func Envelope(wkt string) (string, error) {
	geoGeom := GeomFromWKTStr(wkt)
	g := C.GEOSEnvelope_r(geosContext, geoGeom)
	s, e := ToWKTStr(g)
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return s, e
}

// Crosses takes two geometry objects and returns TRUE if their intersection "spatially cross",
// that is, the geometries have some, but not all interior points in common.
// The intersection of the interiors of the geometries must not be the empty set and must have a dimensionality less than the maximum dimension of the two input geometries.
// Additionally, the intersection of the two geometries must not equal either of the source geometries. Otherwise, it returns FALSE.
func Crosses(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
	}()
	c := C.GEOSCrosses_r(geosContext, geom1, geom2)
	return boolFromC(c)
}

//Returns TRUE if geometry A is completely inside geometry B.
// For this function to make sense, the source geometries must both be of the same coordinate projection,
// having the same SRID.
func Within(g1 string, g2 string) (bool, error) {

	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
	}()
	c := C.GEOSWithin_r(geosContext, geom1, geom2)
	return boolFromC(c)
}

// Geometry A contains Geometry B if and only if no points of B lie in the exterior of A,
// and at least one point of the interior of B lies in the interior of A.
// An important subtlety of this definition is that A does not contain its boundary, but A does contain itself.
//Returns TRUE if geometry B is completely inside geometry A.
// For this function to make sense, the source geometries must both be of the same coordinate projection, having the same SRID.
func Contains(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
	}()
	c := C.GEOSContains_r(geosContext, geom1, geom2)
	return boolFromC(c)
}

// UniquePoints return all distinct vertices of input geometry as a MultiPoint.
func UniquePoints(g string) (string, error) {
	geom := GeomFromWKTStr(g)
	c := C.GEOSGeom_extractUniquePoints_r(geosContext, geom)
	if c == nil {
		return "", errors.New("UniquePoints return null")
	}
	wkt, e := ToWKTStr(c)
	if e != nil {
		return "", e
	}
	return wkt, nil

}

// Returns a collection containing paths shared by the two input geometries.
// Those going in the same direction are in the first element of the collection, those going in the opposite direction are in the second element.
// The paths themselves are given in the direction of the first geometry.
func SharedPaths(g1 string, g2 string) (string, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	g := C.GEOSSharedPaths_r(geosContext, geom1, geom2)
	wkt, e := ToWKTStr(g)
	if e != nil {
		return "", e
	}
	C.GEOSGeom_destroy_r(geosContext, geom1)
	C.GEOSGeom_destroy_r(geosContext, geom2)
	return wkt, nil
}

// Snaps the vertices and segments of a geometry to another Geometry's vertices.
// A snap distance tolerance is used to control where snapping is performed.
// The result geometry is the input geometry with the vertices snapped.
// If no snapping occurs then the input geometry is returned unchanged.
func Snap(input string, reference string, tolerance float64) (string, error) {
	inputGeom := GeomFromWKTStr(input)
	referenceGeom := GeomFromWKTStr(reference)
	g := C.GEOSSnap_r(geosContext, inputGeom, referenceGeom, C.double(tolerance))
	s, e := ToWKTStr(g)
	if e != nil {
		return "", e
	}
	return s, nil
}

// Buffer returns a geometry that represents all points whose distance from
// this Geometry is less than or equal to distance.
func Buffer(g string, width float64, quadsegs int32) (wkt string, err error) {
	geom := GeomFromWKTStr(g)
	defer C.GEOSGeom_destroy_r(geosContext, geom)
	bufferGeom := C.GEOSBuffer_r(geosContext, geom, C.double(width), C.int(quadsegs))
	if wkt, err = ToWKTStr(bufferGeom); err != nil {
		wkt = ""
	}
	return
}

// EqualsExact returns true if both geometries are Equal, as evaluated by their
// points being within the given tolerance.
func EqualsExact(g1 string, g2 string, tolerance float64) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
	}()
	c := C.GEOSEqualsExact_r(geosContext, geom1, geom2, C.double(tolerance))
	return boolFromC(c)
}

// NGeometry returns the number of component geometries.
func NGeometry(g string) (int, error) {
	geom := GeomFromWKTStr(g)
	defer C.GEOSGeom_destroy_r(geosContext, geom)
	c := C.GEOSGetNumGeometries_r(geosContext, geom)
	return intFromC(c, -1)
}

// Overlaps returns true if one geometry overlaps the other.
func Overlaps(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	pGeom := C.GEOSPrepare_r(geosContext, geom1)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
		C.GEOSPreparedGeom_destroy_r(geosContext, pGeom)
	}()
	c := C.GEOSPreparedOverlaps_r(geosContext, pGeom, geom2)
	return boolFromC(c)
}

// Equals returns true if the two geometries have at least one point in common.
func Equals(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
	}()
	c := C.GEOSEquals_r(geosContext, geom1, geom2)
	return boolFromC(c)
}

// Covers computes whether the prepared geometry covers the other.
func Covers(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	pGeom := C.GEOSPrepare_r(geosContext, geom1)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
		C.GEOSPreparedGeom_destroy_r(geosContext, pGeom)
	}()
	c := C.GEOSPreparedCovers_r(geosContext, pGeom, geom2)
	return boolFromC(c)
}

// CoveredBy computes whether the prepared geometry is covered by the other.
func CoversBy(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	pGeom := C.GEOSPrepare_r(geosContext, geom1)
	defer func() {
		C.GEOSGeom_destroy_r(geosContext, geom1)
		C.GEOSGeom_destroy_r(geosContext, geom2)
		C.GEOSPreparedGeom_destroy_r(geosContext, pGeom)
	}()
	c := C.GEOSPreparedCoveredBy_r(geosContext, pGeom, geom2)
	return boolFromC(c)
}

// IsRing returns true if the lineal geometry has the ring property.
func IsRing(g string) (bool, error) {
	geoGeom := GeomFromWKTStr(g)
	defer C.GEOSGeom_destroy_r(geosContext, geoGeom)
	c := C.GEOSisRing_r(geosContext, geoGeom)
	return boolFromC(c)
}

// IsClosed returns true if the geometry is closed
func IsClosed(g string) (bool, error) {
	geoGeom := GeomFromWKTStr(g)
	defer C.GEOSGeom_destroy_r(geosContext, geoGeom)
	c := C.GEOSisClosed_r(geosContext, geoGeom)
	return boolFromC(c)
}