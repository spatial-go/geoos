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
	"fmt"
)

type GEOSContext C.GEOSContextHandle_t

type GEOSGeometry *C.GEOSGeometry
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

func IsSimple(wkt string) (bool, error) {
	geoGeom := GeomFromWKTStr(wkt)
	c := C.GEOSisSimple_r(geosContext, geoGeom)
	b, e := boolFromC(c)
	if e != nil {
		return false, e
	}
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return b, nil
}

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

func IsEmpty(g string) (bool, error) {
	geoGeom := GeomFromWKTStr(g)
	c := C.GEOSisEmpty_r(geosContext, geoGeom)
	b, e := boolFromC(c)
	if e != nil {
		return false, e
	}
	C.GEOSGeom_destroy_r(geosContext, geoGeom)
	return b, nil
}

func Crosses(g1 string, g2 string) (bool, error) {
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	c := C.GEOSCrosses_r(geosContext, geom1, geom2)
	b, e := boolFromC(c)
	if e != nil {
		return false, e
	}
	C.GEOSGeom_destroy_r(geosContext, geom1)
	C.GEOSGeom_destroy_r(geosContext, geom2)
	return b, nil

}

func Within(g1 string, g2 string) (bool, error) {

	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	c := C.GEOSWithin_r(geosContext, geom1, geom2)
	b, e := boolFromC(c)
	if e != nil {
		return false, e
	}
	C.GEOSGeom_destroy_r(geosContext, geom1)
	C.GEOSGeom_destroy_r(geosContext, geom2)
	return b, nil

}
func Contains(g1 string,g2 string)(bool,error){
	geom1 := GeomFromWKTStr(g1)
	geom2 := GeomFromWKTStr(g2)
	c := C.GEOSContains_r(geosContext, geom1, geom2)
	b, e := boolFromC(c)
	if e != nil {
		return false, e
	}
	C.GEOSGeom_destroy_r(geosContext, geom1)
	C.GEOSGeom_destroy_r(geosContext, geom2)
	return b, nil
}