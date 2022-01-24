// Package simplify define Douglas Peucker and Topology Preserving simplify.
package simplify

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// DouglasPeuckerLineSimplifier  Simplifies a linestring (sequence of points) using
// the standard Douglas-Peucker algorithm.
type DouglasPeuckerLineSimplifier struct {
	pts               []matrix.Matrix
	usePt             []bool
	distanceTolerance float64
	seg0              matrix.Matrix
	seg1              matrix.Matrix
}

// Simplify  Simplifies a linestring (sequence of points) using
// the standard Douglas-Peucker algorithm.
func (d *DouglasPeuckerLineSimplifier) Simplify() []matrix.Matrix {
	d.usePt = make([]bool, len(d.pts))
	for i := 0; i < len(d.pts); i++ {
		d.usePt[i] = true
	}
	d.simplifySection(0, len(d.pts)-1)
	newPts := []matrix.Matrix{}
	for i := 0; i < len(d.pts); i++ {
		if d.usePt[i] {
			newPts = append(newPts, d.pts[i])
		}
	}
	return newPts
}

func (d *DouglasPeuckerLineSimplifier) simplifySection(i, j int) {
	if (i + 1) == j {
		return
	}
	d.seg0 = d.pts[i]
	d.seg1 = d.pts[j]
	maxDistance := -1.0
	maxIndex := i
	for k := i + 1; k < j; k++ {
		distance := measure.PlanarDistance(d.pts[k], matrix.LineMatrix{d.seg0, d.seg1})
		if distance > maxDistance {
			maxDistance = distance
			maxIndex = k
		}
	}
	if maxDistance <= d.distanceTolerance {
		for k := i + 1; k < j; k++ {
			d.usePt[k] = false
		}
	} else {
		d.simplifySection(i, maxIndex)
		d.simplifySection(maxIndex, j)
	}
}
