package measure

import (
	"testing"

	"github.com/spatial-go/geoos"
)

func TestDistance(t *testing.T) {
	fromPoint := geoos.Point{12, 15}
	toPoint := geoos.Point{13, 15}

	wantResult := 107405.96007592858
	type args struct {
		fromPoint geoos.Point
		toPoint   geoos.Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "testDistance", args: args{fromPoint: fromPoint, toPoint: toPoint}, want: wantResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.fromPoint, tt.args.toPoint); got != tt.want {
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
