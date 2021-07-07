package overlay

import (
	"errors"
	"math"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Atherton is a func of overlay.
type Atherton func(walkings []*Edge, pol *Polygon, start *Point, which bool) *Point

// SliceContains Returns index of slice.
func SliceContains(list []Point, p *Point) (int, error) {
	for i, v := range list {
		if v.X() == p.X() && v.Y() == p.Y() {
			return i, nil
		}
	}
	return len(list), errors.New("point not in this slice")
}

// CrossProduct Returns cross product of a,b point.
func CrossProduct(a, b *Point) float64 {
	return a.X()*b.Y() - a.Y()*b.X()
}

// InLine returns true if spot in ab,false else.
func InLine(spot, a, b *Point) bool {
	x := spot.X() <= math.Max(a.X(), b.X()) && spot.X() >= math.Min(a.X(), b.X())
	y := spot.Y() <= math.Max(a.Y(), b.Y()) && spot.Y() >= math.Min(a.Y(), b.Y())
	return x && y
}

// Distance returns Distance of pq.
func Distance(p, q *Point) float64 {
	return math.Sqrt((p.X()-q.X())*(p.X()-q.X()) + (p.Y()-q.Y())*(p.Y()-q.Y()))
}

// Merge merge polygon.
func Merge(walkings []*Edge, pol *Polygon, start *Point, which bool) *Point {
	return Overlay(walkings, pol, start, which, MERGE)
}

// Clip clip polygon.
func Clip(walkings []*Edge, pol *Polygon, start *Point, which bool) *Point {
	return Overlay(walkings, pol, start, which, CLIP)
}

// Overlay overlay polygon.
func Overlay(walkings []*Edge, pol *Polygon, start *Point, which bool, kind int) *Point {
	// find in each edge
	for _, w := range walkings {
		if iter, err := SliceContains(w.points, start); err == nil {
			for {
				pol.AddPointWhich(&w.points[iter], which)
				switch kind {
				case CLIP:
					if w.isClockwise {
						iter++
					} else {
						iter--
					}
				case MERGE:
					if w.isClockwise {
						iter--
					} else {
						iter++
					}
				}

				// 循环列表
				if iter == len(w.points) {
					iter = 0
				}
				if iter == -1 {
					iter = len(w.points) - 1
				}

				if w.points[iter].isIntersectionPoint {
					break
				}
			}
			return &w.points[iter]
		}
	}
	// should not happend
	return &Point{}
}

// Intersection returns intersection of a and b.
func Intersection(aStart, aEnd, bStart, bEnd *Point) (mark bool, p *Point) {
	a1 := aEnd.Y() - aStart.Y()
	b1 := aStart.X() - aEnd.X()
	c1 := -aStart.X()*a1 - b1*aStart.Y()
	a2 := bEnd.Y() - bStart.Y()
	b2 := bStart.X() - bEnd.X()
	c2 := -a2*bStart.X() - b2*bStart.Y()

	var u, v *Point
	u = aEnd.Sub(aStart)
	v = bEnd.Sub(bStart)

	determinant := CrossProduct(u, v)

	if determinant == 0 {
		mark = false
	} else {
		p = &Point{}
		p.matrix = matrix.Matrix{(b1*c2 - b2*c1) / determinant, (a2*c1 - a1*c2) / determinant}

		// check if point belongs to segment
		if InLine(p, aStart, aEnd) && InLine(p, bStart, bEnd) {
			p.isIntersectionPoint = true
			// determine if the point is entering by determinant
			p.isEntering = determinant < 0
			mark = true
		} else {
			mark = false
		}
	}
	return
}

// Weiler Weiler overlay.
func Weiler(subject, clipping *Polygon, ath Atherton) *Polygon {
	var pol *Polygon = &Polygon{}

	var enteringPoints, exitingPoints []Point

	var mark bool

	for _, v := range subject.lines {
		for _, vClip := range clipping.lines {
			ip := &Point{}
			mark, ip = Intersection(v.start, v.end, vClip.start, vClip.end)

			if mark {
				if ip.isEntering {
					enteringPoints = append(enteringPoints, *ip)
				} else {
					exitingPoints = append(exitingPoints, *ip)
				}
				AddPointToVertexSlice(subject.rings, v.start, v.end, ip)
				AddPointToVertexSlice(clipping.rings, vClip.start, vClip.end, ip)
			}
		}
	}

	for _, iterPoints := range exitingPoints {
		if iterPoints.isChecked {
			continue
		}
		edge := &Edge{}
		pol.edge = edge
		pol.rings = append(pol.rings, edge)

		start := &iterPoints
		next := &Point{matrix: matrix.Matrix{start.X(), start.Y()}}
		start.isChecked = true

		for {
			next = ath(subject.rings, pol, next, true)
			next = ath(clipping.rings, pol, next, false)
			if where, err := SliceContains(exitingPoints, next); err == nil {
				exitingPoints[where].isChecked = true
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
func AddPointToVertexSlice(edges []*Edge, pStart, pEnd, ip *Point) {
	for _, v := range edges {

		if start, err := SliceContains(v.points, pStart); err == nil {
			end, _ := SliceContains(v.points, pEnd)

			it := start
			distFromStart := Distance(ip, &v.points[it])

			// 处理多个交点
			for it != end && it != len(v.points) {
				if Distance(&v.points[it], &v.points[start]) >= distFromStart {
					break
				}
				it++
			}

			circ := v.points[it:]
			v.points = append([]Point{}, v.points[:it]...)
			v.points = append(v.points, *ip)
			v.points = append(v.points, circ...)
			break

		}
	}
}
