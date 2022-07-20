package quadedge

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// TriangleVisitor ...
type TriangleVisitor interface {
	Visit(triEdges []*QuadEdge)
}

// TriangleCircumcentreVisitor ...
type TriangleCircumcentreVisitor struct {
}

// Visit ...
func (t *TriangleCircumcentreVisitor) Visit(triEdges []*QuadEdge) {
	var (
		a = triEdges[0].Origin()
		b = triEdges[1].Origin()
		c = triEdges[2].Origin()
	)

	cc := circumcentrePF(a, b, c)
	// save the circumcentre as the origin for the dual edges originating in this triangle
	for i := 0; i < 3; i++ {
		triEdges[i].rot.SetOrigin(cc)
	}
}

// circumcentrePF Returns the circumcentre of the triangle.
// The circumcentre is the centre of the circumcircle.
func circumcentrePF(a, b, c matrix.Matrix) matrix.Matrix {
	var (
		ax    = calc.ValueOf(a[0]).Subtract(c[0], 0)
		ay    = calc.ValueOf(a[1]).Subtract(c[1], 0)
		bx    = calc.ValueOf(b[0]).Subtract(c[0], 0)
		by    = calc.ValueOf(b[1]).Subtract(c[1], 0)
		denom = calc.DeterminantPair(ax, ay, bx, by).Multiply(2, 0)
		asqr  = ax.Pow2().AddPair(ay.Pow2())
		bsqr  = bx.Pow2().AddPair(by.Pow2())
		numx  = calc.DeterminantPair(ay, asqr, by, bsqr)
		numy  = calc.DeterminantPair(ax, asqr, bx, bsqr)
		ccx   = calc.ValueOf(c[0]).SubtractPair(numx.DividePair(denom)).Value()
		ccy   = calc.ValueOf(c[1]).AddPair(numy.DividePair(denom)).Value()
	)
	return matrix.Matrix{ccx, ccy}
}
