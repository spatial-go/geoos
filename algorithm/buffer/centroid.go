package buffer

import (
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Centroid Computes the centroid of a  Geometry of any dimension.
// For collections the centroid is computed for the collection of
// non-empty elements of highest dimension.
// The centroid of an empty geometry is nil.
type Centroid struct {
	AreaBasePt    matrix.Matrix // the point all triangles are based at
	TriangleCent3 matrix.Matrix // temporary variable to hold centroid of triangle
	Areasum2      float64       /* Partial area sum */
	Cg3           matrix.Matrix // partial centroid sum

	// data for linear centroid computation, if needed
	LineCentSum matrix.Matrix
	TotalLength float64

	PtCount   int
	PtCentSum matrix.Matrix
}

// GetCentroid Gets the computed centroid.
// returns he computed centroid, or nil if the input is empty
func (c *Centroid) GetCentroid() matrix.Matrix {
	/**
	 * The centroid is computed from the highest dimension components present in the input.
	 * I.e. areas dominate lineal geometry, which dominates points.
	 * Degenerate geometry are computed using their effective dimension
	 * (e.g. areas may degenerate to lines or points)
	 */
	cent := matrix.Matrix{0, 0}
	if math.Abs(c.Areasum2) > 0.0 {
		/**
		* Input contains areal geometry
		 */
		cent[0] = c.Cg3[0] / 3 / c.Areasum2
		cent[1] = c.Cg3[1] / 3 / c.Areasum2
	} else if c.TotalLength > 0.0 {
		/**
		* Input contains lineal geometry
		 */
		cent[0] = c.LineCentSum[0] / c.TotalLength
		cent[1] = c.LineCentSum[1] / c.TotalLength
	} else if c.PtCount > 0 {
		/**
		* Input contains puntal geometry only
		 */
		cent[0] = c.PtCentSum[0] / float64(c.PtCount)
		cent[1] = c.PtCentSum[1] / float64(c.PtCount)
	} else {
		return nil
	}
	return cent
}
