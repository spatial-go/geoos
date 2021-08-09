package snap

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestSnap(t *testing.T) {
	type args struct {
		g0            matrix.Steric
		g1            matrix.Steric
		snapTolerance float64
	}
	tests := []struct {
		name         string
		args         args
		wantSnapGeom matrix.Collection
	}{
		{name: "snap", args: args{
			g0: matrix.Matrix{0.05, 0.05}, g1: matrix.Matrix{0, 0}, snapTolerance: 0.1,
		},
			wantSnapGeom: matrix.Collection{matrix.LineMatrix{{0, 0}}, matrix.LineMatrix{{0, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSnapGeom := Snap(tt.args.g0, tt.args.g1, tt.args.snapTolerance); !reflect.DeepEqual(gotSnapGeom, tt.wantSnapGeom) {
				t.Errorf("Snap() = %v, want %v", gotSnapGeom, tt.wantSnapGeom)
			}
		})
	}
}
