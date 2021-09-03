package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/filter"
)

func TestRing_Filter(t *testing.T) {

	var f filter.Filter = &filter.UniqueArrayFilter{}
	tests := []struct {
		name string
		r    Ring
		want Ring
	}{
		{"ring filter", Ring{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}},
			Ring{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.Filter(f)
			if !got.Equals(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
