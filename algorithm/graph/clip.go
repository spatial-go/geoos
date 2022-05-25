package graph

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Clip ...
type Clip struct {
	Arg      []matrix.Steric // the arg(s) of the operation
	ArgGraph []Graph
	IsDesc   bool
}

// ClipHandle handle graph with m1 and m2,returns graph of intersection , union, difference and sym  difference.
func ClipHandle(m0, m1 matrix.Steric) *Clip {
	arg := []matrix.Steric{m0, m1}

	if p0, ok := m0.(matrix.PolygonMatrix); ok {
		if p1, ok := m1.(matrix.PolygonMatrix); ok {
			polygonHandle(p0, p1)
			arg = []matrix.Steric{p0, p1}
		}
	}
	clip := &Clip{Arg: arg, ArgGraph: make([]Graph, 2)}

	for i, v := range clip.Arg {
		clip.ArgGraph[i], _ = GenerateGraph(v)
	}

	if err := IntersectionHandle(clip.Arg[0], clip.Arg[1], clip.ArgGraph[0], clip.ArgGraph[1]); err != nil {
		return nil
	}

	return clip
}

// Intersection  Computes the Intersection of two Graph.
func (c *Clip) Intersection() (Graph, error) {
	return c.ArgGraph[0].Intersection(c.ArgGraph[1])
}

// Union  Computes the Union of two Graph.
func (c *Clip) Union() (Graph, error) {
	return c.ArgGraph[0].Union(c.ArgGraph[1])
}

// Difference returns a Graph that represents that part of Graph A that does not intersect with Graph B.
// One can think of this as GraphA - Intersection(A,B).
func (c *Clip) Difference() (Graph, error) {
	return c.ArgGraph[0].Difference(c.ArgGraph[1])
}

// SymDifference returns a Graph that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
//
// One can think of this as Union(A,B) - Intersection(A,B).
func (c *Clip) SymDifference() (Graph, error) {
	return c.ArgGraph[0].SymDifference(c.ArgGraph[1])
}

func polygonHandle(m0, m1 matrix.PolygonMatrix) {
	for _, v1 := range m0 {
		for _, v2 := range v1 {
			for _, u1 := range m1 {
				for _, u2 := range u1 {
					if matrix.Matrix(v2).Proximity(matrix.Matrix(u2)) &&
						!matrix.Matrix(v2).Equals(matrix.Matrix(u2)) {
						v2[0] = u2[0]
						v2[1] = u2[1]
					}
				}
			}
		}
	}
}
