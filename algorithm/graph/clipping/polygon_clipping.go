package clipping

import (
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/graph/de9im"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay"
)

// PolygonClipping  Computes the overlay of two geometries,either or both of which may be nil.
type PolygonClipping struct {
	*PointClipping
	subjectPlane, clippingPlane *overlay.Plane
}

// Union  Computes the Union of two geometries,either or both of which may be nil.
func (p *PolygonClipping) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.PolygonMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix, matrix.LineMatrix:
			return Union(pc, ps), nil
		case matrix.PolygonMatrix:
			switch im := de9im.IM(ps, pc); {
			case im.IsCovers():
				return ps, nil
			case im.IsCoveredBy():
				return pc, nil
			case !im.IsIntersects():
				return matrix.Collection{ps, pc}, nil
			default:
				if ok, _ = im.Matches("FF**0****"); ok {
					return matrix.Collection{ps, pc}, nil
				}
			}

			result := matrix.PolygonMatrix{shellUnion(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[0]})}
			for i := 1; i < len(pc); i++ {
				result = append(result, holeUnion(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[i]}))
			}
			for i := 1; i < len(ps); i++ {
				result = append(result, holeUnion(matrix.PolygonMatrix{pc[0]}, matrix.PolygonMatrix{ps[i]}))
			}
			return result, nil
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// Intersection  Computes the Intersection of two geometries,either or both of which may be nil.
func (p *PolygonClipping) Intersection() (matrix.Steric, error) {
	if res, ok := p.intersectionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.PolygonMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix, matrix.LineMatrix:
			return Intersection(pc, ps)
		default:
			switch im := de9im.IM(ps, pc); {
			case im.IsContains():
				return pc, nil
			case im.IsWithin():
				return ps, nil
			case !im.IsIntersects():
				return matrix.PolygonMatrix{}, nil
			}

			result := matrix.PolygonMatrix{edgeIntersection(ps, pc.(matrix.PolygonMatrix))}
			return result, nil
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// Difference returns a geometry that represents that part of geometry A that does not intersect with geometry B.
// One can think of this as GeometryA - Intersection(A,B).
// If A is completely contained in B then an empty geometry collection is returned.
func (p *PolygonClipping) Difference() (matrix.Steric, error) {
	if res, ok := p.differenceCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.PolygonMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix, matrix.LineMatrix:
			// return Difference(pc, ps), nil
		case matrix.PolygonMatrix:
			switch im := de9im.IM(ps, pc); {
			case im.IsCoveredBy():
				return matrix.PolygonMatrix{}, nil
			case im.IsCovers():
				if ok, _ = im.Matches("****1****"); !ok {
					result := matrix.PolygonMatrix{matrix.LineMatrix(ps[0])}

					result = append(result, pc[0])

					for i := 1; i < len(ps); i++ {
						result = append(result, ps[i])
					}
					return result, nil
				}
			default:
				if ok, _ = im.Matches("FF**0****"); ok {
					return ps, nil
				}

			}

			result := matrix.PolygonMatrix{edgeDifference(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[0]})}
			for i := 1; i < len(pc); i++ {
				result = append(result, edgeDifference(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[i]}))
			}
			for i := 1; i < len(ps); i++ {
				result = append(result, edgeDifference(matrix.PolygonMatrix{pc[0]}, matrix.PolygonMatrix{ps[i]}))
			}
			return result, nil
		}
	}
	return nil, algorithm.ErrNotMatchType
}

// DifferenceReverse returns a geometry that represents reverse that part of geometry A that does not intersect with geometry B .
// One can think of this as GeometryB - Intersection(A,B).
// If B is completely contained in A then an empty geometry collection is returned.
func (p *PolygonClipping) DifferenceReverse() (matrix.Steric, error) {
	newPoly := &PolygonClipping{PointClipping: &PointClipping{Subject: p.Clipping, Clipping: p.Subject}}
	return newPoly.Difference()
}

// SymDifference returns a geometry that represents the portions of A and B that do not intersect.
// It is called a symmetric difference because SymDifference(A,B) = SymDifference(B,A).
// One can think of this as Union(geomA,geomB) - Intersection(A,B).
func (p *PolygonClipping) SymDifference() (matrix.Steric, error) {
	result := matrix.Collection{}
	if res, err := p.Difference(); err == nil && !res.IsEmpty() {
		result = append(result, res)
	}
	if res, err := p.DifferenceReverse(); err == nil && !res.IsEmpty() {
		result = append(result, res)
	}
	switch len(result) {
	case 0:
		return nil, nil
	case 1:
		return result[0], nil
	default:
		return Union(result[0], result[1]), nil
	}
}

// shellUnion union shell of two polygonal geometries.
func shellUnion(ps, pc matrix.PolygonMatrix) matrix.LineMatrix {
	clip := ClipHandle(ps, pc)
	gu, _ := clip.Union()
	guNodes := gu.Nodes()
	for _, v := range guNodes {
		if v.NodeType == graph.PNode {
			gu.DeleteNode(v)
		}
		psDe9im := de9im.IM(v.Value, ps)
		pcDe9im := de9im.IM(v.Value, pc)
		if psDe9im.IsWithin() {
			gu.DeleteNode(v)
		}
		if pcDe9im.IsWithin() {
			gu.DeleteNode(v)
		}
		if psDe9im.IsCoveredBy() && pcDe9im.IsCoveredBy() {
			gu.DeleteNode(v)
		}
	}
	return link(gu)
}

// holeUnion union holes of two polygonal geometries.
func holeUnion(ps, pc matrix.PolygonMatrix) matrix.LineMatrix {
	clip := ClipHandle(ps, pc)
	gu, _ := clip.Union()
	guNodes := gu.Nodes()
	for _, v := range guNodes {
		if v.NodeType == graph.PNode {
			gu.DeleteNode(v)
		}
		if de9im.IM(v.Value, ps).IsWithin() {
			gu.DeleteNode(v)
		}
		if !de9im.IM(v.Value, pc).IsCoveredBy() {
			gu.DeleteNode(v)
		}
	}
	return link(gu)
}

// edgeDifference difference edge two polygonal geometries.
func edgeDifference(ps, pc matrix.PolygonMatrix) matrix.LineMatrix {
	clip := ClipHandle(ps, pc)
	gu, _ := clip.Union()
	guNodes := gu.Nodes()
	for _, v := range guNodes {
		if v.NodeType == graph.PNode {
			gu.DeleteNode(v)
		}
		psDe9im := de9im.IM(v.Value, ps)
		pcDe9im := de9im.IM(v.Value, pc)
		if pcDe9im.IsWithin() {
			gu.DeleteNode(v)
		}
		if !psDe9im.IsCoveredBy() {
			gu.DeleteNode(v)
		}
		if psDe9im.IsCoveredBy() && pcDe9im.IsCoveredBy() && !psDe9im.IsWithin() {
			gu.DeleteNode(v)
		}
	}
	return link(gu)
}

// edgeIntersection intersection edge two polygonal geometries.
func edgeIntersection(ps, pc matrix.PolygonMatrix) matrix.LineMatrix {
	clip := ClipHandle(ps, pc)
	gu, _ := clip.Union()
	guNodes := gu.Nodes()
	gi := &graph.MatrixGraph{}
	for _, v := range guNodes {
		if v.NodeType == graph.PNode {
			continue
		}
		psDe9im := de9im.IM(v.Value, ps)
		pcDe9im := de9im.IM(v.Value, pc)
		if psDe9im.IsCoveredBy() && pcDe9im.IsCoveredBy() {
			gi.AddNode(v)
		}
		if psDe9im.IsWithin() {
			gi.AddNode(v)
		}
		if pcDe9im.IsWithin() {
			gi.AddNode(v)
		}
	}
	return link(gi)
}

// link returns edge by link nodes
func link(g graph.Graph) (result matrix.LineMatrix) {
	result = matrix.LineMatrix{}

	for {
		guNodes := g.Nodes()
		for _, v := range guNodes {
			if v.NodeType == graph.CNode || v.NodeType == graph.LNode {
				line := v.Value.(matrix.LineMatrix)
				startPoint := matrix.Matrix(line[0])
				lastPoint := matrix.Matrix(line[len(line)-1])
				if len(result) == 0 {
					for _, point := range line {
						result = append(result, point)
					}
					g.DeleteNode(v)
					break
				} else {
					if matrix.Matrix(result[len(result)-1]).Equals(startPoint) {
						for i, point := range line {
							if i == 0 {
								continue
							}
							result = append(result, point)
						}
						g.DeleteNode(v)
						break
					}
					if matrix.Matrix(result[len(result)-1]).Equals(lastPoint) {
						for i, point := range line.Reverse() {
							if i == 0 {
								continue
							}
							result = append(result, point)
						}
						g.DeleteNode(v)
						break
					}
				}

			}
		}
		if result.IsClosed() {
			break
		}
	}
	return result
}
