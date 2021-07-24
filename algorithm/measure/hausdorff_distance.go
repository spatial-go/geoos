package measure

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// HausdorffDistance An algorithm for computing a distance metric
// which is an approximation to the Hausdorff Distance
// based on a discretization of the input  Geometry.
// The algorithm computes the Hausdorff distance restricted to discrete points
// for one of the geometries.
// The points can be either the vertices of the geometries (the default),
// or the geometries with line segments densified by a given fraction.
// Also determines two points of the Geometries which are separated by the computed distance.
// This algorithm is an approximation to the standard Hausdorff distance.
// Specifically,
//    for all geometries a, b:    DHD(a, b) >= HD(a, b)
// The approximation can be made as close as needed by densifying the input geometries.
// In the limit, this value will approach the true Hausdorff distance:
// <pre>
//    DHD(A, B, densifyFactor) ->; HD(A, B) as densifyFactor ->; 0.0
// </pre>
// The default approximation is exact or close enough for a large subset of useful cases.
// Examples of these are:
// <ul>
// <li>computing distance between Linestrings that are roughly parallel to each other,
// and roughly equal in length.  This occurs in matching linear networks.
// <li>Testing similarity of geometries.
// </ul>
// An example where the default approximation is not close is:
// <pre>
//   A = LINESTRING (0 0, 100 0, 10 100, 10 100)
//   B = LINESTRING (0 100, 0 10, 80 10)
//
//   DHD(A, B) = 22.360679774997898
//   HD(A, B) ~= 47.8
// </pre>
type HausdorffDistance struct {
	g0, g1 matrix.Steric
	ptDist *PointPairDistance
	//Value of 0.0 indicates that no densification should take place
	densifyFrac float64
}

// Distance ...
func (h *HausdorffDistance) Distance(g0, g1 matrix.Steric) float64 {
	dist := &HausdorffDistance{g0: g0, g1: g1, ptDist: &PointPairDistance{}}
	return dist.distance()
}

// DistanceDensifyFrac ...
func (h *HausdorffDistance) DistanceDensifyFrac(g0, g1 matrix.Steric, densifyFrac float64) (float64, error) {
	if densifyFrac > 1.0 || densifyFrac <= 0.0 {
		return 0, algoerr.ErrWrongFractionRange
	}
	dist := &HausdorffDistance{g0: g0, g1: g1, densifyFrac: densifyFrac, ptDist: &PointPairDistance{}}
	return dist.distance(), nil
}

func (h *HausdorffDistance) distance() float64 {
	h.compute()
	return h.ptDist.Distance
}

func (h *HausdorffDistance) orientedDistance() float64 {
	h.computeOrientedDistance(h.g0, h.g1, h.ptDist)
	return h.ptDist.Distance
}

func (h *HausdorffDistance) compute() {
	h.computeOrientedDistance(h.g0, h.g1, h.ptDist)
	h.computeOrientedDistance(h.g1, h.g0, h.ptDist)
}

func (h *HausdorffDistance) computeOrientedDistance(g0, g1 matrix.Steric, ptDist *PointPairDistance) {
	distFilter := &MaxPointDistanceFilter{geom: g1}
	distFilter.MaxPtDist, distFilter.MinPtDist = &PointPairDistance{}, &PointPairDistance{}
	distFilter.euclideanDist = &DistanceToPoint{}
	distFilter.apply(matrix.TransMatrixs(g0))
	ptDist.setMaximum(distFilter.MaxPtDist)

	if h.densifyFrac > 0 {
		fracFilter := &MaxDensifiedByFractionDistanceFilter{geom: g1}
		fracFilter.MaxPtDist, fracFilter.MinPtDist = &PointPairDistance{}, &PointPairDistance{}
		fracFilter.numSubSegs = int((1.0 / h.densifyFrac))
		fracFilter.apply(matrix.TransMatrixs(g0))
		ptDist.setMaximum(fracFilter.MaxPtDist)

	}
}

// MaxPointDistanceFilter ...
type MaxPointDistanceFilter struct {
	MaxPtDist, MinPtDist *PointPairDistance
	euclideanDist        *DistanceToPoint
	geom                 matrix.Steric
}

func (m *MaxPointDistanceFilter) filter(pt matrix.Matrix) {
	m.MinPtDist.IsNil = true
	(&DistanceToPoint{}).computeDistance(m.geom, pt, m.MinPtDist)
	m.MaxPtDist.setMaximum(m.MinPtDist)
}

func (m *MaxPointDistanceFilter) apply(pts []matrix.Matrix) {
	for i := 0; i < len(pts); i++ {
		m.filter(pts[i])
	}
}

// MaxDensifiedByFractionDistanceFilter ...
type MaxDensifiedByFractionDistanceFilter struct {
	MaxPtDist, MinPtDist *PointPairDistance
	geom                 matrix.Steric
	numSubSegs           int

	// numSubSegs = (int) Math.rint(1.0/fraction);
}

func (m *MaxDensifiedByFractionDistanceFilter) apply(pts []matrix.Matrix) {

	if len(pts) == 0 {
		return
	}
	for i := 0; i < len(pts); i++ {
		m.filter(pts, i)
	}
}

func (m *MaxDensifiedByFractionDistanceFilter) filter(pts []matrix.Matrix, index int) {
	/**
	 * This logic also handles skipping Point geometries
	 */
	if index == 0 {
		return
	}

	p0 := pts[index-1]
	p1 := pts[index]

	delx := (p1[0] - p0[0]) / float64(m.numSubSegs)
	dely := (p1[1] - p0[1]) / float64(m.numSubSegs)

	for i := 0; i < m.numSubSegs; i++ {
		x := p0[0] + float64(i)*delx
		y := p0[1] + float64(i)*dely
		pt := matrix.Matrix{x, y}
		m.MinPtDist.IsNil = true
		(&DistanceToPoint{}).computeDistance(m.geom, pt, m.MinPtDist)
		m.MaxPtDist.setMaximum(m.MinPtDist)
	}

}
