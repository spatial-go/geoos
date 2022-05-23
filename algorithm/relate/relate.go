// Package relate the  criteria  for  judging  the  various
// topological relations between points, line and surface entities in DE-9IM model are given.
package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Relationship  be used during the relate computation.
type Relationship struct {
	// The operation args into an array so they can be accessed by index
	Arg            []matrix.Steric // the arg(s) of the operation
	IntersectBound bool
	relateComputer *Computer
}

// Relate Gets the relate string for the spatial relationship
// between the input geometries.
func Relate(g0, g1 matrix.Steric, intersectBound bool) string {

	im := IM(g0, g1, intersectBound)
	return im.ToString()
}

// IM Gets the relate  for the spatial relationship
// between the input geometries.
func IM(g0, g1 matrix.Steric, intersectBound bool) *matrix.IntersectionMatrix {
	rs := &Relationship{
		Arg:            []matrix.Steric{g0, g1},
		IntersectBound: intersectBound,
		relateComputer: &Computer{},
	}
	return rs.IntersectionMatrix()
}

// IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (r *Relationship) IntersectionMatrix() *matrix.IntersectionMatrix {
	return r.relateComputer.computeIM(r.Arg, r.IntersectBound)
}
