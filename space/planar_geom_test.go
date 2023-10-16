package space

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/algorithm/measure"
)

func Test_Centroid(t *testing.T) {
	for _, tt := range TestsCentroid {

		t.Run(tt.name, func(t *testing.T) {

			got := tt.args.g.Centroid()
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
		f    measure.DistanceFunc
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
			got, err := tt.args.from.Distance(tt.args.to)
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
		{"case1", args{Point{2.073333263397217, 48.81027603149414}, 100, 4},
			Polygon{{{2.07469731736744, 48.81027603149415},
				{2.0745934849415466, 48.80993226430165},
				{2.0742977952094663, 48.809640830701994},
				{2.0738552642524732, 48.8094461000523},
				{2.073333263397217, 48.80937771956217},
				{2.0728112625419604, 48.8094461000523},
				{2.0723687315849677, 48.809640830701994},
				{2.072073041852887, 48.80993226430165},
				{2.0719692094269937, 48.81027603149415},
				{2.072073041852887, 48.810619796329775},
				{2.0723687315849677, 48.81091122423942},
				{2.0728112625419604, 48.811105949199124},
				{2.073333263397217, 48.81117432733234},
				{2.0738552642524732, 48.811105949199124},
				{2.0742977952094663, 48.81091122423942},
				{2.0745934849415466, 48.810619796329775},
				{2.07469731736744, 48.81027603149415},
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.geometry.BufferInMeter(tt.args.width, tt.args.quadsegs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BufferInMeter() = %v, want %v", got.ToMatrix(), tt.want)
			}
		})
	}
}
