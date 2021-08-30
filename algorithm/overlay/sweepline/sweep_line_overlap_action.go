package sweepline

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// OverlapAction An action taken when a SweepLineIndex detects that two Interval.
type OverlapAction interface {
	Overlap(s0, s1 *Interval) bool
}

// Interval ...
type Interval struct {
	Min, Max float64
	Item     *matrix.LineSegment
}

// CoordinatesOverlapAction ...
type CoordinatesOverlapAction struct {
}

// Overlap ...
func (c *CoordinatesOverlapAction) Overlap(s0, s1 *Interval) bool {
	mark, ips := relate.IntersectionLineSegment(s0.Item, s1.Item)
	if mark {
		for _, ip := range ips {
			if ip.IsIntersectionPoint {
				return true
			}
		}
	}
	return false
}
