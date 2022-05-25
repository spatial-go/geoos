package simplify

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Simplify Simplifies a geometry using a given tolerance.
func Simplify(geom matrix.Steric, distanceTolerance float64) matrix.Steric {
	tss := &DouglasPeuckerSimplifier{geom, distanceTolerance, true}
	return tss.getResultGeometry()
}

// DouglasPeuckerSimplifier Simplifies a  Geometry using the Douglas-Peucker algorithm.
//  Ensures that any polygonal geometries returned are valid.
//  Simple lines are not guaranteed to remain simple after simplification.
//  All geometry types are handled.
//  Empty and point geometries are returned unchanged.
//  Empty geometry components are deleted.
type DouglasPeuckerSimplifier struct {
	inputGeom             matrix.Steric
	distanceTolerance     float64
	isEnsureValidTopology bool
}

// Gets the simplified geometry.
func (d *DouglasPeuckerSimplifier) getResultGeometry() matrix.Steric {
	// empty input produces an empty result
	if d.inputGeom.IsEmpty() {
		return d.inputGeom
	}

	result, _ := (&DPTransformer{&Transformer{d.inputGeom, true, false}, d.isEnsureValidTopology, d.distanceTolerance}).Transform(d.inputGeom)
	return result
}
