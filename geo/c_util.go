package geo
/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

// various wrappers around C API

type unaryTopo func(*C.GEOSGeometry) *C.GEOSGeometry
type unaryPred func(*C.GEOSGeometry) C.char
type float64Getter func(*C.GEOSGeometry, *C.double) C.int

func  UnaryTopo(name string, cfn unaryTopo,g GEOSGeometry) (GEOSGeometry, error) {
	return geomFromC(name, cfn(g))
}

func UnaryPred(name string, cfn unaryPred,g GEOSGeometry) (bool, error) {
	return boolFromC(name, cfn(g))
}

type binaryTopo func(*C.GEOSGeometry, *C.GEOSGeometry) *C.GEOSGeometry
type binaryPred func(*C.GEOSGeometry, *C.GEOSGeometry) C.char

func  BinaryTopo(name string, cfn binaryTopo, g GEOSGeometry,other GEOSGeometry) (GEOSGeometry, error) {
	return geomFromC(name, cfn(g, other))
}

func  BinaryPred(name string, cfn binaryPred,g GEOSGeometry, other GEOSGeometry) (bool, error) {
	return boolFromC(name, cfn(g, other))
}

//func geomFromCoordSeq(cs *coordSeq, name string, cfn func(*C.GEOSCoordSequence) *C.GEOSGeometry) (GEOSGeometry, error) {
//	return geomFromC(name, cfn(cs.c))
//}

func emptyGeom(name string, cfn func() *C.GEOSGeometry) (GEOSGeometry, error) {
	return geomFromC(name, cfn())
}

func geomFromC(name string, ptr *C.GEOSGeometry) (GEOSGeometry, error) {
	if ptr == nil {
		return nil, Error()
	}
	return geomFromPtr(ptr), nil
}

func boolFromC(name string, c C.char) (bool, error) {
	if c == 2 {
		return false, Error()
	}
	return c == 1, nil
}

func intFromC(name string, i C.int, exception C.int) (int, error) {
	if i == exception {
		return 0, Error()
	}
	return int(i), nil
}

func float64FromC(name string, rv C.int, d C.double) (float64, error) {
	if rv == 0 {
		return 0.0, Error()
	}
	return float64(d), nil
}

type binaryFloatGetter func(*C.GEOSGeometry, *C.GEOSGeometry, *C.double) C.int

func  binaryFloat(name string, cfn binaryFloatGetter,g GEOSGeometry, other GEOSGeometry) (float64, error) {
	var d C.double
	return float64FromC(name, cfn(g, other, &d), d)
}

func  simplify(name string, cfn func(*C.GEOSGeometry, C.double) *C.GEOSGeometry,g GEOSGeometry, d float64) (GEOSGeometry, error) {
	return geomFromC(name, cfn(g, C.double(d)))
}

// geomFromPtr returns a new Geometry that's been initialized with a C pointer
// to the GEOS C API object.
//
// This constructor should be used when the caller has ownership of the
// underlying C object.
func geomFromPtr(ptr *C.GEOSGeometry) GEOSGeometry {
	return nil
}
