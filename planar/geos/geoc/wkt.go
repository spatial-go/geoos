package geoc

//
// #cgo LDFLAGS: -lgeos_c
// #include "geos.h"
//
import "C"
import (
	"errors"
	"unsafe"
)

// GEOSWKTReader ...
type GEOSWKTReader *C.GEOSWKTReader

// GEOSWKTWriter ...
type GEOSWKTWriter *C.GEOSWKTWriter

// ToWKTStr convert GEOSGeometry to WKT string
func ToWKTStr(g GEOSGeometry) (string, error) {
	w := WKTWriterFactory()
	defer WKTWriterDestroy(w)
	return DecodeWKTToStr(w, g)
}

// GeomFromWKTStr convert WKT string to GEOSGeometry
func GeomFromWKTStr(wktstr string) GEOSGeometry {
	reader := WKTReaderFactory()
	defer WKTReaderDestroy(reader)
	return EncodeWKTToGeom(reader, wktstr)
}

// WKTWriterFactory ...
func WKTWriterFactory() GEOSWKTWriter {
	writer := C.GEOSWKTWriter_create_r(geosContext)
	return GEOSWKTWriter(writer)
}

// WKTWriterDestroy ...
func WKTWriterDestroy(writer GEOSWKTWriter) {
	C.GEOSWKTWriter_destroy_r(geosContext, writer)
}

// WKTReaderFactory ...
func WKTReaderFactory() GEOSWKTReader {
	reader := C.GEOSWKTReader_create_r(geosContext)
	return GEOSWKTReader(reader)
}

// WKTReaderDestroy ...
func WKTReaderDestroy(reader GEOSWKTReader) {
	C.GEOSWKTReader_destroy_r(geosContext, reader)
}

// DecodeWKTToStr decode GEOSGeometry to WKT string
func DecodeWKTToStr(writer GEOSWKTWriter, g GEOSGeometry) (string, error) {
	cstr := C.GEOSWKTWriter_write_r(geosContext, writer, g)
	defer C.free(unsafe.Pointer(cstr))
	if cstr == nil {
		return "", errors.New("writer to wkt is null")
	}
	return C.GoString(cstr), nil
}

// EncodeWKTToGeom encode WKT string to GEOSGeometry
func EncodeWKTToGeom(reader GEOSWKTReader, wktStr string) GEOSGeometry {
	cs := C.CString(wktStr)
	defer C.free(unsafe.Pointer(cs))
	g := C.GEOSWKTReader_read_r(geosContext, reader, cs)
	return GEOSGeometry(g)
}
