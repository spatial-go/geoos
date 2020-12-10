package geo
/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

// various wrappers around C API

//=========================== Go与C 转换帮助===============================
// c语言字符转成go 字符串
func CStrToGo(c GEOSChar) string {
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


func boolFromC( c C.char) (bool, error) {
	if c == 2 {
		return false, Error()
	}
	return c == 1, nil
}

func intFromC( i C.int, exception C.int) (int, error) {
	if i == exception {
		return 0, Error()
	}
	return int(i), nil
}

func float64FromC( rv C.int, d C.double) (float64, error) {
	if rv == 0 {
		return 0.0, Error()
	}
	return float64(d), nil
}

