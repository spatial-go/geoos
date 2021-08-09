package angle

import (
	"math"
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestAngle(t *testing.T) {
	type args struct {
		p0 matrix.Matrix
		p1 matrix.Matrix
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"angle", args{matrix.Matrix{100, 100}, matrix.Matrix{100, 200}}, math.Pi / 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Angle(tt.args.p0, tt.args.p1); got != tt.want {
				t.Errorf("Angle() = %v, want %v", got, tt.want)
			}
		})
	}
}
