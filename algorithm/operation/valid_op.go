// Package operation define valid func for geometries.
package operation

import (
	"container/ring"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// ValidOP describes a geographic Element Valid
type ValidOP struct {
	matrix.Steric
}

// IsSimple Computes simplicity for geometries.
func (el *ValidOP) IsSimple() bool {
	switch matr := el.Steric.(type) {
	case matrix.Matrix:
		return true
	case matrix.LineMatrix:
		return el.isSimpleLine(matr)
	case matrix.PolygonMatrix:
		return el.isSimplePolygon(matr)
	case matrix.Collection:
		return el.isSimpleCollection(matr)
	default:
		return true

	}
}

// isSimpleMultiSteric Computes simplicity for MultiPoint geometries.
func (el *ValidOP) isSimpleMultiSteric(matr matrix.Collection) bool {
	points := ring.New(len(matr))
	nonSimplePts := true
	for _, v := range matr {
		points.Do(func(i interface{}) {
			if v.Equals(i.(matrix.Steric)) {
				nonSimplePts = false
			}
		})
		if nonSimplePts {
			points.Value = v
			points = points.Next()
		} else {
			return false
		}
	}
	return true
}

// isSimplePolygon Computes simplicity for polygonal geometries.
// Polygonal geometries are simple if and only if
//  all of their component rings are simple.
func (el *ValidOP) isSimplePolygon(matr matrix.PolygonMatrix) bool {
	for _, ring := range matr {
		elem := ValidOP{matrix.LineMatrix(ring)}
		if !elem.IsSimple() {
			return false
		}
	}
	return true
}

// isSimpleCollection Computes simplicity for collection geometries.
//  geometries are simple if and only if
//  all geometries are simple.
func (el *ValidOP) isSimpleCollection(matr matrix.Collection) bool {
	el.isSimpleMultiSteric(matr)
	for _, g := range matr {
		elem := ValidOP{g}
		if !elem.IsSimple() {
			return false
		}
	}
	return true
}

// isSimpleLine Computes simplicity for LineString geometries.
// geometries are simple if they do not self-intersect at interior points
// (i.e. points other than the endpoints)..
func (el *ValidOP) isSimpleLine(matr matrix.LineMatrix) bool {
	lines := matr.ToLineArray()
	numLine := len(lines)
	for i, line1 := range lines {
		for j, line2 := range lines {
			if i == j || j-i == 1 || i-j == 1 {
				continue
			}
			if relate.IsIntersectionLineSegment(line1, line2) {
				if (i == 0 && j == numLine-1) ||
					(j == 0 && i == numLine-1) {
					_, ips := relate.IntersectionLineSegment(line1, line2)
					isIPoint := true
					for _, ip := range ips {
						if !ip.EqualsExact(lines[0].P0, 0.000001) &&
							!ip.EqualsExact(lines[numLine-1].P1, 0.000001) {
							isIPoint = false
						}
					}
					if isIPoint {
						continue
					}
				}
				return false
			}
		}
	}
	return true
}
