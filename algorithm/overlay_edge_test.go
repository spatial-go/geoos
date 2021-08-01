package algorithm

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestVertex_Sub(t *testing.T) {
	type fields struct {
		Matrix              matrix.Matrix
		IsIntersectionPoint bool
		IsEntering          bool
		IsChecked           bool
	}
	type args struct {
		point *Vertex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vertex
	}{
		{"bertex sub", fields{Matrix: matrix.Matrix{2, 2}},
			args{&Vertex{Matrix: matrix.Matrix{1, 1}}}, &Vertex{Matrix: matrix.Matrix{1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vertex{
				Matrix:              tt.fields.Matrix,
				IsIntersectionPoint: tt.fields.IsIntersectionPoint,
				IsEntering:          tt.fields.IsEntering,
				IsChecked:           tt.fields.IsChecked,
			}
			if got := v.Sub(tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vertex.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
