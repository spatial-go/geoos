// package de9im the  criteria  for  judging  the  various
// topological relations between points, line and surface entities in DE-9IM model are given.
// the method of graph operation and geometric calculation to deduce
// the specific topological relations among entities by analyzing the structure of points and line
// vertices  and  the  combination  of  points  and  lines
package de9im

import (
	"github.com/spatial-go/geoos/algorithm/graph"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// RelateAlgorithm  the entity be used during the relate computation.
type RelateAlgorithm func(arg []matrix.Steric) Relationship

// RelateClipAlgorithm  the entity be used during the relate computation.
type RelateClipAlgorithm func(clip *graph.Clip) Relationship

// Relationship  be used during the relate computation.
type Relationship interface {
	ComputeIM() *matrix.IntersectionMatrix
}

// GetRelationship returns  algorithm by new Algorithm.
func GetRelationship(f RelateAlgorithm, arg []matrix.Steric) Relationship {
	return f(arg)
}

// GetRelationshipByClip returns  algorithm by new Algorithm.
func GetRelationshipByClip(f RelateClipAlgorithm, clip *graph.Clip) Relationship {
	return f(clip)
}

func withDegreesAndClip(clip *graph.Clip) Relationship {
	return &RelationshipByDegrees{
		Clip:          clip,
		gIntersection: &graph.MatrixGraph{},
		gUnion:        &graph.MatrixGraph{},
		IM:            matrix.IntersectionMatrixDefault(),
		IsClosed:      []bool{false, false},
	}
}

func withDegrees(arg []matrix.Steric) Relationship {
	clip := graph.ClipHandle(arg[0], arg[1])
	return withDegreesAndClip(clip)
}

func withStructure(arg []matrix.Steric) Relationship {
	clip := graph.ClipHandle(arg[0], arg[1])
	return withStructureAndClip(clip)
}

func withStructureAndClip(clip *graph.Clip) Relationship {
	return &RelationshipByStructure{
		Clip:          clip,
		gIntersection: &graph.MatrixGraph{},
		gUnion:        &graph.MatrixGraph{},
		IM:            matrix.IntersectionMatrixDefault(),
		maxDlPoint:    0,
		sumDlPoint:    0,
		maxDlLine:     0,
		IsClosed:      []bool{false, false},
	}
}

// Relate Gets the relate string for the spatial relationship
// between the input geometries.
func Relate(m0, m1 matrix.Steric) string {

	im := IM(m0, m1)
	return im.ToString()
}

// IM Gets the relate  for the spatial relationship
// between the input geometries.
func IM(m0, m1 matrix.Steric) *matrix.IntersectionMatrix {
	return IMByRelationship(m0, m1, withDegrees)
}

// IMByClip Gets the relate  for the spatial relationship
// between the input geometries.
func IMByClip(clip *graph.Clip) *matrix.IntersectionMatrix {
	rs := GetRelationshipByClip(withDegreesAndClip, clip)
	im := rs.ComputeIM()
	return im
}

// IM Gets the relate  for the spatial relationship
// between the input geometries.
func IMStructure(m0, m1 matrix.Steric) *matrix.IntersectionMatrix {
	return IMByRelationship(m0, m1, withStructure)
}

// IMByRelationship Gets the relate  for the spatial relationship
// between the input geometries.
func IMByRelationship(m0, m1 matrix.Steric, f RelateAlgorithm) *matrix.IntersectionMatrix {
	arg := []matrix.Steric{m0, m1}
	if m0.Dimensions() > m1.Dimensions() {
		arg = []matrix.Steric{m1, m0}
	}
	rs := GetRelationship(f, arg)
	im := rs.ComputeIM()
	if m0.Dimensions() > m1.Dimensions() {
		im = im.Transpose()
	}
	return im
}
