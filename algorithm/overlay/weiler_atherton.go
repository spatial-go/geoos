package overlay

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/algoerr"
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/algorithm/relate"
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
	return len(list), algoerr.ErrNotInSlice
}

// Merge merge polygon.
func Merge(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool) *algorithm.Vertex {
	return Overlay(walkings, pol, start, which, calc.MERGE)
}

// Clip clip polygon.
func Clip(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool) *algorithm.Vertex {
	return Overlay(walkings, pol, start, which, calc.CLIP)
}

// Overlay overlay polygon.
func Overlay(walkings []*algorithm.Edge, pol *algorithm.Plane, start *algorithm.Vertex, which bool, kind int) *algorithm.Vertex {
	// find in each edge
	for _, w := range walkings {
		if iter, err := SliceContains(w.Vertexs, start); err == nil {
			for {
				pol.AddPointWhich(&w.Vertexs[iter], which)
				switch kind {
				case calc.CLIP:
					if w.IsClockwise {
						iter++
					} else {
						iter--
					}
				case calc.MERGE:
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

// Weiler Weiler overlay.
func Weiler(subject, clipping *algorithm.Plane, ath Atherton) *algorithm.Plane {
	var pol *algorithm.Plane = &algorithm.Plane{}

	var enteringPoints, exitingPoints []algorithm.Vertex

	var mark bool

	for _, v := range subject.Lines {
		for _, vClip := range clipping.Lines {
			ip := &algorithm.Vertex{}
			mark, ip.Matrix, ip.IsIntersectionPoint, ip.IsEntering =
				relate.Intersection(v.Start.Matrix, v.End.Matrix, vClip.Start.Matrix, vClip.End.Matrix)

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
