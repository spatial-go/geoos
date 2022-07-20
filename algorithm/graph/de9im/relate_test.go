package de9im

import (
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/graph/graphtests"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

func TestRelate(t *testing.T) {

	for _, tt := range graphtests.TestRelateData {
		if !geoos.GeoosTestTag &&
			tt.Name != "LinePoly 6" {
			continue
		}

		t.Run(tt.Name, func(t *testing.T) {
			intersectBound := false
			env1 := envelope.Bound(tt.Args[0].Bound())
			env2 := envelope.Bound(tt.Args[1].Bound())
			if env1.IsIntersects(env2) {
				intersectBound = true
			}
			if env1.Contains(env2) || env2.Contains(env1) {
				intersectBound = true
			}
			if got := Relate(tt.Args[0], tt.Args[1]); got != tt.Want {
				t.Errorf("Relate()%v = %v, want %v  %v", tt.Name, got, tt.Want, intersectBound)
			}
		})
	}
}
