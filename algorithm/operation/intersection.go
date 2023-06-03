package operation

import (
	"math"
	"sort"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// IsIntersectedLineSegment returns intersection of a and other.
func IsIntersectedLineSegment(l, o *matrix.LineSegment) bool {
	mark, _ := FindIntersection(l.P0, l.P1, o.P0, o.P1)
	return mark
}

// FindIntersectionLineSegment returns intersection of a and other.
func FindIntersectionLineSegment(l, o *matrix.LineSegment) (bool, IntersectionArray) {
	mark, ips := FindIntersection(l.P0, l.P1, o.P0, o.P1)
	return mark, ips
}

// FindIntersection returns intersection of a and b.
func FindIntersection(aStart, aEnd, bStart, bEnd matrix.Matrix) (mark bool, ips IntersectionArray) {
	u := matrix.Matrix{aEnd[0] - aStart[0], aEnd[1] - aStart[1]}
	v := matrix.Matrix{bEnd[0] - bStart[0], bEnd[1] - bStart[1]}

	determinant := CrossProduct(aStart, aEnd, bStart, bEnd)

	accuracy := calc.DefaultTolerance * math.Max((aStart[0]+aEnd[0])/2.0, (aStart[1]+aEnd[1])/2.0)

	if math.Abs(determinant) < accuracy {
		isEnter := true
		if (u[0] > 0 && v[0] > 0) || (u[1] > 0 && v[1] > 0) {
			isEnter = false
		} else if (u[0] < 0 && v[0] < 0) || (u[1] < 0 && v[1] < 0) {
			isEnter = false
		}
		if in, isVex := InLineSegment(bStart, aStart, aEnd); in {
			if isVex {
				ips = append(ips, Intersection{bStart, false, isEnter, true, true})
			} else {
				ips = append(ips, Intersection{bStart, true, isEnter, true, true})
			}
			mark = true
		}
		if in, isVex := InLineSegment(bEnd, aStart, aEnd); in {
			if isVex {
				ips = append(ips, Intersection{bEnd, false, isEnter, true, true})
			} else {
				ips = append(ips, Intersection{bEnd, true, isEnter, true, true})
			}
			mark = true
		}
		if in, isVex := InLineSegment(aStart, bStart, bEnd); in && !aStart.Equals(bStart) && !aStart.Equals(bEnd) {
			if isVex {
				ips = append(ips, Intersection{aStart, false, isEnter, true, true})
			} else {
				ips = append(ips, Intersection{aStart, true, isEnter, true, true})
			}
			mark = true
		}
		if in, isVex := InLineSegment(aEnd, bStart, bEnd); in && !aEnd.Equals(bStart) && !aEnd.Equals(bEnd) {
			if isVex {
				ips = append(ips, Intersection{aEnd, false, isEnter, true, true})
			} else {
				ips = append(ips, Intersection{aEnd, true, isEnter, true, true})
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
		ina, _ := InLineSegment(ip, aStart, aEnd)
		if inb, _ := InLineSegment(ip, bStart, bEnd); ina && inb {
			isIntersectionPoint := true
			//todo
			// if isVexA && inVexB {
			// 	//isIntersectionPoint = false
			// }
			isOriginal := false
			if ip.EqualsExact(aStart, accuracy) {
				ip = aStart
				isOriginal = true
			} else if ip.EqualsExact(aEnd, accuracy) {
				ip = aEnd
				isOriginal = true
			} else if ip.EqualsExact(bStart, accuracy) {
				ip = bStart
				isOriginal = true
			} else if ip.EqualsExact(bEnd, accuracy) {
				ip = bEnd
				isOriginal = true
			}
			ips = append(ips, Intersection{ip, isIntersectionPoint, determinant < 0, isOriginal, false})

			mark = true
		} else {
			mark = false
		}
	}
	return
}

// CrossProduct Returns cross product of a,b Matrix.
func CrossProduct(aStart, aEnd, bStart, bEnd matrix.Matrix) float64 {
	aa0 := calc.ValueOf(aEnd[0]).SelfSubtractPair(calc.ValueOf(aStart[0]))
	aa1 := calc.ValueOf(aEnd[1]).SelfSubtractPair(calc.ValueOf(aStart[1]))
	bb0 := calc.ValueOf(bEnd[0]).SelfSubtractPair(calc.ValueOf(bStart[0]))
	bb1 := calc.ValueOf(bEnd[1]).SelfSubtractPair(calc.ValueOf(bStart[1]))
	return calc.DeterminantPair(aa0, aa1, bb0, bb1).Value()
}

// IsIntersectedLineMatrix returns intersection of edge a and b.
func IsIntersectedLineMatrix(aLine, bLine matrix.LineMatrix) (mark bool) {
	mark, _ = FindIntersectionLineMatrix(aLine, bLine)
	return
}

// FindIntersectionLineMatrix returns intersection of edge a and b.
func FindIntersectionLineMatrix(aLine, bLine matrix.LineMatrix) (mark bool, ps IntersectionArray) {
	mark = false
	for i := range aLine {
		for j := range bLine {
			if i < len(aLine)-1 && j < len(bLine)-1 {
				markInter, ips := FindIntersection(matrix.Matrix(aLine[i]),
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
	sort.Sort(ps)
	filt := &UniqueIntersectionArrayFilter{}
	for _, v := range ps {
		filt.Filter(v)
	}
	ps = filt.Ips
	return
}
