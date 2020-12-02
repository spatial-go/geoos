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

type GeosContextHandler struct {
	handler C.GEOSContextHandle_t
}

func InitGeosContext() GeosContextHandler{
	handler := C.geos_initGEOS()
	return GeosContextHandler{handler:handler}
}
func FinishGeosContext(c GeosContextHandler) {
	C.finishGEOS_r(c.handler)
}

func Version() string {
	return C.GoString(C.GEOSversion())
}

func Error() error {
	return fmt.Errorf("geo: %s", C.GoString(C.geos_get_last_error()))
}
