// Package algorithm defines Specifies Computational Geometric and algorithm err.
// Specifies and implements various fundamental Computational Geometric algorithms.
package algorithm

import (
	"fmt"
)

// ErrNotGraphCollection ...
var ErrNotGraphCollection = fmt.Errorf("Operation does not support not match type ,please try CreateGraphCollection")

// ErrNotGraph ...
var ErrNotGraph = fmt.Errorf("Operation does not support not match type ,please try CreateGraph")

// ErrQuadrant ...
var ErrQuadrant = fmt.Errorf("Cannot compute the quadrant for point")

// ErrNilSteric ...
var ErrNilSteric = fmt.Errorf("Steric is nil")

// ErrNotInSlice ...
var ErrNotInSlice = fmt.Errorf("Steric not in slice")

// ErrNotMatchType ...
var ErrNotMatchType = fmt.Errorf("Operation does not support not match type arguments")

// ErrNotSupportCollection ...
var ErrNotSupportCollection = fmt.Errorf("Operation does not support StericCollection arguments")

// ErrWrongUsageFunc ...
var ErrWrongUsageFunc = fmt.Errorf("Wrong usage function")

// ErrWrongFractionRange ...
var ErrWrongFractionRange = fmt.Errorf("Fraction is not in range (0.0 - 1.0]")

// ErrWrongTolerance ...
var ErrWrongTolerance = fmt.Errorf("Tolerance must be non-negative")

// ErrWrongExponent ...
var ErrWrongExponent = fmt.Errorf("Exponent out of bounds")

// ErrComputeOffsetZero ...
var ErrComputeOffsetZero = fmt.Errorf("Cannot compute offset from zero-length line segment")

// ErrWrongLink ...
var ErrWrongLink = fmt.Errorf("Cannot link lines")

// ErrBoundBeNil ...
var ErrBoundBeNil = fmt.Errorf("boundary should be nil")

// ErrUnknownType ...
func ErrUnknownType(obj ...interface{}) error {
	return fmt.Errorf("Unknown Geometry subtype: %v", obj...)
}

// ErrorDimension create new ErrorDimension by object.
func ErrorDimension(obj ...interface{}) error {
	return fmt.Errorf("Steric Should be Dimension: %v", obj...)
}

// ErrorShouldBeLength9 create new ErrorshouldBeLength9 by object.
func ErrorShouldBeLength9(obj ...interface{}) error {
	return fmt.Errorf("Should be length 9: %v", obj...)
}

// ErrorUnknownDimension create new ErrorUnknownDimension by object.
func ErrorUnknownDimension(obj ...interface{}) error {
	return fmt.Errorf("Unknown dimension value: %v", obj...)
}

// Error create new error by object.
func Error(str string, obj ...interface{}) error {
	return fmt.Errorf(str, obj...)
}
