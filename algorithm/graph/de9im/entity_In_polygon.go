// package de9im ...

package de9im

import (
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
