// package graph ...

package graph

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// in or out polygon.
const (
	OnlyInPolygon    = 3
	OnlyInLine       = 2
	OnlyOutPolygon   = -1
	DefaultInPolygon = 0
	BothPolygon      = 1
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
				entityInPolygon = BothPolygon
			}
		}
	}

	return pointInPolygon, entityInPolygon
}

func lineInPolygon(m matrix.LineMatrix, poly matrix.PolygonMatrix) (int, int) {
	pointInPolygon := DefaultInPolygon
	entityInPolygon := DefaultInPolygon
	for _, p := range m {
		pInPolygon := OnlyOutPolygon
		lInPolygon := OnlyOutPolygon
		if inShell, onShell := pointInRing(p, poly[0]); inShell {
			pInPolygon = OnlyInPolygon
			lInPolygon = OnlyInPolygon

			for i := 1; i < len(poly); i++ {
				if inHoles, onHoles := pointInRing(p, poly[i]); inHoles {
					pInPolygon = OnlyOutPolygon
					lInPolygon = OnlyOutPolygon

					break
				} else if onHoles {
					pInPolygon = OnlyInLine
					lInPolygon = entityInPolygon
					break
				}
			}
		} else if onShell {
			pInPolygon = OnlyInLine
			lInPolygon = entityInPolygon
		}

		if entityInPolygon == OnlyInPolygon {
			break
		}
		pointInPolygon = calcInPolygon(pointInPolygon, pInPolygon)
		entityInPolygon = calcInPolygon(entityInPolygon, lInPolygon)
	}
	return pointInPolygon, entityInPolygon
}

func pointInRing(p matrix.Matrix, r matrix.LineMatrix) (bool, bool) {
	if !r.IsClosed() {
		return false, false
	}
	if relate.InLineMatrix(p, r) {
		return false, true
	}
	if relate.InPolygon(p, r) {
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
		}
		return BothPolygon
	case OnlyOutPolygon:
		if new == OnlyOutPolygon {
			return new
		}
		return BothPolygon
	case BothPolygon:
		return BothPolygon
	case OnlyInLine:
		if new == OnlyInLine {
			return new
		}
		return BothPolygon
	}
	return DefaultInPolygon
}
