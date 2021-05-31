package wkb

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"testing"

	"github.com/spatial-go/geoos"
)

// AllGeometries lists all possible types and values that a geometry
// interface can be. It should be used only for testing to verify
// functions that accept a Geometry will work in all cases.
var AllGeometries = []geoos.Geometry{
	nil,
	geoos.Point{},
	geoos.MultiPoint{},
	geoos.LineString{},
	geoos.MultiLineString{},
	geoos.Ring{},
	geoos.Polygon{},
	geoos.MultiPolygon{},
	geoos.Bound{},
	geoos.Collection{},

	// nil values
	geoos.MultiPoint(nil),
	geoos.LineString(nil),
	geoos.MultiLineString(nil),
	geoos.Ring(nil),
	geoos.Polygon(nil),
	geoos.MultiPolygon(nil),
	geoos.Collection(nil),

	// Collection of Collection
	geoos.Collection{geoos.Collection{geoos.Point{}}},
}

func TestMarshal(t *testing.T) {
	for _, g := range AllGeometries {
		Marshal(g, binary.BigEndian)
	}
}

func TestMustMarshal(t *testing.T) {
	for _, g := range AllGeometries {
		MustMarshal(g, binary.BigEndian)
	}
}

func TestGeoFromWKBHexStr(t *testing.T) {
	hexStr := `0101000020E61000008EAF3DB324E05C40DC12B9E00C704340`
	g0 := geoos.Point{115.50224, 38.875393}
	g1, _ := GeoFromWKBHexStr(hexStr)
	if g0 != g1 {
		t.Errorf("GeoFromWKBHexStr() got = %v, want %v", g0, g1)
	}
}
func BenchmarkEncode_Point(b *testing.B) {
	g := geoos.Point{1, 2}
	e := NewEncoder(ioutil.Discard)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Encode(g)
	}
}

func BenchmarkEncode_LineString(b *testing.B) {
	g := geoos.LineString{
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
		{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5},
	}
	e := NewEncoder(ioutil.Discard)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e.Encode(g)
	}
}

func compare(t testing.TB, e geoos.Geometry, b []byte) {
	t.Helper()

	// Decoder
	g, err := NewDecoder(bytes.NewReader(b)).Decode()
	if err != nil {
		t.Fatalf("decoder: read error: %v", err)
	}

	if !geoos.Equal(g, e) {
		t.Errorf("decoder: incorrect geometry: %v != %v", g, e)
	}

	g, err = Unmarshal(b)
	if err != nil {
		t.Fatalf("unmarshal: read error: %v", err)
	}

	if !geoos.Equal(g, e) {
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
		t.Errorf("marshal: incorrent encoding")
	}

	// preallocation
	if len(data) != geomLength(e) {
		t.Errorf("prealloc length: %v != %v", len(data), geomLength(e))
	}

	// Scanner
	var sg geoos.Geometry

	switch e.(type) {
	case geoos.Point:
		var p geoos.Point
		err = Scanner(&p).Scan(b)
		sg = p
	case geoos.MultiPoint:
		var mp geoos.MultiPoint
		err = Scanner(&mp).Scan(b)
		sg = mp
	case geoos.LineString:
		var ls geoos.LineString
		err = Scanner(&ls).Scan(b)
		sg = ls
	case geoos.MultiLineString:
		var mls geoos.MultiLineString
		err = Scanner(&mls).Scan(b)
		sg = mls
	case geoos.Polygon:
		var p geoos.Polygon
		err = Scanner(&p).Scan(b)
		sg = p
	case geoos.MultiPolygon:
		var mp geoos.MultiPolygon
		err = Scanner(&mp).Scan(b)
		sg = mp
	case geoos.Collection:
		var c geoos.Collection
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

	if !geoos.Equal(sg, e) {
		t.Errorf("scan: incorrect geometry: %v != %v", sg, e)
	}
}
