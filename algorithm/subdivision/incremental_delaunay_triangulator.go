package subdivision

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/subdivision/quadedge"
)

// IncrementalDelaunayTriangulator ...
type IncrementalDelaunayTriangulator struct {
	subdivision *quadedge.Subdivision
}

// NewIncrementalDelaunayTriangulator ...
func NewIncrementalDelaunayTriangulator(subdivision *quadedge.Subdivision) *IncrementalDelaunayTriangulator {
	return &IncrementalDelaunayTriangulator{
		subdivision: subdivision,
	}
}

func (t *IncrementalDelaunayTriangulator) insertSites(sites []matrix.Matrix) {
	for _, site := range sites {
		t.insertSite(site)
	}
}

// insertSite insert a new point to subdivision representing a Delaunay Triangulator
func (t *IncrementalDelaunayTriangulator) insertSite(v matrix.Matrix) *quadedge.QuadEdge {
	e := t.subdivision.Locate(v)

	if t.subdivision.IsVertexOfEdge(e, v) {
		// point is already in subdivision.
		return e
	} else if t.subdivision.IsOnEdge(e, v) {
		e = e.OPrev()
		t.subdivision.Delete(e.ONext())
	}

	base := t.subdivision.MakeEdge(e.Origin(), v)
	quadedge.Splice(base, e)
	startEdge := base
	done := false
	for !done || e.LNext() != startEdge {
		base = t.subdivision.Connect(e, base.Sym())
		e = base.OPrev()
		done = true
	}
	for {
		pe := e.OPrev()
		if quadedge.IsCCW(pe.Destination(), e.Destination(), e.Origin()) && quadedge.IsInCircle(v, e.Origin(), pe.Destination(), e.Destination()) {
			quadedge.Swap(e)
			e = e.OPrev()
		} else if e.ONext() == startEdge {
			break
		} else {
			e = e.ONext().LPrev()
		}
	}
	return base
}
