package simplify

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestTopologyPreservingSimplifier_Simplify(t *testing.T) {
	type fields struct {
		InputGeom      matrix.Steric
		lineSimplifier *TaggedLinesSimplifier
		linestrings    []*TaggedLineString
	}
	type args struct {
		geom              matrix.Steric
		distanceTolerance float64
	}
	geom := matrix.LineMatrix{{0, 0}, {1, 1}, {0, 2}, {1, 3}, {0, 4}, {1, 5}}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   matrix.Steric
	}{
		{"topolog Preserving simplifier", fields{InputGeom: geom}, args{geom, 1.0}, matrix.LineMatrix{{0, 0}, {1, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TopologyPreservingSimplifier{
				InputGeom:      tt.fields.InputGeom,
				lineSimplifier: tt.fields.lineSimplifier,
				linestrings:    tt.fields.linestrings,
			}
			if got := tr.Simplify(tt.args.geom, tt.args.distanceTolerance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopologyPreservingSimplifier.Simplify() = %v, want %v", got, tt.want)
			}
		})
	}
}
