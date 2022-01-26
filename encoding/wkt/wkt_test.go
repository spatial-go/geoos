package wkt

import (
	"testing"

	"github.com/spatial-go/geoos/space"
)

func TestMarshalString(t *testing.T) {
	type args struct {
		geom space.Geometry
	}
	geomCoord, _ := space.CreateElementValidWithCoordSys(space.LineString{{50, 100}, {50, 200}}, 4326)
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "marshal string", args: args{space.LineString{{50, 100}, {50, 200}}},
			want: "SRID=104326;LINESTRING(50 100,50 200)",
		},
		{name: "marshal string coord", args: args{geomCoord},
			want: "SRID=4326;LINESTRING(50 100,50 200)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MarshalString(tt.args.geom); got != tt.want {
				t.Errorf("MarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
