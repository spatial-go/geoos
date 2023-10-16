package operation

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

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

// PnPolygonByCross returns true if pt is inside pg.
//
// Segments of the polygon are allowed to cross.  InPolygon this case they divide the
// polygon into multiple regions.  The function returns true for points in
// regions on the perimeter of the polygon.  The return value for interior
// regions is determined by a two coloring of the regions.
//
// If point is exactly on a segment or vertex of polygon, the method may return true or
// false.
func PnPolygonByCross(point matrix.Matrix, poly matrix.LineMatrix) bool {
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
