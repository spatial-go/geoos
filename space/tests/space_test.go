package test

import (
	"testing"

	"github.com/spatial-go/geoos/space"
)

func Test_Centroid(t *testing.T) {
	for _, tt := range TestsCentroid {
		t.Run(tt.name, func(t *testing.T) {

			got := space.Centroid(tt.args.g)
			if got == nil && tt.want == nil {
				return
			}
			if got == nil {
				t.Errorf("Centroid() got%v = %v, want %v", tt.name, got, tt.want)
			}
			if !got.Equals(tt.want) {
				t.Errorf("Centroid() got %v = %v, want %v; type %T want %T", tt.name, got, tt.want, got, tt.want)
			}
		})
	}
}
