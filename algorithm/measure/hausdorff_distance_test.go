package measure

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestHausdorffDistance_Distance(t *testing.T) {
	g1, g2 := matrix.LineMatrix{{0, 0}, {2, 0}}, matrix.LineMatrix{{0, 1}, {1, 0}, {2, 1}}
	g3, g4 := matrix.LineMatrix{{130, 0}, {0, 0}, {0, 150}}, matrix.LineMatrix{{10, 10}, {10, 150}, {130, 10}}
	type args struct {
		g0 matrix.Steric
		g1 matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Hausdorff Distance1", args: args{
			g0: g1, g1: g2,
		}, want: 1},
		{name: "Hausdorff Distance2", args: args{
			g0: g3, g1: g4,
		}, want: 14.142135623730951},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HausdorffDistance{}
			if got := h.Distance(tt.args.g0, tt.args.g1); got != tt.want {
				t.Errorf("HausdorffDistance.Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
