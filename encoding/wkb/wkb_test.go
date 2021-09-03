package wkb

import (
	"bytes"
	"io/ioutil"
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
		_, _ = Marshal(g, bigEndian)
	}
}

func TestMustMarshal(t *testing.T) {
	for _, g := range AllGeometries {
		MustMarshal(g, bigEndian)
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
		data, err = Marshal(g, bigEndian)
	} else {
		data, err = Marshal(g, littleEndian)
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
		{" GeomFromWKBHexStr point1 ", args{"0101000020E610000000000020D8135D400000004072054440"},
			space.Point{116.310066223145, 40.0425491333008}, false},
		{" GeomFromWKBHexStr point11 ", args{"0101000020e610000021000020d8135d400300004072054440"},
			space.Point{116.310066223145, 40.0425491333008}, false},
		{" GeomFromWKBHexStr point2 ", args{"0101000020E6100000A9E2F33378145D4088C78C29E2064440"},
			space.Point{116.319836605234, 40.0537769257926}, false},

		{" GeomFromWKBHexStr line1 ", args{"0102000020E610000004000000F7FFFF7F20155D40C9D9B446F6F843400F000020B51C5D409241C66566F94340DDFFFFFF791D5D40336A670189F04340E8FFFF5FA7175D409DF9A3B974EF4340"},
			space.LineString{
				{116.33010864257814, 39.94501575308417},
				{116.44855499267578, 39.948437425427215},
				{116.4605712890625, 39.87918107556866},
				{116.36959075927736, 39.87074966913789}},
			false},
		{" GeomFromWKBHexStr line2 ", args{"0102000020E610000003000000E6FFFF7F81865D406F22C76A1FDD4340ECFFFF1F87865D40B2357F117DD14340070000A0C18F5D40BECAC40CB1D14340"},
			space.LineString{
				{118.10165405273438, 39.72752127383585},
				{118.10199737548828, 39.63662928306019},
				{118.2461929321289, 39.63821563347804}},
			false},

		{" GeomFromWKBHexStr poly1 ", args{"0103000020E610000001000000050000001F00004092885D403B5CC8079ADB4340E1FFFFBFE9835D403AFA86B354D4434000000000388C5D40A40C9887CCCD434009000080FF8F5D407E533CEFF6D843401F00004092885D403B5CC8079ADB4340"},
			space.Polygon{
				{{118.13392639160155, 39.715638134796336},
					{118.06114196777345, 39.658834877879094},
					{118.19091796875, 39.607804249995105},
					{118.24996948242188, 39.69503584333047},
					{118.13392639160155, 39.715638134796336}}},
			false},
		{" GeomFromWKBHexStr poly2 ", args{"0103000020E61000000100000005000000F7FFFF7FA4535E401A5D7FAD5B0F3F400000000098675E401A5D7FAD5B0F3F400000000098675E403E11FC5905503F40F7FFFF7FA4535E403E11FC5905503F40F7FFFF7FA4535E401A5D7FAD5B0F3F40"},
			space.Polygon{
				{{121.30691528320312, 31.05999264106237},
					{121.61865234375, 31.05999264106237},
					{121.61865234375, 31.312581657447687},
					{121.30691528320312, 31.312581657447687},
					{121.30691528320312, 31.05999264106237}}},
			false},

		{" GeomFromWKBHexStr multipoint1 ", args{"0104000020E610000002000000010100000060D7C48B03165D40E5E77531550044400101000000D7133C7E521D5D40F9512DBA77004440"},
			space.MultiPoint{
				{116.343966428973715, 40.002599890300026},
				{116.458159979505325, 40.003653785828554}}, false},
		{" GeomFromWKBHexStr multipoint2 ", args{"0101000020E6100000A9E2F33378145D4088C78C29E2064440"}, space.Point{116.319836605234, 40.0537769257926}, false},

		{" GeomFromWKBHexStr  multiline1 ", args{"0105000020E61000000100000001020000000400000071F3D48E17045D40C3FC5C0DF12B444028912F857FF25C40459494346B094440520B351CBE095D40FE5D20575B014440520B351CBE095D40FE5D20575B014440"},
			space.MultiLineString{
				{{116.063937862357662, 40.343293829349477},
					{115.789033218815135, 40.073584148929967},
					{116.15222840480925, 40.010599985889741},
				}}, false},
		{" GeomFromWKBHexStr  multiline2 ", args{"0105000020E6100000020000000102000000030000001BCA4628A40A5D40DFE069B3ACE64340F4F4BBA8F1F85C40CC5CB05207D8434080C663FFED0C5D404DADF8271BA14340010200000004000000C01C4E0089495D4041D2B9E9EE874340388C2C22F02B5D40CF63FFED6C8D434032D11E0D031A5D403A8E62DBC7074440D099B8C5C8425D40D28E746EBAFB4340"},
			space.MultiLineString{
				{{116.166269368295588, 39.802145411203838},
					{115.889749702026222, 39.687723480333759},
					{116.202026221692492, 39.258641239570942}},
				{{117.148986889153733, 39.061978545887996},
					{116.686531585220479, 39.104886769964274},
					{116.406436233611402, 40.060786650774759},
					{117.04350417163289, 39.966626936829584}}}, false},

		{" GeomFromWKBHexStr  multipoly1 ", args{"0106000020E61000000200000001030000000100000006000000DE7AABC26E235D40B984CBF3970244405F3017ACD4215D40DBB5F4325FF843408240EB9B77285D4045BC313BD2F9434051FC938D50285D40266E31B28800444051FC938D50285D40266E31B288004440DE7AABC26E235D40B984CBF39702444001030000000100000005000000743F50BF8B1A5D403BB288C0AF00444001BE67F4A9155D407451E3B1FFF64340CFEC48E7A61B5D403B3BD2B9E9EE4340CFEC48E7A61B5D403B3BD2B9E9EE4340743F50BF8B1A5D403BB288C0AF004440"},
			space.MultiPolygon{{
				{{116.553635280095321, 40.020262216924934},
					{116.528605482717495, 39.940405244338521},
					{116.632300357568511, 39.951728247914211},
					{116.62991656734205, 40.00417163289633},
					{116.553635280095321, 40.020262216924934}}},
				{{{116.414779499404034, 40.005363528009561},
					{116.338498212157305, 39.92967818831945},
					{116.43206197854586, 39.866507747318259},
					{116.414779499404034, 40.005363528009561}}}}, false},

		{" GeomFromWKBHexStr collection1 ", args{"0107000020E610000002000000010200000002000000000000000000F03F0000000000000040000000000000084000000000000010400101000000000000000000F03F0000000000000040"},
			space.Collection{space.LineString{{1, 2}, {3, 4}}, space.Point{1, 2}}, false},
		{" GeomFromWKBHexStr collection2 ", args{"0107000020E610000002000000010200000002000000EFD5255C9E8F5D40253878DA2CD34340DAF78C3EB5925D404A94169A41D74340010100000047CF5F0A32945D4020C5B68F0FD34340"},
			space.Collection{space.LineString{{118.244040524434, 39.6498063170441}, {118.29231227652, 39.681689511323}}, space.Point{118.315554231228, 39.6489123957092}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeomFromWKBHexStr(tt.args.wkbHex)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeomFromWKBHexStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.EqualsExact(tt.want, 0.00000000001) {
				t.Errorf("GeomFromWKBHexStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeomToWKBHexStr(t *testing.T) {
	type args struct {
		geom space.Geometry
	}
	tests := []struct {
		name       string
		args       args
		wantWkbHex string
		wantErr    bool
	}{

		{name: "Point0",
			args:       args{space.Point{116.310066223145, 40.0425491333008}},
			wantWkbHex: "0101000020e610000021000020d8135d400300004072054440",
			wantErr:    false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWkbHex, err := GeomToWKBHexStr(tt.args.geom)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeomToWKBHexStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWkbHex != tt.wantWkbHex {
				t.Errorf("GeomToWKBHexStr() = %v, want %v", gotWkbHex, tt.wantWkbHex)
			}
		})
	}
}
