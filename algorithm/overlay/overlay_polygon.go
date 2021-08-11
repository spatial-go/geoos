package overlay

import (
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// PolygonOverlay  Computes the overlay of two geometries,either or both of which may be nil.
type PolygonOverlay struct {
	*PointOverlay
	subjectPlane, clippingPlane *Plane
}

// Union  Computes the Union of two geometries,either or both of which may be nil.
func (p *PolygonOverlay) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if _, ok := p.Subject.(matrix.PolygonMatrix); ok {
		if _, ok := p.Clipping.(matrix.PolygonMatrix); ok {
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
	var poly matrix.PolygonMatrix
	if p, ok := p.Subject.(matrix.PolygonMatrix); ok {
		poly = p
	} else {
		return nil, algoerr.ErrNotMatchType
	}
	switch c := p.Clipping.(type) {
	case matrix.Matrix:

		inter := envelope.Bound(poly.Bound()).IsIntersects(envelope.Bound(c.Bound()))
		if mark := relate.IM(poly, c, inter).IsContains(); mark {
			return c, nil
		}
		return nil, nil
	case matrix.LineMatrix:
		result := matrix.Collection{}
		for _, ring := range poly {
			for _, il := range IntersectLine(c, ring) {
				if len(il.Ips) > 1 {
					var ipLine matrix.LineMatrix
					for _, v := range il.Ips {
						ipLine = append(ipLine, v.Matrix)
					}
					result = append(result, ipLine)
				} else {
					result = append(result, il.Ips[0].Matrix)
				}
			}
		}
		return LineMerge(result), nil
	case matrix.PolygonMatrix:

		// inter := envelope.Bound(poly.Bound()).IsIntersects(envelope.Bound(c.Bound()))
		// im := relate.IM(poly, c, inter)
		// if mark := im.IsContains(); mark {
		// 	return c, nil
		// }
		// if mark := im.IsWithin(); mark {
		// 	return poly, nil
		// }

		cpo := &ComputeClipOverlay{p}

		cpo.prepare()
		_, exitingPoints := cpo.Weiler()
		result := ToPolygonMatrix(cpo.ComputePolygon(exitingPoints, cpo))
		return result, nil
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
	if poly, ok := p.Subject.(matrix.PolygonMatrix); ok {
		if c, ok := p.Clipping.(matrix.PolygonMatrix); ok {

			inter := envelope.Bound(poly.Bound()).IsIntersects(envelope.Bound(c.Bound()))
			im := relate.IM(poly, c, inter)
			if mark := im.IsWithin(); mark {
				return matrix.PolygonMatrix{}, nil
			}

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
	newPoly := &PolygonOverlay{PointOverlay: &PointOverlay{Subject: p.Clipping, Clipping: p.Subject}}
	return newPoly.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *PolygonOverlay) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil && !res.IsEmpty() {
		result = append(result, res)
	}
	if res, err := p.DifferenceReverse(); err == nil && !res.IsEmpty() {
		result = append(result, res)
	}
	return result, nil
}

// prepare prepare two polygonal geometries.
func (p *PolygonOverlay) prepare() {
	p.subjectPlane = &Plane{}
	for _, v2 := range p.Subject.(matrix.PolygonMatrix) {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				p.subjectPlane.AddPoint(&Vertex{Matrix: matrix.Matrix(v1)})
			}
		}
		p.subjectPlane.CloseRing()
		p.subjectPlane.Rank = calc.OverlayMain
	}
	p.clippingPlane = &Plane{}
	for _, v2 := range p.Clipping.(matrix.PolygonMatrix) {
		for i, v1 := range v2 {
			if i < len(v2)-1 {
				p.clippingPlane.AddPoint(&Vertex{Matrix: matrix.Matrix(v1)})
			}
		}
		p.clippingPlane.CloseRing()
		p.clippingPlane.Rank = calc.OverlayCut
	}
}

// Weiler Weiler overlay.
func (p *PolygonOverlay) Weiler() (enteringPoints, exitingPoints []Vertex) {

	// TODO overlay ...
	for _, v := range p.subjectPlane.Lines {
		for _, vClip := range p.clippingPlane.Lines {

			mark, ips :=
				relate.Intersection(v.Start.Matrix, v.End.Matrix, vClip.Start.Matrix, vClip.End.Matrix)
			for _, ip := range ips {
				if ip.IsCollinear {
					continue
				}
				inV, _ := relate.InLineVertex(ip.Matrix, matrix.LineMatrix{v.Start.Matrix, v.End.Matrix})
				inVClip, _ := relate.InLineVertex(ip.Matrix, matrix.LineMatrix{vClip.Start.Matrix, vClip.End.Matrix})
				if inV && inVClip {
					continue
				}
				ipVer := &Vertex{}
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

	filt := &UniqueVertexFilter{}
	for _, v := range enteringPoints {
		filt.Filter(v)
	}
	enteringPoints = filt.Ips
	filt = &UniqueVertexFilter{}
	for _, v := range exitingPoints {
		filt.Filter(v)
	}
	exitingPoints = filt.Ips

	return
}

// UniqueVertexFilter  A Filter that extracts a unique array.
type UniqueVertexFilter struct {
	Ips []Vertex
}

// Filter Performs an operation with the provided .
func (u *UniqueVertexFilter) Filter(ip Vertex) {
	u.add(ip)
}

func (u *UniqueVertexFilter) add(ip Vertex) {
	hasMatrix := false
	for _, v := range u.Ips {
		if v.Matrix.Equals(ip.Matrix) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.Ips = append(u.Ips, ip)
	}
}

// ComputePolygon compute overlay.
func (p *PolygonOverlay) ComputePolygon(exitingPoints []Vertex, cpo ComputePolyOverlay) *Plane {
	var pol *Plane = &Plane{}
	for _, iterPoints := range exitingPoints {
		if iterPoints.IsChecked {
			continue
		}
		edge := &Edge{}
		pol.Edge = edge
		pol.Rings = append(pol.Rings, edge)

		start := &iterPoints
		next := &Vertex{Matrix: matrix.Matrix{start.X(), start.Y()}}
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
func ToPolygonMatrix(poly *Plane) matrix.PolygonMatrix {
	result := matrix.PolygonMatrix{}
	for _, v2 := range poly.Rings {
		var edge matrix.LineMatrix
		for _, v1 := range v2.Vertexes {
			edge = append(edge, v1.Matrix)
		}
		if !matrix.Matrix(edge[len(edge)-1]).Equals(matrix.Matrix(edge[0])) {
			edge = append(edge, edge[0])
		}
		result = append(result, edge)
	}

	return result
}
