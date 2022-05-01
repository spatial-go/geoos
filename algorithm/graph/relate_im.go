// package graph ...

package graph

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// Relatelgorithm  the entity be used during the relate computation.
type Relatelgorithm func(arg []matrix.Steric) Relationship

// Relationship  be used during the relate computation.
type Relationship interface {
	ComputeIM() *matrix.IntersectionMatrix
}

// GetRelationship returns  algorithm by new Algorithm.
func GetRelationship(f Relatelgorithm, arg []matrix.Steric) Relationship {
	return f(arg)
}

func withDegrees(arg []matrix.Steric) Relationship {
	return &RelationshipByDegrees{
		Arg:           arg,
		graph:         []Graph{&MatrixGraph{}, &MatrixGraph{}},
		gIntersection: &MatrixGraph{},
		gUnion:        &MatrixGraph{},
		IM:            matrix.IntersectionMatrixDefault(),
		IsClosed:      []bool{false, false},
	}
}

func withStructure(arg []matrix.Steric) Relationship {
	return &RelationshipByStructure{
		Arg:           arg,
		graph:         []Graph{&MatrixGraph{}, &MatrixGraph{}},
		gIntersection: &MatrixGraph{},
		gUnion:        &MatrixGraph{},
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

// IMByRelationship Gets the relate  for the spatial relationship
// between the input geometries.
func IMByRelationship(m0, m1 matrix.Steric, f Relatelgorithm) *matrix.IntersectionMatrix {
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
