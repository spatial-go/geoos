package relate

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Line  straight line  .
type Line struct {
	Start, End matrix.Matrix
}

// LineArray returns the LineArray
func LineArray(l matrix.LineMatrix) (lines []Line) {
	for i := 0; i < len(l)-1; i++ {
		lines = append(lines, Line{matrix.Matrix(l[i]), matrix.Matrix(l[i+1])})
	}
	return
}

// IsIntersection returns intersection of a and other.
func (l *Line) IsIntersection(o *Line) bool {
	mark, _, _, _ := Intersection(l.Start, l.End, o.Start, o.End)
	return mark
}

// Intersection returns intersection of a and other.
func (l *Line) Intersection(o *Line) (bool, matrix.Matrix) {
	mark, ip, _, _ := Intersection(l.Start, l.End, o.Start, o.End)
	return mark, ip
}

// Intersection returns intersection of a and b.
func Intersection(aStart, aEnd, bStart, bEnd matrix.Matrix) (mark bool, p matrix.Matrix, isIntersectionPoint, isEntering bool) {
	a1 := aEnd[1] - aStart[1]
	b1 := aStart[0] - aEnd[0]
	c1 := -aStart[0]*a1 - b1*aStart[1]
	a2 := bEnd[1] - bStart[1]
	b2 := bStart[0] - bEnd[0]
	c2 := -a2*bStart[0] - b2*bStart[1]

	u := matrix.Matrix{aEnd[0] - aStart[0], aEnd[1] - aStart[1]}
	v := matrix.Matrix{bEnd[0] - bStart[0], bEnd[1] - bStart[1]}

	determinant := CrossProduct(u, v)

	if determinant == 0 {
		mark = false
	} else {
		p := matrix.Matrix{(b1*c2 - b2*c1) / determinant, (a2*c1 - a1*c2) / determinant}
		// check if point belongs to segment
		if InLine(p, aStart, aEnd) && InLine(p, bStart, bEnd) {
			isIntersectionPoint = true
			// determine if the point is entering by determinant
			isEntering = determinant < 0
			mark = true
		} else {
			mark = false
		}
	}
	return
}

// CrossProduct Returns cross product of a,b Matrix.
func CrossProduct(a, b matrix.Matrix) float64 {
	return a[0]*b[1] - a[1]*b[0]
}

// InLine returns true if spot in ab,false else.
func InLine(spot, a, b matrix.Matrix) bool {
	x := spot[0] <= math.Max(a[0], b[0]) && spot[0] >= math.Min(a[0], b[0])
	y := spot[1] <= math.Max(a[1], b[1]) && spot[1] >= math.Min(a[1], b[1])
	return x && y
}

// InLineMatrix returns true if spot in LineMatrix,false else..
func InLineMatrix(spot matrix.Matrix, matr matrix.LineMatrix) bool {
	lines := LineArray(matr)
	for _, line := range lines {
		if InLine(spot, line.Start, line.End) {
			return true
		}
	}
	return false
}

// InRing returns true if spot in ring,false else..
func InRing(spot matrix.Matrix, matr matrix.LineMatrix) bool {
	lines := LineArray(matr)
	for _, line := range lines {
		if InLine(spot, line.Start, line.End) {
			return true
		}
	}
	return false
}

// IsIntersectionEdge returns intersection of edge a and b.
func IsIntersectionEdge(aLine, bLine matrix.LineMatrix) (mark bool) {
	mark, _ = IntersectionEdge(aLine, bLine)
	return
}

// IntersectionEdge returns intersection of edge a and b.
func IntersectionEdge(aLine, bLine matrix.LineMatrix) (mark bool, ps []matrix.Matrix) {
	mark = false
	for i := range aLine {
		for j := range bLine {
			if i < len(aLine)-1 && j < len(bLine)-1 {
				markInter, ip, _, _ := Intersection(matrix.Matrix(aLine[i]),
					matrix.Matrix(aLine[i+1]), matrix.Matrix(bLine[i]), matrix.Matrix(bLine[i+1]))
				if markInter {
					mark = markInter
					ps = append(ps, ip)
				}
			}
		}
	}
	return
}
