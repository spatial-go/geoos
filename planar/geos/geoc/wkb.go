package geoc

//
// #cgo LDFLAGS: -lgeos_c
// #include "geos.h"
//
import "C"

import (
	"encoding/hex"
	"errors"
	"unsafe"
)

// GEOSWKBReader ...
type GEOSWKBReader *C.GEOSWKBReader

// GEOSWKBWriter ...
type GEOSWKBWriter *C.GEOSWKBWriter

// GeomFromWKBStr convert wkb byte array to GEOSGeometry
func GeomFromWKBStr(wkbByte []byte) (GEOSGeometry, error) {
	cwkb := GoByteArrayToCCharArray(wkbByte)
	reader := WKBReaderFactory()
	defer WKBReaderDestroy(reader)
	return EncodeWKBToGeom(reader, cwkb)
}

// ToWKB convert GEOSGeometry to wkb byte array
func ToWKB(g GEOSGeometry) ([]byte, error) {
	writer := WKBWriterFactory()
	defer WKBWriterDestroy(writer)
	return DecodeWKBToArray(writer, g)
}

// ToWKBHex convert GEOSGeometry to hex string
func ToWKBHex(g GEOSGeometry) (string, error) {
	w := WKBWriterFactory()
	defer WKBWriterDestroy(w)
	return DecodeWKBToHexStr(w, g)
}

// GeomFromWKBHexStr convert hex string to GEOSGeometry
func GeomFromWKBHexStr(wkbHex string) (GEOSGeometry, error) {
	wkbstr, err := hex.DecodeString(wkbHex)
	if err != nil {
		return nil, err
	}
	array := GoByteArrayToCCharArray(wkbstr)
	r := WKBReaderFactory()
	defer WKBReaderDestroy(r)
	return EncodeHexToGeom(r, array)

}

// EncodeWKBToGeom ...
func EncodeWKBToGeom(reader GEOSWKBReader, cwkb []C.uchar) (GEOSGeometry, error) {

	g := C.GEOSWKBReader_read_r(geosContext, reader, &cwkb[0], C.size_t(len(cwkb)))
	if g == nil {
		return nil, errors.New("C.GEOSGeometry is null")
	}
	return GEOSGeometry(g), nil
}

// EncodeHexToGeom ...
func EncodeHexToGeom(reader GEOSWKBReader, cwkb []C.uchar) (GEOSGeometry, error) {
	g := C.GEOSWKBReader_read_r(geosContext, reader, &cwkb[0], C.size_t(len(cwkb)))
	if g == nil {
		return nil, errors.New("C.GEOSGeometry is null")
	}
	return g, nil
}

// DecodeWKBToArray ...
func DecodeWKBToArray(writer GEOSWKBWriter, g GEOSGeometry) ([]byte, error) {
	var size C.size_t
	bytes := C.GEOSWKBWriter_write_r(geosContext, writer, g, &size)
	if bytes == nil {
		return nil, errors.New("toWKBHex bytes is null")
	}
	ptr := unsafe.Pointer(bytes)
	defer C.free(ptr)
	l := int(size)
	var out []byte
	for i := 0; i < l; i++ {
		el := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(C.uchar(0))*uintptr(i))
		out = append(out, byte(*(*C.uchar)(el)))
	}
	return out, nil
}

// DecodeWKBToHexStr ...
func DecodeWKBToHexStr(writer GEOSWKBWriter, g GEOSGeometry) (string, error) {
	var size C.size_t
	bytes := C.GEOSWKBWriter_writeHEX_r(geosContext, writer, g, &size)
	if bytes == nil {
		return "", errors.New("toWKBHex bytes is null")
	}
	ca := (*C.char)(unsafe.Pointer(bytes))
	s := C.GoString(ca)
	return s, nil
}

// WKBReaderFactory ...
func WKBReaderFactory() GEOSWKBReader {
	reader := C.GEOSWKBReader_create_r(geosContext)
	return GEOSWKBReader(reader)
}

// WKBReaderDestroy ...
func WKBReaderDestroy(reader GEOSWKBReader) {
	C.GEOSWKBReader_destroy_r(geosContext, reader)
}

// WKBWriterFactory ...
func WKBWriterFactory() GEOSWKBWriter {
	writer := C.GEOSWKBWriter_create_r(geosContext)
	return GEOSWKBWriter(writer)
}

// WKBWriterDestroy ...
func WKBWriterDestroy(writer GEOSWKBWriter) {
	C.GEOSWKBWriter_destroy_r(geosContext, writer)
}
