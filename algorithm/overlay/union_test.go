package overlay

import (
	"testing"

	"github.com/spatial-go/geoos"
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

		{name: "line line", args: args{g1: matrix.LineMatrix{{50, 100}, {50, 200}}, g2: matrix.LineMatrix{{50, 50}, {50, 150}}},
			want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}}, matrix.LineMatrix{{50, 150}, {50, 200}}, matrix.LineMatrix{{50, 50}, {50, 100}}},
		},
		{name: "line line1", args: args{g1: matrix.LineMatrix{{50, 100}, {50, 200}}, g2: matrix.LineMatrix{{30, 150}, {80, 150}}},
			want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}}, matrix.LineMatrix{{50, 150}, {50, 200}}, matrix.LineMatrix{{30, 150}, {50, 150}}, matrix.LineMatrix{{50, 150}, {80, 150}}},
		},
		{name: "line line2", args: args{g1: matrix.LineMatrix{{50, 100}, {50, 200}}, g2: matrix.LineMatrix{{30, 30}, {30, 150}}},
			want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{30, 30}, {30, 150}}},
		},
		{name: "line line3", args: args{g1: matrix.LineMatrix{{50, 100}, {50, 200}}, g2: matrix.LineMatrix{{30, 30}, {30, 100}}},
			want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.LineMatrix{{30, 30}, {30, 100}}},
		},

		{name: "line poly", args: args{g1: matrix.LineMatrix{{50, 100}, {50, 200}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 200}}, matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
		},
		{name: "line poly1", args: args{g1: matrix.LineMatrix{{110, 100}, {600, 1000}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.Collection{matrix.LineMatrix{{327.7777777777777715, 500}, {600, 1000}}, matrix.LineMatrix{{110, 100}, {300, 448.9795918367346985}}, matrix.LineMatrix{{300, 448.9795918367346985}, {327.77777777777777, 500}},
				matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
		},
		{name: "line poly2", args: args{g1: matrix.LineMatrix{{500, 500}, {600, 1000}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.Collection{matrix.LineMatrix{{500, 500}, {600, 1000}}, matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
		},
		{name: "line poly3", args: args{g1: matrix.LineMatrix{{400, 400}, {600, 600}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.Collection{matrix.LineMatrix{{500, 500}, {600, 600}}, matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
		},
		{name: "line poly4", args: args{g1: matrix.LineMatrix{{350, 400}, {450, 450}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}},
		},
		{name: "line poly5", args: args{g1: matrix.LineMatrix{{200, 300}, {500, 300}, {600, 600}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.Collection{matrix.LineMatrix{{500, 300}, {600, 600}}, matrix.LineMatrix{{200, 300}, {300, 300}},
				matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}},
			},
		},
		{name: "line poly6", args: args{g1: matrix.LineMatrix{{200, 300}, {500, 300}, {500, 600}, {800, 900}}, g2: matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}}},
			want: matrix.Collection{matrix.LineMatrix{{500, 500}, {500, 600}, {800, 900}}, matrix.LineMatrix{{200, 300}, {300, 300}},
				matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}},
			},
		},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "line poly6" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Union(tt.args.g1, tt.args.g2); !tt.want.Equals(gotResult) {
				t.Errorf("Union()%v = %v, want %v", tt.name, gotResult, tt.want)
			}
		})
	}
}
