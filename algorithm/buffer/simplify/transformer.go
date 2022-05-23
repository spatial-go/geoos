package simplify

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/buffer"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Trans  A framework for processes which transform an input  Geometry into
//  an output  Geometry, possibly changing its structure and type(s).
type Trans interface {
	// Transform ...
	Transform(inputGeom matrix.Steric) (matrix.Steric, error)

	transformCoordinates(pts []matrix.Matrix, parent matrix.Steric) []matrix.Matrix
	transformPoint(geom matrix.Matrix, parent matrix.Steric) matrix.Steric
	transformLine(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric
	transformRing(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric
	transformPolygon(geom matrix.PolygonMatrix, parent matrix.Steric) matrix.Steric
	transformCollection(geom matrix.Collection, parent matrix.Steric) matrix.Steric
}

// Transformer  A framework for processes which transform an input  Geometry into
//  an output  Geometry, possibly changing its structure and type(s).
type Transformer struct {
	InputGeom                        matrix.Steric
	pruneEmptyGeometry, preserveType bool
}

// Transform ...
func (t *Transformer) Transform(inputGeom matrix.Steric) (matrix.Steric, error) {
	t.InputGeom = inputGeom
	switch m := inputGeom.(type) {
	case matrix.Matrix:
		return t.transformPoint(m, nil), nil
	case matrix.LineMatrix:
		return t.transformLine(m, nil), nil
	case matrix.PolygonMatrix:
		return t.transformPolygon(m, nil), nil
	case matrix.Collection:
		return t.transformCollection(m, nil), nil
	default:
		return nil, algorithm.ErrUnknownType(m)
	}
}

func (t *Transformer) transformCoordinates(pts []matrix.Matrix, parent matrix.Steric) []matrix.Matrix {
	return pts
}

func (t *Transformer) transformPoint(geom matrix.Matrix, parent matrix.Steric) matrix.Steric {
	pts := t.transformCoordinates(matrix.TransMatrixes(geom), parent)
	return pts[0]
}

func (t *Transformer) transformLine(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	pts := t.transformCoordinates(matrix.TransMatrixes(geom), parent)
	ml := matrix.LineMatrix{}
	for _, v := range pts {
		ml = append(ml, v)
	}
	return ml
}

// Transforms a LinearRing.
// The transformation of a LinearRing may result in a coordinate sequence
// which does not form a structurally valid ring (i.e. a degenerate ring of 3 or fewer points).
// In this case a LineString is returned.
func (t *Transformer) transformRing(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	pts := t.transformCoordinates(matrix.TransMatrixes(geom), geom)
	if pts == nil || len(pts) <= 0 {
		return nil
	}
	size := len(pts)
	// ensure a valid LinearRing
	if size > 0 && size < 4 && !t.preserveType {
		ml := matrix.LineMatrix{}
		for _, v := range pts {
			ml = append(ml, v)
		}
		return ml
	}
	return geom
}

func (t *Transformer) transformPolygon(geom matrix.PolygonMatrix, parent matrix.Steric) matrix.Steric {
	result := matrix.PolygonMatrix{}
	for i, v := range geom {
		line := t.transformRing(v, geom)
		if line == nil || line.IsEmpty() {
			if i == 0 {
				return nil
			}
			continue
		}
		result = append(result, line.(matrix.LineMatrix))
	}

	return result
}

func (t *Transformer) transformCollection(geom matrix.Collection, parent matrix.Steric) matrix.Steric {

	transGeoms := matrix.Collection{}
	for _, v := range geom {
		transformGeom, _ := t.Transform(v)
		if transformGeom == nil {
			continue
		}
		if t.pruneEmptyGeometry && transformGeom.IsEmpty() {
			continue
		}
		transGeoms = append(transGeoms, transformGeom)
	}
	return transGeoms

}

// LineStringTransformer Transformer  A framework for processes which transform an input linestring into
//  an output Geometry, possibly changing its structure and type(s).
type LineStringTransformer struct {
	Transformer
	linestrings []*TaggedLineString
}

// Transform ...
func (l *LineStringTransformer) Transform(inputGeom matrix.Steric) (matrix.Steric, error) {
	l.InputGeom = inputGeom
	switch m := inputGeom.(type) {
	case matrix.Matrix:
		return l.transformPoint(m, inputGeom), nil
	case matrix.LineMatrix:
		return l.transformLine(m, inputGeom), nil
	case matrix.PolygonMatrix:
		return l.transformPolygon(m, inputGeom), nil
	case matrix.Collection:
		return l.transformCollection(m, inputGeom), nil
	default:
		return nil, algorithm.ErrUnknownType(m)
	}
}

func (l *LineStringTransformer) transformLine(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	pts := l.transformCoordinates(matrix.TransMatrixes(geom), parent)
	ml := matrix.LineMatrix{}
	for _, v := range pts {
		ml = append(ml, v)
	}
	return ml
}

func (l *LineStringTransformer) findLineString(parent matrix.LineMatrix) (*TaggedLineString, error) {
	for _, v := range l.linestrings {
		if v.ParentLine.Equals(parent) {
			return v, nil
		}
	}
	return nil, algorithm.ErrNotInSlice
}

func (l *LineStringTransformer) transformCoordinates(pts []matrix.Matrix, parent matrix.Steric) []matrix.Matrix {
	if len(pts) <= 0 {
		return nil
	}
	// for linear components (including rings), simplify the linestring
	if pl, ok := parent.(matrix.LineMatrix); ok {
		if taggedLine, err := l.findLineString(pl); err == nil {
			return taggedLine.GetResultMatrixes()
		}
		return pts
	}
	// for anything else (e.g. points) just copy the coordinates
	return pts
}

// DPTransformer ...
type DPTransformer struct {
	*Transformer
	isEnsureValidTopology bool
	distanceTolerance     float64
}

// Transform ...
func (d *DPTransformer) Transform(inputGeom matrix.Steric) (matrix.Steric, error) {
	d.InputGeom = inputGeom
	switch m := inputGeom.(type) {
	case matrix.Matrix:
		return d.transformPoint(m, nil), nil
	case matrix.LineMatrix:
		return d.transformLine(m, nil), nil
	case matrix.PolygonMatrix:
		return d.transformPolygon(m, nil), nil
	case matrix.Collection:
		return d.transformCollection(m, nil), nil
	default:
		return nil, algorithm.ErrUnknownType(m)
	}
}

func (d *DPTransformer) transformLine(line matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	var newPts matrix.LineMatrix
	if len(line) == 0 {
		newPts = matrix.LineMatrix{matrix.Matrix{0, 0}}
	} else {
		pts := (&DouglasPeuckerLineSimplifier{pts: matrix.TransMatrixes(line), distanceTolerance: d.distanceTolerance}).
			Simplify()
		for _, v := range pts {
			newPts = append(newPts, v)
		}
	}
	return newPts
}

// Simplifies a polygon, fixing it if required.
func (d *DPTransformer) transformPolygon(geom matrix.PolygonMatrix, parent matrix.Steric) matrix.Steric {
	// empty geometries are simply removed
	if geom.IsEmpty() {
		return nil
	}

	rawGeom := matrix.PolygonMatrix{}
	for i, v := range geom {
		line := d.transformRing(v, geom)
		if line.IsEmpty() {
			if i == 0 {
				return nil
			}
			continue
		}
		rawGeom = append(rawGeom, line.(matrix.LineMatrix))
	}

	// don't try and correct if the parent is going to do this
	if _, ok := parent.(matrix.MultiPolygonMatrix); ok {
		return rawGeom
	}
	return d.createValidArea(rawGeom)
}

// Simplifies a LinearRing.  If the simplification results in a degenerate ring, remove the component.
func (d *DPTransformer) transformRing(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	if _, ok := parent.(matrix.PolygonMatrix); ok {
		simpResult := d.transformLine(geom, parent)
		return simpResult
	}
	return geom
}

//  Creates a valid area geometry from one that possibly has bad topology (i.e. self-intersections).
//  Since buffer can handle invalid topology, but always returns valid geometry, constructing a 0-width buffer "corrects" the  topology.
func (d *DPTransformer) createValidArea(rawAreaGeom matrix.Steric) matrix.Steric {
	isValidArea := rawAreaGeom.Dimensions() == 2
	// if geometry is invalid then make it valid
	if d.isEnsureValidTopology && !isValidArea {
		return buffer.Buffer(rawAreaGeom, 0.0, 8)
	}
	return rawAreaGeom
}
