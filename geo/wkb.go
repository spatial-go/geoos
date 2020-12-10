package geo

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"
import (
	"encoding/hex"
	"errors"
	"unsafe"
)

type GEOSWKBReader *C.GEOSWKBReader
type GEOSWKBWriter *C.GEOSWKBWriter

// 根据wkb 二进制生成集合对象
func GeomFromWKBStr(wkbByte []byte) (GEOSGeometry, error) {
	cwkb := GoByteArrayToCCharArray(wkbByte)
	reader := WKBReaderFactory()
	defer WKBReaderDestroy(reader)
	return EncodeWKBToGeom(reader, cwkb)
}

// 生成WKB格式字符串表示
func ToWKB(g GEOSGeometry) ([]byte, error) {
	writer := WKBWriterFactory()
	defer WKBWriterDestroy(writer)
	return DecodeWKBToArray(writer, g)
}

// 生成WKB Hex格式字符串表达
func ToWKBHex(g GEOSGeometry) (string, error) {
	w := WKBWriterFactory()
	defer WKBWriterDestroy(w)
	return DecodeWKBToHexStr(w, g)
}

// 根据hex格式字符串生成集合对象
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

//  ====================Reader\Writer 初始化、销毁方法===============================

func WKBReaderFactory() GEOSWKBReader {
	reader := C.GEOSWKBReader_create_r(geosContext)
	return GEOSWKBReader(reader)
}
func WKBReaderDestroy(reader GEOSWKBReader) {
	C.GEOSWKBReader_destroy_r(geosContext, reader)
}

func WKBWriterFactory() GEOSWKBWriter {
	writer := C.GEOSWKBWriter_create_r(geosContext)
	return GEOSWKBWriter(writer)
}
func WKBWriterDestroy(writer GEOSWKBWriter) {
	C.GEOSWKBWriter_destroy_r(geosContext, writer)
}
