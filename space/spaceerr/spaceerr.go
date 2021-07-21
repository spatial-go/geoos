package spaceerr

import (
	"fmt"
)

// ErrNilGeometry ...
var ErrNilGeometry = fmt.Errorf("Geometry is nil")

// ErrNotPolygon UnaryUnion parameter is not polygon
var ErrNotPolygon = fmt.Errorf("Geometry is not polygon")

// ErrNotSupportCollection ...
var ErrNotSupportCollection = fmt.Errorf("Operation does not support GeometryCollection arguments")

// ErrNotSupportBound ...
var ErrNotSupportBound = fmt.Errorf("Operation does not support Bound arguments")

// ErrWrongUsageFunc ...
var ErrWrongUsageFunc = fmt.Errorf("Wrong usage function")

// ErrBoundBeNil ...
var ErrBoundBeNil = fmt.Errorf("boundary should be nil")

// ErrorUsageFunc create new ErrorUsageFunc by object.
func ErrorUsageFunc(obj ...interface{}) error {
	return fmt.Errorf("Wrong usage function :%v", obj...)
}

// Error create new error by object.
func Error(str string, obj ...interface{}) error {
	return fmt.Errorf(str, obj...)
}
