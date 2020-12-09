package geo

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

// 点初始化
func NewPoint(coords []Coordinate) (GEOSGeometry, error) {
	if len(coords) == 0 {
		g := C.GEOSGeom_createEmptyPoint_r(geosContext)
		return GEOSGeometry(g), nil
	}
	coordinateSequence, e := makeCodinateSequence(coords)
	if e != nil {
		return nil, e
	}
	g := C.GEOSGeom_createPoint_r(geosContext, coordinateSequence)
	return GEOSGeometry(g), nil
}
