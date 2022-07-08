package angle

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestQuadrantFloat(t *testing.T) {
	type args struct {
		dx float64
		dy float64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"TestQuadrantFloat", args{3, 2}, NE, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QuadrantFloat(tt.args.dx, tt.args.dy)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuadrantFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QuadrantFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuadrant(t *testing.T) {
	type args struct {
		p0 matrix.Matrix
		p1 matrix.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"TestQuadrantFloat", args{matrix.Matrix{3, 2}, matrix.Matrix{2, 1}}, SW, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Quadrant(tt.args.p0, tt.args.p1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Quadrant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Quadrant() = %v, want %v", got, tt.want)
			}
		})
	}
}
