/**
WKT 格式相关转换方法
*/
package geo

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type GEOSWKTReader *C.GEOSWKTReader
type GEOSWKTWriter *C.GEOSWKTWriter

// 生成WKT格式字符串表示
func ToWKTStr(g GEOSGeometry) (string, error) {
	w := WKTWriterFactory()
	defer WKTWriterDestroy(w)
	return DecodeWKTToStr(w, g)
}

// 根据WKT格式字符串生成集合对象
func GeomFromWKTStr(wktstr string) GEOSGeometry {
	reader := WKTReaderFactory()
	defer WKTReaderDestroy(reader)
	return EncodeWKTToGeom(reader, wktstr)
}

/**
Writer factory
*/
func WKTWriterFactory() GEOSWKTWriter {
	writer := C.GEOSWKTWriter_create_r(geosContext)
	return GEOSWKTWriter(writer)
}
func WKTWriterDestroy(writer GEOSWKTWriter) {
	C.GEOSWKTWriter_destroy_r(geosContext, writer)
}

/**
Reader factory
*/
func WKTReaderFactory() GEOSWKTReader {
	reader := C.GEOSWKTReader_create_r(geosContext)
	return GEOSWKTReader(reader)
}
func WKTReaderDestroy(reader GEOSWKTReader) {
	C.GEOSWKTReader_destroy_r(geosContext, reader)
}

//===========================几何对象、字符串之间编解码方法================================
// 解码生成wkt格式字符串
func DecodeWKTToStr(writer GEOSWKTWriter, g GEOSGeometry) (string, error) {
	cstr := C.GEOSWKTWriter_write_r(geosContext, writer, g)
	if cstr == nil {
		return "", errors.New("writer to wkt is null")
	}
	return C.GoString(cstr), nil
}

// 编码WKT字符串生成几何对象
func EncodeWKTToGeom(reader GEOSWKTReader, wktStr string) GEOSGeometry {
	cs := C.CString(wktStr)
	g := C.GEOSWKTReader_read_r(geosContext, reader, cs)
	C.free(unsafe.Pointer(cs))
	//TODO:: 写入这里 嵌套调用时候不成功
	//defer C.GEOSGeom_destroy_r(geosContext, g)
	return GEOSGeometry(g)
}
