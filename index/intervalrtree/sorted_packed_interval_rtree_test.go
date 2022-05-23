// Package intervalrtree Contains structs to implement an R-tree index for one-dimensional intervals.
package intervalrtree

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

var indexTree *SortedPackedIntervalRTree

func TestMain(m *testing.M) {

	fmt.Println("test start")
	buildTree()
	code := m.Run()
	os.Exit(code)
	fmt.Println("test end")
}
func buildTree() *SortedPackedIntervalRTree {
	indexTree = &SortedPackedIntervalRTree{}
	var ms matrix.Collection = matrix.Collection{
		matrix.Matrix{1, 1},
		matrix.Matrix{1, 1},
		matrix.Matrix{2, 1},
		matrix.Matrix{2, 2},
		matrix.Matrix{3, 1},
		matrix.Matrix{3, 2},
	}
	for i := 0; i < len(ms); i++ {
		if err := indexTree.Insert(envelope.Matrix(ms[i].(matrix.Matrix)), ms[i]); err != nil {
			log.Println(err)
		}
	}
	return indexTree
}
func TestSortedPackedIntervalRTree_Insert(t *testing.T) {
	type args struct {
		queryEnv *envelope.Envelope
		item     interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"insert：", args{envelope.Matrix(matrix.Matrix{1, 8}), matrix.Matrix{1, 8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := indexTree
			if err := s.Insert(tt.args.queryEnv, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("SortedPackedIntervalRTree.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSortedPackedIntervalRTree_QueryVisitor(t *testing.T) {
	type args struct {
		queryEnv *envelope.Envelope
		visitor  index.ItemVisitor
	}
	visitor := &index.ArrayVisitor{}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"query", args{envelope.Matrix(matrix.Matrix{1, 8}), visitor}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := indexTree
			if err := s.QueryVisitor(tt.args.queryEnv, tt.args.visitor); (err != nil) != tt.wantErr {
				t.Errorf("SortedPackedIntervalRTree.QueryVisitor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
