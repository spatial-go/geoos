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
		{name: "wkb Point0",
			args: args{space.Point{116.310066223145, 40.0425491333008}, WKB},
			want: []byte("0101000020e610000021000020d8135d400300004072054440"),
		},
		{name: "geojson Point0",
			args: args{space.Point{116.310066223145, 40.0425491333008}, GeoJSON},
			want: []byte("{\"type\":\"Point\",\"coordinates\":[116.310066223145,40.0425491333008]}"),
		},
		{name: "geocsv Points",
			args: args{space.Collection{space.Point{116.310066223145, 40.0425491333008},
				space.Point{116.31, 40.04}}, GeoCSV},
			want: []byte("way_id,pt_id,x,y\n0,0,116.310066223145,40.0425491333008\n1,1,116.31,40.04\n"),
		},
		{name: "geobuf Point0",
			args: args{space.Point{116.310066223145, 40.0425491333008}, Geobuf},
			want: []byte("dimensions:2  precision:1  geometry:{}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.g, tt.args.codeType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", string(got), string(tt.want))
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
		// {name: "Decode string", args: args{
		// 	[]byte(`GEOMETRYCOLLECTION(MULTILINESTRING((126 156.25,126 125),(101 150,90 161),(90 161,76 175)),MULTILINESTRING EMPTY)`), WKT},
		// 	want: space.Collection{space.MultiLineString{{{126, 156.25}, {126, 125}}, {{101, 150}, {90, 161}}, {{90, 161}, {76, 175}}},
		// 		space.MultiLineString{}},
		// },
		// {" wkb line1 ", args{
		// 	[]byte("0102000020E610000004000000F7FFFF7F20155D40C9D9B446F6F843400F000020B51C5D409241C66566F94340DDFFFFFF791D5D40336A670189F04340E8FFFF5FA7175D409DF9A3B974EF4340"),
		// 	WKB},
		// 	space.LineString{
		// 		{116.33010864257814, 39.94501575308417},
		// 		{116.44855499267578, 39.948437425427215},
		// 		{116.4605712890625, 39.87918107556866},
		// 		{116.36959075927736, 39.87074966913789}},
		// 	false},
		// {name: "geojson string", args: args{
		// 	[]byte("{\"type\":\"Point\",\"coordinates\":[116.310066223145,40.0425491333008]}"), GeoJSON},
		// 	want: space.Point{116.310066223145, 40.0425491333008},
		// },
		// {name: "geocsv string", args: args{
		// 	[]byte("way_id,pt_id,x,y\n0,0,116.310066223145,40.0425491333008\n1,1,116.31,40.04\n"), GeoCSV},
		// 	want: space.Collection{space.Point{116.310066223145, 40.0425491333008},
		// 		space.Point{116.31, 40.04}},
		// },
		{name: "geobuf string", args: args{
			[]byte("{\"type\":\"Point\",\"coordinates\":[116.310066223145,40.0425491333008]}"), Geobuf},
			want: space.Point{116.310066223145, 40.0425491333008},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.s, tt.args.codeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() %v error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !got.EqualsExact(tt.want, 0.000001) {
				t.Errorf("Decode()%v %T= %v, want %v", tt.name, got, got, tt.want)
			}
		})
	}
}
