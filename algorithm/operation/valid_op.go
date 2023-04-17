// Package operation define valid func for geometries.
package operation

import (
	"container/ring"

	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/graph/de9im"
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
			if mark, ips := relate.IntersectionLineSegment(line1, line2); mark {
				if (i == 0 && j == numLine-1) ||
					(j == 0 && i == numLine-1) {
					isIPoint := true
					for _, ip := range ips {
						if !ip.EqualsExact(lines[0].P0, calc.DefaultTolerance) &&
							!ip.EqualsExact(lines[numLine-1].P1, calc.DefaultTolerance) {
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

// CorrectPolygonMatrixSelfIntersect correct self intersect for polygon.
func CorrectPolygonMatrixSelfIntersect(ms matrix.Steric) matrix.Steric {
	if p, ok := ms.(matrix.PolygonMatrix); ok {
		if p.IsEmpty() {
			return p
		}
		shell := matrix.LineMatrix(p[0])
		if shell.IsEmpty() {
			return p
		}

		if !shell.IsClosed() {
			shell = append(shell, shell[0])
		}
		mulitPoly := matrix.Collection{}
		for {
			if res, ok := CorrectRingSelfIntersect(shell); ok {
				mulitPoly = append(mulitPoly, matrix.PolygonMatrix{res[0]})
				shell = res[1]
			} else {
				mulitPoly = append(mulitPoly, matrix.PolygonMatrix{shell})
				break
			}
		}
		if mulitPoly.IsEmpty() {
			return p
		}
		for _, s := range mulitPoly {
			poly := s.(matrix.PolygonMatrix)
			for i := 1; i < len(p); i++ {
				if de9im.IM(poly, matrix.LineMatrix(p[i])).IsCovers() {
					poly = append(poly, matrix.LineMatrix(p[i]))
				}
			}
		}
		return mulitPoly
	}
	return ms
}

// CorrectRingSelfIntersect correct self intersect for ring.
func CorrectRingSelfIntersect(shell matrix.LineMatrix) ([]matrix.LineMatrix, bool) {
	numShell := len(shell)
	result := make([]matrix.LineMatrix, 2)
	for i := 0; i < numShell-1; i++ {
		for j := 0; j < numShell-1; j++ {
			if i == j || j-i == 1 || i-j == 1 {
				continue
			}
			selfip := relate.IntersectionPoint{}
			if ok, ips := relate.IntersectionLineSegment(&matrix.LineSegment{P0: shell[i], P1: shell[i+1]},
				&matrix.LineSegment{P0: shell[j], P1: shell[j+1]}); ok {
				if (i == 0 && j == numShell-2) ||
					(j == 0 && i == numShell-2) {
					isIPoint := true
					for _, ip := range ips {
						if !ip.EqualsExact(matrix.Matrix(shell[0]), calc.DefaultTolerance) &&
							!ip.EqualsExact(matrix.Matrix(shell[numShell-1]), calc.DefaultTolerance) {
							isIPoint = false
							selfip = ip
						}
					}
					if isIPoint {
						continue
					}
				}
				first := 0
				second := 0
				if i < j {
					first, second = i, j
				} else {
					first, second = j, i
				}
				if selfip.Matrix == nil {
					selfip = ips[0]
				}
				if selfip.EqualsExact(matrix.Matrix(shell[first]), calc.DefaultTolerance) {
					result[0] = append(result[0], shell[:first]...)
					result[1] = append(result[1], shell[first:second+1]...)
				} else if selfip.EqualsExact(matrix.Matrix(shell[first+1]), calc.DefaultTolerance) {
					result[0] = append(result[0], shell[:first+1]...)
					result[1] = append(result[1], shell[first+1:second+1]...)
				} else {
					result[0] = append(result[0], shell[:first]...)
					result[0] = append(result[0], selfip.Matrix)
					result[1] = append(result[1], selfip.Matrix)
					result[1] = append(result[1], shell[first+1:second+1]...)
				}
				if selfip.EqualsExact(matrix.Matrix(shell[second]), calc.DefaultTolerance) {
					result[0] = append(result[0], shell[second:]...)
				} else if selfip.EqualsExact(matrix.Matrix(shell[second+1]), calc.DefaultTolerance) {
					result[0] = append(result[0], shell[second+1:]...)
					result[1] = append(result[1], shell[second+1])
				} else {
					result[0] = append(result[0], selfip.Matrix)
					result[0] = append(result[0], shell[second+1:]...)
					result[1] = append(result[1], selfip.Matrix)
				}

				return result, true
			}
		}
	}
	return nil, false
}
