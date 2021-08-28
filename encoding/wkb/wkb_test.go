package wkb

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/space"
)

// AllGeometries lists all possible types and values that a geometry
// interface can be. It should be used only for testing to verify
// functions that accept a Geometry will work in all cases.
var AllGeometries = []space.Geometry{
	nil,
	space.Point{},
	space.MultiPoint{},
	space.LineString{},
	space.MultiLineString{},
	space.Ring{},
	space.Polygon{},
	space.MultiPolygon{},
	space.Bound{},
	space.Collection{},

	// nil values
	space.MultiPoint(nil),
	space.LineString(nil),
	space.MultiLineString(nil),
	space.Ring(nil),
	space.Polygon(nil),
	space.MultiPolygon(nil),
	space.Collection(nil),

	// Collection of Collection
	space.Collection{space.Collection{space.Point{}}},
}

func TestMarshal(t *testing.T) {
	for _, g := range AllGeometries {
		_, _ = Marshal(g, binary.BigEndian)
	}
	type args struct {
		wkbHex space.Geometry
	}
	wkbStr, _ := hex.DecodeString("0101000020E610000000000020D8135D400000004072054440")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{" GeomFromWKBHexStr ", args{space.Point{116.310066223145, 40.0425491333008}}, wkbStr, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.wkbHex, binary.BigEndian)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeomFromWKBHexStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeomFromWKBHexStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustMarshal(t *testing.T) {
	for _, g := range AllGeometries {
		MustMarshal(g, binary.BigEndian)
	}
}
func BenchmarkEncode_Point(b *testing.B) {
	g := space.Point{1, 2}
	e := NewEncoder(ioutil.Discard)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.Encode(g)
	}
}

func BenchmarkEncode_LineString(b *testing.B) {
	g := space.LineString{
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
	}
	e := NewEncoder(ioutil.Discard)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = e.Encode(g)
	}
}

func compare(t testing.TB, e space.Geometry, b []byte) {
	t.Helper()

	// Decoder
	g, err := NewDecoder(bytes.NewReader(b)).Decode()
	if err != nil {
		t.Fatalf("decoder: read error: %v", err)
	}

	if !g.Equals(e) {
		t.Errorf("decoder: incorrect geometry: %v != %v", g, e)
	}

	g, err = Unmarshal(b)
	if err != nil {
		t.Fatalf("unmarshal: read error: %v", err)
	}

	if !g.Equals(e) {
		t.Errorf("unmarshal: incorrect geometry: %v != %v", g, e)
	}

	var data []byte
	if b[0] == 0 {
		data, err = Marshal(g, binary.BigEndian)
	} else {
		data, err = Marshal(g, binary.LittleEndian)
	}
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	if !bytes.Equal(data, b) {
		t.Logf("%v", data)
		t.Logf("%v", b)
		t.Errorf("marshal: incorrect encoding")
	}

	// preallocation
	if len(data) != geomLength(e) {
		t.Errorf("preallot length: %v != %v", len(data), geomLength(e))
	}

	// Scanner
	var sg space.Geometry

	switch e.(type) {
	case space.Point:
		var p space.Point
		err = Scanner(&p).Scan(b)
		sg = p
	case space.MultiPoint:
		var mp space.MultiPoint
		err = Scanner(&mp).Scan(b)
		sg = mp
	case space.LineString:
		var ls space.LineString
		err = Scanner(&ls).Scan(b)
		sg = ls
	case space.MultiLineString:
		var mls space.MultiLineString
		err = Scanner(&mls).Scan(b)
		sg = mls
	case space.Polygon:
		var p space.Polygon
		err = Scanner(&p).Scan(b)
		sg = p
	case space.MultiPolygon:
		var mp space.MultiPolygon
		err = Scanner(&mp).Scan(b)
		sg = mp
	case space.Collection:
		var c space.Collection
		err = Scanner(&c).Scan(b)
		sg = c
	default:
		t.Fatalf("unknown type: %T", e)
	}

	if err != nil {
		t.Errorf("scan error: %v", err)
	}

	if sg.GeoJSONType() != e.GeoJSONType() {
		t.Errorf("scanning to wrong type: %v != %v", sg.GeoJSONType(), e.GeoJSONType())
	}

	if !sg.Equals(e) {
		t.Errorf("scan: incorrect geometry: %v != %v", sg, e)
	}
}

func TestGeomFromWKBHexStr(t *testing.T) {
	type args struct {
		wkbHex string
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{" GeomFromWKBHexStr ", args{"0101000020E610000000000020D8135D400000004072054440"}, space.Point{116.310066223145, 40.0425491333008}, false},
		{" GeomFromWKBHexStr ", args{"0101000020E6100000A9E2F33378145D4088C78C29E2064440"}, space.Point{116.319836605234, 40.0537769257926}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeomFromWKBHexStr(tt.args.wkbHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeomFromWKBHexStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.EqualsExact(tt.want, 0.00000001) {
				t.Errorf("GeomFromWKBHexStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
