package measure

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestSpheroidDistance(t *testing.T) {
	fromPoint := matrix.Matrix{12, 15}
	toPoint := matrix.Matrix{13, 15}
	line0 := matrix.LineMatrix{{116.40495300292967, 39.926785883895654}, {116.3975715637207, 39.9295502919}}
	line1 := matrix.LineMatrix{{116.37310981750488, 39.92099342895789}, {116.39928817749023, 39.9174387253541}}

	wantResult := 107405.96007592858
	wantResult1 := 1147.420777283722
	type args struct {
		from matrix.Steric
		to   matrix.Steric
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "testDistance", args: args{from: fromPoint, to: toPoint}, want: wantResult},
		{name: "testDistanceLine", args: args{from: line0, to: line1}, want: wantResult1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpheroidDistance(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMercatorDistance(t *testing.T) {
	type args struct {
		dis float64
		lat float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test lat=0 MercatorDistance", args: args{dis: 1, lat: 0}, want: 1},
		{name: "test lat=18.2454 MercatorDistance", args: args{dis: 1, lat: 18.2454}, want: 1.0529348778624306},
		{name: "test lat=39.886051 MercatorDistance", args: args{dis: 1, lat: 39.886051}, want: 1.3032230353739989},
		{name: "test lat=52.987939 MercatorDistance", args: args{dis: 1, lat: 52.987939}, want: 1.6611523970517712},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MercatorDistance(tt.args.dis, tt.args.lat); got != tt.want {
				t.Errorf("MercatorDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
