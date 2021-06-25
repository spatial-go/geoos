package planar

import (
	"testing"

	"github.com/spatial-go/geoos"
)

var geom geoos.Geometry = geoos.Polygon{geoos.Ring{geoos.Point{-1, -1}, geoos.Point{1, -1}, geoos.Point{1, 1}, geoos.Point{-1, 1}, geoos.Point{-1, -1}}}

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
			b.Run(tt.name, func(b *testing.B) {
				got, err := NormalStrategy().Area(tt.args.g)
				if (err != nil) != tt.wantErr {
					b.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					b.Errorf("Area() got = %v, want %v", got, tt.want)
				}
			})
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
			b.Run(tt.name, func(b *testing.B) {
				got, err := GetStrategy(newGEOAlgorithm).Area(tt.args.g)
				if (err != nil) != tt.wantErr {
					b.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					b.Errorf("Area() got = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
