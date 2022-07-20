package de9im

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
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

// IsInPolygon returns true if Steric point and entity is inside pg  .
//
// Segments of the polygon are allowed to cross.  InPolygon this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If point is exactly on a segment or vertex of polygon, the method may return true or
// false.
func IsInPolygon(arg matrix.Steric, poly matrix.PolygonMatrix) (int, int) {
	pointInPolygon := DefaultInPolygon
	entityInPolygon := DefaultInPolygon

	switch m := arg.(type) {
	case matrix.Matrix:
		if in, on := pointInRing(m, poly[0]); in {
			pointInPolygon = OnlyInPolygon
			for i := 1; i < len(poly); i++ {
				if relate.InPolygon(m, poly[i]) {
					pointInPolygon = OnlyOutPolygon
					break
				}
			}
		} else if on {
			pointInPolygon = OnlyInLine
		} else {
			pointInPolygon = OnlyOutPolygon
		}
		entityInPolygon = pointInPolygon
	case matrix.LineMatrix:
		pointInPolygon, entityInPolygon = lineInPolygon(m, poly)
	case matrix.PolygonMatrix:
		pointInPolygon, entityInPolygon = lineInPolygon(matrix.LineMatrix(m[0]), poly)
		if entityInPolygon == OnlyOutPolygon {
			if _, v := lineInPolygon(matrix.LineMatrix(poly[0]), m); v == OnlyInPolygon {
				entityInPolygon = IncludePolygon
			}
		}
	}

	return pointInPolygon, entityInPolygon
}

func lineInPolygon(m matrix.LineMatrix, poly matrix.PolygonMatrix) (int, int) {
	pointInPolygon := DefaultInPolygon
	entityInPolygon := DefaultInPolygon
	if b, err := m.Boundary(); err == nil {
		for _, p := range b.(matrix.Collection) {
			pInPolygon := OnlyOutPolygon
			if inShell, onShell := pointInRing(p.(matrix.Matrix), poly[0]); inShell {
				pInPolygon = OnlyInPolygon
				for i := 1; i < len(poly); i++ {
					if inHoles, onHoles := pointInRing(p.(matrix.Matrix), poly[i]); inHoles {
						pInPolygon = OnlyOutPolygon
						break
					} else if onHoles {
						pInPolygon = OnlyInLine
						break
					}
				}
			} else if onShell {
				pInPolygon = OnlyInLine
			}
			pointInPolygon = calcInPolygon(pointInPolygon, pInPolygon)
		}
	}
	for _, p := range m {
		lInPolygon := OnlyOutPolygon
		if inShell, onShell := pointInRing(p, poly[0]); inShell {
			lInPolygon = OnlyInPolygon

			for i := 1; i < len(poly); i++ {
				if inHoles, onHoles := pointInRing(p, poly[i]); inHoles {
					lInPolygon = OnlyOutPolygon

					break
				} else if onHoles {
					lInPolygon = entityInPolygon
					break
				}
			}
		} else if onShell {
			lInPolygon = entityInPolygon
		}
		entityInPolygon = calcInPolygon(entityInPolygon, lInPolygon)
	}
	// if entityInPolygon == DefaultInPolygon {
	// 	entityInPolygon = pointInPolygon
	// }
	return pointInPolygon, entityInPolygon
}

func pointInRing(p matrix.Matrix, r matrix.LineMatrix) (bool, bool) {
	if !r.IsClosed() {
		return false, false
	}
	if relate.InLineMatrix(p, r) {
		return false, true
	}
	if InPolygon(p, r) {
		return true, false
	}
	return false, false
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

// InPolygon returns true if pt is inside pg.
//
// Segments of the polygon are allowed to cross.  InPolygon this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If point is exactly on a segment or vertex of polygon, the method may return true or
// false.
func InPolygon(point matrix.Matrix, poly matrix.LineMatrix) bool {
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

//OddEvenFill:
// Specifies that the region is filled using the odd even fill rule.
// With this rule, we determine whether a point is inside the shape
// by using the following method. Draw a horizontal line from the point
// to a location outside the shape, and count the number of intersections.
// If the number of intersections is an odd number, the point is inside the shape.
// This mode is the default.

// WindingFill:
// Specifies that the region is filled using the non zero winding rule.
// With this rule, we determine whether a point is inside the shape by
// using the following method. Draw a horizontal line from the point to
// a location outside the shape. Determine whether the direction of the
// line at each intersection point is up or down. The winding number is
// determined by summing the direction of each intersection. If the number
// is non zero, the point is inside the shape. This fill mode can also in most
// cases be considered as the intersection of closed shapes.
const (
	OddEvenFill = iota
	WindingFill
)

// PointInPolygon returns true if pt is inside pg.
//
// Segments of the polygon are allowed to cross.  InPolygon this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If point is exactly on a segment or vertex of polygon, the method may return true or
// false.
func PointInPolygon(point matrix.Matrix, poly matrix.LineMatrix) bool {
	//  https://wrf.ecse.rpi.edu/Research/Short_Notes/pnpoly.html
	//  https://stackoverflow.com/questions/217578/how-can-i-determine-whether-a-2d-point-is-within-a-polygon
	//  Arrays containing the x- and y-coordinates of the polygon's vertices.
	pointNum := len(poly)
	intersectCount := 0 //cross points count of x
	precision := 2e-10
	p := point

	p1 := poly[0] //left vertex
	for i := 0; i < pointNum; i++ {
		if p[1] == p1[1] && p[0] == p1[0] {
			return true
		}
		p2 := poly[i%pointNum]
		if p[1] < math.Min(p1[1], p2[1]) || p[1] > math.Max(p1[1], p2[1]) {
			p1 = p2
			continue //next ray left point
		}

		if p[1] > math.Min(p1[1], p2[1]) && p[1] < math.Max(p1[1], p2[1]) {
			if p[0] <= math.Max(p1[0], p2[0]) { //x is before of ray
				if p1[1] == p2[1] && p[0] >= math.Min(p1[0], p2[0]) {
					return true
				}

				if p1[0] == p2[0] { //ray is vertical
					if p1[0] == p[0] { //overlies on a vertical ray
						return true
					}
					//before ray
					intersectCount++

				} else { //cross point on the left side
					xinters := (p[1]-p1[1])*(p2[0]-p1[0])/(p2[1]-p1[1]) + p1[0]
					if math.Abs(p[0]-xinters) < precision {
						return true
					}

					if p[0] < xinters { //before ray
						intersectCount++
					}
				}
			}
		} else { //special case when ray is crossing through the vertex
			if p[1] == p2[1] && p[0] <= p2[0] { //p crossing over p2
				p3 := poly[(i+1)%pointNum]
				if p[1] >= math.Min(p1[1], p3[1]) && p[1] <= math.Max(p1[1], p3[1]) {
					intersectCount++
				} else {
					intersectCount += 2
				}
			}
		}
		p1 = p2 //next ray left point
	}
	return intersectCount%2 != 0
}
