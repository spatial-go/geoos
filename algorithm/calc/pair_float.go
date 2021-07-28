package calc

// PairFloat A DoubleDouble uses a representation containing two double-precision values.
// A number x is represented as a pair of doubles, x.hi and x.lo
type PairFloat struct {
	Hi, Lo float64
}

// ValueOf Converts the  argument to a  number.
func ValueOf(x float64) *PairFloat {
	return &PairFloat{x, 0.0}
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

// SelfSubtract Subtracts the argument from the value of DD.
// To prevent altering constants,
// this method must only be used on values known to
// be newly created.
func (d *PairFloat) SelfSubtract(yhi, ylo float64) *PairFloat {
	return d.SelfAdd(-yhi, -ylo)
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
func (d *PairFloat) SelfDivide(yhi, ylo float64) *PairFloat {
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
