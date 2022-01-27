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

func TestUnmarshalString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{

		{name: "unmarshal string", args: args{`SRID=104326;GEOMETRYCOLLECTION(SRID=104326;MULTILINESTRING((126 156.25,126 125),(101 150,90 161),(90 161,76 175)),SRID=104326;MULTILINESTRING EMPTY)`},
			want: space.Collection{space.MultiLineString{{{126, 156.25}, {126, 125}}, {{101, 150}, {90, 161}}, {{90, 161}, {76, 175}}}, space.MultiLineString{}},
		},
		{name: "unmarshal string", args: args{"SRID=104326;LINESTRING(50 100,50 200)"},
			want: space.LineString{{50, 100}, {50, 200}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("UnmarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
