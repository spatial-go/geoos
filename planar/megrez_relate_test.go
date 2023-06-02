package planar

import (
	"fmt"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/geoencoding/wkt"
	"github.com/spatial-go/geoos/space"
	"github.com/spatial-go/geoos/space/topograph"
)

type args struct {
	g1 space.Geometry
	g2 space.Geometry
}
type TestStruct struct {
	name    string
	args    args
	want    bool
	wantErr bool
}

var polyTestscase = []struct {
	name string
	args args
}{
	{name: "poly 1",
		args: args{g1: space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			g2: space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}},
		},
	},
	{name: "poly 2",
		args: args{g1: space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			g2: space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}},
		},
	},

	{name: "poly 3",
		args: args{g1: space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			g2: space.Polygon{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}},
		},
	},

	{name: "poly 4",
		args: args{g1: space.Polygon{{{1, 2}, {3, 2}, {3, 3}, {1, 3}, {1, 2}}},
			g2: space.Polygon{{{2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 1}}},
		},
	},

	{name: "poly 5",
		args: args{g1: space.Polygon{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			g2: space.Polygon{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}},
		},
	},

	{name: "poly 6",
		args: args{g1: space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			g2: space.Polygon{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
		},
	},
}

func TestAlgorithm_Contains(t *testing.T) {
	const polygon0 = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const polygon1 = `POLYGON((1 1, 5 1, 5 5, 1 5, 1 1))`

	const point1 = `POINT(3 3)`
	const point2 = `POINT(-1 35)`

	const polygon2 = `POLYGON((113.48581807062143 23.33621329259057,113.48785155569199 23.336383160940287,113.48792833916376 23.335028970144833,113.48582512451681 23.33493907732756,113.48581807062143 23.33621329259057))`
	const polygon3 = `POLYGON((113.48668269733025 23.335774475286513,113.48720314737876 23.3358047169134,113.4873120145494 23.335572135724533,113.48676140749377 23.33538486387792,113.48668269733025 23.335774475286513))`

	p1, _ := wkt.UnmarshalString(point1)
	p2, _ := wkt.UnmarshalString(point2)

	poly0, _ := wkt.UnmarshalString(polygon0)
	poly1, _ := wkt.UnmarshalString(polygon1)

	poly2, _ := wkt.UnmarshalString(polygon2)
	poly3, _ := wkt.UnmarshalString(polygon3)

	tests := []TestStruct{
		{name: "contain", args: args{
			g1: poly0,
			g2: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly0,
			g2: p2,
		}, want: false, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly0,
			g2: poly1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly2,
			g2: poly3,
		}, want: true, wantErr: false},
	}
	for i, v := range polyTestscase {
		want := true
		if i != 4 {
			want = false
		}
		tests = append(tests, TestStruct{v.name, v.args, want, false})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Contains(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Contains() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Covers(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const point1 = `POINT(3 3)`
	const point2 = `POINT(-1 35)`

	p1, _ := wkt.UnmarshalString(point1)
	p2, _ := wkt.UnmarshalString(point2)
	poly, _ := wkt.UnmarshalString(polygon)
	tests := []TestStruct{
		{name: "contain", args: args{
			g1: poly,
			g2: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly,
			g2: p2,
		}, want: false, wantErr: false},
	}
	for i, v := range polyTestscase {
		want := true
		if i != 4 {
			want = false
		}
		tests = append(tests, TestStruct{v.name, v.args, want, false})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Covers(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Covers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Covers()%v got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestAlgorithm_CoveredBy(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const point1 = `POINT(3 3)`
	const point2 = `POINT(-1 35)`

	p1, _ := wkt.UnmarshalString(point1)
	p2, _ := wkt.UnmarshalString(point2)
	poly, _ := wkt.UnmarshalString(polygon)

	tests := []TestStruct{
		{name: "contain", args: args{
			g2: poly,
			g1: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g2: poly,
			g1: p2,
		}, want: false, wantErr: false},
	}
	for i, v := range polyTestscase {
		want := true
		if i != 5 {
			want = false
		}
		tests = append(tests, TestStruct{v.name, v.args, want, false})
	}
	for i, tt := range tests {
		if i < 0 {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.CoveredBy(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("CoveredBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CoveredBy()%v got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Crosses(t *testing.T) {
	const g1 = `LINESTRING(0 0, 10 10)`
	const g2 = `LINESTRING(10 0, 0 10)`

	geom1, _ := wkt.UnmarshalString(g1)
	geom2, _ := wkt.UnmarshalString(g2)
	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "crosses", args: args{
			g1: geom1,
			g2: geom2,
		}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Crosses(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crosses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Crosses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Disjoint(t *testing.T) {
	point01, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line01, _ := wkt.UnmarshalString(`LINESTRING ( 2 0, 0 2 )`)

	point02, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING ( 0 0, 0 2 )`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "disjoint", args: args{g1: point01, g2: line01}, want: true},
		{name: "not disjoint", args: args{g1: point02, g2: line02}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Disjoint(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Disjoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Disjoint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Intersects(t *testing.T) {
	point01, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line01, _ := wkt.UnmarshalString(`LINESTRING ( 0 0, 0 2 )`)

	point02, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING ( 2 1, 1 2 )`)

	const a1 = "LINESTRING (119.3039141328787309 26.0688567207556332, 119.3036083610503510 26.0681869226817753, 119.3033186824765153 26.0679845368985639)"
	const a2 = "POLYGON ((119.3172630350001100 26.0702955380000390, 119.3172497400000793 26.0703321550000737, 119.3172290960000055 26.0703549800000474, 119.3171979140000758 26.0703685560000622, 119.3171192560000691 26.0703590920000465, 119.3150906640000812 26.0699619740000230, 119.3149878960000478 26.0699494980000281, 119.3143222760000981 26.0698098000000300, 119.3139883680000821 26.0697879690000605, 119.3132765640000343 26.0697358590000476, 119.3132011550001153 26.0697304560000589, 119.3129190370000288 26.0697096750000696, 119.3111563600000409 26.0695852310000760, 119.3109965890000694 26.0695196390000774, 119.3107551660000354 26.0694689840000251, 119.3106054270000413 26.0694540860000643, 119.3105945770000744 26.0694530400000417, 119.3105856270000231 26.0694516730000601, 119.3105761370001119 26.0694508760000758, 119.3105671780000421 26.0694500390000599, 119.3105579580001177 26.0694489820000399, 119.3105487380000795 26.0694481760000372, 119.3105397880000282 26.0694470890000503, 119.3105305480000879 26.0694462820000581, 119.3105216080000446 26.0694454650000580, 119.3105123880000065 26.0694443880000790, 119.3105031680000820 26.0694435710000789, 119.3104942080000228 26.0694427650000762, 119.3104849880000984 26.0694416580000734, 119.3104681690000461 26.0694400640000481, 119.3104369790000874 26.0694370640000557, 119.3103078550000191 26.0693934490000743, 119.3102112770000076 26.0693755690000444, 119.3099221230000921 26.0693252120000238, 119.3088639410000269 26.0691093130000695, 119.3086276680000992 26.0690527240000733, 119.3084014460000617 26.0689920200000529, 119.3078483490000963 26.0688210030000391, 119.3075865890000387 26.0687429470000325, 119.3073801560000220 26.0686865570000350, 119.3070351180000443 26.0685920170000713, 119.3069911800000682 26.0685800900000686, 119.3069032920000154 26.0685624780000467, 119.3066070560000753 26.0685197070000640, 119.3064738660000330 26.0685132670000712, 119.3055648530000781 26.0684864910000442, 119.3053117620000876 26.0684792470000275, 119.3007346390000976 26.0683380540000371, 119.3006188190000785 26.0683266760000265, 119.3005433790000325 26.0683196640000574, 119.3005168190001086 26.0683183240000744, 119.3001297180001075 26.0682964810000612, 119.2996200020000970 26.0682958950000625, 119.2993834490000609 26.0683003320000353, 119.2991243780001014 26.0682819690000542, 119.2983976580001126 26.0682839130000730, 119.2981323310000334 26.0683084390000772, 119.2977460480000218 26.0683579540000778, 119.2974045160000287 26.0684075110000322, 119.2965745140000990 26.0685377980000226, 119.2957896610000716 26.0685491330000332, 119.2954264850000072 26.0685515730000361, 119.2948474920000308 26.0685097760000417, 119.2948101140000290 26.0686241440000686, 119.2953998390000834 26.0686505420000572, 119.2958137050001142 26.0686534930000562, 119.2965880250000055 26.0686596380000424, 119.2974137140000721 26.0685480800000278, 119.2977308180001046 26.0684904100000381, 119.2981231100000059 26.0684443940000392, 119.2983781080000654 26.0684130890000461, 119.2990806990001147 26.0684019210000315, 119.2993546810000680 26.0684102540000708, 119.2995790230000921 26.0684185770000454, 119.3000887390001026 26.0684145510000462, 119.3005572210000764 26.0684423330000641, 119.3005629010000348 26.0684431410000457, 119.3006383100000676 26.0684536950000734, 119.3006567700000460 26.0684561280000366, 119.3006977290000350 26.0684620840000321, 119.3007180780000454 26.0684740200000533, 119.3007218740000326 26.0684981680000760, 119.3053149980000853 26.0686393630000453, 119.3055637490000436 26.0686441480000326, 119.3064667840000084 26.0686581810000462, 119.3065940230001161 26.0686689850000448, 119.3067073910000317 26.0686865780000403, 119.3068311080000967 26.0687150200000701, 119.3069512740000846 26.0687434730000405, 119.3069881630000282 26.0687543130000563, 119.3073364600001014 26.0688575640000408, 119.3074984030000678 26.0689052580000293, 119.3077775230000270 26.0689833080000426, 119.3085641160000705 26.0691944590000730, 119.3088245570000936 26.0692634930000509, 119.3104337000000896 26.0695586580000622, 119.3104681500000197 26.0695621660000256, 119.3105988980000802 26.0695759610000550, 119.3107415770000443 26.0695908320000740, 119.3109618480000336 26.0696518090000495, 119.3111069680001037 26.0697228380000752, 119.3128978550000738 26.0698554030000764, 119.3131954130000167 26.0698761580000564, 119.3132700220000970 26.0698810410000306, 119.3139812860000575 26.0699304300000563, 119.3143030130000852 26.0699590790000570, 119.3149478770000087 26.0700692340000728, 119.3150706170000603 26.0700736100000654, 119.3170849910001152 26.0704565680000542, 119.3172141570000804 26.0705490010000744, 119.3172494210000423 26.0705875310000579, 119.3172630350001100 26.0702955380000390))"
	g1, _ := wkt.UnmarshalString(a1)
	g2, _ := wkt.UnmarshalString(a2)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "intersects", args: args{g1: point01, g2: line01}, want: true},
		{name: "not intersects", args: args{g1: point02, g2: line02}, want: false},
		{name: "issue 89", args: args{g1: g1, g2: g2}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Intersects(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Intersects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Intersects() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Relate(t *testing.T) {

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	type TestStruct struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	tests := []TestStruct{}

	tests = append(tests, TestStruct{fmt.Sprintf("Disjoint%v", "00"),
		args{space.Point{0, 0}, space.LineString{{2, 0}, {0, 2}}}, "FF0FFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("inter%v", "3-6"),
		args{space.Point{3, 3}, space.Polygon{{{0, 0}, {6, 0}, {6, 6}, {0, 6}, {0, 0}}}}, "0FFFFF212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("linepoint%v", "00"),
		args{space.Point{1, 1}, space.LineString{{0, 0}, {1, 1}, {0, 2}}}, "0FFFFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("linepoint%v", "01"),
		args{space.Point{0, 2}, space.LineString{{0, 0}, {1, 1}, {0, 2}}}, "F0FFFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "0f"),
		args{space.Polygon{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			space.Polygon{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, "2FF11F212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f"),
		args{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			space.Polygon{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}}}, "2FF11F212", false})

	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f1"),
		args{space.Polygon{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			space.Polygon{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}}, "212FF1FF2", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f2"),
		args{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, "FF2F11212", false})

	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f3"),
		args{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			space.Polygon{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}}}, "FF2F01212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "f1"),
		args{space.Polygon{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			space.Polygon{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}}, "212F11FF2", false})

	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "LineLine0" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			//G := GEOAlgorithm{}
			G := NormalStrategy()
			got, err := G.Relate(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("%v error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {

				s, d := tt.args.g1, tt.args.g2
				intersectBound := s.Bound().IntersectsBound(d.Bound())
				if s.Bound().ContainsBound(d.Bound()) || d.Bound().ContainsBound(s.Bound()) {
					intersectBound = true
				}

				t.Errorf("%v got = %v, want %v intersect %v", tt.name, got, tt.want, intersectBound)
				return
			}
		})
	}
}

func TestAlgorithm_Touches(t *testing.T) {
	line01, _ := wkt.UnmarshalString(`LINESTRING(0 0, 1 1, 0 2)`)
	point01, _ := wkt.UnmarshalString(`POINT(0 2)`)

	line02, _ := wkt.UnmarshalString(`LINESTRING(0 0, 1 1, 0 2)`)
	point02, _ := wkt.UnmarshalString(`POINT(1 1)`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "touches", args: args{g1: line01, g2: point01}, want: true},
		{name: "not touches", args: args{g1: line02, g2: point02}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Touches(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Touches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Touches() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlgorithm_Within(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const point1 = `POINT(3 3)`
	const point2 = `POINT(-1 35)`

	p1, _ := wkt.UnmarshalString(point1)
	p2, _ := wkt.UnmarshalString(point2)
	poly, _ := wkt.UnmarshalString(polygon)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "in", args: args{
			g1: p1,
			g2: poly,
		}, want: true, wantErr: false},
		{name: "not in", args: args{
			g1: p2,
			g2: poly,
		}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Within(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Within() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Within() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Overlaps(t *testing.T) {
	type args struct {
		A space.Geometry
		B space.Geometry
	}
	tests := []struct {
		name    string
		g       *megrezAlgorithm
		args    args
		want    bool
		wantErr bool
	}{
		{name: "polypoly",
			args: args{space.Polygon{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
				space.Polygon{{{100, 100}, {100, 102}, {102, 102}, {102, 100}, {100, 100}}},
			},
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Overlaps(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("MegrezAlgorithm.Overlaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MegrezAlgorithm.Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_megrezAlgorithm_Adjacency(t *testing.T) {
	type fields struct {
		topog topograph.Relationship
	}
	type args struct {
		A space.Geometry
		B space.Geometry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &megrezAlgorithm{
				topog: tt.fields.topog,
			}
			got, err := g.Adjacency(tt.args.A, tt.args.B)
			if (err != nil) != tt.wantErr {
				t.Errorf("megrezAlgorithm.Adjacency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("megrezAlgorithm.Adjacency() = %v, want %v", got, tt.want)
			}
		})
	}
}
