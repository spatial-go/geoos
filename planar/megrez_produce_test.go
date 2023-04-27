package planar

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/debugtools"
	"github.com/spatial-go/geoos/geoencoding/wkt"
	"github.com/spatial-go/geoos/space"
)

func TestAlgorithm_Boundary(t *testing.T) {
	const sourceLine = `LINESTRING(1 1,0 0, -1 1)`
	const expectLine = `MULTIPOINT(1 1,-1 1)`
	sLine, _ := wkt.UnmarshalString(sourceLine)
	eLine, _ := wkt.UnmarshalString(expectLine)

	const sourcePolygon = `POLYGON((1 1,0 0, -1 1, 1 1))`
	const expectPolygon = `LINESTRING(1 1,0 0,-1 1,1 1)`
	sPolygon, _ := wkt.UnmarshalString(sourcePolygon)
	ePolygon, _ := wkt.UnmarshalString(expectPolygon)

	// const multiPolygon = `POLYGON (( 10 130, 50 190, 110 190, 140 150, 150 80, 100 10, 20 40, 10 130 ),
	// ( 70 40, 100 50, 120 80, 80 110, 50 90, 70 40 ))`
	// const expectMultiPolygon = `MULTILINESTRING((10 130,50 190,110 190,140 150,150 80,100 10,20 40,10 130),
	// (70 40,100 50,120 80,80 110,50 90,70 40))`

	sMultiPolygon, _ := wkt.UnmarshalString(sourceLine)
	eMultiPolygon, _ := wkt.UnmarshalString(expectLine)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "line", args: args{g: sLine}, want: eLine, wantErr: false},
		{name: "polygon", args: args{g: sPolygon}, want: ePolygon, wantErr: false},
		{name: "multiPolygon", args: args{g: sMultiPolygon}, want: eMultiPolygon, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Boundary(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Boundary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(wkt.MarshalString(got))
			t.Log(wkt.MarshalString(tt.want))
			if !got.Equals(tt.want) {
				t.Errorf("Boundary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Buffer(t *testing.T) {
	geometry, _ := wkt.UnmarshalString("POINT(100 90)")
	expectGeometry, _ := wkt.UnmarshalString("POLYGON((150 90,146.193976625564 70.8658283817455,135.355339059327 54.6446609406727,119.134171618255 43.8060233744357,100 40,80.8658283817456 43.8060233744356,64.6446609406727 54.6446609406725,53.8060233744357 70.8658283817454,50 89.9999999999998,53.8060233744356 109.134171618254,64.6446609406725 125.355339059327,80.8658283817453 136.193976625564,99.9999999999998 140,119.134171618254 136.193976625564,135.355339059327 125.355339059328,146.193976625564 109.134171618255,150 90))")

	line, _ := wkt.UnmarshalString("LINESTRING(2.073333263397217 48.81027603149414, 1.5225944519042969 48.45795440673828,1.361388921737671 48.705833435058594,2.073333263397217 48.81027603149414,2.0408332347869873 49.0966682434082)")
	expectline, _ := wkt.UnmarshalString("POLYGON((264.6446609406726 335.3553390593274,280.8658283817456 346.19397662556435,300.00000000000006 350,319.1341716182545 346.19397662556435,335.3553390593274 335.3553390593274,346.19397662556435 319.1341716182545,350 300.00000000000006,346.19397662556435 280.8658283817456,335.3553390593274 264.6446609406726,135.35533905932738 64.64466094067262,119.13417161825444 53.80602337443564,99.99999999999997 50,80.8658283817455 53.80602337443566,64.64466094067262 64.64466094067262,53.806023374435675 80.86582838174549,50 99.99999999999994,53.80602337443563 119.13417161825441,64.64466094067262 135.35533905932738,264.6446609406726 335.3553390593274))")

	type args struct {
		g        space.Geometry
		width    float64
		quadsegs int
	}
	tests := []struct {
		name string
		args args
		want space.Geometry
	}{
		{name: "buffer", args: args{g: geometry, width: 50, quadsegs: 4}, want: expectGeometry},
		{name: "line buffer", args: args{g: line, width: 0.01, quadsegs: 10}, want: expectline},
		{name: "point buffer", args: args{
			g:        space.Point{100, 90},
			width:    50,
			quadsegs: 4,
		}, want: space.Polygon{
			{{150, 90}, {146.193976625564, 70.8658283817455}, {135.355339059327, 54.6446609406727},
				{119.134171618255, 43.8060233744357}, {100, 40}, {80.8658283817456, 43.8060233744356}, {64.6446609406727, 54.6446609406725},
				{53.8060233744357, 70.8658283817454}, {50, 89.9999999999998}, {53.8060233744356, 109.134171618254}, {64.6446609406725, 125.355339059327},
				{80.8658283817453, 136.193976625564}, {99.9999999999998, 140}, {119.134171618254, 136.193976625564}, {135.355339059327, 125.355339059328},
				{146.193976625564, 109.134171618255}, {150, 90}},
		},
		},
		{name: "polygon buffer", args: args{
			g: space.Polygon{{{12695002.300434208, 2578061.0407920154}, {12695011.873910416, 2577943.999133326},
				{12695182.081411814, 2577948.097385522}, {12695210.133923491, 2578044.286065328}, {12695102.042697946, 2578089.728780503},
				{12695002.300434208, 2578061.0407920154}}},
			width:    108,
			quadsegs: 8,
		}, want: space.Polygon{
			{{1.2694972447614357e+07, 2.578164832935971e+06}, {1.2694952801158171e+07, 2.5781570294458857e+06}, {1.2694935051145962e+07, 2.5781455483847535e+06},
				{1.2694919877626132e+07, 2.5781308296214156e+06}, {1.2694907861935072e+07, 2.578113437069374e+06}, {1.2694899464424662e+07, 2.5780940370818204e+06}, {1.2694895006825024e+07, 2.578073372921938e+06},
				{1.269489465991821e+07, 2.57805223628659e+06}, {1.2694904233394418e+07, 2.5779351946279006e+06}, {1.2694907808887677e+07, 2.577915111920888e+06}, {1.2694915096776368e+07, 2.5778960597306876e+06},
				{1.269492583707376e+07, 2.5778787177214357e+06}, {1.2694939646632573e+07, 2.577863704548606e+06}, {1.269495603281328e+07, 2.577851555789213e+06}, {1.2694974411058454e+07, 2.5778427048357427e+06},
				{1.269499412574614e+07, 2.5778374674353916e+06}, {1.2695014473578395e+07, 2.577836030426163e+06},
				{1.2695184681079794e+07, 2.577840128678359e+06}, {1.2695207182815429e+07, 2.5778430549201253e+06}, {1.2695228576484643e+07, 2.5778506181148915e+06},
				{1.2695247917693771e+07, 2.5778624843959813e+06}, {1.2695264352652138e+07, 2.5778781299429825e+06}, {1.2695277155861448e+07, 2.577896864105062e+06}, {1.26952857621419e+07, 2.5779178598887883e+06},
				{1.2695313814653577e+07, 2.5780140485685943e+06}, {1.2695317824740773e+07, 2.5780361198009993e+06}, {1.2695317188721681e+07, 2.578058543350649e+06}, {1.2695311934036078e+07, 2.5780803517980585e+06},
				{1.2695302287386933e+07, 2.578100604261106e+06}, {1.2695288664959753e+07, 2.578118426987522e+06}, {1.2695271654467084e+07, 2.578133051051196e+06}, {1.2695251989792826e+07, 2.5781438455259646e+06},
				{1.269514389856728e+07, 2.5781892882411396e+06}, {1.26951205234274e+07, 2.578196135841541e+06}, {1.2695096208269283e+07, 2.578197571070177e+06}, {1.2695072189878095e+07, 2.5781935209244587e+06}, {1.2694972447614357e+07, 2.578164832935971e+06}},
		},
		},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "line buffer" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry := G.Buffer(tt.args.g, tt.args.width, tt.args.quadsegs)
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Buffer() = %v\n, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))

				// gotGeometry, _ = G.SimplifyP(gotGeometry, 0.9)

				debugtools.WriteGeom("buffer_test.geojson", gotGeometry)
			}
			// t.Log(wkt.MarshalString(gotGeometry))
		})
	}
}

func TestAlgorithm_Centroid(t *testing.T) {
	const multipoint = `MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`
	geometry, _ := wkt.UnmarshalString(multipoint)
	const pointresult = `POINT(2.3076923076923075 3.3076923076923075)`

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "point", args: args{g: geometry}, want: pointresult, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Centroid(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Centroid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			s := wkt.MarshalString(got)
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Centroid() got = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestAlgorithm_ConvexHull(t *testing.T) {
	polygon, _ := wkt.UnmarshalString(`POLYGON((1 1, 3 1, 2 2, 3 3, 1 3, 1 1))`)
	expectPolygon, _ := wkt.UnmarshalString(`POLYGON ((1 1, 1 3, 3 3, 3 1, 1 1))`)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "ConvexHull Polygon", args: args{g: polygon}, want: expectPolygon, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.ConvexHull(tt.args.g)

			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Envelope() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_Envelope(t *testing.T) {
	point, _ := wkt.UnmarshalString(`POINT(1 3)`)
	expectPoint, _ := wkt.UnmarshalString(`POINT(1 3)`)

	line, _ := wkt.UnmarshalString(`LINESTRING(0 0, 1 3)`)
	expectPolygon0, _ := wkt.UnmarshalString(`POLYGON((0 0,1 0,1 3,0 3,0 0))`)

	polygon1, _ := wkt.UnmarshalString(`POLYGON((0 0, 0 1, 1.0000001 1, 1.0000001 0, 0 0))`)
	expectPolygon1, _ := wkt.UnmarshalString(`POLYGON((0 0,1.0000001 0,1.0000001 1,0 1,0 0))`)

	polygon2, _ := wkt.UnmarshalString(`POLYGON((0 0, 0 1, 1.0000000001 1, 1.0000000001 0, 0 0))`)
	expectPolygon2, _ := wkt.UnmarshalString(`POLYGON((0 0,1.0000000001 0,1.0000000001 1,0 1,0 0))`)

	type args struct {
		g space.Geometry
	}

	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "envelope Point", args: args{g: point}, want: expectPoint, wantErr: false},
		{name: "envelope LineString", args: args{g: line}, want: expectPolygon0, wantErr: false},
		{name: "envelope Polygon", args: args{g: polygon1}, want: expectPolygon1, wantErr: false},
		{name: "envelope Polygon", args: args{g: polygon2}, want: expectPolygon2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.Envelope(tt.args.g)

			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Envelope() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_PointOnSurface(t *testing.T) {
	point, _ := wkt.UnmarshalString(`POINT(0 5)`)
	expectPoint0, _ := wkt.UnmarshalString(`POINT(0 5)`)

	lineString, _ := wkt.UnmarshalString(`LINESTRING(0 5, 0 10)`)
	expectPoint1, _ := wkt.UnmarshalString(`POINT(0 5)`)

	polygon, _ := wkt.UnmarshalString(`POLYGON((0 0, 0 5, 5 5, 5 0, 0 0))`)
	expectPoint2, _ := wkt.UnmarshalString(`POINT(2.5 2.5)`)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "PointOnSurface Point", args: args{g: point}, want: expectPoint0, wantErr: false},
		{name: "PointOnSurface LineString0", args: args{g: lineString}, want: expectPoint1, wantErr: false},
		{name: "PointOnSurface Polygon", args: args{g: polygon}, want: expectPoint2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.PointOnSurface(tt.args.g)

			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Envelope() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_Simplify(t *testing.T) {
	lineString, _ := wkt.UnmarshalString(`LINESTRING(0 0, 1 1, 0 2, 1 3, 0 4, 1 5)`)
	expectLine, _ := wkt.UnmarshalString(`LINESTRING (0 0, 1 5)`)

	type args struct {
		g         space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "Simplify Point", args: args{g: lineString, tolerance: 1.0}, want: expectLine, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.Simplify(tt.args.g, tt.args.tolerance)

			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Simplify() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_SimplifyP(t *testing.T) {
	lineString, _ := wkt.UnmarshalString(`LINESTRING(0 0, 1 1, 0 2, 1 3, 0 4, 1 5)`)
	expectLine, _ := wkt.UnmarshalString(`LINESTRING (0 0, 1 5)`)

	type args struct {
		g         space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "SimplifyP Point", args: args{g: lineString, tolerance: 1.0}, want: expectLine, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.SimplifyP(tt.args.g, tt.args.tolerance)

			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Envelope() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_Snap(t *testing.T) {
	type args struct {
		input     space.Geometry
		reference space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "snap", args: args{
			input:     space.Point{0.05, 0.05},
			reference: space.Point{0, 0},
			tolerance: 0.1,
		}, want: space.Point{0, 0}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Snap(tt.args.input, tt.args.reference, tt.args.tolerance)
			if (err != nil) != tt.wantErr {
				t.Errorf("Snap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("Snap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_UniquePoints(t *testing.T) {
	const polygon = `POLYGON((0 0, 6 0, 6 6, 0 6, 0 0))`
	const multipoint = `MULTIPOINT((0 0),(6 0),(6 6),(0 6))`

	poly, _ := wkt.UnmarshalString(polygon)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "uniquepoints", args: args{g: poly}, want: multipoint, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.UniquePoints(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniquePoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			res := wkt.MarshalString(got)
			t.Log(res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("UniquePoints() got = %v, want %v", res, tt.want)
			}
		})
	}
}

func TestAlgorithm_BufferInMeter(t *testing.T) {
	wantGeometry, _ := wkt.UnmarshalString("POLYGON((110.00117265646337 40.00000000000001,110.00115012419823 39.99982474877957,110.00108339330515 39.999656231941024,110.00097502821495 39.99950092556494,110.00082919333724 39.999364798097254,110.00065149302459 39.99925308096896,110.00044875620038 39.99917006753527,110.00022877392705 39.99911894806501,110.00000000000001 39.99910168712361,109.99977122607295 39.99911894806501,109.99955124379962 39.99917006753527,109.9993485069754 39.99925308096896,109.99917080666276 39.999364798097254,109.99902497178505 39.99950092556494,109.99891660669483 39.999656231941024,109.99884987580177 39.99982474877957,109.99882734353663 40.00000000000001,109.99884987580177 40.00017525077067,109.99891660669483 40.00034376632828,109.99902497178505 40.00049907078736,109.99917080666276 40.000635195993794,109.9993485069754 40.00074691086084,109.99955124379962 40.00082992237753,109.99977122607295 40.00088104056688,110.00000000000001 40.00089830105848,110.00022877392705 40.00088104056688,110.00044875620038 40.00082992237753,110.00065149302459 40.00074691086084,110.00082919333724 40.000635195993794,110.00097502821495 40.00049907078736,110.00108339330515 40.00034376632828,110.00115012419823 40.00017525077067,110.00117265646337 40.00000000000001))")
	wantGeometry2, _ := wkt.UnmarshalString("POLYGON((110.09906815774535 40.10054562342642,110.09922522259059 40.10067419589715,110.09941206169721 40.10077685932706,110.09962149494307 40.10084966854274,110.09984547392598 40.10088982562589,110.1000753912593 40.10089578742074,110.10030241134876 40.100867324827774,110.10051780993972 40.10080553160651,110.10071330938504 40.10071278234935,110.10088139675076 40.10059264124038,110.10101561253364 40.10044972510394,110.10111079889639 40.10028952600239,110.1011632978805 40.10011820019832,110.10117109197948 40.09994233158786,110.1011338816704 40.09976867869421,110.10105309692463 40.09960391494264,110.10093184225462 40.09945437219815,110.00093184225464 39.99945357112295,110.00077477740938 39.99932480757869,110.00058793830277 39.999221991229994,110.00037850505693 39.99914907337598,110.00015452607398 39.99910885630793,109.9999246087407 39.99910288560356,109.9996975886512 39.999131390722596,109.99948219006026 39.999193276187356,109.99928669061492 39.999286163687735,109.99911860324923 39.999406483491484,109.99898438746635 39.999549611644895,109.99888920110361 39.999710047688836,109.99883670211949 39.999881626057096,109.99882890802051 40.00005775303085,109.99886611832957 40.00023166014029,109.99894690307536 40.00039666427477,109.99906815774537 40.000546424504286,110.09906815774535 40.10054562342642))")
	geom1, _ := wkt.UnmarshalString("POLYGON((114.041146 22.553089,114.041232 22.552118,114.042761 22.552152,114.043013 22.55295,114.042042 22.553327,114.041146 22.553089))")
	want1, _ := wkt.UnmarshalString("POLYGON((114.04087713193498 22.55395230733391,114.04070018702038 22.55388740077672,114.04054032234225 22.55379190548216,114.04040366272567 22.553669480060506,114.04029544395028 22.55352481487839,114.04021981215449 22.553363452368803,114.04017966498598 22.55319157469546,114.04017654058555 22.553015766905304,114.04026254058554 22.552044766389916,114.04029474310306 22.551877723456286,114.04036038114127 22.551719251916776,114.04045711314154 22.55157500509319,114.04058148830525 22.551450128870737,114.04072906969695 22.55134907811753,114.04089459252671 22.55127545775416,114.04107215196521 22.55123189414466,114.04125541379148 22.551219941397644,114.04278441379147 22.551253941618953,114.04298707465064 22.551278281488642,114.04317975575995 22.55134119054798,114.0433539514808 22.551439891714242,114.04350197218628 22.551570027891145,114.04361728370894 22.551725854319198,114.04369479578263 22.551900492183982,114.04394679578263 22.55269849363868,114.04398291245069 22.552882075526004,114.04397718417364 22.553068587632264,114.0439298580866 22.553249983267712,114.04384297597957 22.55341843651319,114.04372028620841 22.55356667984813,114.04356708197939 22.55368831768277,114.04338997298458 22.553778102270698,114.04241897298458 22.554155100007772,114.04220844584977 22.554212055717667,114.04198945248345 22.55422399339518,114.04177313193497 22.554190305844568,114.04087713193498 22.55395230733391))")
	type args struct {
		geom     space.Geometry
		width    float64
		quadsegs int
	}
	tests := []struct {
		name string
		args args
		want space.Geometry
	}{
		{name: "BufferInMeter point", args: args{geom: space.Point{110, 40}, width: 100, quadsegs: 8}, want: wantGeometry},
		{name: "BufferInMeter linestring", args: args{geom: space.LineString{{110, 40}, {110.1, 40.1}}, width: 100, quadsegs: 8}, want: wantGeometry2},
		{name: "issue87", args: args{geom: geom1, width: 100, quadsegs: 8}, want: want1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &megrezAlgorithm{}
			gotGeometry := g.BufferInMeter(tt.args.geom, tt.args.width, tt.args.quadsegs)
			isEqual, _ := g.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("MegrezAlgorithm.BufferInMeter() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}
