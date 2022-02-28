// package graph ...
package graph

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestIntersectionHandle(t *testing.T) {
	type args struct {
		name   string
		m1     matrix.Steric
		m2     matrix.Steric
		g1, g2 Graph
	}
	type graphs struct {
		g1 Graph
		g2 Graph
	}
	testMatrix := []args{
		{"matrix matrix", matrix.Matrix{100, 100}, matrix.Matrix{100, 100},
			&MatrixGraph{nodes: []*Node{{Index: 0, Value: matrix.Matrix{100, 100}, stat: true, inGraph: true, NodeType: PNode}}, edges: []map[int]int{{}}},
			&MatrixGraph{nodes: []*Node{{Index: 0, Value: matrix.Matrix{100, 100}, stat: true, inGraph: true, NodeType: PNode}}, edges: []map[int]int{{}}},
		},
		{"matrix line", matrix.Matrix{100, 100}, matrix.LineMatrix{{100, 100}, {100, 200}},
			&MatrixGraph{nodes: []*Node{{Index: 0, Value: matrix.Matrix{100, 100}, stat: true, inGraph: true, NodeType: PNode}}, edges: []map[int]int{{}}},
			&MatrixGraph{nodes: []*Node{{Index: 0, stat: true, inGraph: true, NodeType: LNode, Value: matrix.LineMatrix{{100, 100}, {100, 200}}},
				{Index: 1, stat: true, inGraph: true, NodeType: PNode, Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{1: PointLine}, {0: PointLine}}},
		},
		{"matrix poly", matrix.Matrix{100, 100}, matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}}, {Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{1: PointPoly}, {0: PointPoly}}},
		},

		{"line matrix", matrix.LineMatrix{{100, 100}, {100, 200}}, matrix.Matrix{100, 100},
			&MatrixGraph{nodes: []*Node{{Value: matrix.LineMatrix{{100, 100}, {100, 200}}}, {Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{1: PointLine}, {0: PointLine}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{}}},
		},
		{"line line", matrix.LineMatrix{{100, 100}, {100, 202}}, matrix.LineMatrix{{100, 100}, {100, 200}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.LineMatrix{{100, 100}, {100, 202}}}, {Value: matrix.Matrix{100, 100}}, {Value: matrix.LineMatrix{{100, 100}, {100, 200}}},
				{Value: matrix.Matrix{100, 200}}, {Value: matrix.LineMatrix{{100, 200}, {100, 202}}},
			}, edges: []map[int]int{{}, {2: PointLine}, {1: PointLine, 3: PointLine}, {2: PointLine, 4: PointLine}, {3: PointLine}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.LineMatrix{{100, 100}, {100, 200}}}, {Value: matrix.Matrix{100, 100}}, {Value: matrix.Matrix{100, 200}}},
				edges: []map[int]int{{1: PointLine, 2: PointLine}, {0: PointLine}, {0: PointLine}}},
		},
		{"line poly", matrix.LineMatrix{{100, 100}, {100, 200}}, matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.LineMatrix{{100, 100}, {100, 200}}}, {Value: matrix.Matrix{100, 100}}, {Value: matrix.Matrix{100, 200}}},
				edges: []map[int]int{{1: PointLine, 2: PointLine}, {0: PointLine}, {0: PointLine}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}},
				{Value: matrix.Matrix{100, 100}}, {Value: matrix.LineMatrix{{100, 100}, {100, 200}}}, {Value: matrix.Matrix{100, 200}},
				{Value: matrix.LineMatrix{{100, 200}, {200, 200}, {200, 100}, {100, 100}}},
			},
				edges: []map[int]int{{}, {2: PointLine, 4: PointLine}, {1: PointLine, 3: PointLine}, {2: PointLine, 4: PointLine}, {1: PointLine, 3: PointLine}}},
		},

		{"poly matrix", matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}, matrix.Matrix{100, 100},
			&MatrixGraph{nodes: []*Node{{Value: matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}}, {Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{1: PointLine}, {0: PointLine}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.Matrix{100, 100}}}, edges: []map[int]int{{}}},
		},
		{"poly line", matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}, matrix.LineMatrix{{100, 100}, {100, 200}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}},
				{Value: matrix.Matrix{100, 100}}, {Value: matrix.LineMatrix{{100, 100}, {100, 200}}}, {Value: matrix.Matrix{100, 200}},
				{Value: matrix.LineMatrix{{100, 200}, {200, 200}, {200, 100}, {100, 100}}}},
				edges: []map[int]int{{}, {2: PointLine, 4: PointLine}, {1: PointLine, 3: PointLine}, {2: PointLine, 4: PointLine}, {1: PointLine, 3: PointLine}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.LineMatrix{{100, 100}, {100, 200}}}, {Value: matrix.Matrix{100, 100}}, {Value: matrix.Matrix{100, 200}}},
				edges: []map[int]int{{1: PointLine, 2: PointLine}, {0: PointLine}, {0: PointLine}}},
		},
		{"poly poly", matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}, matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}}}, edges: []map[int]int{{}}},
			&MatrixGraph{nodes: []*Node{{Value: matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}}}}, edges: []map[int]int{{}}},
		},
	}
	type testCase struct {
		args       args
		graphs     graphs
		wantErr    bool
		wantGraphs graphs
	}
	tests := []testCase{}

	for _, v := range testMatrix {
		g1, _ := GenerateGraph(v.m1)
		g2, _ := GenerateGraph(v.m2)
		for _, g := range v.g1.Nodes() {
			g.stat = true
		}
		for _, g := range v.g2.Nodes() {
			g.stat = true
		}
		tests = append(tests, testCase{v, graphs{g1, g2}, false, graphs{v.g1, v.g2}})
	}

	for _, tt := range tests {
		if tt.args.name != "matrix poly" {
			continue
		}
		t.Run(tt.args.name, func(t *testing.T) {
			if err := IntersectionHandle(tt.args.m1, tt.args.m2, tt.graphs.g1, tt.graphs.g2); (err != nil) != tt.wantErr {
				t.Errorf("IntersectionHandle() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.graphs.g1.Equals(tt.wantGraphs.g1) || !tt.graphs.g2.Equals(tt.wantGraphs.g2) {
				t.Errorf("IntersectionHandle() %v\nreturn = %v, \n  want %v \nreturn = %v, \n  want %v",
					tt.args.name, tt.graphs.g1, tt.wantGraphs.g1, tt.graphs.g2, tt.wantGraphs.g2)
			}
		})
	}
}
