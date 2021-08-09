package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

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
	return (a[1] > p[1]) != (b[1] > p[1]) &&
		p[0] < (b[0]-a[0])*(p[1]-a[1])/(b[1]-a[1])+a[0]
}
