package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/filter"
)

func TestCollection_Filter(t *testing.T) {
	var f filter.Filter = &filter.UniqueArrayFilter{}
	tests := []struct {
		name string
		c    Collection
		want Collection
	}{
		{"Collection filter", Collection{Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
			Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
			MultiLineString{LineString{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}, LineString{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}}},
			LineString{{1, 1}, {1, 2}, {2, 2}, {2, 2}, {2, 1}, {1, 1}},
		},
			Collection{Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
				Polygon{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}, {{1.5, 1.5}, {1.5, 2}, {2, 2}, {2, 1.5}, {1.5, 1.5}}},
				MultiLineString{LineString{{1, 1}, {1, 2}, {2, 2}, {2, 1}}, LineString{{1, 1}, {1, 2}, {2, 2}, {2, 1}}},
				LineString{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Filter(f)
			if !got.Equals(tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
