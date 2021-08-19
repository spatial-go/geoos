package buffer

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestConvexHull(t *testing.T) {
	type args struct {
		geom matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want matrix.Steric
	}{
		{"convexHull", args{matrix.PolygonMatrix{{{1, 1}, {3, 1}, {2, 2}, {3, 3}, {1, 3}, {1, 1}}}},
			matrix.PolygonMatrix{{{1, 1}, {1, 3}, {3, 3}, {3, 1}, {1, 1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvexHull(tt.args.geom); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvexHull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvexHullComputer_ConvexHull(t *testing.T) {
	type fields struct {
		inputPts []matrix.Matrix
	}
	tests := []struct {
		name   string
		fields fields
		want   matrix.Steric
	}{
		{"  Computer convexHull", fields{[]matrix.Matrix{{1, 1}, {3, 1}, {2, 2}, {3, 3}, {1, 3}, {1, 1}}},
			matrix.PolygonMatrix{{{1, 1}, {1, 3}, {3, 3}, {3, 1}, {1, 1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConvexHullComputer{
				inputPts: tt.fields.inputPts,
			}
			if got := c.ConvexHull(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvexHullComputer.ConvexHull() = %v, want %v", got, tt.want)
			}
		})
	}
}
