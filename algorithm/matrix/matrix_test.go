package matrix

import (
	"testing"
)

func TestLineMatrix_Bound(t *testing.T) {
	tests := []struct {
		name string
		l    LineMatrix
		want []Matrix
	}{
		{"line bound", LineMatrix{
			{0, 0}, {0, 5}, {5, 5}, {5, 0}, {0, 0},
		}, []Matrix{{0, 0}, {5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Bound(); !got.Equals(tt.want) {
				t.Errorf("LineMatrix.Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}
