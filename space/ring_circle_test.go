package space

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/calc"
)

func TestCreateCircle(t *testing.T) {
	type args struct {
		centre Point
		radius float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Circle
		wantErr bool
	}{
		{"create circle", args{Point{10, 10}, 5.5}, &Circle{Centre: Point{10, 10}, Radius: 5.5, Segments: calc.QuadrantSegments}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateCircle(tt.args.centre, tt.args.radius)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCircle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("CreateCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
