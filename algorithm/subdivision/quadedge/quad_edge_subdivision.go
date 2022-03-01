package quadedge

import (
	"errors"
	"fmt"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/algorithm/measure"
	"github.com/spatial-go/geoos/utils"
)

// EdgeCoincidenceTolFactor ...
const EdgeCoincidenceTolFactor = 1000

// Subdivision ...
type Subdivision struct {
	tolerance                float64
	edgeCoincidenceTolerance float64
	quadEdge                 *QuadEdge
	quadEdgeList             []*QuadEdge
	startingEdge             *QuadEdge
	frameEnv                 *envelope.Envelope
	frameVertex              []matrix.Matrix
	locator                  Locator
	triEdges                 []*QuadEdge
	visitedKey               int
}

// NewQuadEdgeSubdivision ...
func NewQuadEdgeSubdivision(env *envelope.Envelope, tolerance float64) *Subdivision {
	q := &Subdivision{}
	q.triEdges = make([]*QuadEdge, 3, 3)
	q.tolerance = tolerance
	q.edgeCoincidenceTolerance = tolerance / EdgeCoincidenceTolFactor
	q.createFrame(env)
	q.initSubdivision()
	q.locator = NewLastFoundQuadEdgeLocator(q)
	return q
}

func (q *Subdivision) createFrame(env *envelope.Envelope) {
	var (
		deltaX = env.Width()
		deltaY = env.Height()
		offset = 0.0
	)
	if deltaX > deltaY {
		offset = deltaX * 10.0
	} else {
		offset = deltaY * 10.0
	}
	q.frameVertex = make([]matrix.Matrix, 3, 3)
	q.frameVertex[0] = matrix.Matrix{(env.MaxX + env.MinX) / 2.0, env.MaxY + offset}
	q.frameVertex[1] = matrix.Matrix{env.MinX - offset, env.MinY - offset}
	q.frameVertex[2] = matrix.Matrix{env.MaxX + offset, env.MinY - offset}

	q.frameEnv = envelope.Bound(q.frameVertex[0:2])
	q.frameEnv.ExpandToIncludeMatrix(q.frameVertex[2])
}

// MakeEdge ...
func (q *Subdivision) MakeEdge(o matrix.Matrix, d matrix.Matrix) *QuadEdge {
	qe := NewQuadEdge(o, d)
	q.quadEdgeList = append(q.quadEdgeList, qe)
	return qe
}

func (q *Subdivision) initSubdivision() {
	var (
		ea = q.MakeEdge(q.frameVertex[0], q.frameVertex[1])
		eb = q.MakeEdge(q.frameVertex[1], q.frameVertex[2])
		ec = q.MakeEdge(q.frameVertex[2], q.frameVertex[0])
	)
	Splice(ea.Sym(), eb)
	Splice(eb.Sym(), ec)
	Splice(ec.Sym(), ea)
	q.startingEdge = ea
}

// Connect ...
func (q *Subdivision) Connect(a *QuadEdge, b *QuadEdge) *QuadEdge {
	e := Connect(a, b)
	q.quadEdgeList = append(q.quadEdgeList, e)
	return e
}

// Delete ...
func (q *Subdivision) Delete(e *QuadEdge) {
	Splice(e, e.OPrev())
	Splice(e.Sym(), e.Sym().OPrev())
	var (
		eSym    = e.Sym()
		eRot    = e.rot
		eRotSym = e.rot.Sym()
	)

	q.removeEdge(e)
	q.removeEdge(eSym)
	q.removeEdge(eRot)
	q.removeEdge(eRotSym)

	e.rot = nil
	eSym.rot = nil
	eRot.rot = nil
	eRotSym.rot = nil
}

func (q *Subdivision) removeEdge(e *QuadEdge) {
	var (
		index = -1
	)
	for i := range q.quadEdgeList {
		if q.quadEdgeList[i] == e {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	q.quadEdgeList = append(q.quadEdgeList[:index], q.quadEdgeList[index+1:]...)
}

// IsVertexOfEdge ...
func (q *Subdivision) IsVertexOfEdge(e *QuadEdge, v matrix.Matrix) bool {
	return (v.EqualsExact(e.Origin(), q.tolerance)) || (v.EqualsExact(e.Destination(), q.tolerance))
}

// IsOnEdge ...
func (q *Subdivision) IsOnEdge(e *QuadEdge, p matrix.Matrix) bool {
	dist := measure.PlanarDistance(p, matrix.LineMatrix{e.Origin(), e.Destination()})
	return dist < q.edgeCoincidenceTolerance
}

// IsFrameVertex ...
func (q *Subdivision) IsFrameVertex(v matrix.Matrix) bool {
	if v.Equals(q.frameVertex[0]) {
		return true
	} else if v.Equals(q.frameVertex[1]) {
		return true
	} else if v.Equals(q.frameVertex[2]) {
		return true
	}
	return false
}

func (q *Subdivision) getPrimaryEdges(includeFrame bool) []*QuadEdge {
	q.visitedKey++

	var (
		edges        = make([]*QuadEdge, 0, 0)
		edgeStack    = utils.NewStack()
		visitedEdges = make(map[*QuadEdge]struct{})
	)
	edgeStack.Push(q.startingEdge)

	for !edgeStack.Empty() {
		edge := edgeStack.Pop().(*QuadEdge)
		if _, found := visitedEdges[edge]; !found {
			priQE := edge.Primary()

			if includeFrame || !q.IsFrameEdge(priQE) {
				edges = append(edges, priQE)
			}

			edgeStack.Push(edge.ONext())
			edgeStack.Push(edge.Sym().ONext())

			visitedEdges[edge] = struct{}{}
			visitedEdges[edge.Sym()] = struct{}{}
		}
	}
	return edges
}

// IsFrameEdge whether a QuadEdge is an edge incident on a frame triangle vertex
func (q *Subdivision) IsFrameEdge(e *QuadEdge) bool {
	return (q.IsFrameVertex(e.Origin())) || (q.IsFrameVertex(e.Destination()))
}

func (q *Subdivision) getVertexUniqueEdges(includeFrame bool) []*QuadEdge {
	var (
		edges           []*QuadEdge
		visitedVertices = make(map[string]struct{})
	)
	for _, edge := range q.quadEdgeList {
		v := edge.Origin()
		key := fmt.Sprintf("%f_%f", v[0], v[1])
		if _, found := visitedVertices[key]; !found {
			visitedVertices[key] = struct{}{}
			if includeFrame || !q.IsFrameVertex(v) {
				edges = append(edges, edge)
			}
		}
		qd := edge.Sym()
		vd := qd.Origin()
		key = fmt.Sprintf("%f_%f", vd[0], vd[1])
		if _, found := visitedVertices[key]; !found {
			visitedVertices[key] = struct{}{}
			if includeFrame || !q.IsFrameVertex(vd) {
				edges = append(edges, qd)
			}
		}
	}

	return edges
}

func (q *Subdivision) locateFromEdge(v matrix.Matrix, startEdge *QuadEdge) (*QuadEdge, error) {
	var (
		iter    = 0
		maxIter = len(q.quadEdgeList)
		qe      = startEdge
	)

	for {
		iter++

		if iter > maxIter {
			return nil, errors.New("locate failed")
		}

		if (v.Equals(qe.Origin())) || (v.Equals(qe.Destination())) {
			break
		} else if IsCCW(v, qe.Destination(), qe.Origin()) {
			qe = qe.Sym()
		} else if !IsCCW(v, qe.ONext().Destination(), qe.ONext().Origin()) {
			qe = qe.ONext()
		} else if !IsCCW(v, qe.DPrev().Destination(), qe.DPrev().Origin()) {
			qe = qe.DPrev()
		} else {
			// on edge or in triangle containing edge
			break
		}
	}
	return qe, nil
}

func (q *Subdivision) visitTriangles(triVisitor TriangleVisitor, includeFrame bool) {
	q.visitedKey++
	var (
		edgeStack    = utils.NewStack()
		visitedEdges = make(map[*QuadEdge]struct{})
	)
	edgeStack.Push(q.startingEdge)

	for !edgeStack.Empty() {
		edge := edgeStack.Pop().(*QuadEdge)
		if _, found := visitedEdges[edge]; !found {
			triEdges := q.fetchTriangleToVisit(edge, edgeStack, includeFrame, visitedEdges)
			if triEdges != nil {
				triVisitor.Visit(triEdges)
			}
		}
	}
}

func (q *Subdivision) fetchTriangleToVisit(edge *QuadEdge, edgeStack *utils.Stack, includeFrame bool,
	visitedEdges map[*QuadEdge]struct{}) []*QuadEdge {
	var (
		curr      = edge
		edgeCount = 0
		isFrame   = false
		done      = false
	)
	for !done || curr != edge {
		if edgeCount >= len(q.triEdges) {
			break
		}
		q.triEdges[edgeCount] = curr

		if q.IsFrameEdge(curr) {
			isFrame = true
		}

		sym := curr.Sym()
		if _, found := visitedEdges[sym]; !found {
			edgeStack.Push(sym)
		}

		visitedEdges[curr] = struct{}{}

		edgeCount++
		curr = curr.LNext()
		done = true
	}

	if isFrame && !includeFrame {
		return nil
	}
	return q.triEdges

}

// GetVoronoiCellPolygons ...
func (q *Subdivision) GetVoronoiCellPolygons() []matrix.PolygonMatrix {
	q.visitTriangles(&TriangleCircumcentreVisitor{}, true)

	var (
		cells = make([]matrix.PolygonMatrix, 0)
		edges = q.getVertexUniqueEdges(false)
	)
	for _, edge := range edges {
		cells = append(cells, getVoronoiCellPolygon(edge))
	}
	return cells
}

func getVoronoiCellPolygon(qe *QuadEdge) matrix.PolygonMatrix {
	var (
		startQE    = qe
		lineMatrix = matrix.LineMatrix{}
		done       = false
	)
	for !done || qe != startQE {
		// use previously computed circumcentre
		cc := qe.rot.Origin()
		lineMatrix = append(lineMatrix, cc)
		// move to next triangle CW around vertex
		qe = qe.OPrev()
		done = true
	}
	if len(lineMatrix) < 4 {
		lineMatrix = append(lineMatrix, lineMatrix[len(lineMatrix)-1])
	}
	if !lineMatrix.IsClosed() {
		lineMatrix = append(lineMatrix, lineMatrix[0])
	}
	cellPoly := matrix.PolygonMatrix{lineMatrix}
	return cellPoly
}

// Locate ...
func (q *Subdivision) Locate(v matrix.Matrix) *QuadEdge {
	return q.locator.locate(v)
}

// Edges ...
func (q *Subdivision) Edges() []*QuadEdge {
	return q.quadEdgeList
}

// SetEdges ...
func (q *Subdivision) SetEdges(edges []*QuadEdge) {
	q.quadEdgeList = edges
}
