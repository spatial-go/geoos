package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/filter"
)

func TestLineString_Filter(t *testing.T) {
	var f filter.Filter = &filter.UniqueArrayFilter{}
	tests := []struct {
		name string
		ls   LineString
		want LineString
	}{
		{"line filter", LineString{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}},
			LineString{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ls.Filter(f)
			if !got.Equals(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
