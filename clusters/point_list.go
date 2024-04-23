package clusters

import (
	"fmt"
	"sort"

	"github.com/spatial-go/geoos/planar"
	"github.com/spatial-go/geoos/space"
)

// PointList is a slice of Point
type PointList []space.Point

// Center returns the center coordinates of a set of Points
func (points PointList) Center() (p space.Point, err error) {
	p = make(space.Point, 2)
	var l = len(points)
	if l == 0 {
		return p, fmt.Errorf("there is no mean for an empty set of points")
	}

	for _, point := range points {
		for j, v := range point {
			p[j] += v
		}
	}

	p[0] = p[0] / float64(l)
	p[1] = p[1] / float64(l)

	return p, nil
}

// AverageDistance returns the average distance between o and all PointList
func AverageDistance(point space.Point, points PointList) float64 {
	var d float64
	var l int
	G := planar.NormalStrategy()

	for _, observation := range points {
		dist, _ := G.Distance(point, observation)
		if dist == 0 {
			continue
		}

		l++
		d += dist
	}

	if l == 0 {
		return 0
	}
	return d / float64(l)
}

// ConvexHull returns the convex hull of a list of points
// Implementation of Andrew's Monotone Chain algorithm as specified on https://en.wikibooks.org/wiki/Algorithm_Implementation/Geometry/Convex_hull/Monotone_chain
func (points PointList) ConvexHull() PointList {
	// empty slice, single point, and two points are already their own convex hull
	if len(points) <= 2 {
		return points
	}

	// sort points by x coordinates (ties stay in original order)
	sort.SliceStable(points, func(i, j int) bool {
		return points[i].X() < points[j].X() || (points[i].X() == points[j].X() && points[i].Y() < points[j].Y())
	})

	// build lower hull
	lowerHull := PointList{}
	for _, p := range points {
		for len(lowerHull) >= 2 && cross(lowerHull[len(lowerHull)-2], lowerHull[len(lowerHull)-1], p) <= 0 {
			// pop last point
			lowerHull = lowerHull[:len(lowerHull)-1]
		}
		// add p
		lowerHull = append(lowerHull, p)
	}
	// build upper hull
	upperHull := PointList{}
	// iterate in reverse order
	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(upperHull) >= 2 && cross(upperHull[len(upperHull)-2], upperHull[len(upperHull)-1], p) <= 0 {
			// pop last point
			upperHull = upperHull[:len(upperHull)-1]
		}
		// add p
		upperHull = append(upperHull, p)
	}

	// concatenate lower and upper hull to build convexHull.
	// omit the last point of lowerHull as it is the beginning of upperHull
	// keep last point of upperHull to get a closed polygon shape (last point = first point)
	hull := append(lowerHull[:len(lowerHull)-1], upperHull...)

	if len(hull) == 3 {
		// for lines, remove last point (so no closed polygon shape)
		hull = hull[:len(hull)-1]
	}
	return hull
}

// cross is a helper function for ConvexHull()
// cross product of OA and OB vectors.
// returns positive value, if OAB makes a counter-clockwise turn,
// negative for clockwise turn, and zero if the points are colinear.
func cross(o, a, b space.Point) float64 {
	return (a.X()-o.X())*(b.Y()-o.Y()) - (a.Y()-o.Y())*(b.X()-o.X())
}
