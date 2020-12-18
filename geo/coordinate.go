package geo
/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"
import (
	"fmt"
)

// 3维坐标点
type Coordinate struct {
	X float64
	Y float64
	Z float64
}

// 2维点的字符串表达
func (c Coordinate) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}

func PrepareGeometry(g GEOSGeometry) GEOSPreparedGeometry {
	ptr := C.GEOSPrepare_r(geosContext, g)
	return ptr
}

