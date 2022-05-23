package hprtree

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

// MaxLevel ...
const MaxLevel = 16

// HilbertEncoder ...
type HilbertEncoder struct {
	level   int
	minx    float64
	miny    float64
	strideX float64
	strideY float64
}

// NewHilbertEncoder ...
func NewHilbertEncoder(level int, extent *envelope.Envelope) *HilbertEncoder {
	h := &HilbertEncoder{}
	h.level = level
	hSide := math.Pow(2.0, float64(h.level)) - 1

	h.minx = extent.MinX
	extentX := extent.Width()
	h.strideX = extentX / hSide

	h.miny = extent.MinX
	extentY := extent.Height()
	h.strideY = extentY / hSide
	return h
}

func (h *HilbertEncoder) encode(env *envelope.Envelope) int {
	midX := env.Width()/2 + env.MinX
	x := int((midX - h.minx) / h.strideX)

	midY := env.Height()/2 + env.MinY
	y := int((midY - h.miny) / h.strideY)

	return encode(h.level, x, y)
}

// Encodes a point (x,y)
//   in the range of the the Hilbert curve at a given level
//   as the index of the point along the curve.
//   The index will lie in the range [0, 2<sup>level + 1</sup>].
// Returns the index of the point along the Hilbert curve
func encode(level, x, y int) int {
	// Fast Hilbert curve algorithm by http://threadlocalmutex.com/
	// Ported from C++ https://github.com/rawrunprotected/hilbert_curves (public
	// domain)

	lvl1 := math.Max(float64(level), 1)
	lvl := int(math.Min(lvl1, float64(MaxLevel)))

	x = x << (16 - lvl)
	y = y << (16 - lvl)

	a := x ^ y
	b := 0xFFFF ^ a
	c := 0xFFFF ^ (x | y)
	d := x & (y ^ 0xFFFF)

	A := a | (b >> 1)
	B := (a >> 1) ^ a
	C := ((c >> 1) ^ (b & (d >> 1))) ^ c
	D := ((a & (c >> 1)) ^ (d >> 1)) ^ d

	a = A
	b = B
	c = C
	d = D
	A = ((a & (a >> 2)) ^ (b & (b >> 2)))
	B = ((a & (b >> 2)) ^ (b & ((a ^ b) >> 2)))
	C ^= ((a & (c >> 2)) ^ (b & (d >> 2)))
	D ^= ((b & (c >> 2)) ^ ((a ^ b) & (d >> 2)))

	a = A
	b = B
	c = C
	d = D
	A = ((a & (a >> 4)) ^ (b & (b >> 4)))
	B = ((a & (b >> 4)) ^ (b & ((a ^ b) >> 4)))
	C ^= ((a & (c >> 4)) ^ (b & (d >> 4)))
	D ^= ((b & (c >> 4)) ^ ((a ^ b) & (d >> 4)))

	a = A
	b = B
	c = C
	d = D
	C ^= ((a & (c >> 8)) ^ (b & (d >> 8)))
	D ^= ((b & (c >> 8)) ^ ((a ^ b) & (d >> 8)))

	a = C ^ (C >> 1)
	b = D ^ (D >> 1)

	i0 := x ^ y
	i1 := b | (0xFFFF ^ (i0 | a))

	i0 = (i0 | (i0 << 8)) & 0x00FF00FF
	i0 = (i0 | (i0 << 4)) & 0x0F0F0F0F
	i0 = (i0 | (i0 << 2)) & 0x33333333
	i0 = (i0 | (i0 << 1)) & 0x55555555

	i1 = (i1 | (i1 << 8)) & 0x00FF00FF
	i1 = (i1 | (i1 << 4)) & 0x0F0F0F0F
	i1 = (i1 | (i1 << 2)) & 0x33333333
	i1 = (i1 | (i1 << 1)) & 0x55555555

	index := ((i1 << 1) | i0) >> (32 - 2*lvl)
	return index
}
