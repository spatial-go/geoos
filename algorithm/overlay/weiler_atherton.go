package overlay

import (
	"errors"
	"math"

	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
)

// Atherton is a func of overlay.
type Atherton func(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool) *algorithm.Vertex

// SliceContains Returns index of slice.
func SliceContains(list []algorithm.Vertex, p *algorithm.Vertex) (int, error) {
	for i, v := range list {
		if v.X() == p.X() && v.Y() == p.Y() {
			return i, nil
		}
	}
	return len(list), errors.New("point is not in slice")
}

// CrossProduct Returns cross product of a,b point.
func CrossProduct(a, b *algorithm.Vertex) float64 {
	return a.X()*b.Y() - a.Y()*b.X()
}

// InLine returns true if spot in ab,false else.
func InLine(spot, a, b *algorithm.Vertex) bool {
	x := spot.X() <= math.Max(a.X(), b.X()) && spot.X() >= math.Min(a.X(), b.X())
	y := spot.Y() <= math.Max(a.Y(), b.Y()) && spot.Y() >= math.Min(a.Y(), b.Y())
	return x && y
}

// Merge merge polygon.
func Merge(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool) *algorithm.Vertex {
	return Overlay(walkings, pol, start, which, algorithm.MERGE)
}

// Clip clip polygon.
func Clip(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool) *algorithm.Vertex {
	return Overlay(walkings, pol, start, which, algorithm.CLIP)
}

// Overlay overlay polygon.
func Overlay(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool, kind int) *algorithm.Vertex {
	// find in each edge
	for _, w := range walkings {
		if iter, err := SliceContains(w.Vertexs, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexs[iter], which)
				switch kind {
				case algorithm.CLIP:
					if w.IsClockwise {
						iter++
					} else {
						iter--
					}
				case algorithm.MERGE:
					if w.IsClockwise {
						iter--
					} else {
						iter++
					}
				}

				// 循环列表
				if iter == len(w.Vertexs) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.Vertexs) - 1
				}

				if w.Vertexs[iter].IsIntersectionPoint {
					break
				}
			}
			return &w.Vertexs[iter]
		}
	}
	// should not happend
	return &algorithm.Vertex{}
}

// Intersection returns intersection of a and b.
func Intersection(aStart, aEnd, bStart, bEnd *algorithm.Vertex) (mark bool, p *algorithm.Vertex) {
	a1 := aEnd.Y() - aStart.Y()
	b1 := aStart.X() - aEnd.X()
	c1 := -aStart.X()*a1 - b1*aStart.Y()
	a2 := bEnd.Y() - bStart.Y()
	b2 := bStart.X() - bEnd.X()
	c2 := -a2*bStart.X() - b2*bStart.Y()

	var u, v *algorithm.Vertex
	u = aEnd.Sub(aStart)
	v = bEnd.Sub(bStart)

	determinant := CrossProduct(u, v)

	if determinant == 0 {
		mark = false
	} else {
		p = &algorithm.Vertex{}
		p.Matrix = matrix.Matrix{(b1*c2 - b2*c1) / determinant, (a2*c1 - a1*c2) / determinant}

		// check if point belongs to segment
		if InLine(p, aStart, aEnd) && InLine(p, bStart, bEnd) {
			p.IsIntersectionPoint = true
			// determine if the point is entering by determinant
			p.IsEntering = determinant < 0
			mark = true
		} else {
			mark = false
		}
	}
	return
}

// IsIntersectionEdge returns intersection of edge a and b.
func IsIntersectionEdge(aLine, bLine algorithm.Edge) (mark bool) {
	mark = false
	for i, v := range aLine.Vertexs {
		for j, vClip := range bLine.Vertexs {
			if i < len(aLine.Vertexs)-1 && j < len(bLine.Vertexs)-1 {
				markInter, _ := Intersection(&v, &aLine.Vertexs[i+1], &vClip, &aLine.Vertexs[i+1])
				if markInter {
					return
				}
			}
		}
	}
	return
}

// IntersectionEdge returns intersection of edge a and b.
func IntersectionEdge(aLine, bLine algorithm.Edge) (mark bool, ps []*algorithm.Vertex) {
	mark = false
	for i, v := range aLine.Vertexs {
		for j, vClip := range bLine.Vertexs {
			if i < len(aLine.Vertexs)-1 && j < len(bLine.Vertexs)-1 {
				markInter, ip := Intersection(&v, &aLine.Vertexs[i+1], &vClip, &aLine.Vertexs[i+1])
				if markInter {
					mark = markInter
					ps = append(ps, ip)
				}
			}
		}
	}
	return
}

// Weiler Weiler overlay.
func Weiler(subject, clipping *algorithm.Plane, ath Atherton) *algorithm.Plane {
	var pol *algorithm.Plane = &algorithm.Plane{}

	var enteringPoints, exitingPoints []algorithm.Vertex

	var mark bool

	for _, v := range subject.Lines {
		for _, vClip := range clipping.Lines {
			ip := &algorithm.Vertex{}
			mark, ip = Intersection(v.Start, v.End, vClip.Start, vClip.End)

			if mark {
				if ip.IsEntering {
					enteringPoints = append(enteringPoints, *ip)
				} else {
					exitingPoints = append(exitingPoints, *ip)
				}
				AddPointToVertexSlice(subject.Rings, v.Start, v.End, ip)
				AddPointToVertexSlice(clipping.Rings, vClip.Start, vClip.End, ip)
			}
		}
	}

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
			next = ath(subject.Rings, pol, next, true)
			next = ath(clipping.Rings, pol, next, false)
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

// AddPointToVertexSlice add point to vertex slice
func AddPointToVertexSlice(edges []*algorithm.Edge, pStart, pEnd, ip *algorithm.Vertex) {
	for _, v := range edges {

		if start, err := SliceContains(v.Vertexs, pStart); err == nil {
			end, _ := SliceContains(v.Vertexs, pEnd)

			it := start
			distFromStart := measure.PlanarDistance(ip.Matrix, v.Vertexs[it].Matrix)

			// 处理多个交点
			for it != end && it != len(v.Vertexs) {
				if measure.PlanarDistance(v.Vertexs[it].Matrix, v.Vertexs[start].Matrix) >= distFromStart {
					break
				}
				it++
			}

			circ := v.Vertexs[it:]
			v.Vertexs = append([]algorithm.Vertex{}, v.Vertexs[:it]...)
			v.Vertexs = append(v.Vertexs, *ip)
			v.Vertexs = append(v.Vertexs, circ...)
			break

		}
	}
}
