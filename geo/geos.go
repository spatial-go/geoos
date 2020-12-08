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
	"unsafe"
)

type GEOSContext C.GEOSContextHandle_t

type GEOSGeometry *C.GEOSGeometry
type GEOSCoordSequence C.GEOSCoordSequence
type GEOSBufferParams C.GEOSBufferParams
type GEOSSTRtree C.GEOSSTRtree

type GEOSWKTReader *C.GEOSWKTReader
type GEOSWKTWriter *C.GEOSWKTWriter
type GEOSWKBReader *C.GEOSWKBReader
type GEOSWKBWriter *C.GEOSWKBWriter

type GEOSChar C.char

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
//  ====================Reader\Writer 初始化、销毁方法===============================
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

func WKBReaderFactory() GEOSWKBReader {
	reader := C.GEOSWKBReader_create_r(geosContext)
	return GEOSWKBReader(reader)
}
func WKBReaderDestroy(reader GEOSWKBReader) {
	C.GEOSWKBReader_destroy_r(geosContext, reader)
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
func WKBWriterFactory() GEOSWKBWriter {
	writer := C.GEOSWKBWriter_create_r(geosContext)
	return GEOSWKBWriter(writer)
}
func WKBWriterDestroy(writer GEOSWKBWriter) {
	C.GEOSWKBWriter_destroy_r(geosContext, writer)
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

// 解码生成wkt格式字符串
func DecodeWKBToArray(writer GEOSWKBWriter, g GEOSGeometry) ([]byte, error) {
	var size C.size_t
	bytes := C.GEOSWKBWriter_write_r(geosContext, writer, g, &size)
	if bytes == nil {
		return nil, errors.New("toWKBHex bytes is null")
	}
	// TODO: 指针类型之间进行转换?
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

// 解码生成wkt格式hex字符串
func DecodeWKBToHexStr(writer GEOSWKBWriter, g GEOSGeometry) (string, error) {
	var size C.size_t
	bytes := C.GEOSWKBWriter_writeHEX_r(geosContext, writer, g, &size)
	if bytes == nil {
		return "", errors.New("toWKBHex bytes is null")
	}
	// 指针类型之间进行转换
	ca := (*C.char)(unsafe.Pointer(bytes))
	s := C.GoString(ca)
	return s, nil
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

// 编码WKB字节数组生成几何对象
func EncodeWKBToGeom(reader GEOSWKBReader, cwkb []C.uchar) (GEOSGeometry, error) {

	g := C.GEOSWKBReader_read_r(geosContext, reader, &cwkb[0], C.size_t(len(cwkb)))
	if g == nil {
		return nil, errors.New("C.GEOSGeometry is null")
	}
	//defer C.GEOSGeom_destroy_r(geosContext, g)
	return GEOSGeometry(g), nil
}

// 编码hex字符串生成几何图形
func EncodeHexToGeom(reader GEOSWKBReader, cwkb []C.uchar) (GEOSGeometry, error) {
	g := C.GEOSWKBReader_read_r(geosContext, reader, &cwkb[0], C.size_t(len(cwkb)))
	if g == nil {
		return nil, errors.New("C.GEOSGeometry is null")
	}
	//defer C.GEOSGeom_destroy_r(geosContext, g)
	return g, nil
}

//=========================== Go与C 转换帮助===============================
// c语言字符转成go 字符串
func CStrToGo(c *C.char) string {
	return C.GoString(c)
}

// go字符串转成C字符
func GoStrTOC(str string) *C.char {
	return C.CString(str)
}

// go字节数组转换成C数组
func GoByteArrayToCCharArray(array []byte) []C.uchar {
	var cArray []C.uchar
	for i := range array {
		cArray = append(cArray, C.uchar(array[i]))
	}
	return cArray
}
