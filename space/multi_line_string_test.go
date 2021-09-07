package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestMultiLineString_Nums(t *testing.T) {
	multiLineString := MultiLineString{
		{{10, 130}, {50, 190}, {110, 190}, {140, 150}, {150, 80}, {100, 10}, {20, 40}, {10, 130}},
		{{70, 40}, {100, 50}, {120, 80}, {80, 110}, {50, 90}, {70, 40}},
	}
	tests := []struct {
		name string
		mls  MultiLineString
		want int
	}{
		{name: "geometry multiLineString", mls: multiLineString, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mls.Nums(); got != tt.want {
				t.Errorf("MultiLineString.Nums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiLineString_Filter(t *testing.T) {
	var f matrix.Filter = &matrix.UniqueArrayFilter{}
	tests := []struct {
		name string
		mls  MultiLineString
		want MultiLineString
	}{
		{"multi line filter", MultiLineString{LineString{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, LineString{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}},
			MultiLineString{LineString{{1, 1}, {1, 2}, {2, 2}, {2, 1}}, LineString{{1, 1}, {1, 2}, {2, 2}, {2, 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mls.Filter(f)
			if !got.Equals(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
