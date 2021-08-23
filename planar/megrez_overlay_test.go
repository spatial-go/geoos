package planar

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

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
		{name: "intersection error", args: args{g1: space.Collection{point02}, g2: space.Collection{line02}}, want: nil, wantErr: true},
		{name: "intersection error", args: args{g1: point02, g2: space.Collection{line02}}, want: nil, wantErr: true},
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

func TestAlgorithm_LineMerge(t *testing.T) {
	multiLineString0, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33),(-45 -33,-46 -32))`)
	expectLine0, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33,-46 -32))`)

	multiLineString1, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33),(-45.2 -33.2,-46 -32))`)
	expectMultiLineString, _ := wkt.UnmarshalString(`MULTILINESTRING((-29 -27,-30 -29.7,-36 -31,-45 -33),(-45.2 -33.2,-46 -32))`)

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
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
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want, 0.000001)
			if !isEqual {
				t.Errorf(" Error got = %v, want %v", wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want))
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

func TestAlgorithm_SharedPaths(t *testing.T) {
	const mullinestring = `MULTILINESTRING((26 125,26 200,126 200,126 125,26 125),
	   (51 150,101 150,76 175,51 150))`
	const linestring = `LINESTRING(151 100,126 156.25,126 125,90 161, 76 175)`
	const res = `GEOMETRYCOLLECTION(MULTILINESTRING((126 156.25,126 125),(101 150,90 161),(90 161,76 175)),MULTILINESTRING EMPTY)`

	mulline, _ := wkt.UnmarshalString(mullinestring)
	line, _ := wkt.UnmarshalString(linestring)

	type args struct {
		g1 space.Geometry
		g2 space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "sharepath", args: args{
			g1: line,
			g2: mulline,
		}, want: res, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.SharedPaths(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("SharedPaths() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SharedPaths() got = %v, want %v", got, tt.want)
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
	expectMultiline, _ := wkt.UnmarshalString(`MULTILINESTRING((50 100,50 150),(50 150,50 200),(50 50,50 100))`)

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
