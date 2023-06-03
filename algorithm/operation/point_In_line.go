package operation

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// InLineSegment returns true if spot in ab,false else.
func InLineSegment(spot, a, b matrix.Matrix) (in bool, isVertex bool) {
	// x := spot[0] <= math.Max(a[0], b[0]) && spot[0] >= math.Min(a[0], b[0])
	// y := spot[1] <= math.Max(a[1], b[1]) && spot[1] >= math.Min(a[1], b[1])
	accuracy := calc.DefaultTolerance * math.Max(spot[0], spot[1])

	if spot.EqualsExact(a, accuracy) || spot.EqualsExact(b, accuracy) {
		return true, true
	}

	ax := calc.ValueOf(spot[0]).SelfSubtractPair(calc.ValueOf(a[0])).SelfMultiplyPair(calc.ValueOf(a[1]).SelfSubtractPair(calc.ValueOf(b[1])))
	bx := calc.ValueOf(a[0]).SelfSubtractPair(calc.ValueOf(b[0])).SelfMultiplyPair(calc.ValueOf(spot[1]).SelfSubtractPair(calc.ValueOf(a[1])))
	dAB := ax.SubtractPair(bx).Value()

	// ax = (spot[0] - a[0]) * (a[1] - b[1])
	// bx = (a[0] - b[0]) * (spot[1] - a[1])
	// dAB = ax - bx

	// dAB = (spot[0]-a[0])*(a[1]-b[1]) - (a[0]-b[0])*(spot[1]-a[1])

	if math.Abs(dAB) < accuracy &&
		(spot[0]+accuracy >= math.Min(a[0], b[0]) && spot[0]-accuracy <= math.Max(a[0], b[0])) &&
		(spot[1]+accuracy >= math.Min(a[1], b[1]) && spot[1]-accuracy <= math.Max(a[1], b[1])) {
		return true, false
	}
	return false, false
}

// IsLineMatrixVertex returns true if spot in LineVertex,false else..
func IsLineMatrixVertex(spot matrix.Matrix, matr matrix.LineMatrix) (isVertex, isEndPoint bool) {
	for i, v := range matr {
		if spot.EqualsExact(matrix.Matrix(v), calc.DefaultTolerance) {
			if i == 0 || i == len(matr)-1 {
				return true, true
			}
			return true, false
		}
	}
	return false, false
}

// InLineMatrix returns true if spot in LineMatrix,false else..
func InLineMatrix(spot matrix.Matrix, matr matrix.LineMatrix) bool {
	//lines := matr.ToLineArray()
	for i := 0; i < len(matr)-1; i++ {
		if in, _ := InLineSegment(spot, matr[i], matr[i+1]); in {
			return true
		}
	}
	return false
}
