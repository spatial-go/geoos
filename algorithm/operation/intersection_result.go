package operation

import (
	"github.com/spatial-go/geoos/algorithm/filter"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Crossroads ...
type Crossroads struct {
	Pos        int
	Start, End int
	Line       matrix.LineSegment
	Ips        IntersectionArray
}

// Intersection overlay point.
type Intersection struct {
	matrix.Matrix
	IsIntersectionPoint, IsEntering, IsOriginal, IsCollinear bool
}

// X Returns x  .
func (ip *Intersection) X() float64 {
	return ip.Matrix[0]
}

// Y Returns y  .
func (ip *Intersection) Y() float64 {
	return ip.Matrix[1]
}

// Compare Returns Compare of  IntersectionPoint.
func (ip *Intersection) Compare(other *Intersection, tes int) bool {
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

// IntersectionArray overlay point array.
type IntersectionArray []Intersection

// IsOriginal returns line overlays.
func (ips IntersectionArray) IsOriginal() bool {
	for _, v := range ips {
		if v.IsOriginal {
			return true
		}
	}
	return false
}

// Len ...
func (ips IntersectionArray) Len() int {
	return len(ips)
}

// Less ...
func (ips IntersectionArray) Less(i, j int) bool {
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
func (ips IntersectionArray) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}

// UniqueIntersectionArrayFilter  A Filter that extracts a unique array.
type UniqueIntersectionArrayFilter struct {
	Ips IntersectionArray
}

// Filter Performs an operation with the provided .
func (u *UniqueIntersectionArrayFilter) Filter(ip interface{}) bool {
	return u.add(ip)
}

func (u *UniqueIntersectionArrayFilter) add(ip interface{}) bool {
	hasMatrix := false
	for _, v := range u.Ips {
		if v.Matrix.Proximity(ip.(Intersection).Matrix) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.Ips = append(u.Ips, ip.(Intersection))
		return true
	}
	return false
}

// Entities  Returns the gathered Matrixes.
func (u *UniqueIntersectionArrayFilter) Entities() interface{} {
	return u.Ips
}

// compile time checks
var (
	_ filter.Filter = &UniqueIntersectionArrayFilter{}
)
