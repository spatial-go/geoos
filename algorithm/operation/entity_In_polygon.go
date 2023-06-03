package operation

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// in or out polygon.
const (
	PartInPolygon    = 4
	OnlyInPolygon    = 3
	OnlyInLine       = 2
	OnlyOutPolygon   = -1
	PartOutPolygon   = -2
	DefaultInPolygon = 0
	BothPolygon      = 1
	IncludePolygon   = -3
)

// InPolygon returns true if Steric point and entity is inside pg  .
//
// Segments of the polygon are allowed to cross.  InPolygon this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If point is exactly on a segment or vertex of polygon, the method may return true or
// false.
func InPolygon(arg matrix.Steric, poly matrix.PolygonMatrix) (int, false int) {
	pointInPolygon := DefaultInPolygon
	entityInPolygon := DefaultInPolygon

	switch m := arg.(type) {
	case matrix.Matrix:
		if in := pointInRing(m, poly[0]); in == OnlyInPolygon {
			pointInPolygon = OnlyInPolygon
			for i := 1; i < len(poly); i++ {
				if IsPnPolygon(m, poly[i]) {
					pointInPolygon = OnlyOutPolygon
					break
				}
			}
		} else if in == OnlyInLine {
			pointInPolygon = OnlyInLine
		} else {
			pointInPolygon = OnlyOutPolygon
		}
		entityInPolygon = pointInPolygon
	case matrix.LineMatrix:
		pointInPolygon, entityInPolygon = lineInPolygon(m, poly)
	case matrix.PolygonMatrix:
		pointInPolygon, entityInPolygon = lineInPolygon(matrix.LineMatrix(m[0]), poly)
		if entityInPolygon == OnlyOutPolygon || entityInPolygon == PartOutPolygon {
			if _, v := lineInPolygon(matrix.LineMatrix(poly[0]), m); v == OnlyInPolygon || v == PartInPolygon {
				entityInPolygon = IncludePolygon
			}
		}
	}

	return pointInPolygon, entityInPolygon
}

func lineInPolygon(m matrix.LineMatrix, poly matrix.PolygonMatrix) (int, int) {
	pointInPolygon := DefaultInPolygon
	entityInPolygon := DefaultInPolygon

	for i, p := range m {
		if i == 0 || i == len(m)-1 {
			pInPolygon := OnlyOutPolygon
			if inShell := pointInRing(matrix.Matrix(p), poly[0]); inShell == OnlyInPolygon {
				pInPolygon = OnlyInPolygon
				for i := 1; i < len(poly); i++ {
					if inHoles := pointInRing(matrix.Matrix(p), poly[i]); inHoles == OnlyInPolygon {
						pInPolygon = OnlyOutPolygon
						break
					} else if inHoles == OnlyInLine {
						pInPolygon = OnlyInLine
						break
					}
				}
			} else if inShell == OnlyInLine {
				pInPolygon = OnlyInLine
			}
			pointInPolygon = calcInPolygon(pointInPolygon, pInPolygon)
		}
	}

	for _, l := range m.ToLineArray() {
		lInPolygon := 0
		if inShell := lineSegmentInRing(l, poly[0]); inShell == OnlyInPolygon {
			lInPolygon = OnlyInPolygon

			for i := 1; i < len(poly); i++ {
				if inHoles := lineSegmentInRing(l, poly[i]); inHoles == OnlyInPolygon {
					lInPolygon = OnlyOutPolygon
					break
				} else if inHoles == OnlyInLine {
					lInPolygon = OnlyInLine
					break
				}
			}
		} else {
			lInPolygon = inShell
		}
		entityInPolygon = calcInPolygon(entityInPolygon, lInPolygon)
	}

	return pointInPolygon, entityInPolygon
}

func pointInRing(p matrix.Matrix, r matrix.LineMatrix) int {
	if !r.IsClosed() {
		return DefaultInPolygon
	}
	if InLineMatrix(p, r) {
		return OnlyInLine
	}
	if IsPnPolygon(p, r) {
		return OnlyInPolygon
	}
	return OnlyOutPolygon
}

func lineSegmentInRing(l *matrix.LineSegment, r matrix.LineMatrix) int {
	if !r.IsClosed() {
		return DefaultInPolygon
	}
	p0InPolygon, p1InPolygon := 0, 0
	if InLineMatrix(l.P0, r) {
		p0InPolygon = OnlyInLine
	} else {
		if IsPnPolygon(l.P0, r) {
			p0InPolygon = OnlyInPolygon
		} else {
			p0InPolygon = OnlyOutPolygon
		}
	}
	if InLineMatrix(l.P1, r) {
		p1InPolygon = OnlyInLine
	} else {
		if IsPnPolygon(l.P1, r) {
			p1InPolygon = OnlyInPolygon
		} else {
			p1InPolygon = OnlyOutPolygon
		}
	}
	switch {
	case p0InPolygon == OnlyInLine && p1InPolygon == OnlyInLine:
		_, ips := FindIntersectionLineMatrix(matrix.LineMatrix{l.P0, l.P1}, r)
		isCollinear := true
		for _, ip := range ips {
			if !ip.IsCollinear {
				isCollinear = false
			}
		}
		if isCollinear {
			return OnlyInLine
		}
		if IsPnPolygon(matrix.Matrix{(l.P0[0] + l.P1[0]) / 2.0, (l.P0[1] + l.P1[1]) / 2.0}, r) {
			return OnlyInPolygon
		}
		return OnlyOutPolygon

	case p0InPolygon == OnlyInPolygon && p1InPolygon == OnlyInPolygon:
		return OnlyInPolygon
	case p0InPolygon == OnlyOutPolygon && p1InPolygon == OnlyOutPolygon:
		if mark, _ := FindIntersectionLineMatrix(matrix.LineMatrix{l.P0, l.P1}, r); mark {
			return BothPolygon
		}
		return OnlyOutPolygon
	default:
		return p0InPolygon + p1InPolygon - OnlyInLine
	}
}

func calcInPolygon(old, new int) int {
	switch old {
	case DefaultInPolygon:
		return new
	case OnlyInPolygon:
		if new == OnlyInPolygon {
			return new
		} else if new == OnlyInLine {
			return PartInPolygon
		}
		return BothPolygon
	case OnlyOutPolygon:
		if new == OnlyOutPolygon {
			return new
		} else if new == OnlyInLine {
			return PartOutPolygon
		}
		return BothPolygon
	case BothPolygon:
		return BothPolygon
	case OnlyInLine:
		if new == OnlyInLine {
			return new
		} else if new == OnlyInPolygon {
			return PartInPolygon
		} else if new == OnlyOutPolygon {
			return PartOutPolygon
		}
		return PartOutPolygon
	case PartOutPolygon:
		if new == OnlyInPolygon {
			return BothPolygon
		}
		return PartOutPolygon
	case PartInPolygon:
		if new == OnlyOutPolygon {
			return BothPolygon
		}
		return PartInPolygon
	}
	return DefaultInPolygon
}

// IsPnPolygon returns true if pt is inside pg.
//
// Segments of the polygon are allowed to cross.  IsPnPolygon this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If point is exactly on a segment or vertex of polygon, the method may return true or
// false.
func IsPnPolygon(point matrix.Matrix, poly matrix.LineMatrix) bool {
	if len(poly) < 3 {
		return false
	}
	a := poly[0]
	in := rayIntersectsSegment(point, poly[len(poly)-1], a)
	for _, b := range poly[1:] {
		if rayIntersectsSegment(point, a, b) {
			in = !in
		}
		a = b
	}
	return in
}

// Segment intersect expression from
// https://www.ecse.rpi.edu/Homepages/wrf/Research/Short_Notes/pnpoly.html
//
// Currently the compiler in lines the function by default.
func rayIntersectsSegment(p, a, b matrix.Matrix) bool {

	////c := (b[0]-a[0])*(p[1]-a[1])/(b[1]-a[1]) + a[0]
	ax := calc.ValueOf(b[0]).Subtract(a[0], 0)
	bx := calc.ValueOf(p[1]).Subtract(a[1], 0)
	by := calc.ValueOf(b[1]).Subtract(a[1], 0)
	cc := ax.MultiplyPair(bx).DividePair(by).Add(a[0], 0).Value()
	return (a[1] > p[1]) != (b[1] > p[1]) &&
		p[0] < cc

	// return (a[1] > p[1]) != (b[1] > p[1]) &&
	// 	p[0] < (b[0]-a[0])*(p[1]-a[1])/(b[1]-a[1])+a[0]
}
