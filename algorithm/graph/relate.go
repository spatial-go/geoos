// package graph ...

package graph

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/relate"
)

// relation Symbol
const (
	Disjoint = iota + 1
	Touch
	Cross
	Within
	Overlap
	Equal
	Contain
	Cover
	CoveredBy
)

// edge cost
const (
	PointPoint = PNode + PNode
	PointLine  = PNode + LNode
	PointCLine = PNode + CNode
	PointPoly  = PNode + ANode

	LineLine  = LNode + LNode
	LineCLine = LNode + CNode
	LinePoly  = LNode + ANode
)

// RelationshipByStructure  be used during the relate computation.
type RelationshipByStructure struct {
	// The operation args into an array so they can be accessed by index
	Arg                    []matrix.Steric // the arg(s) of the operation
	graph                  []Graph
	gIntersection, gUnion  Graph
	IM                     *matrix.IntersectionMatrix
	relationshipSymbol     int
	maxDlPoint, sumDlPoint int
	maxDlLine              int
	degrees                []int
	haveIntersectionVertex []int

	//nPoint number of point node, nLine number of line node, nCompositeLine number of CompositeLine node,
	nPoint, nLine, nCompositeLine int
	IsClosed                      []bool
}

// ComputeIM IntersectionMatrix Gets the IntersectionMatrix for the spatial relationship
// between the input geometries.
func (r *RelationshipByStructure) ComputeIM() *matrix.IntersectionMatrix {
	for i, v := range r.Arg {
		r.graph[i], _ = GenerateGraph(v)
	}

	if err := IntersectionHandle(r.Arg[0], r.Arg[1], r.graph[0], r.graph[1]); err != nil {
		return r.IM
	}

	var err error
	if r.gIntersection, err = r.graph[0].Intersection(r.graph[1]); err != nil {
		return r.IM
	}
	if r.gUnion, err = r.graph[0].Union(r.graph[1]); err != nil {
		return r.IM
	}

	relateType := 0
	if r.Arg[0].Equals(r.Arg[1]) {
		relateType = Equal
	} else {
		if r.gIntersection.Order() == 0 {
			relateType = Disjoint
		} else {
			r.degrees = make([]int, r.gIntersection.Order())
			r.haveIntersectionVertex = []int{0, 0}
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

	r.IM.SetString(RelateStringsTransposeByRing(RelateStrings[r.relationshipSymbol], inputRing))
	return r.IM
}

func (r *RelationshipByStructure) handleNode() {

	for i, n := range r.gIntersection.Nodes() {
		r.degrees[i] = r.gUnion.Degree(n.Index)
		if n.NodeType == PNode {
			r.nPoint++
			for j, v := range r.Arg {
				if boundary, err := v.Boundary(); err == nil {
					for _, b := range boundary.(matrix.Collection) {
						if n.Value.Equals(b) || n.Reverse.Equals(b) {
							r.haveIntersectionVertex[j]++
						}
					}
				}
			}
			indexUnion, _ := r.gUnion.NodeIndex(n)
			dl := 0
			for _, v := range r.gUnion.Edges()[indexUnion] {
				if v == PointLine {
					dl++
				}
			}
			if dl > r.maxDlPoint {
				r.maxDlPoint = dl
			}
			r.sumDlPoint += dl
		}
		if n.NodeType == LNode {
			r.nLine++
			indexUnion, _ := r.gUnion.NodeIndex(n)
			var pIndex []int
			var maxDlPoints = []int{0, 0}

			for k, v := range r.gUnion.Edges()[indexUnion] {
				if v == PointLine {
					pIndex = append(pIndex, k)
				}
			}
			for j, index := range pIndex {
				for _, v := range r.gUnion.Edges()[index] {
					if v == PointLine {
						maxDlPoints[j]++
					}
				}
			}

			dl := 0
			switch {
			case maxDlPoints[0]+maxDlPoints[1] == 2:
				dl = 1
			case maxDlPoints[0]+maxDlPoints[1] == 3:
				dl = 2
			case maxDlPoints[0] == 2 && maxDlPoints[1] == 2:
				dl = 4
			case maxDlPoints[0]+maxDlPoints[1] == 4:
				dl = 3
			case maxDlPoints[0]+maxDlPoints[1] == 5:
				dl = 5
			case maxDlPoints[0]+maxDlPoints[1] == 6:
				dl = 6
			}

			if r.maxDlLine < dl {
				r.maxDlLine = dl
			}
		}
		if n.NodeType == CNode {
			r.nCompositeLine++
		}
	}
}

func (r *RelationshipByStructure) matrixIM(p matrix.Matrix, RelateType int) {
	switch RelateType {
	case Equal:
		r.relationshipSymbol = RPP2
	case Disjoint:
		switch m := r.Arg[1].(type) {
		case matrix.Matrix:
			r.relationshipSymbol = RPP1
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			r.relationshipSymbol = RPL1
		case matrix.PolygonMatrix:
			pointInPolygon, _ := IsInPolygon(p, m)
			if pointInPolygon == OnlyInPolygon {
				r.relationshipSymbol = RPA2
			} else {
				r.relationshipSymbol = RPA1
			}
		}
	default:
		switch m := r.Arg[1].(type) {
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			if r.degrees[0] >= 2 {
				r.relationshipSymbol = RPL2
			} else {
				r.relationshipSymbol = RPL3
			}
		case matrix.PolygonMatrix:
			r.relationshipSymbol = RPA3
		}
	}
}

func (r *RelationshipByStructure) lineIM(l matrix.LineMatrix, RelateType int) {
	switch RelateType {
	case Equal:
		r.relationshipSymbol = RLL25
	case Disjoint:
		switch m := r.Arg[1].(type) {
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			r.relationshipSymbol = RLL1
		case matrix.PolygonMatrix:
			pointInPolygon, _ := IsInPolygon(l, m)
			if pointInPolygon == OnlyInPolygon {
				r.relationshipSymbol = RLA2
			} else {
				r.relationshipSymbol = RLA1
			}
		}
	default:
		switch m := r.Arg[1].(type) {
		case matrix.LineMatrix:
			if m.IsClosed() {
				r.IsClosed[1] = true
			}
			r.lineAnalyse(DefaultInPolygon, DefaultInPolygon)
		case matrix.PolygonMatrix:
			pointInPolygon, entityInPolygon := IsInPolygon(l, m)
			r.lineAnalyse(pointInPolygon, entityInPolygon)
		}

	}
}

func (r *RelationshipByStructure) polygonIM(p matrix.PolygonMatrix, RelateType int) {
	switch RelateType {
	case Equal:
		r.relationshipSymbol = RAA9
	case Disjoint:
		switch r.Arg[1].(type) {
		case matrix.PolygonMatrix:
			poly := r.Arg[1].(matrix.PolygonMatrix)
			if relate.InPolygon(p[0][0], poly[0]) {
				r.relationshipSymbol = RAA6
				for i := 1; i < len(poly); i++ {
					if relate.InPolygon(p[0][0], poly[i]) {
						r.relationshipSymbol = RAA1
					}
				}
				return
			}
			if relate.InPolygon(poly[0][0], p[0]) {
				r.relationshipSymbol = RAA5
				for i := 1; i < len(p); i++ {
					if relate.InPolygon(poly[0][0], p[i]) {
						r.relationshipSymbol = RAA1
					}
				}
				return
			}
			r.relationshipSymbol = RAA1
		}
	default:
		pointInPolygon, entityInPolygon := IsInPolygon(p, r.Arg[1].(matrix.PolygonMatrix))
		r.polygonAnalyse(pointInPolygon, entityInPolygon)
	}
}
