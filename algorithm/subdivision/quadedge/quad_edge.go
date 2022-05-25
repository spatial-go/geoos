// Package quadedge a topological subdivision of quadeges, to support creating triangulations and Voronoi diagrams.
package quadedge

import (
	"fmt"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

// IsCCW ...
func IsCCW(a matrix.Matrix, b matrix.Matrix, c matrix.Matrix) bool {
	return (b[0]-a[0])*(c[1]-a[1])-(b[1]-a[1])*(c[0]-a[0]) > 0
}

// IsInCircle ...
func IsInCircle(v matrix.Matrix, a matrix.Matrix, b matrix.Matrix, c matrix.Matrix) bool {
	return isInCircleRobust(a, b, c, v)
}

// QuadEdge ...
type QuadEdge struct {
	vertex matrix.Matrix
	next   *QuadEdge
	rot    *QuadEdge
}

// EmptyQuadEdge ...
func EmptyQuadEdge() *QuadEdge {
	return &QuadEdge{}
}

// NewQuadEdge return a QuadEdge
func NewQuadEdge(o matrix.Matrix, d matrix.Matrix) *QuadEdge {
	var (
		q0 = EmptyQuadEdge()
		q1 = EmptyQuadEdge()
		q2 = EmptyQuadEdge()
		q3 = EmptyQuadEdge()
	)

	q0.rot = q1
	q1.rot = q2
	q2.rot = q3
	q3.rot = q0

	q0.next = q0
	q1.next = q3
	q2.next = q2
	q3.next = q1

	qe := q0
	qe.vertex = o
	qe.Sym().vertex = d
	return qe
}

// Connect ...
func Connect(a, b *QuadEdge) *QuadEdge {
	e := NewQuadEdge(a.Destination(), b.Origin())
	Splice(e, a.LNext())
	Splice(e.Sym(), b)
	return e
}

// Splice splice two edges together or apart
func Splice(a *QuadEdge, b *QuadEdge) {
	var (
		alpha = a.ONext().rot
		beta  = b.ONext().rot
		t1    = b.ONext()
		t2    = a.ONext()
		t3    = beta.ONext()
		t4    = alpha.ONext()
	)
	a.next = t1
	b.next = t2
	alpha.next = t3
	beta.next = t4
}

// Swap ....
func Swap(e *QuadEdge) {
	var (
		a = e.OPrev()
		b = e.Sym().OPrev()
	)
	Splice(e, a)
	Splice(e.Sym(), b)
	Splice(e, a.LNext())
	Splice(e.Sym(), b.LNext())

	e.SetOrigin(a.Destination())
	e.SetDestination(b.Destination())
}

// isLive whether this edge has been deleted
func (q *QuadEdge) isLive() bool {
	return q.rot != nil
}

// Origin ...
func (q *QuadEdge) Origin() matrix.Matrix {
	return q.vertex
}

// SetOrigin ...
func (q *QuadEdge) SetOrigin(v matrix.Matrix) {
	q.vertex = v
}

// Destination ...
func (q *QuadEdge) Destination() matrix.Matrix {
	return q.Sym().Origin()
}

// SetDestination ...
func (q *QuadEdge) SetDestination(v matrix.Matrix) {
	q.Sym().SetOrigin(v)
}

// InvRot return the inverse rotated edge.
func (q *QuadEdge) InvRot() *QuadEdge {
	return q.rot.Sym()
}

// Sym return the sym of the edge
func (q *QuadEdge) Sym() *QuadEdge {
	return q.rot.rot
}

// ONext return next edge
func (q *QuadEdge) ONext() *QuadEdge {
	return q.next
}

// OPrev return prev edge
func (q *QuadEdge) OPrev() *QuadEdge {
	return q.rot.next.rot
}

// DNext return next destination edge
func (q *QuadEdge) DNext() *QuadEdge {
	return q.Sym().ONext().Sym()
}

// DPrev return prev destination edge
func (q *QuadEdge) DPrev() *QuadEdge {
	return q.InvRot().ONext().InvRot()
}

// LNext return next left face edge
func (q *QuadEdge) LNext() *QuadEdge {
	return q.InvRot().ONext().rot
}

// LPrev return previous left face edge
func (q *QuadEdge) LPrev() *QuadEdge {
	return q.next.Sym()
}

// RNext return next right face edge
func (q *QuadEdge) RNext() *QuadEdge {
	return q.rot.next.InvRot()
}

// RPrev return previous right face edge
func (q *QuadEdge) RPrev() *QuadEdge {
	return q.Sym().ONext()
}

// Primary get the primary edge of this quadedge
func (q *QuadEdge) Primary() *QuadEdge {
	value, _ := q.Origin().Compare(q.Destination())
	if value <= 0 {
		return q
	}
	return q.Sym()

}

// ToString ...
func (q *QuadEdge) ToString() string {
	return fmt.Sprintf("%v, %v", q.Origin(), q.Destination())
}
