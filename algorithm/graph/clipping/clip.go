package clipping

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Clip ...
type Clip struct {
	Arg   []matrix.Steric // the arg(s) of the operation
	graph []graph.Graph
}

// ClipHandle handle graph with m1 and m2,returns graph of intersection , union, difference and sym  difference.
func ClipHandle(m0, m1 matrix.Steric) *Clip {
	clip := &Clip{Arg: []matrix.Steric{m0, m1}, graph: make([]graph.Graph, 2)}
	for i, v := range clip.Arg {
		clip.graph[i], _ = graph.GenerateGraph(v)
	}

	if err := graph.IntersectionHandle(clip.Arg[0], clip.Arg[1], clip.graph[0], clip.graph[1]); err != nil {
		return nil
	}

	return clip
}

// Intersection  Computes the Intersection of two Graph.
func (c *Clip) Intersection() (graph.Graph, error) {
	return c.graph[0].Intersection(c.graph[1])
}

// Union  Computes the Union of two Graph.
func (c *Clip) Union() (graph.Graph, error) {
	return c.graph[0].Union(c.graph[1])
}

// Difference returns a Graph that represents that part of Graph A that does not intersect with Graph B.
// One can think of this as GraphA - Intersection(A,B).
func (c *Clip) Difference() (graph.Graph, error) {
	return c.graph[0].Difference(c.graph[1])
}

// SymDifference returns a Graph that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
//
// One can think of this as Union(A,B) - Intersection(A,B).
func (c *Clip) SymDifference() (graph.Graph, error) {
	return c.graph[0].SymDifference(c.graph[1])
}

// UnaryUnion returns a Geometry containing the union.
//	or an empty atomic geometry, or an empty GEOMETRYCOLLECTION
func UnaryUnion(matrix4 matrix.Steric) matrix.Steric {
	if c, ok := matrix4.(matrix.Collection); ok {
		return UnaryUnionByHalf(c, 0, len(c))
	}
	return nil
}

// UnaryUnionByHalf returns Unions a section of a list using a recursive binary union on each half of the section.
func UnaryUnionByHalf(matrix4 matrix.Collection, start, end int) matrix.Steric {
	if matrix4 == nil {
		return nil
	}
	if end-start <= 1 {
		return Union(matrix4[start], nil)
	} else if end-start == 2 {
		return Union(matrix4[start], matrix4[start+1])
	} else {
		mid := (end + start) / 2
		g0 := UnaryUnionByHalf(matrix4, start, mid)
		g1 := UnaryUnionByHalf(matrix4, mid, end)
		return Union(g0, g1)
	}
}

// Union  Computes the Union of two geometries,either or both of which may be null.
func Union(m0, m1 matrix.Steric) (result matrix.Steric) {
	switch m := m0.(type) {
	case matrix.Matrix:
		over := &PointClipping{Subject: m, Clipping: m1}
		result, _ = over.Union()
	case matrix.LineMatrix:
		over := &LineClipping{PointClipping: &PointClipping{Subject: m0, Clipping: m1}}
		result, _ = over.Union()
	case matrix.PolygonMatrix:
		polyOver := &PolygonClipping{PointClipping: &PointClipping{Subject: m0, Clipping: m1}}
		result, _ = polyOver.Union()
	case matrix.Collection:
		result = Union(UnaryUnion(m), m1)
	}
	return result
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be null.
func Intersection(m0, m1 matrix.Steric) (matrix.Steric, error) {
	switch m := m0.(type) {
	case matrix.Matrix:
		over := &PointClipping{Subject: m, Clipping: m1}
		return over.Intersection()
	case matrix.LineMatrix:
		var err error
		newLine := &LineClipping{PointClipping: &PointClipping{Subject: m, Clipping: m1}}
		if result, err := newLine.Intersection(); err == nil {
			return result, nil
		}
		return nil, err
	case matrix.PolygonMatrix:
		polyOver := &PolygonClipping{PointClipping: &PointClipping{Subject: m, Clipping: m1}}
		return polyOver.Intersection()
	default:
		return nil, algorithm.ErrNotSupportCollection

	}
}
