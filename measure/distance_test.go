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
		fromPoint *geoos.Point
		toPoint   *geoos.Point
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "testDistance", args: args{fromPoint: &fromPoint, toPoint: &toPoint}, want: wantResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.fromPoint, tt.args.toPoint); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
