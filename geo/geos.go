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

// 计算面积
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
