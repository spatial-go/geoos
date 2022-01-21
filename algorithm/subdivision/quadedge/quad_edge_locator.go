package quadedge

import "github.com/spatial-go/geoos/algorithm/matrix"

type Locator interface {
	locate(v matrix.Matrix) *QuadEdge
}

type LastFoundQuadEdgeLocator struct {
	subdivision *QuadEdgeSubdivision
	lastEdge    *QuadEdge
}

func NewLastFoundQuadEdgeLocator(subdivision *QuadEdgeSubdivision) *LastFoundQuadEdgeLocator {
	locator := &LastFoundQuadEdgeLocator{subdivision: subdivision}
	locator.init()
	return locator
}

func (l *LastFoundQuadEdgeLocator) init() {
	l.lastEdge = l.findEdge()
}

func (l *LastFoundQuadEdgeLocator) findEdge() *QuadEdge {
	edges := l.subdivision.Edges()
	return edges[0]
}

func (l *LastFoundQuadEdgeLocator) locate(v matrix.Matrix) *QuadEdge {
	if !l.lastEdge.isLive() {
		l.init()
	}

	edge, err := l.subdivision.locateFromEdge(v, l.lastEdge)
	if err == nil {
		l.lastEdge = edge
	} else {
		print(err)
	}
	return l.lastEdge
}
