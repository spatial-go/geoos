package geoc

// #cgo LDFLAGS: -lgeos_c
// #include "geos.h"
import "C"

// GoByteArrayToCCharArray convert go byte array to C char array
func GoByteArrayToCCharArray(array []byte) []C.uchar {
	var cArray []C.uchar
	for i := range array {
		cArray = append(cArray, C.uchar(array[i]))
	}
	return cArray
}

func boolFromC(c C.char) (bool, error) {
	if c == 2 {
		return false, Error()
	}
	return c == 1, nil
}

func intFromC(i C.int, exception C.int) (int, error) {
	if i == exception {
		return 0, Error()
	}
	return int(i), nil
}

func float64FromC(c C.int, d C.double) (float64, error) {
	if c == 0 {
		return 0.0, Error()
	}
	return float64(d), nil
}
