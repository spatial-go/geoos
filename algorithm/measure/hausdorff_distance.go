package measure

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// HausdorffDistance An algorithm for computing a distance metric
// which is an approximation to the Hausdorff Distance
type HausdorffDistance struct {
	g0, g1      matrix.Steric
	ptDist      *PointPairDistance
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
	distFilter.apply(matrix.TransMatrixes(g0))
	ptDist.setMaximum(distFilter.MaxPtDist)

	if h.densifyFrac > 0 {
		fracFilter := &MaxDensifiedByFractionDistanceFilter{geom: g1}
		fracFilter.MaxPtDist, fracFilter.MinPtDist = &PointPairDistance{}, &PointPairDistance{}
		fracFilter.numSubSegs = int((1.0 / h.densifyFrac))
		fracFilter.apply(matrix.TransMatrixes(g0))
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
	//This logic also handles skipping Point geometries
	if index == 0 {
		return
	}

	p0 := pts[index-1]
	p1 := pts[index]

	delX := (p1[0] - p0[0]) / float64(m.numSubSegs)
	delY := (p1[1] - p0[1]) / float64(m.numSubSegs)

	for i := 0; i < m.numSubSegs; i++ {
		x := p0[0] + float64(i)*delX
		y := p0[1] + float64(i)*delY
		pt := matrix.Matrix{x, y}
		m.MinPtDist.IsNil = true
		(&DistanceToPoint{}).computeDistance(m.geom, pt, m.MinPtDist)
		m.MaxPtDist.setMaximum(m.MinPtDist)
	}

}
