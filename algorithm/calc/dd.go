package calc

// const
const (
	// The smallest representable relative difference between two  values.
	EPS   = 1.23259516440783e-32 /* = 2^-106 */
	SPLIT = 134217729.0          // 2^27+1, for IEEE

	MAXPRINTDIGITS = 32

	SCINOTEXPONENTCHAR = "E"
	SCINOTZERO         = "0.0E0"
)

var (
	// PI The value nearest to the constant Pi.
	PI = &DD{3.141592653589793116e+00,
		1.224646799147353207e-16}
	// TWOPI The value nearest to the constant 2 * Pi.
	TWOPI = &DD{
		6.283185307179586232e+00,
		2.449293598294706414e-16}
	// PI2 The value nearest to the constant Pi / 2.
	PI2 = &DD{
		1.570796326794896558e+00,
		6.123233995736766036e-17}
	//E  The value nearest to the constant e (the natural logarithm base).
	E = &DD{
		2.718281828459045091e+00,
		1.445646891729250158e-16}
)

// DD A DoubleDouble uses a representation containing two double-precision values.
// A number x is represented as a pair of doubles, x.hi and x.lo
type DD struct {
	Hi, Lo float64
}

// ValueOf Converts the  argument to a  number.
func ValueOf(x float64) *DD {
	return &DD{x, 0.0}
}

// SelfAdd  Adds the argument to the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *DD) SelfAdd(yhi, ylo float64) *DD {
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

// SelfSubtract Subtracts the argument from the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *DD) SelfSubtract(yhi, ylo float64) *DD {
	return d.SelfAdd(-yhi, -ylo)
}

// SelfMultiply Multiplies this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *DD) SelfMultiply(yhi, ylo float64) *DD {
	var hx, tx, hy, ty, C, c float64
	C = SPLIT * d.Hi
	hx = C - d.Hi
	c = SPLIT * yhi
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

// SelfDivide Divides this object by the argument, returning DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *DD) SelfDivide(yhi, ylo float64) *DD {
	var hc, tc, hy, ty, C, c, U, u float64
	C = d.Hi / yhi
	c = SPLIT * C
	hc = c - C
	u = SPLIT * yhi
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
func (d *DD) Signum() int {
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

// IsZero Tests whether this value is equal to 0.
func (d *DD) IsZero() bool {
	return d.Hi == 0.0 && d.Lo == 0.0
}

// Equals Tests whether this value is equal to another value.
func (d *DD) Equals(y *DD) bool {
	return d.Hi == y.Hi && d.Lo == y.Lo
}

// Gt Tests whether this value is greater than another value.
func (d *DD) Gt(y *DD) bool {
	return (d.Hi > y.Hi) || (d.Hi == y.Hi && d.Lo > y.Lo)
}

// Ge Tests whether this value is greater than or equals to another value.
func (d *DD) Ge(y *DD) bool {
	return (d.Hi > y.Hi) || (d.Hi == y.Hi && d.Lo >= y.Lo)
}

// Lt Tests whether this value is less than another  value.
func (d *DD) Lt(y *DD) bool {
	return (d.Hi < y.Hi) || (d.Hi == y.Hi && d.Lo < y.Lo)
}

// Le Tests whether this value is less than or equal to another  value.
func (d *DD) Le(y *DD) bool {
	return (d.Hi < y.Hi) || (d.Hi == y.Hi && d.Lo <= y.Lo)
}

// CompareTo Compares two  objects numerically.
func (d *DD) CompareTo(other *DD) int {

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
