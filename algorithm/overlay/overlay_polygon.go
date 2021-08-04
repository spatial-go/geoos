package overlay

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// PolygonOverlay  Computes the overlay of two geometries,either or both of which may be nil.
type PolygonOverlay struct {
	*PointOverlay
	subjectPlane, clippingPlane *algorithm.Plane
}

// Union  Computes the Union of two geometries,either or both of which may be nil.
func (p *PolygonOverlay) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if _, ok := p.subject.(matrix.PolygonMatrix); ok {
		if _, ok := p.clipping.(matrix.PolygonMatrix); ok {
			cpo := &ComputeMergeOverlay{p}

			cpo.prepare()
			_, exitingPoints := cpo.Weiler()
			result := ToPolygonMatrix(cpo.ComputePolygon(exitingPoints, cpo))
			return result, nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *PolygonOverlay) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	if _, ok := p.subject.(matrix.PolygonMatrix); ok {
		if _, ok := p.clipping.(matrix.PolygonMatrix); ok {
			cpo := &ComputeClipOverlay{p}

			cpo.prepare()
			_, exitingPoints := cpo.Weiler()
			result := ToPolygonMatrix(cpo.ComputePolygon(exitingPoints, cpo))
			return result, nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *PolygonOverlay) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}
	if _, ok := p.subject.(matrix.PolygonMatrix); ok {
		if _, ok := p.clipping.(matrix.PolygonMatrix); ok {
			cpo := &ComputeMainOverlay{p}

			cpo.prepare()
			_, exitingPoints := cpo.Weiler()
			result := ToPolygonMatrix(cpo.ComputePolygon(exitingPoints, cpo))
			return result, nil
		}
	}
	return nil, algoerr.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *PolygonOverlay) DifferenceReverse() (matrix.Steric, error) {
	newPoly := &PolygonOverlay{PointOverlay: &PointOverlay{subject: p.clipping, clipping: p.subject}}
	return newPoly.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *PolygonOverlay) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil {
		result = append(result, res)
	}
	if res, err := p.DifferenceReverse(); err == nil {
		result = append(result, res)
	}
	return result, nil
}

// prepare prepare two polygonal geometries.
func (p *PolygonOverlay) prepare() {
	p.subjectPlane = &algorithm.Plane{}
	for _, v2 := range p.subject.(matrix.PolygonMatrix) {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				p.subjectPlane.AddPoint(&algorithm.Vertex{Matrix: matrix.Matrix(v1)})
			}
		}
		p.subjectPlane.CloseRing()
		p.subjectPlane.Rank = calc.MAIN
	}
	p.clippingPlane = &algorithm.Plane{}
	for _, v2 := range p.clipping.(matrix.PolygonMatrix) {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				p.clippingPlane.AddPoint(&algorithm.Vertex{Matrix: matrix.Matrix(v1)})
			}
		}
		p.clippingPlane.CloseRing()
		p.clippingPlane.Rank = calc.CUT
	}
}

// Weiler Weiler overlay.
func (p *PolygonOverlay) Weiler() (enteringPoints, exitingPoints []algorithm.Vertex) {
	for _, v := range p.subjectPlane.Lines {
		for _, vClip := range p.clippingPlane.Lines {

			mark, ips :=
				relate.Intersection(v.Start.Matrix, v.End.Matrix, vClip.Start.Matrix, vClip.End.Matrix)
			for _, ip := range ips {
				ipVer := &algorithm.Vertex{}
				ipVer.Matrix = ip.Matrix
				ipVer.IsIntersectionPoint = ip.IsIntersectionPoint
				ipVer.IsEntering = ip.IsEntering
				if mark {
					if ipVer.IsEntering {
						enteringPoints = append(enteringPoints, *ipVer)
					} else {
						exitingPoints = append(exitingPoints, *ipVer)
					}
					AddPointToVertexSlice(p.subjectPlane.Rings, v.Start, v.End, ipVer)
					AddPointToVertexSlice(p.clippingPlane.Rings, vClip.Start, vClip.End, ipVer)
				}
			}
		}
	}
	return
}

// ComputePolygon compute overlay.
func (p *PolygonOverlay) ComputePolygon(exitingPoints []algorithm.Vertex, cpo ComputePolyOverlay) *algorithm.Plane {
	var pol *algorithm.Plane = &algorithm.Plane{}
	for _, iterPoints := range exitingPoints {
		if iterPoints.IsChecked {
			continue
		}
		edge := &algorithm.Edge{}
		pol.Edge = edge
		pol.Rings = append(pol.Rings, edge)

		start := &iterPoints
		next := &algorithm.Vertex{Matrix: matrix.Matrix{start.X(), start.Y()}}
		start.IsChecked = true

		for {
			next = cpo.Next(pol, next)
			if where, err := SliceContains(exitingPoints, next); err == nil {
				exitingPoints[where].IsChecked = true
			}
			if next.X() == start.X() && next.Y() == start.Y() {
				pol.CloseRing()
				break
			}
		}
	}

	return pol
}

// ToPolygonMatrix ...
func ToPolygonMatrix(poly *algorithm.Plane) matrix.PolygonMatrix {
	var result matrix.PolygonMatrix
	for _, v2 := range poly.Rings {
		var edge matrix.LineMatrix
		for _, v1 := range v2.Vertexs {
			edge = append(edge, v1.Matrix)
		}
		if !matrix.Matrix(edge[len(edge)-1]).Equals(matrix.Matrix(edge[0])) {
			edge = append(edge, edge[0])
		}
		result = append(result, edge)
	}

	return result
}
