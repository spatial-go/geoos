package geoc

// #cgo LDFLAGS: -lgeos_c
// #include "geos.h"
import "C"
import (
	"fmt"
)

// Coordinate coord
type Coordinate struct {
	X float64
	Y float64
	Z float64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}

// PrepareGeometry ...
func PrepareGeometry(g GEOSGeometry) GEOSPreparedGeometry {
	ptr := C.GEOSPrepare_r(geosContext, g)
	return ptr
}
