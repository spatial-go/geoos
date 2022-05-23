// Package encoding  is a library for encoding and decoding into Go structs using the geometries.
package encoding

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/space"
)

func TestEncode(t *testing.T) {
	type args struct {
		g        space.Geometry
		codeType int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "Encode string", args: args{space.LineString{{50, 100}, {50, 200}}, WKT},
			want: []byte("LINESTRING(50 100,50 200)"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.g, tt.args.codeType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s        []byte
		codeType int
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "Decode string", args: args{
			[]byte(`GEOMETRYCOLLECTION(MULTILINESTRING((126 156.25,126 125),(101 150,90 161),(90 161,76 175)),MULTILINESTRING EMPTY)`), WKT},
			want: space.Collection{space.MultiLineString{{{126, 156.25}, {126, 125}}, {{101, 150}, {90, 161}}, {{90, 161}, {76, 175}}},
				space.MultiLineString{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.s, tt.args.codeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode()%v = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
