package snap

import "github.com/spatial-go/geoos/algorithm/matrix"

// Snap two geometries together with a given tolerance.
func Snap(g0, g1 matrix.Steric, snapTolerance float64) (snapGeom matrix.Collection) {
	snapper0 := &Snapper{g0, g1, snapTolerance}
	snapGeom = append(snapGeom, snapper0.SnapTo(g1, snapTolerance))

	/**
	 * Snap the second geometry to the snapped first geometry
	 * (this strategy minimizes the number of possible different points in the result)
	 */
	snapper1 := &Snapper{g1, g0, snapTolerance}
	snapGeom = append(snapGeom, snapper1.SnapTo(snapGeom[0], snapTolerance))

	return snapGeom
}

// Snapper Snaps the vertices and segments of a Geometry
// to another Geometry's vertices.
type Snapper struct {
	srcGeom, snapGeom matrix.Steric
	snapTolerance     float64
}

// SnapTo Snaps the vertices in the component {@link LineString}s
// of the source geometry
// to the vertices of the given snap geometry.
func (s *Snapper) SnapTo(snapGeom matrix.Steric, snapTolerance float64) matrix.Steric {
	return s.transform(snapGeom)
}

func (s *Snapper) transform(inputGeom matrix.Steric) matrix.Steric {

	switch m := inputGeom.(type) {
	case matrix.Matrix:
		return s.transformPoint(m, nil)
	case matrix.LineMatrix:
		return s.transformLine(m, nil)
	case matrix.PolygonMatrix:
		return s.transformPolygon(m, nil)
	case matrix.Collection:
		return s.transformCollection(m, nil)

	}
	return nil
}

func (s *Snapper) transformMatrix(srcPt matrix.Matrix, parent matrix.Steric) matrix.Steric {
	newPts := s.snapLine(matrix.LineMatrix{srcPt}, s.snapGeom)
	return newPts
}
func (s *Snapper) transformMatrixs(srcPts matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	newPts := s.snapLine(srcPts, s.snapGeom)
	return newPts
}

func (s *Snapper) snapLine(srcPts matrix.LineMatrix, snapPts matrix.Steric) matrix.Steric {
	snapper := &LineSnapper{srcPts, snapPts, s.snapTolerance}
	return snapper.snapTo()
}

func (s *Snapper) transformPoint(pt matrix.Matrix, parent matrix.Steric) matrix.Steric {
	return s.transformMatrix(pt, pt)
}

// transformLine Transforms a LinearRing.
// The transformation of a LinearRing may result in a coordinate sequence
// which does not form a structurally valid ring (i.e. a degenerate ring of 3 or fewer points).
// In this case a LineString is returned.
func (s *Snapper) transformLine(geom matrix.LineMatrix, parent matrix.Steric) matrix.Steric {
	seq := s.transformMatrixs(geom, geom)
	return seq
}

func (s *Snapper) transformPolygon(geom matrix.PolygonMatrix, parent matrix.Steric) matrix.Steric {
	poly := matrix.PolygonMatrix{}

	for _, v := range geom {
		poly = append(poly, s.transformLine(v, geom).(matrix.LineMatrix))
	}

	return poly
}

func (s *Snapper) transformCollection(geom matrix.Collection, parent matrix.Steric) matrix.Collection {
	transGeoms := matrix.Collection{}
	for _, v := range geom {
		transformGeom := s.transform(v)
		if transformGeom == nil {
			continue
		}

		transGeoms = append(transGeoms, transformGeom)
	}
	return transGeoms
}
