package planar

import (
	"fmt"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/space"
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
			g := &megrezAlgorithm{}
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
