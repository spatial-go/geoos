package geo

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

type GEOSCoordinateSequence *C.GEOSCoordSequence

// 3维坐标点
type Coordinate struct {
	X float64
	Y float64
	Z float64
}
type CoordinateSequece struct {
}

func NewCoordinate(x, y float64) Coordinate {
	return Coordinate{x, y, 0}
}

// 坐标设置
func (cs CoordinateSequece) setX(seq GEOSCoordinateSequence, index int, val float64) error {

	res := C.GEOSCoordSeq_setX_r(geosContext, seq, index, val)
	if res == 0 {
		return errors.New("method C.GEOSCoordSeq_setX_r error")
	}
	return nil
}

func (cs CoordinateSequece) setY(seq GEOSCoordinateSequence, index int, val float64) error {

	res := C.GEOSCoordSeq_setY_r(geosContext, seq, index, val)
	if res == 0 {
		return errors.New("method C.GEOSCoordSeq_setY_r error")
	}
	return nil
}

func (cs CoordinateSequece) setZ(seq GEOSCoordinateSequence, index int, val float64) error {

	res := C.GEOSCoordSeq_setZ_r(geosContext, seq, index, val)
	if res == 0 {
		return errors.New("method C.GEOSCoordSeq_setZ_r error")
	}
	return nil
}

func (cs CoordinateSequece) getX(seq GEOSCoordinateSequence, idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getX(seq, C.uint(idx), &val)
	if i == 0 {
		return 0.0, errors.New("method C.GEOSCoordSeq_getX error")
	}
	return float64(val), nil
}

func (cs CoordinateSequece) getY(seq GEOSCoordinateSequence, idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getY(seq, C.uint(idx), &val)
	if i == 0 {
		return 0.0, errors.New("method C.GEOSCoordSeq_getY error")
	}
	return float64(val), nil
}

func (cs CoordinateSequece) getZ(seq GEOSCoordinateSequence, idx int) (float64, error) {
	var val C.double
	i := C.GEOSCoordSeq_getZ(seq, C.uint(idx), &val)
	if i == 0 {
		return 0.0, errors.New("method C.GEOSCoordSeq_getZ error")
	}
	return float64(val), nil
}

func (cs CoordinateSequece) seqSize(seq GEOSCoordinateSequence) (int, error) {
	var val C.uint
	i := C.GEOSCoordSeq_getSize(seq, &val)
	if i == 0 {
		return 0, errors.New("method GEOSCoordSeq_getSize error")
	}
	return int(val), nil
}

func (cs CoordinateSequece) seqDims(seq GEOSCoordinateSequence) (int, error) {
	var val C.uint
	i := C.GEOSCoordSeq_getDimensions(seq, &val)
	if i == 0 {
		return 0, errors.New("method GEOSCoordSeq_getDimensions error")
	}
	return int(val), nil
}

// 2维点的字符串表达
func (c Coordinate) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}
// 坐标数组创建坐标序列
func makeCodinateSequence(coords []Coordinate) (GEOSCoordinateSequence, error) {
	seq := C.GEOSCoordSeq_create(C.uint(len(coords)), C.uint(2))
	if seq == nil {
		return nil, Error()
	}
	cs := new(CoordinateSequece)
	for i, c := range coords {
		if err := cs.setX(seq, i, c.X); err != nil {
			return nil, err
		}
		if err := cs.setY(seq, i, c.Y); err != nil {
			return nil, err
		}
	}
	return GEOSCoordinateSequence(seq), nil
}
