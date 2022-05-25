package clipping

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/graph/graphtests"
)

func TestLineClipping_Union(t *testing.T) {
	for _, tt := range graphtests.TestsLineUnion {
		if !geoos.GeoosTestTag &&
			tt.Name != "line poly2" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &LineClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.Union()
			if (err != nil) != tt.WantErr {
				t.Errorf("LineOverlay.Union() %v error = %v, wantErr %v", tt.Name, err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("LineOverlay.Union() %v = %v, want %v", tt.Name, got, tt.Want)
			}
		})
	}
}

func TestLineClipping_Intersection(t *testing.T) {

	for _, tt := range graphtests.TestsLineIntersecation {
		if !geoos.GeoosTestTag && tt.Name != "line poly5" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &LineClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.Intersection()
			if (err != nil) != tt.WantErr {
				t.Errorf("LineOverlay.Intersection() %v error = %v, wantErr %v", tt.Name, err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("LineOverlay.Intersection() %v = %v, want %v", tt.Name, got, tt.Want)
			}
		})
	}
}

func TestLineClipping_Difference(t *testing.T) {

	for _, tt := range graphtests.TestsLineDifference {
		if !geoos.GeoosTestTag &&
			tt.Name != "line line7" {
			continue
		}
		t.Run(tt.Name, func(t *testing.T) {
			p := &LineClipping{
				PointClipping: &PointClipping{tt.Fields[0], tt.Fields[1]},
			}
			got, err := p.Difference()
			if (err != nil) != tt.WantErr {
				t.Errorf("LineOverlay.Difference() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if !got.Equals(tt.Want) {
				t.Errorf("LineOverlay.Difference()%v = %v, want %v", tt.Name, got, tt.Want)
			}
		})
	}
}
