package clusters

import (
	"fmt"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/planar"
)

// Points is a slice of Point
type Points []geoos.Point

// Center returns the center coordinates of a set of Points
func (points Points) Center() (p geoos.Point, err error) {
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

// AverageDistance returns the average distance between o and all Points
func AverageDistance(point geoos.Point, points Points) float64 {
	var d float64
	var l int
	G := planar.GEOAlgorithm{}

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
