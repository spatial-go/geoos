package planar

import (
	"testing"

	"github.com/spatial-go/geoos"
)

var geom geoos.Geometry = geoos.Polygon{{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}}}

// Benchmark_Megrez test megrez
func Benchmark_Megrez(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type args struct {
			g geoos.Geometry
		}
		tests := []struct {
			name    string
			args    args
			want    float64
			wantErr bool
		}{
			{name: "area", args: args{g: geom}, want: 4.0, wantErr: false},
		}
		for _, tt := range tests {
			got, err := NormalStrategy().Area(tt.args.g)
			if (err != nil) != tt.wantErr {
				b.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				b.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		}
	}
}

func Benchmark_Geos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type args struct {
			g geoos.Geometry
		}
		tests := []struct {
			name    string
			args    args
			want    float64
			wantErr bool
		}{
			{name: "area", args: args{g: geom}, want: 4.0, wantErr: false},
		}
		for _, tt := range tests {
			got, err := GetStrategy(newGEOAlgorithm).Area(tt.args.g)
			if (err != nil) != tt.wantErr {
				b.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				b.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		}
	}
}
