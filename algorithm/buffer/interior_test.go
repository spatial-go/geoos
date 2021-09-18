package buffer

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestInteriorPoint(t *testing.T) {
	type args struct {
		geom matrix.Steric
	}
	tests := []struct {
		name           string
		args           args
		wantInteriorPt matrix.Matrix
	}{
		{"point interior", args{matrix.Matrix{0, 5}}, matrix.Matrix{0, 5}},
		{"line interior", args{matrix.LineMatrix{{0, 5}, {0, 10}}}, matrix.Matrix{0, 5}},
		{"polygon interior", args{matrix.PolygonMatrix{
			{{0, 0}, {0, 5}, {5, 5}, {5, 0}, {0, 0}},
		},
		}, matrix.Matrix{2.5, 2.5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInteriorPt := InteriorPoint(tt.args.geom); !gotInteriorPt.Equals(tt.wantInteriorPt) {
				t.Errorf("InteriorPoint() = %v, want %v", gotInteriorPt, tt.wantInteriorPt)
			}
		})
	}
}
