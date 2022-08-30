package de9im

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// RelationshipByDegrees  be used during the relate computation.
type RelationshipByDegrees struct {
	// The operation args into an array so they can be accessed by index
	*graph.Clip
	gIntersection, gUnion  graph.Graph
	IM                     *matrix.IntersectionMatrix
	degrees                []int
	haveIntersectionVertex []int
	boundary               []matrix.Collection

	//nPoint number of point node, nLine number of line node
	nPoint, nLine int
	IsClosed      []bool
}

// ComputeIM IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (r *RelationshipByDegrees) ComputeIM() *matrix.IntersectionMatrix {

	var err error
	if r.gIntersection, err = r.Intersection(); err != nil {
		return r.IM
	}
	if r.gUnion, err = r.Union(); err != nil {
		return r.IM
	}

	relateType := 0
	if r.Arg[0].Proximity(r.Arg[1]) {
		relateType = Equal
	} else {
		if r.gIntersection.Order() == 0 {
			relateType = Disjoint
		} else {
			r.degrees = make([]int, r.gIntersection.Order())
			r.haveIntersectionVertex = []int{0, 0}
			r.boundary = make([]matrix.Collection, 2)
			for j, v := range r.Arg {
				if boundary, err := v.Boundary(); err == nil {
					r.boundary[j] = boundary.(matrix.Collection)
				}
			}
			r.handleNode()
		}
	}

	switch m := r.Arg[0].(type) {
	case matrix.Matrix:
		r.matrixIM(m, relateType)
	case matrix.LineMatrix:
		if m.IsClosed() {
			r.IsClosed[0] = true
		}
		r.lineIM(m, relateType)
	case matrix.PolygonMatrix:
		r.polygonIM(m, relateType)
	}
	inputRing := -1
	switch {
	case r.IsClosed[0] && r.IsClosed[1]:
		inputRing = 3
	case r.IsClosed[0]:
		inputRing = 1
	case r.IsClosed[1]:
		inputRing = 2
	}

	IMTransposeByRing(r.IM, inputRing)
	return r.IM
}

func (r *RelationshipByDegrees) handleNode() {

	for i, n := range r.gIntersection.Nodes() {
		if n.NodeType == graph.PNode {
			imNode, _ := r.gUnion.Node(n)
			r.degrees[i] = r.gUnion.Degree(imNode.Index)
			r.nPoint++
			for j, v := range r.boundary {
				for _, b := range v {
					if n.Value.EqualsExact(b, calc.DefaultTolerance) || n.Reverse.EqualsExact(b, calc.DefaultTolerance) {
						r.haveIntersectionVertex[j]++
					}
				}
			}
		}

		if n.NodeType == graph.LNode {
			r.nLine++
		}
	}
}

func (r *RelationshipByDegrees) matrixIM(p matrix.Matrix, RelateType int) {
	switch RelateType {
	case Equal:
		r.IM.SetString(RelateStrings[RPP2])
	case Disjoint:
		switch m := r.Arg[1].(type) {
		case matrix.Matrix:
			r.IM.SetString(RelateStrings[RPP1])
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			r.IM.SetString(RelateStrings[RPL1])
		case matrix.PolygonMatrix:
			pointInPolygon, _ := IsInPolygon(p, m)
			if pointInPolygon == OnlyInPolygon {
				r.IM.SetString(RelateStrings[RPA2])
			} else {
				r.IM.SetString(RelateStrings[RPA1])
			}
		}
	default:
		switch m := r.Arg[1].(type) {
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			if r.degrees[0] >= 2 {
				r.IM.SetString(RelateStrings[RPL2])
			} else {
				r.IM.SetString(RelateStrings[RPL3])
			}
		case matrix.PolygonMatrix:
			r.IM.SetString(RelateStrings[RPA3])
		}
	}
}

func (r *RelationshipByDegrees) lineIM(l matrix.LineMatrix, RelateType int) {
	switch RelateType {
	case Equal:
		r.IM.SetString(RelateStrings[RLL25])
	case Disjoint:
		switch m := r.Arg[1].(type) {
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			r.IM.SetString(RelateStrings[RLL1])
		case matrix.PolygonMatrix:
			pointInPolygon, _ := IsInPolygon(l, m)
			if pointInPolygon == OnlyInPolygon {
				r.IM.SetString(RelateStrings[RLA2])
			} else {
				r.IM.SetString(RelateStrings[RLA1])
			}
		}
	default:
		switch m := r.Arg[1].(type) {
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			r.lineAndLineAnalyse(DefaultInPolygon, DefaultInPolygon)
		case matrix.PolygonMatrix:
			pointInPolygon, entityInPolygon := IsInPolygon(l, m)
			lrd := &LineRelationshipByDegrees{r, pointInPolygon, entityInPolygon}
			produce(lrd)
		}

	}
}

func (r *RelationshipByDegrees) polygonIM(p matrix.PolygonMatrix, RelateType int) {
	switch RelateType {
	case Equal:
		r.IM.SetString(RelateStrings[RAA9])
	case Disjoint:
		switch r.Arg[1].(type) {
		case matrix.PolygonMatrix:
			poly := r.Arg[1].(matrix.PolygonMatrix)
			if relate.InPolygon(p[0][0], poly[0]) {
				r.IM.SetString(RelateStrings[RAA6])
				for i := 1; i < len(poly); i++ {
					if relate.InPolygon(p[0][0], poly[i]) {
						r.IM.SetString(RelateStrings[RAA1])
					}
				}
				return
			}
			if relate.InPolygon(poly[0][0], p[0]) {
				r.IM.SetString(RelateStrings[RAA5])
				for i := 1; i < len(p); i++ {
					if relate.InPolygon(poly[0][0], p[i]) {
						r.IM.SetString(RelateStrings[RAA1])
					}
				}
				return
			}
			r.IM.SetString(RelateStrings[RAA1])
		}
	default:
		_, entityInPolygon := IsInPolygon(p, r.Arg[1].(matrix.PolygonMatrix))
		prd := &PolygonRelationshipByDegrees{r, entityInPolygon}
		produce(prd)
		// pointInPolygon, entityInPolygon := IsInPolygon(p, r.Arg[1].(matrix.PolygonMatrix))
		// r.polygonTwoAnalyse(pointInPolygon, entityInPolygon)
	}
}

type producingIM interface {
	setInteriorIM()
	setExteriorIM()
	setBoundaryIM()
	produce()
}

func produce(l producingIM) {
	l.produce()
}
