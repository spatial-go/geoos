package geos

import (
	"fmt"
	"testing"

	"github.com/spatial-go/geoos/encoding/wkt"
	"github.com/spatial-go/geoos/planar"
	"github.com/spatial-go/geoos/space"
)

// Benchmark_MegrezArea test megrez area
func Benchmark_MegrezArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = planar.NormalStrategy().Area(geom)
	}
}

// Benchmark_GeosArea test geos area
func Benchmark_GeosArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = planar.GetStrategy(newGEOAlgorithm).Area(geom)
	}
}

// Benchmark_MegrezUnaryUnion test megrez UnaryUnion
func Benchmark_MegrezUnaryUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
		G := planar.NormalStrategy()
		_, _ = G.UnaryUnion(multiPolygon)

	}
}

// Benchmark_GeosUnaryUnion test geos UnaryUnion
func Benchmark_GeosUnaryUnion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
		G := planar.GetStrategy(newGEOAlgorithm)
		_, _ = G.UnaryUnion(multiPolygon)
	}
}

func Benchmark_Relate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		G := planar.NormalStrategy()
		for _, tt := range tests {
			_, _ = G.Relate(tt.args.g1, tt.args.g2)
		}
	}
}

func Benchmark_GEOSRelate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		G := GEOAlgorithm{}
		for _, tt := range tests {
			_, _ = G.Relate(tt.args.g1, tt.args.g2)
		}
	}
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

var tests []TestStruct

var geom space.Geometry = space.Polygon{{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}}}

func TestMain(m *testing.M) {
	initTestData()
	fmt.Println("test start")
	_ = m.Run()
	fmt.Println("test end")
}

func initTestData() {

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

}
