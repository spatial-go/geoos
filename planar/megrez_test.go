package planar

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

func TestAlgorithm_Area(t *testing.T) {
	const polygon = `POLYGON((-1 -1, 1 -1, 1 1, -1 1, -1 -1))`
	geometry, _ := wkt.UnmarshalString(polygon)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "area", args: args{g: geometry}, want: 4.0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalStrategy().Area(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Length(t *testing.T) {
	const line = `LINESTRING(0 0, 1 1)`
	geometry, _ := wkt.UnmarshalString(line)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "lengh", args: args{g: geometry}, want: 1.4142135623730951, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Length(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Length() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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

	smultiPolygon, _ := wkt.UnmarshalString(sourceLine)
	emultiPolygon, _ := wkt.UnmarshalString(expectLine)

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
		{name: "multiPolygon", args: args{g: smultiPolygon}, want: emultiPolygon, wantErr: false},
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

func TestAlgorithm_UnaryUnion(t *testing.T) {
	multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
	expectPolygon, _ := wkt.UnmarshalString(`POLYGON((5 10,0 10,0 0,10 0,10 5,15 5,15 15,5 15,5 10))`)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		G       GEOAlgorithm
		args    args
		want    []space.Geometry
		wantErr bool
	}{
		{name: "UnaryUnion Polygon", args: args{g: multiPolygon}, want: []space.Geometry{expectPolygon}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.UnaryUnion(tt.args.g)

			if (err != nil) != tt.wantErr {
				t.Errorf("Algorithm UnaryUnion error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want[0], 0.000001)
			if !isEqual {
				t.Errorf("Algorithm UnaryUnion = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want[0]))
			}
		})
	}
}

func TestAlgorithm_Union(t *testing.T) {
	point01, _ := wkt.UnmarshalString(`POINT(1 2)`)
	point02, _ := wkt.UnmarshalString(`POINT(-2 3)`)
	expectMultiPoint, _ := wkt.UnmarshalString(`MULTIPOINT(1 2,-2 3)`)

	line01, _ := wkt.UnmarshalString(`LINESTRING(50 100, 50 200)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING(50 50, 50 150)`)
	expectMultiline, _ := wkt.UnmarshalString(`MULTILINESTRING((50 150,50 200),(50 50,50 100))`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "union", args: args{g1: point01, g2: point02}, want: expectMultiPoint},
		{name: "union line", args: args{g1: line01, g2: line02}, want: expectMultiline},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Union(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Distance(t *testing.T) {
	point01, _ := wkt.UnmarshalString(`POINT(1 3)`)
	point02, _ := wkt.UnmarshalString(`POINT(4 7)`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "distance", args: args{g1: point01, g2: point02}, want: 5, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Distance(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Distance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_SphericalDistance(t *testing.T) {
	point01 := space.Point{116.397439, 39.909177}
	point02 := space.Point{116.397725, 39.903079}
	point03 := space.Point{118.1487, 39.586671}

	type args struct {
		p1 space.Point
		p2 space.Point
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "SphericalDistance", args: args{p1: point01, p2: point02}, want: 678.5053586786567, wantErr: false},
		{name: "SphericalDistance", args: args{p1: point01, p2: point03}, want: 153953.98145619757, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, _ := G.SphericalDistance(tt.args.p1, tt.args.p2)
			if got != tt.want {
				t.Errorf("SphericalDistance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_IsClosed(t *testing.T) {
	const linestring = `LINESTRING(1 1,2 3,3 2,1 2)`
	const closedLinestring = `LINESTRING(1 1,2 3,3 2,1 2,1 1)`
	line, _ := wkt.UnmarshalString(linestring)
	closedLine, _ := wkt.UnmarshalString(closedLinestring)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "line", args: args{g: line}, want: false, wantErr: false},
		{name: "closedLine", args: args{g: closedLine}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.IsClosed(tt.args.g)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("geometry: %s IsClose() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("geometry: %s IsClose() got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestAlgorithm_EqualsExact(t *testing.T) {
	geometry1, _ := wkt.UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	geometry2, _ := wkt.UnmarshalString("POINT(116.309878725564 40.0427783827455)")
	geometry3, _ := wkt.UnmarshalString("POINT(116.309877625564 40.0427783827455)")
	type args struct {
		g1        space.Geometry
		g2        space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		G       GEOAlgorithm
		args    args
		want    bool
		wantErr bool
	}{
		{name: "equals exact", args: args{g1: geometry1, g2: geometry2, tolerance: 0.000001}, want: true, wantErr: false},
		{name: "not equals exact", args: args{g1: geometry1, g2: geometry3, tolerance: 0.000001}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.EqualsExact(tt.args.g1, tt.args.g2, tt.args.tolerance)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.EqualsExact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOAlgorithm.EqualsExact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_Equals(t *testing.T) {
	geometry1, _ := wkt.UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	geometry2, _ := wkt.UnmarshalString("POINT(116.309878725564 40.0427783827455)")
	geometry3, _ := wkt.UnmarshalString("POINT(116.309878625564 40.0427783817455)")
	type args struct {
		g1        space.Geometry
		g2        space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		G       GEOAlgorithm
		args    args
		want    bool
		wantErr bool
	}{
		{name: "equals exact", args: args{g1: geometry1, g2: geometry3}, want: true, wantErr: false},
		{name: "not equals exact", args: args{g1: geometry1, g2: geometry2}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Equals(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.Equals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOAlgorithm.Equals() = %v, want %v", got, tt.want)
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

func TestAlgorithm_IsSimple(t *testing.T) {
	const polygon = `POLYGON((1 2, 3 4, 5 6, 5 3, 1 2))`
	const linestring = `LINESTRING(1 1,2 2,2 3.5,1 3,1 2,2 1)`
	poly, _ := wkt.UnmarshalString(polygon)
	line, _ := wkt.UnmarshalString(linestring)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "polygon", args: args{g: poly}, want: true, wantErr: false},
		{name: "line", args: args{g: line}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.IsSimple(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsSimple() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsSimple() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_IsRing(t *testing.T) {
	const linestring1 = `LINESTRING((1 2, 3 4, 5 6, 5 3, 1 2))`
	const linestring2 = `LINESTRING(1 1,2 2,2 3.5,1 3,1 2,2 1)`
	line1, _ := wkt.UnmarshalString(linestring1)
	line2, _ := wkt.UnmarshalString(linestring2)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{name: "polygon", args: args{g: line1}, want: true, wantErr: false},
		{name: "line", args: args{g: line2}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.IsSimple(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsRing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsRing() got = %v, want %v", got, tt.want)
			}
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

func TestAlgorithm_Buffer(t *testing.T) {
	geometry, _ := wkt.UnmarshalString("POINT(100 90)")
	expectGeometry, _ := wkt.UnmarshalString("POLYGON((150 90,146.193976625564 70.8658283817455,135.355339059327 54.6446609406727,119.134171618255 43.8060233744357,100 40,80.8658283817456 43.8060233744356,64.6446609406727 54.6446609406725,53.8060233744357 70.8658283817454,50 89.9999999999998,53.8060233744356 109.134171618254,64.6446609406725 125.355339059327,80.8658283817453 136.193976625564,99.9999999999998 140,119.134171618254 136.193976625564,135.355339059327 125.355339059328,146.193976625564 109.134171618255,150 90))")
	type args struct {
		g        space.Geometry
		width    float64
		quadsegs int32
	}
	tests := []struct {
		name string
		G    GEOAlgorithm
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
				t.Errorf("GEOAlgorithm.Buffer() = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
			}
		})
	}
}

func TestAlgorithm_NGeometry(t *testing.T) {
	multiPoint, _ := wkt.UnmarshalString(`MULTIPOINT ( -1 0, -1 2, -1 3, -1 4, -1 7, 0 1, 0 3, 1 1, 2 0, 6 0, 7 8, 9 8, 10 6 )`)
	multiLineString, _ := wkt.UnmarshalString(`MULTILINESTRING((10 130,50 190,110 190,140 150,150 80,100 10,20 40,10 130),
	(70 40,100 50,120 80,80 110,50 90,70 40))`)
	multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON (((40 40, 20 45, 45 30, 40 40)),
	((20 35, 10 30, 10 10, 30 5, 45 20, 20 35)),
	((30 20, 20 15, 20 25, 30 20)))`)
	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		G       GEOAlgorithm
		args    args
		want    int
		wantErr bool
	}{
		{name: "ngeometry multiPoint", args: args{g: multiPoint}, want: 13, wantErr: false},
		{name: "ngeometry multiLineString", args: args{g: multiLineString}, want: 2, wantErr: false},
		{name: "ngeometry multiPolygon", args: args{g: multiPolygon}, want: 3, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.NGeometry(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GEOAlgorithm.NGeometry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GEOAlgorithm.NGeometry() = %v, want %v", got, tt.want)
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
		G       GEOAlgorithm
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "PointOnSurface Point", args: args{g: point}, want: expectPoint0, wantErr: false},
		{name: "PointOnSurface LineString0", args: args{g: lineString}, want: expectPoint1, wantErr: false},
		{name: "PointOnSurface Polygon", args: args{g: polygon}, want: expectPoint2, wantErr: false},
	}
	for i, tt := range tests {
		if i == 0 || i == 2 {
			continue
		}
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

func TestAlgorithm_Relate(t *testing.T) {
	g0 := space.Point{3, 3}

	g1 := space.LineString{{3, 3}, {3, 4}}

	g2 := space.Polygon{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}}
	polys := []space.Polygon{
		{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}},
			{{2.5, 2.5}, {4.5, 2.5}, {4.5, 4.5}, {2.5, 4.5}, {2.5, 2.5}}},
		{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}},
		{{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}},
		{{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}},
		{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
	}
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
	pts := []space.Point{{3, 3}, {4, 4}}
	wants1 := []string{"0FFFFFFF2", "FF0FFF0F2"}
	for i, v := range pts {
		tests = append(tests, TestStruct{fmt.Sprintf("Point%v", i), args{g0, v}, wants1[i], false})
	}
	ls := []space.LineString{{{3.5, 2}, {3.5, 4}},
		{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}},
		{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}},
		{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}},
		{{3, 3}, {3, 6}}}
	wants2 := []string{"FF0FFF102", "FF1FF0102", "FF0FFF1F2", "FF1FF01F2", "FF0FFF1F2", "FF1FF01F2", "FF0FFF1F2", "FF1FF01F2", "F0FFFF102", "1FF00F102"}
	for i, v := range ls {
		tests = append(tests, TestStruct{fmt.Sprintf("PointLine%v", i), args{g0, v}, wants2[i*2], false})
		tests = append(tests, TestStruct{fmt.Sprintf("LineLine%v", i), args{g1, v}, wants2[i*2+1], false})
	}

	wants3 := []string{"FF0FFF212", "FF1FF0212", "FF2FF1212",
		"0FFFFF212", "1FF0FF212", "2FF1FF212",
		"FF0FFF212", "FF1FF0212", "212101212",
		"FF0FFF212", "FF1FF0212", "FF2FF1212",
		"F0FFFF212", "F1FF0F212", "2FFF1FFF2",
	}
	for i, v := range polys {
		tests = append(tests, TestStruct{fmt.Sprintf("PointPoly%v", i), args{g0, v}, wants3[i*3], false})
		tests = append(tests, TestStruct{fmt.Sprintf("LinePoly%v", i), args{g1, v}, wants3[i*3+1], false})
		tests = append(tests, TestStruct{fmt.Sprintf("PolyPoly%v", i), args{g2, v}, wants3[i*3+2], false})
	}

	tests = append(tests, TestStruct{fmt.Sprintf("Disjoint%v", "00"),
		args{space.Point{0, 0}, space.LineString{{2, 0}, {0, 2}}}, "FF0FFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("inter%v", "3-6"),
		args{space.Point{3, 3}, space.Polygon{{{0, 0}, {6, 0}, {6, 6}, {0, 6}, {0, 0}}}}, "0FFFFF212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("linepoint%v", "00"),
		args{space.Point{1, 1}, space.LineString{{0, 0}, {1, 1}, {0, 2}}}, "0FFFFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("linepoint%v", "01"),
		args{space.Point{0, 2}, space.LineString{{0, 0}, {1, 1}, {0, 2}}}, "F0FFFF102", false})

	for _, tt := range tests {
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
		{name: "notin", args: args{
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

func TestAlgorithm_Contains(t *testing.T) {
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
		{name: "contain", args: args{
			g1: poly,
			g2: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly,
			g2: p2,
		}, want: false, wantErr: false},
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
		{name: "contain", args: args{
			g1: poly,
			g2: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly,
			g2: p2,
		}, want: false, wantErr: false},
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
				t.Errorf("Covers() got = %v, want %v", got, tt.want)
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
		{name: "contain", args: args{
			g1: poly,
			g2: p1,
		}, want: true, wantErr: false},
		{name: "notcontain", args: args{
			g1: poly,
			g2: p2,
		}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.CoveredBy(tt.args.g2, tt.args.g1)
			if (err != nil) != tt.wantErr {
				t.Errorf("CoveredBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CoveredBy() got = %v, want %v", got, tt.want)
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

func TestAlgorithm_Intersection(t *testing.T) {
	point02, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING ( 0 0, 0 2 )`)
	expectPoint, _ := wkt.UnmarshalString(`POINT(0 0)`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "intersection", args: args{g1: point02, g2: line02}, want: expectPoint, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Intersection(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Intersection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAlgorithm_Intersects(t *testing.T) {
	point01, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line01, _ := wkt.UnmarshalString(`LINESTRING ( 0 0, 0 2 )`)

	point02, _ := wkt.UnmarshalString(`POINT(0 0)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING ( 2 1, 1 2 )`)

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

func TestAlgorithm_LineMerge(t *testing.T) {
	multiLineString0, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33),(-45 -33,-46 -32))`)
	expectLine0, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33,-46 -32))`)

	multiLineString1, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33),(-45.2 -33.2,-46 -32))`)
	expectMultiLineString, _ := wkt.UnmarshalString(`MULTILINESTRING((-45.2 -33.2,-46 -32),(-29 -27,-30 -29.7,-36 -31,-45 -33))`)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		G       GEOAlgorithm
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "LineMerge Point", args: args{g: multiLineString0}, want: expectLine0, wantErr: false},
		{name: "LineMerge LineString0", args: args{g: multiLineString1}, want: expectMultiLineString, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.LineMerge(tt.args.g)

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
	const input = `POINT(0.05 0.05)`
	const refernce = `POINT(0 0)`
	const expect = `POINT(0 0)`

	inputGeom, _ := wkt.UnmarshalString(input)
	referenceGeom, _ := wkt.UnmarshalString(refernce)

	type args struct {
		input     space.Geometry
		reference space.Geometry
		tolerance float64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "snap", args: args{
			input:     inputGeom,
			reference: referenceGeom,
			tolerance: 0.1,
		}, want: expect, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Snap(tt.args.input, tt.args.reference, tt.args.tolerance)

			s := wkt.MarshalString(got)
			t.Log(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Snap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(s, tt.want) {
				t.Errorf("Snap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_HausdorffDistance(t *testing.T) {
	g3 := "LINESTRING (130 0, 0 0, 0 150)"
	g4 := "LINESTRING (10 10, 10 150, 130 10)"

	const g1 = `LINESTRING (0 0, 2 0)`
	const g2 = `MULTIPOINT (0 1, 1 0, 2 1)`
	geom1, _ := wkt.UnmarshalString(g1)
	geom2, _ := wkt.UnmarshalString(g2)
	geom3, _ := wkt.UnmarshalString(g3)
	geom4, _ := wkt.UnmarshalString(g4)
	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "HausdorffDistance", args: args{
			g1: geom1,
			g2: geom2,
		}, want: 1, wantErr: false},
		{name: "HausdorffDistance", args: args{
			g1: geom3,
			g2: geom4,
		}, want: 14.142135623730951, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.HausdorffDistance(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("HausdorffDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HausdorffDistance() got = %v, want %v", got, tt.want)
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
		G       GEOAlgorithm
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
		G       GEOAlgorithm
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

func TestAlgorithm_Difference(t *testing.T) {
	line01, _ := wkt.UnmarshalString(`LINESTRING(50 100, 50 200)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING(50 50, 50 150)`)
	expectLine, _ := wkt.UnmarshalString(`LINESTRING(50 150,50 200)`)
	expectLine2, _ := wkt.UnmarshalString(`LINESTRING(50 50,50 100)`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "difference", args: args{g1: line01, g2: line02}, want: expectLine},
		{name: "difference2", args: args{g2: line01, g1: line02}, want: expectLine2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Difference(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Difference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgorithm_SymDifference(t *testing.T) {
	line01, _ := wkt.UnmarshalString(`LINESTRING(50 100, 50 200)`)
	line02, _ := wkt.UnmarshalString(`LINESTRING(50 50, 50 150)`)
	expectMultiLines, _ := wkt.UnmarshalString(`MULTILINESTRING((50 150,50 200),(50 50,50 100))`)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    space.Geometry
		wantErr bool
	}{
		{name: "symDifference", args: args{g1: line01, g2: line02}, want: expectMultiLines},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.SymDifference(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SymDifference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("SymDifference() got = %v, want %v", got, tt.want)
			}
		})
	}
}
