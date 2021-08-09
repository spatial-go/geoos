package planar

import (
	"fmt"
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
)

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
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "0f"),
		args{space.Polygon{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			space.Polygon{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, "2FF1FF212", false})

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

func TestAlgorithm_Overlaps(t *testing.T) {
	type args struct {
		A space.Geometry
		B space.Geometry
	}
	tests := []struct {
		name    string
		g       *MegrezAlgorithm
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
			g := &MegrezAlgorithm{}
			got, err := g.Overlaps(tt.args.A, tt.args.B)
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
