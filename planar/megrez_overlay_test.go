package planar

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos"
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

	type args struct {
		g space.Geometry
	}
	tests := []struct {
		name    string
		args    args
		want    []space.Geometry
		wantErr bool
	}{
		{name: "UnaryUnion Polygon", args: args{g: space.MultiPolygon{space.Polygon{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			space.Polygon{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}}}}, want: []space.Geometry{space.Polygon{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}}}, wantErr: false},

		{name: "poly 1",
			args: args{space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}},
			},
			},
			want: []space.Geometry{space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
				space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
					space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
			},
			wantErr: false},

		{name: "poly 2",
			args: args{space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}},
			},
			},
			want: []space.Geometry{space.Polygon{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 2}, {1, 2}, {1, 1}, {2, 1}, {5, 1}, {5, 2}, {2, 2}}},
			},
			wantErr: false},

		{name: "poly 3",
			args: args{space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}},
			},
			},
			want: []space.Geometry{space.MultiPolygon{space.Polygon{{{2, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}}},
				space.Polygon{{{2, 2}, {2, 3}, {5, 3}, {5, 2}, {2, 2}}}},
				space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
					space.Polygon{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}}},
			},
			wantErr: false},

		{name: "poly 4",
			args: args{space.MultiPolygon{space.Polygon{{{1, 2}, {3, 2}, {3, 3}, {1, 3}, {1, 2}}},
				space.Polygon{{{2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 1}}},
			},
			},
			want: []space.Geometry{space.Polygon{{{2, 2}, {1, 2}, {1, 3}, {2, 3}, {2, 5}, {5, 5}, {5, 1}, {2, 1}, {2, 2}}},
				space.Polygon{{{2, 3}, {1, 3}, {1, 2}, {2, 2}, {2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 3}}},
			},
			wantErr: false},

		{name: "poly 5",
			args: args{space.MultiPolygon{space.Polygon{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
				space.Polygon{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}},
			},
			},
			want: []space.Geometry{space.Polygon{{{1, 1}, {1, 5}, {5, 5}, {5, 1}, {1, 1}}},
				space.Polygon{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			},
			wantErr: false},

		{name: "poly 6",
			args: args{space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
			},
			},
			want: []space.Geometry{space.Polygon{{{2, 1}, {1, 1}, {1, 2}, {1, 3}, {5, 3}, {5, 1}, {2, 1}}},
				space.Polygon{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "poly 4" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			gotGeometry, err := G.UnaryUnion(tt.args.g)

			if (err != nil) != tt.wantErr {
				t.Errorf("Algorithm UnaryUnion error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(gotGeometry, tt.want[0], 0.000001)
			if len(tt.want) > 1 {
				isEqual1, _ := G.EqualsExact(gotGeometry, tt.want[1], 0.000001)
				isEqual = isEqual || isEqual1
			}
			if !isEqual {
				t.Errorf("Algorithm UnaryUnion %v = %v, want %v", tt.name, wkt.MarshalString(gotGeometry), wkt.MarshalString(tt.want[0]))
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
		want    []space.Geometry
		wantErr bool
	}{
		{name: "union", args: args{g1: point01, g2: point02}, want: []space.Geometry{expectMultiPoint}},
		{name: "union line", args: args{g1: line01, g2: line02}, want: []space.Geometry{expectMultiline}},
		{name: "UnaryUnion Polygon", args: args{g1: space.Polygon{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
			g2: space.Polygon{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}}},
			want:    []space.Geometry{space.Polygon{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}}},
			wantErr: false},

		{name: "poly 1",
			args: args{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}},
			},
			want: []space.Geometry{space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
				space.MultiPolygon{space.Polygon{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
					space.Polygon{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
			},
			wantErr: false},
		{
			name: "union poly2",
			args: args{space.Polygon{{{110.85205078124999, 38.92522904714054}, {110.72021484375, 37.80544394934271}, {113.22509765625, 37.64903402157866},
				{113.818359375, 39.027718840211605}, {112.1484375, 39.57182223734374}, {110.85205078124999, 38.92522904714054}}},
				space.Polygon{{{113.99414062499999, 38.25543637637947}, {112.3681640625, 38.70265930723801}, {112.03857421875, 37.37015718405753},
					{114.01611328125, 36.29741818650811}, {114.43359375, 37.47485808497102}, {113.99414062499999, 38.25543637637947}}},
			},
			want: []space.Geometry{
				space.Polygon{{{112.124551191479, 37.7177543593912}, {110.72021484375, 37.8054439493427}, {110.85205078125, 38.9252290471405}, {112.1484375, 39.5718222373437},
					{113.818359375, 39.0277188402116}, {113.539811295113, 38.3803991213374}, {113.994140625, 38.2554363763795}, {114.43359375, 37.474858084971},
					{114.01611328125, 36.2974181865081}, {112.03857421875, 37.3701571840575}, {112.124551191479, 37.7177543593912}}},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			G := NormalStrategy()
			got, err := G.Union(tt.args.g1, tt.args.g2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Union() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			isEqual, _ := G.EqualsExact(got, tt.want[0], 0.000001)
			if len(tt.want) > 1 {
				isEqual1, _ := G.EqualsExact(got, tt.want[1], 0.000001)
				isEqual = isEqual || isEqual1
			}
			if !isEqual {
				t.Errorf("Algorithm Union %v = %v, want %v", tt.name, wkt.MarshalString(got), wkt.MarshalString(tt.want[0]))
			}
		})
	}
}
