// Package spaceerr A representation of error.
package spaceerr

import (
	"fmt"
)

// ErrNilGeometry ...
var ErrNilGeometry = fmt.Errorf("Geometry is nil")

// ErrNotValidGeometry ...
var ErrNotValidGeometry = fmt.Errorf("Geometry is not valid")

// ErrNotPolygon UnaryUnion parameter is not polygon
var ErrNotPolygon = fmt.Errorf("Geometry is not polygon")

// ErrNotSupportCollection ...
var ErrNotSupportCollection = fmt.Errorf("Operation does not support GeometryCollection arguments")

// ErrNotSupportBound ...
var ErrNotSupportBound = fmt.Errorf("Operation does not support Bound arguments")

// ErrNotSupportGeometry ...
var ErrNotSupportGeometry = fmt.Errorf("Operation does not support arguments")

// ErrWrongUsageFunc ...
var ErrWrongUsageFunc = fmt.Errorf("Wrong usage function")

// ErrBoundBeNil ...
var ErrBoundBeNil = fmt.Errorf("boundary should be nil")

// ErrorUsageFunc create new ErrorUsageFunc by object.
func ErrorUsageFunc(obj ...interface{}) error {
	return fmt.Errorf("Wrong usage function :%v", obj)
}
