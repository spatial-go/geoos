package graph

import (
	"sort"

	"github.com/spatial-go/geoos/algorithm/filter"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
)

// IntersectmultiLines returns a Collection  that represents that part of geometry A intersect with geometry B.
func IntersectmultiLines(m1 matrix.LineMatrix, m2 []matrix.LineMatrix) chain.CorrelationNodeResult {
	smi := &chain.SegmentMutualIntersector{SegmentMutual: m1}
	intersectNodes := make([]chain.IntersectionNodeOfLine, len(m2)+1)
	for i, m := range m2 {
		icd := &chain.IntersectionCorrelation{Edge: m1, Edge1: m}
		smi.Process(m, icd)
		inols := icd.Result()

		if inolsType, ok := inols.([]chain.IntersectionNodeOfLine); ok {
			intersectNodes[0] = append(intersectNodes[0], inolsType[0]...)
			intersectNodes[i+1] = append(intersectNodes[i+1], inolsType[1]...)
		}
	}
	result := make(chain.CorrelationNodeResult, len(m2)+1)
	for i, v := range intersectNodes {
		v = sortIntersectionNode(v)
		m := m1
		if i > 0 {
			m = m2[i-1]
		}
		resultCorr := createCorrelationNode(v, m)
		result[i] = resultCorr

	}

	return result
}

func sortIntersectionNode(ins chain.IntersectionNodeOfLine) chain.IntersectionNodeOfLine {
	sort.Sort(ins)
	filter := IntersectionNodeUniqueArrayFilter{}
	for _, v := range ins {
		filter.Filter(v)
	}
	result := filter.Entities()
	return result.(chain.IntersectionNodeOfLine)
}

// IntersectLinePolygon returns a Collection  that represents that part of geometry A intersect with geometry B.
func IntersectLinePolygon(line matrix.LineMatrix, poly matrix.PolygonMatrix) chain.CorrelationNodeResult {
	multiLine := make([]matrix.LineMatrix, len(poly))
	for i, v := range poly {
		multiLine[i] = v
	}
	return IntersectmultiLines(line, multiLine)
}

// IntersectPolygons returns a Collection  that represents that part of geometry A intersect with geometry B.
func IntersectPolygons(poly1, poly2 matrix.PolygonMatrix) []chain.CorrelationNodeResult {
	intersectNodes := [][]chain.IntersectionNodeOfLine{
		make([]chain.IntersectionNodeOfLine, len(poly1)),
		make([]chain.IntersectionNodeOfLine, len(poly2)),
	}

	for i, m1 := range poly1 {
		smi := &chain.SegmentMutualIntersector{SegmentMutual: m1}
		for j, m2 := range poly2 {

			icd := &chain.IntersectionCorrelation{Edge: m1, Edge1: m2}
			smi.Process(m2, icd)
			inols := icd.Result()

			if inolsType, ok := inols.([]chain.IntersectionNodeOfLine); ok {
				intersectNodes[0][i] = append(intersectNodes[0][i], inolsType[0]...)
				intersectNodes[1][j] = append(intersectNodes[1][j], inolsType[1]...)
			}
		}
	}
	result := []chain.CorrelationNodeResult{
		make(chain.CorrelationNodeResult, len(poly1)),
		make(chain.CorrelationNodeResult, len(poly2)),
	}

	for i, nodes := range intersectNodes {
		poly := poly1
		if i > 0 {
			poly = poly2
		}
		for j, v := range nodes {
			v = sortIntersectionNode(v)
			resultCorr := createCorrelationNode(v, poly[j])

			result[i][j] = resultCorr
		}
	}

	return result
}

// IntersectionNodeUniqueArrayFilter  A Filter that extracts a unique array.
type IntersectionNodeUniqueArrayFilter struct {
	entities chain.IntersectionNodeOfLine
}

// Entities  Returns the gathered Matrixes.
func (u *IntersectionNodeUniqueArrayFilter) Entities() interface{} {
	return u.entities
}

// Filter Performs an operation with the provided .
func (u *IntersectionNodeUniqueArrayFilter) Filter(entity interface{}) bool {
	return u.add(entity)
}

func (u *IntersectionNodeUniqueArrayFilter) add(entity interface{}) bool {
	hasMatrix := false
	for _, v := range u.entities {
		if v.InterNode.Matrix.Equals(entity.(*chain.IntersectionNodeResult).InterNode.Matrix) {
			hasMatrix = true
			break
		}
	}
	if !hasMatrix {
		u.entities = append(u.entities, entity.(*chain.IntersectionNodeResult))
		return true
	}
	return false
}

// compile time checks
var (
	_ filter.Filter = &IntersectionNodeUniqueArrayFilter{}
)
