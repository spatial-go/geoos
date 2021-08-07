package overlay

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Intersection  Computes the Intersection of two geometries,either or both of which may be null.
func Intersection(m0, m1 matrix.Steric) (matrix.Steric, error) {
	switch m := m0.(type) {
	case matrix.Matrix:
		over := &PointOverlay{subject: m, clipping: m1}
		return over.Intersection()
	case matrix.LineMatrix:
		var err error
		newLine := &LineOverlay{PointOverlay: &PointOverlay{subject: m, clipping: m1}}
		if result, err := newLine.Intersection(); err == nil {
			if len(result.(matrix.Collection)) == 1 {
				return result.(matrix.Collection)[0], nil
			}
			return result, nil
		}
		return nil, err
	case matrix.PolygonMatrix:
		polyOver := &PolygonOverlay{PointOverlay: &PointOverlay{subject: m, clipping: m1}}
		return polyOver.Intersection()
	default:
		return nil, algoerr.ErrNotSupportCollection

	}
}
