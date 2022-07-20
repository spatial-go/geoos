package graph

import (
	"github.com/spatial-go/geoos/algorithm/calc"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/overlay/chain"
)

// IntersectLine returns a Collection  that represents that part of geometry A intersect with geometry B.
func IntersectLine(m1, m2 matrix.LineMatrix) chain.CorrelationNodeResult {
	smi := &chain.SegmentMutualIntersector{SegmentMutual: m1}
	icd := &chain.IntersectionCorrelation{Edge: m1, Edge1: m2}
	smi.Process(m2, icd)
	inols := icd.Result()

	m := []matrix.LineMatrix{m1, m2}
	if inolsType, ok := inols.([]chain.IntersectionNodeOfLine); ok {
		result := make(chain.CorrelationNodeResult, len(inolsType))
		for i, inol := range inolsType {
			inol = sortIntersectionNode(inol)
			resultCorr := createCorrelationNode(inol, m[i])
			result[i] = resultCorr
		}
		return result
	}
	return nil
}

// createCorrelationNode handle intersection lines.
func createCorrelationNode(result chain.IntersectionNodeOfLine, edge matrix.LineMatrix) []*chain.IntersectionCorrelationNode {
	lines := []matrix.LineMatrix{}
	totalIps := 0
	ips := chain.IntersectionNodeOfLine{}
	pos := [2]int{0, 0}

	correlationNode := matrix.LineMatrix{}

	var endNode *chain.IntersectionNodeResult

	for i := 0; i < len(result); i++ {
		r0 := result[i]
		pos[1] = r0.Pos
		endNode = r0

		lines, ips = writeLineNode(pos, correlationNode, edge, endNode, lines, ips)
		totalIps++
		pos[0] = endNode.End
		correlationNode = matrix.LineMatrix{}
		correlationNode = append(correlationNode, endNode.InterNode.Matrix)

		if i == result.Len()-1 {
			for j := pos[0]; j < len(edge); j++ {
				if len(correlationNode) == 0 ||
					!matrix.Matrix(correlationNode[len(correlationNode)-1]).
						EqualsExact(matrix.Matrix(edge[j]), calc.DefaultTolerance) {
					correlationNode = append(correlationNode, edge[j])
				}
			}
			lines = append(lines, correlationNode)
		}
	}

	resultCorr := writeIntersectionCorrelationNode(ips, lines)
	return resultCorr
}

func writeIntersectionCorrelationNode(ips chain.IntersectionNodeOfLine, lines []matrix.LineMatrix) []*chain.IntersectionCorrelationNode {
	resultCorr := []*chain.IntersectionCorrelationNode{}
	ip := &chain.IntersectionNodeResult{}

	for i := 0; i < len(ips); i++ {

		if i > 0 && ips[i].InterNode.Matrix.
			EqualsExact(ips[i].Line.P1, calc.DefaultTolerance) &&
			i < len(ips)-1 {
			if ips[i+1].Pos == ips[i].End && ips[i+1].InterNode.IsCollinear {
				lines[i+1] = append(lines[i], lines[i+1]...)
				lines[i+1] = lines[i+1].Filter(&matrix.UniqueArrayFilter{}).(matrix.LineMatrix)
				if i > 0 && ip.InterNode.Matrix == nil {
					ip = ips[i-1]
				}
				continue
			}
		}
		if i > 0 {
			if !matrix.Matrix(lines[i][0]).
				EqualsExact(matrix.Matrix(lines[i][len(lines[i])-1]), calc.DefaultTolerance) {
				if ip.InterNode.Matrix == nil {
					ip = ips[i-1]
				}
				resultCorr = append(resultCorr,
					&chain.IntersectionCorrelationNode{InterNode: ip.InterNode.Matrix, CorrelationNode: lines[i]})
				ip = &chain.IntersectionNodeResult{}
			}
		}
		if !matrix.Matrix(lines[i][0]).
			EqualsExact(matrix.Matrix(lines[i][len(lines[i])-1]), calc.DefaultTolerance) {
			resultCorr = append(resultCorr,
				&chain.IntersectionCorrelationNode{InterNode: ips[i].InterNode.Matrix, CorrelationNode: lines[i]})
		}

		if i == len(ips)-1 {
			if !matrix.Matrix(lines[i+1][0]).
				EqualsExact(matrix.Matrix(lines[i+1][len(lines[i+1])-1]), calc.DefaultTolerance) {
				resultCorr = append(resultCorr,
					&chain.IntersectionCorrelationNode{InterNode: ips[i].InterNode.Matrix, CorrelationNode: lines[i+1]})
			}
		}
	}
	if len(ips) == 1 && len(resultCorr) == 0 {
		resultCorr = append(resultCorr,
			&chain.IntersectionCorrelationNode{InterNode: ips[0].InterNode.Matrix, CorrelationNode: matrix.LineMatrix{}})
	}
	return resultCorr
}

func writeLineNode(pos [2]int, correlationNode, edge matrix.LineMatrix, endNode *chain.IntersectionNodeResult,
	lines []matrix.LineMatrix, ips chain.IntersectionNodeOfLine) ([]matrix.LineMatrix, chain.IntersectionNodeOfLine) {
	for j := pos[0]; j <= pos[1]; j++ {
		if len(correlationNode) == 0 ||
			!matrix.Matrix(correlationNode[len(correlationNode)-1]).
				EqualsExact(matrix.Matrix(edge[j]), calc.DefaultTolerance) {
			correlationNode = append(correlationNode, edge[j])
		}
	}
	if len(correlationNode) == 1 ||
		!matrix.Matrix(correlationNode[len(correlationNode)-1]).
			EqualsExact(endNode.InterNode.Matrix, calc.DefaultTolerance) {
		correlationNode = append(correlationNode, endNode.InterNode.Matrix)
	}
	lines = append(lines, correlationNode)
	ips = append(ips, endNode)

	return lines, ips
}
