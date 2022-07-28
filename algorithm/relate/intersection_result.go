package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// IntersectionResult ...
type IntersectionResult struct {
	Pos        int
	Start, End int
	Line       matrix.LineSegment
	Ips        IntersectionPointLine
}

// IntersectionPoint overlay point.
type IntersectionPoint struct {
	matrix.Matrix
	IsIntersectionPoint, IsEntering, IsOriginal, IsCollinear bool
}

// X Returns x  .
func (ip *IntersectionPoint) X() float64 {
	return ip.Matrix[0]
}

// Y Returns y  .
func (ip *IntersectionPoint) Y() float64 {
	return ip.Matrix[1]
}

// Compare Returns Compare of  IntersectionPoint.
func (ip *IntersectionPoint) Compare(other *IntersectionPoint, tes int) bool {
	if tes > 0 {
		if ip.X() == other.X() {
			if ip.Y() == other.Y() {
				return ip.IsCollinear
			}
			return ip.Y() < other.Y()
		}
		return ip.X() < other.X()
	}
	if ip.X() == other.X() {
		if ip.Y() == other.Y() {
			return ip.IsCollinear
		}
		return ip.Y() > other.Y()
	}
	return ip.X() > other.X()
}

// IntersectionPointLine overlay point array.
type IntersectionPointLine []IntersectionPoint

// IsOriginal returns line overlays.
func (ips IntersectionPointLine) IsOriginal() bool {
	for _, v := range ips {
		if v.IsOriginal {
			return true
		}
	}
	return false
}

// Len ...
func (ips IntersectionPointLine) Len() int {
	return len(ips)
}

// Less ...
func (ips IntersectionPointLine) Less(i, j int) bool {
	if ips[i].Matrix.Proximity(ips[j].Matrix) {
		return ips[i].IsCollinear
	}
	if ips[i].Matrix[0] == ips[j].Matrix[0] {
		if ips[i].Matrix[1] == ips[j].Matrix[1] {
			return ips[i].IsCollinear
		}
		return ips[i].Matrix[1] < ips[j].Matrix[1]
	}
	return ips[i].Matrix[0] < ips[j].Matrix[0]
}

// Swap ...
func (ips IntersectionPointLine) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}
