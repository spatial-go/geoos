package relate

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Relationship  be used during the relate computation.
type Relationship struct {
	// The operation args into an array so they can be accessed by index
	Arg            []matrix.Steric // the arg(s) of the operation
	IntersectBound bool
	relateComputer Computer
}

// IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (r *Relationship) IntersectionMatrix() *matrix.IntersectionMatrix {
	return r.relateComputer.computeIM(r.Arg, r.IntersectBound)
}
