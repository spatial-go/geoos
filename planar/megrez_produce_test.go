package planar

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry := G.Buffer(tt.args.g, tt.args.width, tt.args.quadsegs)
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf("GEOAlgorithm.Buffer() = %v\n, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_Centroid(t *testing.T) {
	const multipoint = `MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`
	geometry, _ := wkt.UnmarshalString(multipoint)
	const pointresult = `SRID=104326;POINT(2.3076923076923075 3.3076923076923075)`

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
	const multipoint = `SRID=104326;MULTIPOINT((0 0),(6 0),(6 6),(0 6))`

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

func TestMegrezAlgorithm_Buffer(t *testing.T) {
	type args struct {
		geom     space.Geometry
		width    float64
		quadsegs int
	}
	tests := []struct {
		name         string
		g            *megrezAlgorithm
		args         args
		wantGeometry space.Geometry
	}{
		{name: "point buffer", args: args{
			geom:     space.Point{100, 90},
			width:    50,
			quadsegs: 4,
		}, wantGeometry: space.Polygon{
			{{150, 90}, {146.193976625564, 70.8658283817455}, {135.355339059327, 54.6446609406727},
				{119.134171618255, 43.8060233744357}, {100, 40}, {80.8658283817456, 43.8060233744356}, {64.6446609406727, 54.6446609406725},
				{53.8060233744357, 70.8658283817454}, {50, 89.9999999999998}, {53.8060233744356, 109.134171618254}, {64.6446609406725, 125.355339059327},
				{80.8658283817453, 136.193976625564}, {99.9999999999998, 140}, {119.134171618254, 136.193976625564}, {135.355339059327, 125.355339059328},
				{146.193976625564, 109.134171618255}, {150, 90}},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &megrezAlgorithm{}
			if gotGeometry := g.Buffer(tt.args.geom, tt.args.width, tt.args.quadsegs); !gotGeometry.EqualsExact(tt.wantGeometry, 0.0000001) {
				t.Errorf("MegrezAlgorithm.Buffer() = %v, \nwant %v", gotGeometry, tt.wantGeometry)
			}
		})
	}
}

func TestMegrezAlgorithm_BufferInMeter(t *testing.T) {
	wantGeometry, _ := wkt.UnmarshalString("POLYGON((110.00117265646337 40.00000000000001,110.00115012419823 39.99982474877957,110.00108339330515 39.999656231941024,110.00097502821495 39.99950092556494,110.00082919333724 39.999364798097254,110.00065149302459 39.99925308096896,110.00044875620038 39.99917006753527,110.00022877392705 39.99911894806501,110.00000000000001 39.99910168712361,109.99977122607295 39.99911894806501,109.99955124379962 39.99917006753527,109.9993485069754 39.99925308096896,109.99917080666276 39.999364798097254,109.99902497178505 39.99950092556494,109.99891660669483 39.999656231941024,109.99884987580177 39.99982474877957,109.99882734353663 40.00000000000001,109.99884987580177 40.00017525077067,109.99891660669483 40.00034376632828,109.99902497178505 40.00049907078736,109.99917080666276 40.000635195993794,109.9993485069754 40.00074691086084,109.99955124379962 40.00082992237753,109.99977122607295 40.00088104056688,110.00000000000001 40.00089830105848,110.00022877392705 40.00088104056688,110.00044875620038 40.00082992237753,110.00065149302459 40.00074691086084,110.00082919333724 40.000635195993794,110.00097502821495 40.00049907078736,110.00108339330515 40.00034376632828,110.00115012419823 40.00017525077067,110.00117265646337 40.00000000000001))")
	wantGeometry2, _ := wkt.UnmarshalString("POLYGON((110.09906815774535 40.10054562342642,110.09922522259059 40.10067419589715,110.09941206169721 40.10077685932706,110.09962149494307 40.10084966854274,110.09984547392598 40.10088982562589,110.1000753912593 40.10089578742074,110.10030241134876 40.100867324827774,110.10051780993972 40.10080553160651,110.10071330938504 40.10071278234935,110.10088139675076 40.10059264124038,110.10101561253364 40.10044972510394,110.10111079889639 40.10028952600239,110.1011632978805 40.10011820019832,110.10117109197948 40.09994233158786,110.1011338816704 40.09976867869421,110.10105309692463 40.09960391494264,110.10093184225462 40.09945437219815,110.00093184225464 39.99945357112295,110.00077477740938 39.99932480757869,110.00058793830277 39.999221991229994,110.00037850505693 39.99914907337598,110.00015452607398 39.99910885630793,109.9999246087407 39.99910288560356,109.9996975886512 39.999131390722596,109.99948219006026 39.999193276187356,109.99928669061492 39.999286163687735,109.99911860324923 39.999406483491484,109.99898438746635 39.999549611644895,109.99888920110361 39.999710047688836,109.99883670211949 39.999881626057096,109.99882890802051 40.00005775303085,109.99886611832957 40.00023166014029,109.99894690307536 40.00039666427477,109.99906815774537 40.000546424504286,110.09906815774535 40.10054562342642)), want POLYGON((110.00117265646337 40.00000000000001,110.00115012419823 39.99982474877957,110.00108339330515 39.999656231941024,110.00097502821495 39.99950092556494,110.00082919333724 39.999364798097254,110.00065149302459 39.99925308096896,110.00044875620038 39.99917006753527,110.00022877392705 39.99911894806501,110.00000000000001 39.99910168712361,109.99977122607295 39.99911894806501,109.99955124379962 39.99917006753527,109.9993485069754 39.99925308096896,109.99917080666276 39.999364798097254,109.99902497178505 39.99950092556494,109.99891660669483 39.999656231941024,109.99884987580177 39.99982474877957,109.99882734353663 40.00000000000001,109.99884987580177 40.00017525077067,109.99891660669483 40.00034376632828,109.99902497178505 40.00049907078736,109.99917080666276 40.000635195993794,109.9993485069754 40.00074691086084,109.99955124379962 40.00082992237753,109.99977122607295 40.00088104056688,110.00000000000001 40.00089830105848,110.00022877392705 40.00088104056688,110.00044875620038 40.00082992237753,110.00065149302459 40.00074691086084,110.00082919333724 40.000635195993794,110.00097502821495 40.00049907078736,110.00108339330515 40.00034376632828,110.00115012419823 40.00017525077067,110.00117265646337 40.00000000000001))")
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
