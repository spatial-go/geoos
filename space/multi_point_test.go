package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/filter"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestMultiPoint_Nums(t *testing.T) {
	multiPoint := MultiPoint{{-1, 0}, {-1, 2}, {-1, 3}, {-1, 4}, {-1, 7}, {0, 1}, {0, 3},
		{1, 1}, {2, 0}, {6, 0}, {7, 8}, {9, 8}, {10, 6}}
	tests := []struct {
		name string
		mp   MultiPoint
		want int
	}{
		{name: "geometry multiLineString", mp: multiPoint, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Nums(); got != tt.want {
				t.Errorf("MultiPoint.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPoint_Filter(t *testing.T) {
	var f filter.Filter[matrix.Matrix] = matrix.CreateFilterMatrix()
	tests := []struct {
		name string
		mp   MultiPoint
		want Geometry
	}{
		{"MultiPoint filter", MultiPoint{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}},
			MultiPoint{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Filter(f); !got.Equals(tt.want) {
				t.Errorf("MultiPoint.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
