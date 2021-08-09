package buffer

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestComputeCentroid(t *testing.T) {
	type args struct {
		geom matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want matrix.Matrix
	}{
		{"centroid multi point", args{matrix.Collection{
			matrix.Matrix{-1, 0}, matrix.Matrix{-1, 2}, matrix.Matrix{-1, 3},
			matrix.Matrix{-1, 4}, matrix.Matrix{-1, 7}, matrix.Matrix{0, 1}, matrix.Matrix{0, 3}, matrix.Matrix{1, 1},
			matrix.Matrix{2, 0}, matrix.Matrix{6, 0}, matrix.Matrix{7, 8}, matrix.Matrix{9, 8}, matrix.Matrix{10, 6},
		},
		},
			matrix.Matrix{2.3076923076923075, 3.3076923076923075},
		},

		{"centroid multi area", args{matrix.PolygonMatrix{
			{{0, 0}, {0, 5}, {5, 5}, {5, 0}, {0, 0}},
		},
		},
			matrix.Matrix{2.5, 2.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Centroid(tt.args.geom); !got.Equals(tt.want) {
				t.Errorf("ComputeCentroid() = %v, want %v", got, tt.want)
			}
		})
	}
}
