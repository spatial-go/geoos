package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/filter"
)

func TestMultiPolygon_Nums(t *testing.T) {
	mp := MultiPolygon{
		{{{40, 40}, {20, 45}, {45, 30}, {40, 40}}},
		{{{20, 35}, {10, 30}, {10, 10}, {30, 5}, {45, 20}, {20, 35}}},
		{{{30, 20}, {20, 15}, {20, 25}, {30, 20}}},
	}

	tests := []struct {
		name string
		mp   MultiPolygon
		want int
	}{
		{name: "nums", mp: mp, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.Nums(); got != tt.want {
				t.Errorf("MultiPolygon.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiPolygon_Filter(t *testing.T) {
	var f filter.Filter = &filter.UniqueArrayFilter{}
	tests := []struct {
		name string
		mp   MultiPolygon
		want MultiPolygon
	}{
		{"multi polygon filter", MultiPolygon{Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
			Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}}},
			MultiPolygon{Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
				Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mp.Filter(f)
			if !got.Equals(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
