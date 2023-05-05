// Package buffer define geomtry matrix conversion.
package buffer

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestVariableBuffer(t *testing.T) {
	type args struct {
		line          matrix.Steric
		startDistance float64
		endDistance   float64
	}
	tests := []struct {
		name string
		args args
		want matrix.Steric
	}{
		{"simple case", args{matrix.LineMatrix{{2.073333263397217, 48.81027603149414}, {1.5225944519042969, 48.45795440673828}}, 0.001, 0.0012},
			matrix.PolygonMatrix{{{2.0738724106666613, 48.80943381998754}, {1.5232414286276303, 48.45694375293036}, {1.5219480936357215, 48.45896545618744},
				{2.072794631506731, 48.81111857270177}, {2.0738724106666613, 48.80943381998754}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VariableInterpolatedBuffer(tt.args.line, tt.args.startDistance, tt.args.endDistance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VariableBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVariableDistancesBuffer(t *testing.T) {
	type args struct {
		line      matrix.Steric
		distances []float64
	}
	tests := []struct {
		name       string
		args       args
		wantBuffer matrix.Steric
	}{
		{"simple case", args{matrix.LineMatrix{{2.073333263397217, 48.81027603149414}, {1.5225944519042969, 48.45795440673828}}, []float64{0.001, 0.0012}},
			matrix.PolygonMatrix{{{2.0738724106666613, 48.80943381998754}, {1.5232414286276303, 48.45694375293036}, {1.5219480936357215, 48.45896545618744},
				{2.072794631506731, 48.81111857270177}, {2.0738724106666613, 48.80943381998754}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBuffer := VariableDistancesBuffer(tt.args.line, tt.args.distances); !reflect.DeepEqual(gotBuffer, tt.wantBuffer) {
				t.Errorf("VariableDistancesBuffer() = %v, want %v", gotBuffer, tt.wantBuffer)
			}
		})
	}
}
