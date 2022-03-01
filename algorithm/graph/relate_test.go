// package graph ...

package graph

import (
	"fmt"
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/matrix"
	"github.com/spatial-go/geoos/algorithm/matrix/envelope"
)

func TestRelate(t *testing.T) {
	type args struct {
		g0 matrix.Steric
		g1 matrix.Steric
	}

	type TestStruct struct {
		name    string
		args    args
		want    string
		wantErr bool
	}
	tests := []TestStruct{}
	g0 := matrix.Matrix{3, 3}

	g1 := matrix.LineMatrix{{3, 3}, {3, 4}}

	g2 := matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}}
	polys := []matrix.PolygonMatrix{
		{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}},
			{{2.5, 2.5}, {4.5, 2.5}, {4.5, 4.5}, {2.5, 4.5}, {2.5, 2.5}}},
		{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}},
		{{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}},
		{{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}},
		{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
	}
	pts := []matrix.Matrix{{3, 3}, {4, 4}}
	wants1 := []string{"0FFFFFFF2", "FF0FFF0F2"}
	for i, v := range pts {
		tests = append(tests, TestStruct{fmt.Sprintf("Point%v", i), args{g0, v}, wants1[i], false})
	}
	ls := []matrix.LineMatrix{{{3.5, 2}, {3.5, 4}},
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
		args{matrix.Matrix{0, 0}, matrix.LineMatrix{{2, 0}, {0, 2}}}, "FF0FFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("inter%v", "3-6"),
		args{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{0, 0}, {6, 0}, {6, 6}, {0, 6}, {0, 0}}}}, "0FFFFF212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("linepoint%v", "00"),
		args{matrix.Matrix{1, 1}, matrix.LineMatrix{{0, 0}, {1, 1}, {0, 2}}}, "0FFFFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("linepoint%v", "01"),
		args{matrix.Matrix{0, 2}, matrix.LineMatrix{{0, 0}, {1, 1}, {0, 2}}}, "F0FFFF102", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "0f"),
		args{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, "2FF11F212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f"),
		args{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}}}, "2FF11F212", false})

	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f1"),
		args{matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}}, "212FF1FF2", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f2"),
		args{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, "FF2F11212", false})

	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "00f3"),
		args{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}}}, "FF2F01212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "f1"),
		args{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}}, "212F11FF2", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "_f2"),
		args{matrix.PolygonMatrix{{{110.85205078124999, 38.92522904714054}, {110.72021484375, 37.80544394934271}, {113.22509765625, 37.64903402157866},
			{113.818359375, 39.027718840211605}, {112.1484375, 39.57182223734374}, {110.85205078124999, 38.92522904714054}}},
			matrix.PolygonMatrix{{{113.99414062499999, 38.25543637637947}, {112.3681640625, 38.70265930723801}, {112.03857421875, 37.37015718405753},
				{114.01611328125, 36.29741818650811}, {114.43359375, 37.47485808497102}, {113.99414062499999, 38.25543637637947}}},
		}, "212101212", false})
	tests = append(tests, TestStruct{fmt.Sprintf("polyPoly%v", "_f2"),
		args{matrix.PolygonMatrix{{{277.0764427576214075, 220.0703895697370172}, {231.7694100104272366, 422.7364493119530380}, {361.1115263458341360, 451.6569778039120138},
			{406.4185590930283070, 248.9906205087900162}, {277.0764427576214075, 220.0703895697370172}}},
			matrix.PolygonMatrix{{{157.0188155006617308, 268.8332448475994170}, {157.0188155006617308, 403.2498672713991255}, {314.0376310013234615, 403.2498672713991255},
				{314.0376310013234615, 268.8332448475994170}, {157.0188155006617308, 268.8332448475994170}}},
		}, "212101212", false})

	for _, tt := range tests {
		if !geoos.GeoosTestTag && tt.name != "polyPoly0f" {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			intersectBound := false
			env1 := envelope.Bound(tt.args.g0.Bound())
			env2 := envelope.Bound(tt.args.g1.Bound())
			if env1.IsIntersects(env2) {
				intersectBound = true
			}
			if env1.Contains(env2) || env2.Contains(env1) {
				intersectBound = true
			}
			if got := Relate(tt.args.g0, tt.args.g1); got != tt.want {
				t.Errorf("Relate()%v = %v, want %v  %v", tt.name, got, tt.want, intersectBound)
			}
		})
	}
}
