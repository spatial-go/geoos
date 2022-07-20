package de9im

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
)

func (r *RelationshipByDegrees) lineAndLineAnalyse(pointInPolygon, entityInPolygon int) {
	switch r.nLine {
	case 0:

		r.IM.Set(2, 2, 2)
		r.IM.Set(0, 2, 1)
		r.IM.Set(2, 0, 1)
		r.calcBoundaryIM()

		switch _, maxVal := calcMaxDegrees(r.degrees...); maxVal {
		case 4:
			r.IM.Set(0, 0, 0)
		default:
			r.IM.Set(0, 0, -1)
		}

	default:
		r.IM.Set(0, 0, 1)
		r.IM.Set(2, 2, 2)
		r.calcBoundaryIM()
		degs := []int{r.degrees[0], r.degrees[1], r.degrees[2]}
		switch r.nPoint {
		case 2:
			switch calcSumDegrees(degs...) {
			case 3:
				if r.haveIntersectionVertex[0] == 2 {
					r.IM.Set(0, 2, -1)
					r.IM.Set(2, 0, 1)
				} else {
					r.IM.Set(0, 2, 1)
					r.IM.Set(2, 0, -1)
				}
			case 4:
				switch _, maxVal := calcMaxDegrees(degs...); maxVal {
				case 3:
					r.IM.Set(0, 2, 1)
					r.IM.Set(2, 0, 1)
				default:
					if r.haveIntersectionVertex[0] == 2 {
						r.IM.Set(0, 2, -1)
						r.IM.Set(2, 0, 1)
					} else {
						r.IM.Set(0, 2, 1)
						r.IM.Set(2, 0, -1)
					}
				}
			default:
				r.IM.SetString(RelateStrings[RLL2])
			}
		default:
			r.IM.Set(0, 2, 1)
			r.IM.Set(2, 0, 1)
		}
	}

}

func (r *RelationshipByDegrees) calcBoundaryIM() {
	switch r.haveIntersectionVertex[0] {
	case 0:
		r.IM.Set(1, 1, -1)
		r.IM.Set(1, 0, -1)
		r.IM.Set(1, 2, 0)
		switch r.haveIntersectionVertex[1] {
		case 0:
			r.IM.Set(0, 1, -1)
			r.IM.Set(2, 1, 0)
		case 1:
			r.IM.Set(0, 1, 0)
			r.IM.Set(2, 1, 0)
		case 2:
			r.IM.Set(0, 1, 0)
			r.IM.Set(2, 1, -1)
		}
	case 1:
		r.IM.Set(1, 2, 0)

		aNums, bNums := r.numsOfBoundaryIM()

		if aNums >= 1 {
			r.IM.Set(1, 1, 0)
			r.IM.Set(1, 0, -1)
		} else {
			r.IM.Set(1, 1, -1)
			r.IM.Set(1, 0, 0)
		}
		switch r.haveIntersectionVertex[1] {
		case 0:
			r.IM.Set(0, 1, -1)
			r.IM.Set(2, 1, 0)
		case 1:
			r.IM.Set(2, 1, 0)
			if bNums >= 1 {
				r.IM.Set(0, 1, -1)

			} else {
				r.IM.Set(0, 1, 0)
			}
		case 2:
			r.IM.Set(2, 1, -1)
			r.IM.Set(0, 1, 0)
		}
	case 2: // && r.haveIntersectionVertex[1] == 2:

		r.IM.Set(1, 2, -1)

		aNums, bNums := r.numsOfBoundaryIM()

		if aNums >= 1 {
			r.IM.Set(1, 1, 0)
			if aNums >= 2 {
				r.IM.Set(1, 0, -1)
			} else {
				r.IM.Set(1, 0, 0)
			}

		} else {
			r.IM.Set(1, 1, -1)
			r.IM.Set(1, 0, 0)
		}
		switch r.haveIntersectionVertex[1] {
		case 0:
			r.IM.Set(0, 1, -1)
			r.IM.Set(2, 1, 0)
		case 1:
			r.IM.Set(2, 1, 0)
			if bNums >= 1 {
				r.IM.Set(0, 1, -1)
			} else {
				r.IM.Set(0, 1, 0)
			}
		case 2:
			r.IM.Set(2, 1, -1)
			if bNums >= 2 {
				r.IM.Set(0, 1, -1)
			} else {
				r.IM.Set(0, 1, 0)
			}
		}
	}
}

func (r *RelationshipByDegrees) numsOfBoundaryIM() (int, int) {
	aNums, bNums := 0, 0
	if r.boundary[0] == nil || r.boundary[1] == nil {
		return aNums, bNums
	}
	if r.boundary[0][0].EqualsExact(r.boundary[1][0], calc.DefaultTolerance) ||
		r.boundary[0][0].EqualsExact(r.boundary[1][1], calc.DefaultTolerance) ||
		r.boundary[0][1].EqualsExact(r.boundary[1][0], calc.DefaultTolerance) ||
		r.boundary[0][1].EqualsExact(r.boundary[1][1], calc.DefaultTolerance) {
		aNums, bNums = 1, 1
	}
	if (r.boundary[0][0].EqualsExact(r.boundary[1][0], calc.DefaultTolerance) ||
		r.boundary[0][0].EqualsExact(r.boundary[1][1], calc.DefaultTolerance)) &&
		(r.boundary[0][1].EqualsExact(r.boundary[1][0], calc.DefaultTolerance) ||
			r.boundary[0][1].EqualsExact(r.boundary[1][1], calc.DefaultTolerance)) {
		aNums = 2
	}

	if (r.boundary[1][0].EqualsExact(r.boundary[0][0], calc.DefaultTolerance) ||
		r.boundary[1][0].EqualsExact(r.boundary[0][1], calc.DefaultTolerance)) &&
		(r.boundary[1][1].EqualsExact(r.boundary[0][0], calc.DefaultTolerance) ||
			r.boundary[1][1].EqualsExact(r.boundary[0][1], calc.DefaultTolerance)) {
		bNums = 2
	}
	return aNums, bNums
}

func calcMaxDegrees(degs ...int) (int, int) {
	maxVal := 0
	minVal := math.MaxInt32
	for _, v := range degs {
		if maxVal < v {
			maxVal = v
		}
		if minVal > v {
			minVal = v
		}
	}
	return maxVal, minVal
}

func calcSumDegrees(degs ...int) int {
	sumVal := 0
	for _, v := range degs {
		sumVal += v
	}
	return sumVal
}
