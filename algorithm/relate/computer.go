package relate

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// IntersectionPoint overlay point.
type IntersectionPoint struct {
	matrix.Matrix
	IsIntersectionPoint, IsEntering, IsOriginal bool
}

// X Returns x  .
func (ip *IntersectionPoint) X() float64 {
	return ip.Matrix[0]
}

// Y Returns y  .
func (ip *IntersectionPoint) Y() float64 {
	return ip.Matrix[1]
}

// IntersectionPointLine overlay point array.
type IntersectionPointLine []IntersectionPoint

// IsOriginal returns line overlays.
func (ips IntersectionPointLine) IsOriginal() bool {
	for _, v := range ips {
		if v.IsOriginal {
			return true
		}
	}
	return false
}

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
	mark, _ := Intersection(l.Start, l.End, o.Start, o.End)
	return mark
}

// Intersection returns intersection of a and other.
func (l *Line) Intersection(o *Line) (bool, IntersectionPointLine) {
	mark, ips := Intersection(l.Start, l.End, o.Start, o.End)
	return mark, ips
}

// Intersection returns intersection of a and b.
func Intersection(aStart, aEnd, bStart, bEnd matrix.Matrix) (mark bool, ips IntersectionPointLine) {
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
		if InLine(bStart, aStart, aEnd) {
			ips = append(ips, IntersectionPoint{bStart, true, determinant < 0, true})
			mark = true
		}
		if InLine(bEnd, aStart, aEnd) {
			ips = append(ips, IntersectionPoint{bEnd, true, determinant < 0, true})
			mark = true
		}
		if InLine(aStart, bStart, bEnd) {
			ips = append(ips, IntersectionPoint{aStart, true, determinant < 0, true})
			mark = true
		}
		if InLine(aEnd, bStart, bEnd) {
			ips = append(ips, IntersectionPoint{aEnd, true, determinant < 0, true})
			mark = true
		}
	} else {
		ip := matrix.Matrix{(b1*c2 - b2*c1) / determinant, (a2*c1 - a1*c2) / determinant}

		// check if point belongs to segment
		if InLine(ip, aStart, aEnd) && InLine(ip, bStart, bEnd) {
			if ip.Equals(aStart) || ip.Equals(aEnd) || ip.Equals(bStart) || ip.Equals(bEnd) {
				ips = append(ips, IntersectionPoint{ip, true, determinant < 0, true})
			} else {
				ips = append(ips, IntersectionPoint{ip, true, determinant < 0, false})
			}
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
	// x := spot[0] <= math.Max(a[0], b[0]) && spot[0] >= math.Min(a[0], b[0])
	// y := spot[1] <= math.Max(a[1], b[1]) && spot[1] >= math.Min(a[1], b[1])

	if ((spot[0]-a[0])*(a[1]-b[1])) == ((a[0]-b[0])*(spot[1]-a[1])) &&
		(spot[0] >= math.Min(a[0], b[0]) && spot[0] <= math.Max(a[0], b[0])) &&
		((spot[1] >= math.Min(a[1], b[1])) && (spot[1] <= math.Max(a[1], b[1]))) {
		return true
	}
	return false
}

// InLineVertex returns true if spot in LineVertex,false else..
func InLineVertex(spot matrix.Matrix, matr matrix.LineMatrix) (bool, bool) {
	for i, v := range matr {
		if spot.Equals(matrix.Matrix(v)) {
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
func IntersectionEdge(aLine, bLine matrix.LineMatrix) (mark bool, ps IntersectionPointLine) {
	mark = false
	for i := range aLine {
		for j := range bLine {
			if i < len(aLine)-1 && j < len(bLine)-1 {
				markInter, ips := Intersection(matrix.Matrix(aLine[i]),
					matrix.Matrix(aLine[i+1]),
					matrix.Matrix(bLine[j]),
					matrix.Matrix(bLine[j+1]))
				if markInter {
					mark = markInter
					ps = append(ps, ips...)
				}
			}
		}
	}
	return
}
