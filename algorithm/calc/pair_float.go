// Package calc methods and parameters be realized ouble precision calculation.
package calc

// const Defined constant variable  calc parameter
const (
	// Split The value to split a double-precision value on during multiplication
	Split = 134217729.0 // 2^27+1, for IEEE
)

// PairFloat A DoubleDouble uses a representation containing two double-precision values.
// A number x is represented as a pair of doubles, x.hi and x.lo
type PairFloat struct {
	Hi, Lo float64
}

// ValueOf Converts the  argument to a  number.
func ValueOf(x float64) *PairFloat {
	return &PairFloat{x, 0.0}
}

// AddOne  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) AddOne(y float64) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfAddOne(y)
	// return selfAdd(y, 0.0);
}

// AddPair  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) AddPair(y *PairFloat) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfAddPair(y)
}

// Add  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) Add(yhi, ylo float64) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfAdd(yhi, ylo)
}

// SelfAddOne  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfAddOne(y float64) *PairFloat {
	var H, h, S, s, e, f float64
	S = d.Hi + y
	e = S - d.Hi
	s = S - e
	s = (y - e) + (d.Hi - s)
	f = s + d.Lo
	H = S + f
	h = f + (S - H)
	d.Hi = H + h
	d.Lo = h + (H - d.Hi)
	return d
	// return selfAdd(y, 0.0);
}

// SelfAddPair  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfAddPair(y *PairFloat) *PairFloat {
	return d.SelfAdd(y.Hi, y.Lo)
}

// SelfAdd  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfAdd(yhi, ylo float64) *PairFloat {
	var H, h, T, t, S, s, e, f float64
	S = d.Hi + yhi
	T = d.Lo + ylo
	e = S - d.Hi
	f = T - d.Lo
	s = S - e
	t = T - f
	s = (yhi - e) + (d.Hi - s)
	t = (ylo - f) + (d.Lo - t)
	e = s + T
	H = S + e
	h = e + (S - H)
	e = t + h

	zhi := H + e
	zlo := e + (H - zhi)
	d.Hi = zhi
	d.Lo = zlo
	return d
}

// SubtractPair Subtracts the argument from the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SubtractPair(y *PairFloat) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfAdd(-y.Hi, -y.Lo)
}

// Subtract Subtracts the argument from the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) Subtract(yhi, ylo float64) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfAdd(-yhi, -ylo)
}

// SelfSubtractPair Subtracts the argument from the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfSubtractPair(y *PairFloat) *PairFloat {
	return d.SelfAdd(-y.Hi, -y.Lo)
}

// SelfSubtract Subtracts the argument from the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfSubtract(yhi, ylo float64) *PairFloat {
	return d.SelfAdd(-yhi, -ylo)
}

// MultiplyPair Multiplies this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) MultiplyPair(y *PairFloat) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfMultiply(y.Hi, y.Lo)
}

// Multiply Multiplies this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) Multiply(yhi, ylo float64) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfMultiply(yhi, ylo)
}

// SelfMultiplyPair Multiplies this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfMultiplyPair(y *PairFloat) *PairFloat {
	return d.SelfMultiply(y.Hi, y.Lo)
}

// SelfMultiply Multiplies this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfMultiply(yhi, ylo float64) *PairFloat {
	var hx, tx, hy, ty, C, c float64
	C = Split * d.Hi
	hx = C - d.Hi
	c = Split * yhi
	hx = C - hx
	tx = d.Hi - hx
	hy = c - yhi
	C = d.Hi * yhi
	hy = c - hy
	ty = yhi - hy
	c = ((((hx*hy - C) + hx*ty) + tx*hy) + tx*ty) + (d.Hi*ylo + d.Lo*yhi)
	zhi := C + c
	hx = C - zhi
	zlo := c + hx
	d.Hi = zhi
	d.Lo = zlo
	return d
}

// DividePair Divides this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) DividePair(y *PairFloat) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfDivide(y.Hi, y.Lo)
}

// Divide Divides this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) Divide(yhi, ylo float64) *PairFloat {
	newPair := &PairFloat{d.Hi, d.Lo}
	return newPair.SelfDivide(yhi, ylo)
}

// SelfDividePair Divides this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfDividePair(y *PairFloat) *PairFloat {
	return d.SelfDivide(y.Hi, y.Lo)
}

// SelfDivide Divides this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfDivide(yhi, ylo float64) *PairFloat {
	var hc, tc, hy, ty, C, c, U, u float64
	C = d.Hi / yhi
	c = Split * C
	hc = c - C
	u = Split * yhi
	hc = c - hc
	tc = C - hc
	hy = u - yhi
	U = C * yhi
	hy = u - hy
	ty = yhi - hy
	u = (((hc*hy - U) + hc*ty) + tc*hy) + tc*ty
	c = ((((d.Hi - U) - u) + d.Lo) - C*ylo) / yhi
	u = C + c

	d.Hi = u
	d.Lo = (C - u) + c
	return d
}

// Signum Returns an integer indicating the sign of this value.
func (d *PairFloat) Signum() int {
	if d.Hi > 0 {
		return 1
	}
	if d.Hi < 0 {
		return -1
	}
	if d.Lo > 0 {
		return 1
	}
	if d.Lo < 0 {
		return -1
	}
	return 0
}

// Pow2 Returns the square of this value.
func (d *PairFloat) Pow2() *PairFloat {
	return d.MultiplyPair(d)
}

// SelfPow2 Squares this object.
// To prevent altering constants,
// this method must only be used on values known to be newly created.
func (d *PairFloat) SelfPow2() *PairFloat {
	return d.SelfMultiplyPair(d)
}

// IsZero Tests whether this value is equal to 0.
func (d *PairFloat) IsZero() bool {
	return d.Hi == 0.0 && d.Lo == 0.0
}

// Equals Tests whether this value is equal to another value.
func (d *PairFloat) Equals(y *PairFloat) bool {
	return d.Hi == y.Hi && d.Lo == y.Lo
}

// Gt Tests whether this value is greater than another value.
func (d *PairFloat) Gt(y *PairFloat) bool {
	return (d.Hi > y.Hi) || (d.Hi == y.Hi && d.Lo > y.Lo)
}

// Ge Tests whether this value is greater than or equals to another value.
func (d *PairFloat) Ge(y *PairFloat) bool {
	return (d.Hi > y.Hi) || (d.Hi == y.Hi && d.Lo >= y.Lo)
}

// Lt Tests whether this value is less than another  value.
func (d *PairFloat) Lt(y *PairFloat) bool {
	return (d.Hi < y.Hi) || (d.Hi == y.Hi && d.Lo < y.Lo)
}

// Le Tests whether this value is less than or equal to another  value.
func (d *PairFloat) Le(y *PairFloat) bool {
	return (d.Hi < y.Hi) || (d.Hi == y.Hi && d.Lo <= y.Lo)
}

// CompareTo Compares two  objects numerically.
func (d *PairFloat) CompareTo(other *PairFloat) int {

	if d.Hi < other.Hi {
		return -1
	}
	if d.Hi > other.Hi {
		return 1
	}
	if d.Lo < other.Lo {
		return -1
	}
	if d.Lo > other.Lo {
		return 1
	}
	return 0
}

// Value Converts this value to the nearest double-precision number.
func (d *PairFloat) Value() float64 {
	return d.Hi + d.Lo
}

// Signum Returns an integer indicating the sign of this value.
func Signum(x float64) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

// Determinant Computes the determinant of the 2x2 matrix with the given entries.
func Determinant(x1, y1, x2, y2 float64) *PairFloat {
	return DeterminantPair(ValueOf(x1), ValueOf(y1), ValueOf(x2), ValueOf(y2))
}

// DeterminantPair Computes the determinant of the 2x2 matrix with the given entries.
func DeterminantPair(x1, y1, x2, y2 *PairFloat) *PairFloat {
	return x1.MultiplyPair(y2).SelfSubtractPair(y1.MultiplyPair(x2))
}
