package measure

import (
	"fmt"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// PointPairDistance Contains a pair of points and the distance between them.
// Provides methods to update with a new point pair with
// either maximum or minimum distance.
type PointPairDistance struct {
	Pt       [2]matrix.Matrix
	Distance float64
	IsNil    bool
}

func (p *PointPairDistance) setMaximum(ppd *PointPairDistance) {
	if p.IsNil {
		p.Pt = [2]matrix.Matrix{ppd.Pt[0], ppd.Pt[1]}
		p.Distance = PlanarDistance(ppd.Pt[0], ppd.Pt[1])
		p.IsNil = false
		return
	}
	dist := PlanarDistance(ppd.Pt[0], ppd.Pt[1])
	if dist > p.Distance {
		p.Pt = [2]matrix.Matrix{ppd.Pt[0], ppd.Pt[1]}
		p.Distance = dist
		p.IsNil = false
	}
}

func (p *PointPairDistance) setMinimum(ppd *PointPairDistance) {
	if p.IsNil {
		p.Pt = [2]matrix.Matrix{ppd.Pt[0], ppd.Pt[1]}
		p.Distance = PlanarDistance(ppd.Pt[0], ppd.Pt[1])
		p.IsNil = false
		return
	}
	dist := PlanarDistance(ppd.Pt[0], ppd.Pt[1])
	if dist < p.Distance {
		p.Pt = [2]matrix.Matrix{ppd.Pt[0], ppd.Pt[1]}
		p.Distance = dist
		p.IsNil = false
	}
}

// ToString ...
func (p *PointPairDistance) ToString() string {
	return fmt.Sprintf("PointPairDistance %v %v", p.Pt[0], p.Pt[1])
}

// DistanceToPoint Computes the Euclidean distance (L2 metric) from a point to a point.
// Also computes two points on the geometry which are separated by the distance found.
type DistanceToPoint struct {
}

func (d *DistanceToPoint) computeDistance(geom matrix.Steric, pt matrix.Matrix, ptDist *PointPairDistance) {
	switch m := geom.(type) {
	case matrix.LineMatrix:
		d.computeDistanceLine(m, pt, ptDist)
	case matrix.PolygonMatrix:
		d.computeDistancePolygon(m, pt, ptDist)
	case matrix.Collection:
		for _, v := range m {
			d.computeDistance(v, pt, ptDist)
		}
	case matrix.Matrix:
		ptDist.setMinimum(&PointPairDistance{Pt: [2]matrix.Matrix{m, pt}})
	}
}

func (d *DistanceToPoint) computeDistanceLine(line matrix.LineMatrix, pt matrix.Matrix, ptDist *PointPairDistance) {
	for i := 0; i < len(line)-1; i++ {
		// this is somewhat inefficient - could do better
		closestPt := ClosestPoint(pt, line[i], line[i+1])
		ptDist.setMinimum(&PointPairDistance{Pt: [2]matrix.Matrix{closestPt, pt}})
	}
}

func (d *DistanceToPoint) computeDistancePolygon(poly matrix.PolygonMatrix, pt matrix.Matrix, ptDist *PointPairDistance) {
	for i := 0; i < len(poly); i++ {
		d.computeDistanceLine(poly[i], pt, ptDist)
	}
}

// ClosestPoint Computes the closest point on this line segment to another point.
func ClosestPoint(p, a, b matrix.Matrix) matrix.Matrix {
	factor := ProjectionFactor(p, a, b)
	if factor > 0 && factor < 1 {
		return Project(p, a, b)
	}
	dist0 := PlanarDistance(a, p)
	dist1 := PlanarDistance(b, p)
	if dist0 < dist1 {
		return a
	}
	return b
}

// Project Compute the projection of a point onto the line determined
func Project(p, a, b matrix.Matrix) matrix.Matrix {
	if p.Equals(a) || p.Equals(b) {
		return p
	}

	r := ProjectionFactor(p, a, b)
	pp := matrix.Matrix{0, 0}
	pp[0] = a[0] + r*(b[0]-a[0])
	pp[1] = a[1] + r*(b[1]-a[1])
	return pp
}

// ProjectionFactor Computes the Projection Factor for the projection of the point p
func ProjectionFactor(p, a, b matrix.Matrix) float64 {
	if p.Equals(a) {
		return 0.0
	}
	if p.Equals(b) {
		return 1.0
	}
	// Otherwise, use comp.graphics.algorithms Frequently Asked Questions method
	//     	      AC dot AB
	// 	r = ---------
	// 		  ||AB||^2
	//  r has the following meaning:
	//  r=0 P = A
	//  r=1 P = B
	//  r<0 P is on the backward extension of AB
	//  r>1 P is on the forward extension of AB
	//  0<r<1 P is interior to AB
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	lenD := dx*dx + dy*dy

	// handle zero-length segments
	if lenD <= 0.0 {
		return 0.0
	}
	return ((p[0]-a[0])*dx + (p[1]-a[1])*dy) / lenD
}

// MaxPointDistanceFilter ...
type MaxPointDistanceFilter struct {
	MaxPtDist, MinPtDist *PointPairDistance
	euclideanDist        *DistanceToPoint
	geom                 matrix.Steric
}

// IsChanged  Returns the true when need change.
func (m *MaxPointDistanceFilter) IsChanged() bool {
	return false
}

// Filter  Performs an operation with the provided .
func (m *MaxPointDistanceFilter) Filter(pt matrix.Matrix) {
	m.MinPtDist.IsNil = true
	(&DistanceToPoint{}).computeDistance(m.geom, pt, m.MinPtDist)
	m.MaxPtDist.setMaximum(m.MinPtDist)
}

// FilterMatrixes  Performs an operation with the provided .
func (m *MaxPointDistanceFilter) FilterMatrixes(pts []matrix.Matrix) {
	for i := 0; i < len(pts); i++ {
		m.Filter(pts[i])
	}
}

// Matrixes  Returns the gathered Matrixes.
func (m *MaxPointDistanceFilter) Matrixes() []matrix.Matrix {
	return matrix.TransMatrixes(m.geom)
}

// Clear  clear Matrixes.
func (m *MaxPointDistanceFilter) Clear() {

}

// MaxDensifiedByFractionDistanceFilter ...
type MaxDensifiedByFractionDistanceFilter struct {
	MaxPtDist, MinPtDist *PointPairDistance
	geom                 matrix.Steric
	numSubSegs           int
}

// IsChanged  Returns the true when need change.
func (m *MaxDensifiedByFractionDistanceFilter) IsChanged() bool {
	return false
}

// Matrixes  Returns the gathered Matrixes.
func (m *MaxDensifiedByFractionDistanceFilter) Matrixes() []matrix.Matrix {
	return matrix.TransMatrixes(m.geom)
}

// FilterMatrixes  Performs an operation with the provided .
func (m *MaxDensifiedByFractionDistanceFilter) FilterMatrixes(pts []matrix.Matrix) {

	if len(pts) == 0 {
		return
	}
	for i := 0; i < len(pts); i++ {
		m.filterTwo(pts, i)
	}
}

// filterTwo  Performs an operation with the provided .
func (m *MaxDensifiedByFractionDistanceFilter) filterTwo(pts []matrix.Matrix, index int) {
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

// Clear  clear Matrixes.
func (m *MaxDensifiedByFractionDistanceFilter) Clear() {

}

// Filter  Performs an operation with the provided .
func (m *MaxDensifiedByFractionDistanceFilter) Filter(pt matrix.Matrix) {
}

// compile time checks
var (
	_ matrix.Filter = &MaxDensifiedByFractionDistanceFilter{}
	_ matrix.Filter = &MaxPointDistanceFilter{}
)
