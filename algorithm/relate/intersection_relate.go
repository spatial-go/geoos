package relate

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/filter"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// IsIntersectionLineSegment returns intersection of a and other.
func IsIntersectionLineSegment(l, o *matrix.LineSegment) bool {
	mark, _ := Intersection(l.P0, l.P1, o.P0, o.P1)
	return mark
}

// IntersectionLineSegment returns intersection of a and other.
func IntersectionLineSegment(l, o *matrix.LineSegment) (bool, IntersectionPointLine) {
	mark, ips := Intersection(l.P0, l.P1, o.P0, o.P1)
	return mark, ips
}

// Intersection returns intersection of a and b.
func Intersection(aStart, aEnd, bStart, bEnd matrix.Matrix) (mark bool, ips IntersectionPointLine) {

	u := matrix.Matrix{aEnd[0] - aStart[0], aEnd[1] - aStart[1]}
	v := matrix.Matrix{bEnd[0] - bStart[0], bEnd[1] - bStart[1]}

	determinant := CrossProduct(u, v)

	if determinant == 0 {
		isEnter := true
		if (u[0] > 0 && v[0] > 0) || (u[1] > 0 && v[1] > 0) {
			isEnter = false
		} else if (u[0] < 0 && v[0] < 0) || (u[1] < 0 && v[1] < 0) {
			isEnter = false
		}
		if in, isVex := InLine(bStart, aStart, aEnd); in {
			if isVex {
				ips = append(ips, IntersectionPoint{bStart, false, isEnter, true, true})
			} else {
				ips = append(ips, IntersectionPoint{bStart, true, isEnter, true, true})
			}
			mark = true
		}
		if in, isVex := InLine(bEnd, aStart, aEnd); in {
			if isVex {
				ips = append(ips, IntersectionPoint{bEnd, false, isEnter, true, true})
			} else {
				ips = append(ips, IntersectionPoint{bEnd, true, isEnter, true, true})
			}
			mark = true
		}
		if in, isVex := InLine(aStart, bStart, bEnd); in && !aStart.Equals(bStart) && !aStart.Equals(bEnd) {
			if isVex {
				ips = append(ips, IntersectionPoint{aStart, false, isEnter, true, true})
			} else {
				ips = append(ips, IntersectionPoint{aStart, true, isEnter, true, true})
			}
			mark = true
		}
		if in, isVex := InLine(aEnd, bStart, bEnd); in && !aEnd.Equals(bStart) && !aEnd.Equals(bEnd) {
			if isVex {
				ips = append(ips, IntersectionPoint{aEnd, false, isEnter, true, true})
			} else {
				ips = append(ips, IntersectionPoint{aEnd, true, isEnter, true, true})
			}
			mark = true
		}
	} else {
		aa1 := calc.ValueOf(aEnd[1]).SelfSubtractPair(calc.ValueOf(aStart[1]))
		bb1 := calc.ValueOf(aStart[0]).SelfSubtractPair(calc.ValueOf(aEnd[0]))
		cc1 := calc.DeterminantPair(calc.ValueOf(-aStart[0]), bb1, calc.ValueOf(aStart[1]), aa1)
		aa2 := calc.ValueOf(bEnd[1]).SelfSubtractPair(calc.ValueOf(bStart[1]))
		bb2 := calc.ValueOf(bStart[0]).SelfSubtractPair(calc.ValueOf(bEnd[0]))
		cc2 := calc.DeterminantPair(calc.ValueOf(0).SelfSubtractPair(aa2), bb2, calc.ValueOf(bStart[1]), calc.ValueOf(bStart[0]))

		x := calc.DeterminantPair(bb1, bb2, cc1, cc2).SelfDividePair(calc.ValueOf(determinant)).Value()
		y := calc.DeterminantPair(aa2, aa1, cc2, cc1).SelfDividePair(calc.ValueOf(determinant)).Value()
		ip := matrix.Matrix{x, y}

		// check if point belongs to segment
		ina, isVexA := InLine(ip, aStart, aEnd)
		inb, inVexB := InLine(ip, bStart, bEnd)
		if ina && inb {
			isIntersectionPoint := true
			if isVexA && inVexB {
				//isIntersectionPoint = false  //todo
			}
			isOriginal := false
			if ip.Equals(aStart) || ip.Equals(aEnd) || ip.Equals(bStart) || ip.Equals(bEnd) {
				isOriginal = true
			}
			ips = append(ips, IntersectionPoint{ip, isIntersectionPoint, determinant < 0, isOriginal, false})

			mark = true
		} else {
			mark = false
		}
	}
	return
}

// CrossProduct Returns cross product of a,b Matrix.
func CrossProduct(a, b matrix.Matrix) float64 {
	return calc.Determinant(a[0], a[1], b[0], b[1]).Value()
}

// InLine returns true if spot in ab,false else.
func InLine(spot, a, b matrix.Matrix) (in bool, isVertex bool) {
	// x := spot[0] <= math.Max(a[0], b[0]) && spot[0] >= math.Min(a[0], b[0])
	// y := spot[1] <= math.Max(a[1], b[1]) && spot[1] >= math.Min(a[1], b[1])

	if spot.Equals(a) || spot.Equals(b) {
		return true, true
	}
	ax := (spot[0] - a[0]) * (a[1] - b[1])
	bx := (a[0] - b[0]) * (spot[1] - a[1])
	if math.Abs(ax-bx) < calc.AccuracyFloat &&
<<<<<<< HEAD
		(spot[0] >= math.Min(a[0], b[0]) && spot[0] <= math.Max(a[0], b[0])) &&
		((spot[1] >= math.Min(a[1], b[1])) && (spot[1] <= math.Max(a[1], b[1]))) {
=======
		(spot[0] + calc.AccuracyFloat >= math.Min(a[0], b[0]) && spot[0] - calc.AccuracyFloat <= math.Max(a[0], b[0])) &&
		(spot[1] + calc.AccuracyFloat >= math.Min(a[1], b[1]) && spot[1] - calc.AccuracyFloat <= math.Max(a[1], b[1])) {
>>>>>>> 7d68fdd0673acd057b734be42fa63ca1c896ad14
		return true, false
	}
	return false, false
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
	lines := matr.ToLineArray()
	for _, line := range lines {
		if in, _ := InLine(spot, line.P0, line.P1); in {
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
	filt := &UniqueIntersectionEdgeFilter{}
	for _, v := range ps {
		filt.Filter(v)
	}
	ps = filt.Ips
	return
}

// UniqueIntersectionEdgeFilter  A Filter that extracts a unique array.
type UniqueIntersectionEdgeFilter struct {
	Ips IntersectionPointLine
}

// Filter Performs an operation with the provided .
func (u *UniqueIntersectionEdgeFilter) Filter(ip interface{}) bool {
	return u.add(ip)
}

func (u *UniqueIntersectionEdgeFilter) add(ip interface{}) bool {
	hasMatrix := false
	for _, v := range u.Ips {
		if v.Matrix.Equals(ip.(IntersectionPoint).Matrix) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.Ips = append(u.Ips, ip.(IntersectionPoint))
		return true
	}
	return false
}

// Entities  Returns the gathered Matrixes.
func (u *UniqueIntersectionEdgeFilter) Entities() interface{} {
	return u.Ips
}

// compile time checks
var (
	_ filter.Filter = &UniqueIntersectionEdgeFilter{}
)
