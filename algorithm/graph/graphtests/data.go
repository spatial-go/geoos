// package graphtests ...

package graphtests

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// TestRelateData ...
var TestRelateData = []struct {
	Name    string
	Args    []matrix.Steric
	Want    string
	WantErr bool
}{
	{"point 1", []matrix.Steric{matrix.Matrix{3, 3}, matrix.Matrix{3, 3}}, "0FFFFFFF2", false},
	{"point 2", []matrix.Steric{matrix.Matrix{3, 3}, matrix.Matrix{3, 4}}, "FF0FFF0F2", false},

	{"PointLine 1", []matrix.Steric{matrix.Matrix{3, 3}, matrix.LineMatrix{{3.5, 2}, {3.5, 4}}}, "FF0FFF102", false},
	{"PointLine 2", []matrix.Steric{matrix.Matrix{3, 3}, matrix.LineMatrix{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}}, "FF0FFF1F2", false},
	{"PointLine 3", []matrix.Steric{matrix.Matrix{3, 3}, matrix.LineMatrix{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}}, "FF0FFF1F2", false},
	{"PointLine 4", []matrix.Steric{matrix.Matrix{3, 3}, matrix.LineMatrix{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}}, "FF0FFF1F2", false},
	{"PointLine 5", []matrix.Steric{matrix.Matrix{3, 3}, matrix.LineMatrix{{3, 3}, {3, 6}}}, "F0FFFF102", false},

	{"LineLine 1", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.LineMatrix{{3.5, 2}, {3.5, 4}}}, "FF1FF0102", false},
	{"LineLine 2", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.LineMatrix{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}}, "FF1FF01F2", false},
	{"LineLine 3", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.LineMatrix{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}}, "FF1FF01F2", false},
	{"LineLine 4", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.LineMatrix{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}}, "FF1FF01F2", false},
	{"LineLine 5", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.LineMatrix{{3, 3}, {3, 6}}}, "1FF00F102", false},

	{"PointPoly 1", []matrix.Steric{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}},
		{{2.5, 2.5}, {4.5, 2.5}, {4.5, 4.5}, {2.5, 4.5}, {2.5, 2.5}}},
	}, "FF0FFF212", false},
	{"PointPoly 2", []matrix.Steric{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}}}, "0FFFFF212", false},
	{"PointPoly 3", []matrix.Steric{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}}}, "FF0FFF212", false},
	{"PointPoly 4", []matrix.Steric{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}}}, "FF0FFF212", false},
	{"PointPoly 5", []matrix.Steric{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}}}, "F0FFFF212", false},

	{"LinePoly 1", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}},
		{{2.5, 2.5}, {4.5, 2.5}, {4.5, 4.5}, {2.5, 4.5}, {2.5, 2.5}}},
	}, "FF1FF0212", false},
	{"LinePoly 2", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}}}, "1FF0FF212", false},
	{"LinePoly 3", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.PolygonMatrix{{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}}}, "FF1FF0212", false},
	{"LinePoly 4", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.PolygonMatrix{{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}}}, "FF1FF0212", false},
	{"LinePoly 5", []matrix.Steric{matrix.LineMatrix{{3, 3}, {3, 4}}, matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}}}, "F1FF0F212", false},
	{"LinePoly 6", []matrix.Steric{matrix.LineMatrix{{5, 9}, {1, 9}, {1, 1}}, matrix.PolygonMatrix{{{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}}}, "F1FF0F212", false},
	{"PolyPoly 1", []matrix.Steric{matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
		matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}},
			{{2.5, 2.5}, {4.5, 2.5}, {4.5, 4.5}, {2.5, 4.5}, {2.5, 2.5}}},
	}, "FF2FF1212", false},
	{"PolyPoly 2", []matrix.Steric{matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
		matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 5}, {2, 5}, {2, 2}}}}, "2FF1FF212", false},
	{"PolyPoly 3", []matrix.Steric{matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
		matrix.PolygonMatrix{{{3.5, 3.5}, {3.5, 4.5}, {4.5, 4.5}, {4.5, 3.5}, {3.5, 3.5}}}}, "212101212", false},
	{"PolyPoly 4", []matrix.Steric{matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
		matrix.PolygonMatrix{{{5, 5}, {5, 6}, {6, 6}, {6, 5}, {5, 5}}}}, "FF2FF1212", false},
	{"PolyPoly 5", []matrix.Steric{matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}},
		matrix.PolygonMatrix{{{3, 3}, {3, 4}, {4, 4}, {4, 3}, {3, 3}}}}, "2FFF1FFF2", false},

	{"Disjoint 00",
		[]matrix.Steric{matrix.Matrix{0, 0}, matrix.LineMatrix{{2, 0}, {0, 2}}}, "FF0FFF102", false},
	{"inter 3-6",
		[]matrix.Steric{matrix.Matrix{3, 3}, matrix.PolygonMatrix{{{0, 0}, {6, 0}, {6, 6}, {0, 6}, {0, 0}}}}, "0FFFFF212", false},
	{"linepoint 00",
		[]matrix.Steric{matrix.Matrix{1, 1}, matrix.LineMatrix{{0, 0}, {1, 1}, {0, 2}}}, "0FFFFF102", false},
	{"linepoint 01",
		[]matrix.Steric{matrix.Matrix{0, 2}, matrix.LineMatrix{{0, 0}, {1, 1}, {0, 2}}}, "F0FFFF102", false},
	{"polyPoly 0f",
		[]matrix.Steric{matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, "2FF11F212", false},
	{"polyPoly 00f",
		[]matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}}}, "2FF11F212", false},

	{"polyPoly 00f1",
		[]matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}}, "212FF1FF2", false},
	{"polyPoly 00f2",
		[]matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, "FF2F11212", false},

	{"polyPoly 00f3",
		[]matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}}}, "FF2F01212", false},
	{"polyPoly f1",
		[]matrix.Steric{matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
			matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}}, "212F11FF2", false},
	{"lineline f1c",
		[]matrix.Steric{matrix.LineMatrix{{0, 0}, {10, 10}},
			matrix.LineMatrix{{0, 10}, {10, 0}}}, "0F1FF0102", false},
	{"polyPoly _f2",
		[]matrix.Steric{matrix.PolygonMatrix{{{110.85205078124999, 38.92522904714054}, {110.72021484375, 37.80544394934271}, {113.22509765625, 37.64903402157866},
			{113.818359375, 39.027718840211605}, {112.1484375, 39.57182223734374}, {110.85205078124999, 38.92522904714054}}},
			matrix.PolygonMatrix{{{113.99414062499999, 38.25543637637947}, {112.3681640625, 38.70265930723801}, {112.03857421875, 37.37015718405753},
				{114.01611328125, 36.29741818650811}, {114.43359375, 37.47485808497102}, {113.99414062499999, 38.25543637637947}}},
		}, "212101212", false},
	{"polyPoly _f2",
		[]matrix.Steric{matrix.PolygonMatrix{{{277.0764427576214075, 220.0703895697370172}, {231.7694100104272366, 422.7364493119530380}, {361.1115263458341360, 451.6569778039120138},
			{406.4185590930283070, 248.9906205087900162}, {277.0764427576214075, 220.0703895697370172}}},
			matrix.PolygonMatrix{{{157.0188155006617308, 268.8332448475994170}, {157.0188155006617308, 403.2498672713991255}, {314.0376310013234615, 403.2498672713991255},
				{314.0376310013234615, 268.8332448475994170}, {157.0188155006617308, 268.8332448475994170}}},
		}, "212101212", false},
	{"linePoly _tjhb",
		[]matrix.Steric{matrix.LineMatrix{{117.61332973211957, 39.35329370155506}, {117.61560683206969, 39.34946055539684}},
			matrix.PolygonMatrix{{{117.60388060798783, 39.35142233161234}, {117.61332973211957, 39.35329370155506}, {117.6156068320697, 39.349460555396846}, {117.62308032712902, 39.350634137974595},
				{117.62308032712902, 39.340634137974595}, {117.60388060798783, 39.340634137974595}, {117.60388060798783, 39.35142233161234}}},
		}, "F1FF0F212", false},
	{"test line poly0", []matrix.Steric{matrix.LineMatrix{{1, 1.5}, {2, 1.5}}, matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}}}, "1FFF0F212", false},
	{"test line poly1", []matrix.Steric{matrix.LineMatrix{{5, 10}, {5, 15}, {15, 15}, {15, 5}, {10, 5}}, matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}}}, "FF1F0F212", false},
	{"test line poly2", []matrix.Steric{matrix.LineMatrix{{5, 10}, {5, 9}}, matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}},
		{{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}}}, "1FFF0F212", false},
	{"test line poly3", []matrix.Steric{matrix.LineMatrix{{0, 0}, {10, 0}, {10, 5}}, matrix.PolygonMatrix{{{0, 0}, {0, 10}, {10, 10}, {10, 0}, {0, 0}}}}, "F1FF0F212", false},
	{"test poly poly0", []matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}}, matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}}}, "212F11FF2", false},

	{"test line poly4", []matrix.Steric{
		matrix.LineMatrix{{111.96859688870948, 38.44363360825604},
			{111.7096710205078, 38.39441521865825},
			{111.8671003424276, 38.22285924285058},
		},
		matrix.PolygonMatrix{{{112.01248168945312, 38.4519755295767},
			{111.7096710205078, 38.39441521865825},
			{111.90948486328125, 38.176671418717746},
			{112.01248168945312, 38.4519755295767}}}},
		"F1FF0F212",
		false,
	},
}
