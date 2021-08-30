package quadtree

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

var indexTree *Quadtree
var lineMatrix = matrix.LineMatrix{{1, 1}, {1, 5}, {5, 5}, {10, 8}, {8, 1}, {2, 3}}

func TestMain(m *testing.M) {
	indexTree = NewQuadtree()
	segs := lineMatrix.ToLineArray()
	for i := 0; i < len(segs); i++ {
		seg := segs[i]
		indexTree.Insert(envelope.TwoMatrix(seg.P0, seg.P1), seg)
	}
	fmt.Println("test start")
	code := m.Run()
	os.Exit(code)
	fmt.Println("test end")
}
func TestQuadtree_Insert(t *testing.T) {
	line := matrix.LineMatrix{{5, 5}, {10, 8}}
	seg := line.ToLineArray()[0]
	type args struct {
		itemEnv *envelope.Envelope
		item    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "quadtree insert", args: args{envelope.TwoMatrix(seg.P0, seg.P1), seg}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("%s ： %s", tt.name, err)
				}
			}()
			q := indexTree
			q.Insert(tt.args.itemEnv, tt.args.item)
		})
	}
}

func TestQuadtree_Remove(t *testing.T) {
	seg := lineMatrix.ToLineArray()[0]
	type args struct {
		itemEnv *envelope.Envelope
		item    interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "quadtree remove", args: args{envelope.TwoMatrix(seg.P0, seg.P1), seg}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := indexTree
			if got := q.Remove(tt.args.itemEnv, tt.args.item); got != tt.want {
				t.Errorf("Quadtree.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuadtree_Query(t *testing.T) {
	seg := lineMatrix.ToLineArray()[1]
	type args struct {
		searchEnv *envelope.Envelope
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "quadtree query", args: args{envelope.TwoMatrix(seg.P0, seg.P1)}, want: &matrix.LineSegment{P0: seg.P0, P1: seg.P1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := indexTree
			got := q.Query(tt.args.searchEnv)
			has := false
			for _, v := range got.([]interface{}) {
				if reflect.DeepEqual(v, tt.want) {
					has = true
				}
			}
			if !has {
				t.Errorf("Quadtree.Query() = %v", got)
				t.Errorf("Quadtree.Query  want %v", tt.want)
			}
		})
	}
}
