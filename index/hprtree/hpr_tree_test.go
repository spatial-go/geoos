package hprtree

import (
	"fmt"
	"os"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
	"github.com/spatial-go/geoos/index"
)

var indexTree *HPRTree

func TestMain(m *testing.M) {

	fmt.Println("test start")
	buildTree()
	code := m.Run()
	os.Exit(code)
	fmt.Println("test end")
}
func buildTree() *HPRTree {
	indexTree = NewHPRTree()
	var ms matrix.Collection = matrix.Collection{
		matrix.Matrix{1, 1},
		matrix.Matrix{1, 1},
		matrix.Matrix{2, 1},
		matrix.Matrix{2, 2},
		matrix.Matrix{3, 1},
		matrix.Matrix{3, 2},
	}
	for i := 0; i < len(ms); i++ {
		indexTree.Insert(envelope.Matrix(ms[i].(matrix.Matrix)), ms[i])
	}
	return indexTree
}

func TestHPRTree_Size(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"size:", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := indexTree
			if got := h.Size(); got != tt.want {
				t.Errorf("HPRTree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHPRTree_Insert(t *testing.T) {
	type args struct {
		itemEnv *envelope.Envelope
		item    interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"insert：", args{envelope.Matrix(matrix.Matrix{1, 8}), matrix.Matrix{1, 8}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := indexTree
			h.Insert(tt.args.itemEnv, tt.args.item)
			if got := h.Size(); got != tt.want {
				t.Errorf("HPRTree.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHPRTree_Query(t *testing.T) {

	type args struct {
		searchEnv *envelope.Envelope
	}
	tests := []struct {
		name string
		args args
	}{
		{"query", args{envelope.Matrix(matrix.Matrix{2, 2})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := indexTree
			if got := h.Query(tt.args.searchEnv); got == nil {
				t.Errorf("HPRTree.Query() = %v, want %v", got, "no nil")
			}
		})
	}
}

func TestHPRTree_QueryVisitor(t *testing.T) {
	type args struct {
		searchEnv *envelope.Envelope
		visitor   index.ItemVisitor
	}
	visitor := &index.ArrayVisitor{}
	tests := []struct {
		name string
		args args
	}{
		{"query", args{envelope.Matrix(matrix.Matrix{1, 8}), visitor}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := indexTree
			h.QueryVisitor(tt.args.searchEnv, tt.args.visitor)
			if len(tt.args.visitor.Items().([]interface{})) <= 0 {
				t.Errorf("HPRTree.Query() = %v, want %v", len(tt.args.visitor.Items().([]interface{})), ">0")
			}
		})
	}
}
