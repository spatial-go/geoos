package clipping

import (
	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/graph/de9im"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/space"
)

const dir = "../graphtests/"

// PolygonClipping  Computes the overlay of two geometries.
type PolygonClipping struct {
	*PointClipping
	clip, shellClip *graph.Clip
	polys           []matrix.PolygonMatrix
}

// Union  Computes the Union of two geometries, if one is encountered.
func (p *PolygonClipping) Union() (matrix.Steric, error) {
	if res, ok := p.unionCheck(); !ok {
		return res, nil
	}
	if ps, ok := p.Subject.(matrix.PolygonMatrix); ok {
		switch pc := p.Clipping.(type) {
		case matrix.Matrix, matrix.LineMatrix:
			return Union(pc, ps)
		case matrix.PolygonMatrix:
			p.clip = graph.ClipHandle(ps, pc)
			if len(ps) == 1 && len(pc) == 1 {
				p.shellClip = p.clip
			}

			switch im := de9im.IMByClip(p.clip); {
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
			if shells, err := p.shellUnion(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[0]}); err == nil {
				p.shellsHandle(shells)
			} else {
				return nil, err
			}
			for i := 1; i < len(pc); i++ {
				if im := de9im.IM(matrix.PolygonMatrix{pc[i]}, ps); im.IsCoveredBy() {
					continue
				}
				if holes, err := holeUnion(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[i]}); err == nil {
					if holes != nil {
						p.holesHandle(holes)
					}
				} else {
					return nil, err
				}
			}
			for i := 1; i < len(ps); i++ {
				if im := de9im.IM(matrix.PolygonMatrix{ps[i]}, pc); im.IsCoveredBy() {
					continue
				}
				if holes, err := holeUnion(matrix.PolygonMatrix{pc[0]}, matrix.PolygonMatrix{ps[i]}); err == nil {
					if holes != nil {
						p.holesHandle(holes)
					}
				} else {
					return nil, err
				}
			}
			return p.Result()
		case matrix.Collection:
			result := matrix.Collection{}
			IsUnion := false
			for _, v := range pc {
				un, err := Union(ps, v)

				if _, ok := un.(matrix.Collection); ok || err != nil {
					result = append(result, v)
				} else {
					IsUnion = true
					result = append(result, un)
				}

			}
			if !IsUnion {
				result = append(result, ps)
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

			if shells, err := edgeIntersection(ps, pc.(matrix.PolygonMatrix)); err == nil {
				p.shellsHandle(shells)
			} else {
				return nil, err
			}
			return p.Result()
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

			if shells, err := edgeDifference(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[0]}); err == nil {
				p.shellsHandle(shells)
			} else {
				return nil, err
			}

			for i := 1; i < len(pc); i++ {
				if im := de9im.IM(matrix.PolygonMatrix{pc[i]}, ps); im.IsCoveredBy() {
					continue
				}
				if holes, err := edgeDifference(matrix.PolygonMatrix{ps[0]}, matrix.PolygonMatrix{pc[i]}); err == nil {
					p.holesHandle(holes)
				} else {
					return nil, algorithm.ErrWrongLink
				}
			}
			for i := 1; i < len(ps); i++ {
				if im := de9im.IM(matrix.PolygonMatrix{ps[i]}, pc); im.IsCoveredBy() {
					p.holesHandle([]matrix.LineMatrix{ps[i]})
				} else {
					if holes, err := edgeDifference(matrix.PolygonMatrix{pc[0]}, matrix.PolygonMatrix{ps[i]}); err == nil {
						p.holesHandle(holes)
					} else {
						return nil, algorithm.ErrWrongLink
					}
				}
			}
			return p.Result()
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
		return Union(result[0], result[1])
	}
}

// shellUnion union shell of two polygonal geometries.
func (p *PolygonClipping) shellUnion(ps, pc matrix.PolygonMatrix) ([]matrix.LineMatrix, error) {

	if p.shellClip == nil {
		p.shellClip = graph.ClipHandle(ps, pc)
	}
	switch im := de9im.IMByClip(p.shellClip); {
	case im.IsCovers():
		return []matrix.LineMatrix{ps[0]}, nil
	case im.IsCoveredBy():
		return []matrix.LineMatrix{pc[0]}, nil
	}

	clip := p.shellClip
	gu, _ := clip.Union()
	gi, _ := clip.Intersection()
	guNodes := gu.Nodes()

	if !geoos.GeoosTestTag {
		geom := space.Collection{}
		for _, v := range gu.Nodes() {
			geom = append(geom, space.TransGeometry(v.Value))
		}
		writeGeom(dir+"data_union_graph.geojson", geom)
	}
	for _, v := range guNodes {
		if v.NodeType == graph.PNode {
			gu.DeleteNode(v)
			continue
		}
		if _, ok := gi.Node(v); !ok {
			psDe9im := de9im.IM(v.Value, ps)

			if psDe9im.IsWithin() {
				gu.DeleteNode(v)
				continue
			}

			pcDe9im := de9im.IM(v.Value, pc)
			if pcDe9im.IsWithin() {
				gu.DeleteNode(v)
				continue
			}
			if psDe9im.IsCoveredBy() && pcDe9im.IsCoveredBy() {
				gu.DeleteNode(v)
				continue
			}
		} else {
			gu.DeleteNode(v)
		}
	}
	if !geoos.GeoosTestTag {
		geom := space.Collection{}
		for _, v := range gu.Nodes() {
			geom = append(geom, space.TransGeometry(v.Value))
		}
		writeGeom(dir+"data_link.geojson", geom)
	}
	return link(gu, gi)
}

// holeUnion union holes of two polygonal geometries.
func holeUnion(ps, pc matrix.PolygonMatrix) ([]matrix.LineMatrix, error) {

	switch im := de9im.IM(ps, pc); {
	case im.IsCovers():
		return nil, nil
	case im.IsCoveredBy():
		return nil, nil
	case !im.IsIntersects():
		return []matrix.LineMatrix{pc[0]}, nil
	default:
		if ok, _ := im.Matches("FF**0****"); ok {
			return []matrix.LineMatrix{pc[0]}, nil
		}
	}
	clip := graph.ClipHandle(ps, pc)
	gu, _ := clip.Union()
	gi, _ := clip.Intersection()
	guNodes := gu.Nodes()
	for _, v := range guNodes {
		if v.NodeType == graph.PNode {
			gu.DeleteNode(v)
			continue
		}
		psDe9im := de9im.IM(v.Value, ps)
		pcDe9im := de9im.IM(v.Value, pc)
		if psDe9im.IsWithin() {
			gu.DeleteNode(v)
		}
		if !pcDe9im.IsCoveredBy() {
			gu.DeleteNode(v)
		}
	}
	return link(gu, gi)
}

// edgeDifference difference edge two polygonal geometries.
func edgeDifference(ps, pc matrix.PolygonMatrix) ([]matrix.LineMatrix, error) {
	clip := graph.ClipHandle(ps, pc)
	gu, _ := clip.Union()
	gi, _ := clip.Intersection()
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
	return link(gu, gi)
}

// edgeIntersection intersection edge two polygonal geometries.
func edgeIntersection(ps, pc matrix.PolygonMatrix) ([]matrix.LineMatrix, error) {
	clip := graph.ClipHandle(ps, pc)
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
	return link(gi, nil)
}

// Result returns result.
func (p *PolygonClipping) Result() (matrix.Steric, error) {
	if len(p.polys) > 1 {
		result := matrix.Collection{}
		for _, poly := range p.polys {
			result = append(result, poly)
		}
		return result, nil
	} else if len(p.polys) == 1 {
		return p.polys[0], nil
	} else {
		return nil, nil
	}
}

func (p *PolygonClipping) shellsHandle(shells []matrix.LineMatrix) {
	p.polys = []matrix.PolygonMatrix{}
	switch len(shells) {
	case 0:
		//TODO
	case 1:
		p.polys = append(p.polys, matrix.PolygonMatrix{shells[0]})
	default:
		maxAreas := 0.0
		maxIndex := 0
		for i, v := range shells {
			if area := measure.Area(v); area > float64(maxAreas) {
				maxAreas = area
				maxIndex = i
			}
		}

		maxShell := matrix.PolygonMatrix{shells[maxIndex]}
		for i, v := range shells {
			if i == maxIndex {
				continue
			}
			switch im := de9im.IM(matrix.PolygonMatrix{v}, maxShell); {
			case im.IsCoveredBy():
				maxShell = append(maxShell, v)
			default:
				p.polys = append(p.polys, matrix.PolygonMatrix{v})
			}
		}
		p.polys = append(p.polys, maxShell)
	}
}

func (p *PolygonClipping) holesHandle(holes []matrix.LineMatrix) {
	switch len(p.polys) {
	case 0:
		//TODO
	case 1:
		for _, v := range holes {
			p.polys[0] = append(p.polys[0], v)
		}
	default:
		for i := range p.polys {
			//TODO
			for _, v := range holes {
				p.polys[i] = append(p.polys[i], v)
			}
		}
	}
}
