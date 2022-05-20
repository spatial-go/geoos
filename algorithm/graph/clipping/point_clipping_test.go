package clipping

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/graph/graphtests"
)

func TestPointClipping_Intersection(t *testing.T) {

	for _, tt := range graphtests.TestsPointIntersecation {
		t.Run(tt.Name, func(t *testing.T) {
			p := &PointClipping{
				Subject:  tt.Fields[0],
				Clipping: tt.Fields[1],
			}
			got, err := p.Intersection()
			if (err != nil) != tt.WantErr {
				t.Errorf("PointOverlay.Intersection() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("PointOverlay.Intersection() = %v, want %v", got, tt.Want)
			}
		})
	}
}
