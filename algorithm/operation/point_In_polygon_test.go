package operation

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestPointInPolygon(t *testing.T) {
	type args struct {
		point matrix.Matrix
		poly  matrix.LineMatrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{matrix.Matrix{5, 5}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, false},
		{"test2", args{matrix.Matrix{5, 7}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, true},
		{"test3", args{matrix.Matrix{5, 6}, matrix.LineMatrix{{1, 1}, {2, 1}, {2, 6}, {7, 6}, {7, 1}, {10, 1}, {10, 10}, {1, 10}, {1, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PnPolygonByCross(tt.args.point, tt.args.poly); got != tt.want {
				t.Errorf("PointInPolygon()%v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
