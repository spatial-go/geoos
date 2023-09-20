// Package sweepline Contains struct which implement a sweepline algorithm for scanning geometric data structures.
package sweepline

import (
	"fmt"
	"os"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

var index = &Index{}

func TestMain(m *testing.M) {

	fmt.Println("test start")
	buildTree()
	code := m.Run()
	fmt.Println("test end")
	os.Exit(code)
}
func buildTree() *Index {
	var ms matrix.LineMatrix = matrix.LineMatrix{{1, 1}, {1.5, 1}, {2, 1}, {2, 2}, {2, 3}, {3, 3}}

	for _, line := range ms.ToLineArray() {
		index.Add(&Interval{line.P0[0], line.P1[0], line})
	}
	return index
}
func TestIndex_Add(t *testing.T) {
	type args struct {
		sweepInt *Interval
	}
	matr := &matrix.LineSegment{P0: matrix.Matrix{1, 2}, P1: matrix.Matrix{1, 3}}
	tests := []struct {
		name string
		args args
	}{
		{"add:", args{&Interval{matr.P0[0], matr.P1[0], matr}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := index
			s.Add(tt.args.sweepInt)
			if index.nOverlaps != 0 {
				t.Errorf("%v = %v, want %v", tt.name, index.nOverlaps, 0)
			}
		})
	}
}

func TestIndex_ComputeOverlaps(t *testing.T) {

	type args struct {
		action OverlapAction
	}
	tests := []struct {
		name string
		args args
	}{
		{"ComputeOverlaps:", args{&CoordinatesOverlapAction{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := index
			s.ComputeOverlaps(tt.args.action)
			if index.nOverlaps == 0 {
				t.Errorf("%v = %v, want %v", tt.name, index.nOverlaps, "no zero")
			}
		})
	}
}
