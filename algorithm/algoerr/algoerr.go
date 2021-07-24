package algoerr

import (
	"fmt"
)

// ErrNilSteric ...
var ErrNilSteric = fmt.Errorf("Steric is nil")

// ErrNotInSlice ...
var ErrNotInSlice = fmt.Errorf("Steric not in slice")

// ErrNotSupportCollection ...
var ErrNotSupportCollection = fmt.Errorf("Operation does not support StericCollection arguments")

// ErrWrongUsageFunc ...
var ErrWrongUsageFunc = fmt.Errorf("Wrong usage function")

// ErrWrongFractionRange ...
var ErrWrongFractionRange = fmt.Errorf("Fraction is not in range (0.0 - 1.0]")

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
