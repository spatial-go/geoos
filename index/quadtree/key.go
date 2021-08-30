package quadtree

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

// Key A Key is a unique identifier for a node in a quadtree.
//  It contains a lower-left point and a level number. The level number
//  is the power of two for the size of the node envelope
type Key struct {
	Pt    matrix.Matrix
	Level int
	Env   *envelope.Envelope
}

// ComputeQuadLevel ...
func ComputeQuadLevel(env *envelope.Envelope) int {
	dx := env.Width()
	dy := env.Height()
	dMax := math.Max(dx, dy)
	_, level := math.Frexp(dMax)
	return level
}

// NewKeyEnv ...
func NewKeyEnv(itemEnv *envelope.Envelope) *Key {
	k := &Key{Pt: matrix.Matrix{0, 0}}
	k.ComputeKey(itemEnv)
	return k
}

// Centre ...
func (k *Key) Centre() matrix.Matrix {
	return matrix.Matrix{
		(k.Env.MinX + k.Env.MaxX) / 2,
		(k.Env.MinY + k.Env.MaxY) / 2,
	}
}

// ComputeKey return a square envelope containing the argument envelope,
// whose extent is a power of two and which is based at a power of 2
func (k *Key) ComputeKey(itemEnv *envelope.Envelope) {
	k.Level = ComputeQuadLevel(itemEnv)
	k.Env = &envelope.Envelope{}
	k.computeKey(k.Level, itemEnv)
	// MD - would be nice to have a non-iterative form of this algorithm
	for !k.Env.Contains(itemEnv) {
		k.Level++
		k.computeKey(k.Level, itemEnv)
	}
}

func (k *Key) computeKey(level int, itemEnv *envelope.Envelope) {
	quadSize := math.Exp2(float64(level))
	x, y := itemEnv.MinX, itemEnv.MinY
	k.Pt[0] = math.Floor(x/quadSize) * quadSize
	k.Pt[1] = math.Floor(y/quadSize) * quadSize
	k.Env = envelope.FourFloat(k.Pt[0], k.Pt[0]+quadSize, k.Pt[1], k.Pt[1]+quadSize)
}
