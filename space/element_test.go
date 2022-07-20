package space

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/measure"
)

func Test_Centroid(t *testing.T) {
	for _, tt := range TestsCentroid {
		t.Run(tt.name, func(t *testing.T) {

			got := Centroid(tt.args.g)
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

func TestDistance(t *testing.T) {
	type args struct {
		from Geometry
		to   Geometry
		f    measure.Distance
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"case1", args{Point{1, 1}, Point{4, 5}, measure.PlanarDistance}, 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Distance(tt.args.from, tt.args.to, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBufferInMeter(t *testing.T) {
	type args struct {
		geometry Geometry
		width    float64
		quadsegs int
	}
	tests := []struct {
		name string
		args args
		want Geometry
	}{
		{"case1", args{Point{1, 1}, 1, 1},
			Polygon{{{1.000008984521167, 0.9999999999999887},
				{0.9999999999999999, 0.9999910168471858},
				{0.9999910154788327, 0.9999999999999887},
				{0.9999999999999999, 1.0000089831527534},
				{1.000008984521167, 0.9999999999999887}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BufferInMeter(tt.args.geometry, tt.args.width, tt.args.quadsegs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BufferInMeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
