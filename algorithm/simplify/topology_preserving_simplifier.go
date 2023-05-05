package simplify

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/index/quadtree"
)

// TopologyPreservingSimplifier Simplifies a geometry and ensures that
//  the result is a valid geometry having the
//  same dimension and number of components as the input,
//  and with the components having the same topological relationship.
type TopologyPreservingSimplifier struct {
	InputGeom      matrix.Steric
	lineSimplifier *TaggedLinesSimplifier
	linestrings    []*TaggedLineString
}

// Simplify ...
func (t *TopologyPreservingSimplifier) Simplify(geom matrix.Steric, distanceTolerance float64) matrix.Steric {
	tss := &TopologyPreservingSimplifier{InputGeom: geom}
	tss.lineSimplifier = &TaggedLinesSimplifier{
		&LineSegmentIndex{quadtree.NewQuadtree()},
		&LineSegmentIndex{quadtree.NewQuadtree()},
		distanceTolerance,
	}
	tss.setDistanceTolerance(distanceTolerance)
	return tss.getResultGeometry()
}

//  Sets the distance tolerance for the simplification.
//  All vertices in the simplified geometry will be within this
//  distance of the original geometry.
//  The tolerance value must be non-negative.  A tolerance value
//  of zero is effectively a no-op.
func (t *TopologyPreservingSimplifier) setDistanceTolerance(distanceTolerance float64) {
	t.lineSimplifier.distanceTolerance = distanceTolerance
}

func (t *TopologyPreservingSimplifier) getResultGeometry() matrix.Steric {
	// empty input produces an empty result
	if t.InputGeom.IsEmpty() {
		return t.InputGeom
	}
	t.InputGeom.Filter(&LineStringMapBuilderFilter{t})

	t.lineSimplifier.Simplify(t.linestrings)

	tr := &LineStringTransformer{linestrings: t.linestrings}
	result, _ := tr.Transform(t.InputGeom)
	return result
}

// LineStringMapBuilderFilter A filter to add linear geometries to the linestring map
//  with the appropriate minimum size constraint.
//  For all other linestrings, the minimum size is 2 points.
type LineStringMapBuilderFilter struct {
	tps *TopologyPreservingSimplifier
}

// Filter linear geometries.
func (l *LineStringMapBuilderFilter) Filter(geom matrix.Matrix) {

}

// IsChanged  Returns the true when need change.
func (l *LineStringMapBuilderFilter) IsChanged() bool {
	return false
}

// Matrixes  Returns the gathered Matrixes.
func (l *LineStringMapBuilderFilter) Matrixes() []matrix.Matrix {
	return nil
}

// FilterMatrixes Performs an operation with the provided .
func (l *LineStringMapBuilderFilter) FilterMatrixes(matrixes []matrix.Matrix) {
	if line, ok := l.tps.InputGeom.(matrix.LineMatrix); ok {
		// skip empty geometries
		if line.IsEmpty() {
			return
		}

		minSize := 2
		if line.IsClosed() {
			minSize = 4
		}
		taggedLine := &TaggedLineString{ParentLine: line, MinimumSize: minSize}
		taggedLine.initTaggedLine()
		l.tps.linestrings = append(l.tps.linestrings, taggedLine)
	}
}

// Clear  clear Matrixes.
func (l *LineStringMapBuilderFilter) Clear() {

}

// compile time checks
var (
	_ matrix.Filter = &LineStringMapBuilderFilter{}
)
