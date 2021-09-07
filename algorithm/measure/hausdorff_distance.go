// Package measure Define spatial measurement function.
package measure

import (
	"github.com/spatial-go/geoos/algorithm"
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
		return 0, algorithm.ErrWrongFractionRange
	}
	dist := &HausdorffDistance{g0: g0, g1: g1, densifyFrac: densifyFrac, ptDist: &PointPairDistance{}}
	return dist.distance(), nil
}

func (h *HausdorffDistance) distance() float64 {
	h.compute()
	return h.ptDist.Distance
}

// OrientedDistance ...
func (h *HausdorffDistance) OrientedDistance() float64 {
	h.computeOrientedDistance(h.g0, h.g1, h.ptDist)
	return h.ptDist.Distance
}

func (h *HausdorffDistance) compute() {
	h.computeOrientedDistance(h.g0, h.g1, h.ptDist)
	h.computeOrientedDistance(h.g1, h.g0, h.ptDist)
}

func (h *HausdorffDistance) computeOrientedDistance(g0, g1 matrix.Steric, ptDist *PointPairDistance) {
	distFilter := &MaxPointDistanceFilter{geom: g1, MaxPtDist: ptDist, MinPtDist: &PointPairDistance{}, euclideanDist: &DistanceToPoint{}}
	g0.Filter(distFilter)

	if h.densifyFrac > 0 {
		fracFilter := &MaxDensifiedByFractionDistanceFilter{geom: g1, MaxPtDist: ptDist, MinPtDist: &PointPairDistance{}, numSubSegs: int((1.0 / h.densifyFrac))}
		g0.Filter(fracFilter)

	}
}
