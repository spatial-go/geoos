package overlay

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestUnion(t *testing.T) {
	type args struct {
		g1 matrix.Steric
		g2 matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want matrix.Steric
	}{
		{name: "union", args: args{g1: matrix.Matrix{1, 2}, g2: matrix.Matrix{3, 4}}, want: matrix.Collection{matrix.Matrix{1, 2}, matrix.Matrix{3, 4}}},

		{name: "point line", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.LineMatrix{{1, 1}, {5, 5}}},
			want: matrix.LineMatrix{{1, 1}, {5, 5}}},
		{name: "point line1", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.LineMatrix{{5, 1}, {5, 5}}},
			want: matrix.Collection{matrix.Matrix{3, 3}, matrix.LineMatrix{{5, 1}, {5, 5}}}},
		{name: "point ring", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.LineMatrix{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			want: matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}}},
		{name: "point ring1", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.LineMatrix{{1, 1}, {3, 1}, {3, 3}, {1, 3}, {1, 1}}},
			want: matrix.PolygonMatrix{{{1, 1}, {3, 1}, {3, 3}, {1, 3}, {1, 1}}}},
		{name: "point ring2", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.LineMatrix{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			want: matrix.Collection{matrix.Matrix{3, 3}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}},
		{name: "point poly", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}}},
			want: matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}}},
		{name: "point poly1", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.PolygonMatrix{{{1, 1}, {3, 1}, {3, 3}, {1, 3}, {1, 1}}}},
			want: matrix.PolygonMatrix{{{1, 1}, {3, 1}, {3, 3}, {1, 3}, {1, 1}}}},
		{name: "point poly2", args: args{g1: matrix.Matrix{3, 3}, g2: matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}},
			want: matrix.Collection{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}}},
		{name: "point poly3", args: args{g1: matrix.Matrix{3, 3},
			g2: matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}, {{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}},
			want: matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}, {{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}},
		{name: "point poly4", args: args{g1: matrix.Matrix{3, 3},
			g2: matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}, {{2, 2}, {4, 2}, {4, 4}, {2, 4}, {2, 2}}}},
			want: matrix.Collection{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}, {{2, 2}, {4, 2}, {4, 4}, {2, 4}, {2, 2}}}},
		},

		// {name: "union line", args: args{g1: matrix.LineMatrix{{50, 100}, {50, 200}}, g2: matrix.LineMatrix{{50, 50}, {50, 150}}},
		// 	want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}}, matrix.LineMatrix{{50, 150}, {50, 200}}, matrix.LineMatrix{{50, 50}, {50, 100}}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Union(tt.args.g1, tt.args.g2); !tt.want.Equals(gotResult) {
				t.Errorf("Union()%v = %v, want %v", tt.name, gotResult, tt.want)
			}
		})
	}
}
