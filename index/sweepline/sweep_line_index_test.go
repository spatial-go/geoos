//Package sweepline Contains struct which implement a sweepline algorithm for scanning geometric data structures.
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
	os.Exit(code)
	fmt.Println("test end")
}
func buildTree() *Index {
	var ms matrix.Collection = matrix.Collection{
		matrix.Matrix{1, 1},
		matrix.Matrix{1, 1},
		matrix.Matrix{2, 1},
		matrix.Matrix{2, 2},
		matrix.Matrix{3, 1},
		matrix.Matrix{3, 2},
	}
	for i := 0; i < len(ms); i++ {
		index.Add(&Interval{ms[i].(matrix.Matrix).Bound()[0][0], ms[i].(matrix.Matrix).Bound()[0][1], ms[i]})
	}
	return index
}
func TestIndex_Add(t *testing.T) {
	type args struct {
		sweepInt *Interval
	}
	matr := matrix.Matrix{1, 2}
	tests := []struct {
		name string
		args args
	}{
		{"add:", args{&Interval{matr.Bound()[0][0], matr.Bound()[0][1], matr}}},
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
