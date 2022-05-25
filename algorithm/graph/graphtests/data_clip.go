// package graphtests  is include test datas.
package graphtests

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// TestCase ...
type TestCase struct {
	Name    string
	Fields  []matrix.Steric
	Want    matrix.Steric
	WantErr bool
}

// TestsPointIntersecation ...
var TestsPointIntersecation = []TestCase{
	{"point point0", []matrix.Steric{matrix.Matrix{100, 100}, matrix.Matrix{100, 100}},
		matrix.Matrix{100, 100}, false},
	{"point point1", []matrix.Steric{matrix.Matrix{100, 100}, matrix.Matrix{100, 101}},
		nil, false},
	{"point line0", []matrix.Steric{matrix.Matrix{100, 100},
		matrix.LineMatrix{{100, 100}, {100, 101}}},
		matrix.Matrix{100, 100}, false},
	{"point line1", []matrix.Steric{matrix.Matrix{100, 100},
		matrix.LineMatrix{{100, 105}, {100, 101}}},
		nil, false},
	{"point poly1", []matrix.Steric{matrix.Matrix{100, 100},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}},
		matrix.Matrix{100, 100}, false},
	{"point poly2", []matrix.Steric{matrix.Matrix{100, 100},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}},
		matrix.Matrix{100, 100}, false},
	{"point poly3", []matrix.Steric{matrix.Matrix{100, 100},
		matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}}},
		nil, false},
}

// TestsLineUnion ...
var TestsLineUnion = []TestCase{
	{"line point0", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.Matrix{100, 100}},
		matrix.LineMatrix{{100, 100}, {100, 101}}, false},
	{"line line0", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.LineMatrix{{100, 100}, {100, 101}}},
		matrix.LineMatrix{{100, 100}, {100, 101}}, false},
	{"line line1", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.LineMatrix{{100, 100}, {90, 102}}},
		matrix.Collection{matrix.LineMatrix{{100, 100}, {100, 101}}, matrix.LineMatrix{{100, 100}, {90, 102}}}, false},
	{"line poly1", []matrix.Steric{matrix.LineMatrix{{100, 100}, {101, 101}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}, false},
	{"line poly2", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}, false},
	{"line poly3", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}}},
		matrix.Collection{matrix.LineMatrix{{100, 100}, {100, 101}},
			matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}}}, false},
}

// TestsLineIntersecation ...
var TestsLineIntersecation = []TestCase{
	{"line point0", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.Matrix{100, 100}},
		matrix.Matrix{100, 100}, false},
	{"line line0", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.LineMatrix{{100, 100}, {100, 101}}},
		matrix.LineMatrix{{100, 100}, {100, 101}}, false},
	{"line line1", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.LineMatrix{{100, 100}, {90, 102}}},
		matrix.Matrix{100, 100}, false},
	{"line poly1", []matrix.Steric{matrix.LineMatrix{{100, 100}, {101, 101}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}},
		matrix.LineMatrix{{100, 100}, {101, 101}}, false},
	{"line poly2", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}}},
		matrix.LineMatrix{{100, 100}, {100, 101}}, false},
	{"line poly3", []matrix.Steric{matrix.LineMatrix{{100, 100}, {100, 101}},
		matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}}},
		matrix.Collection{}, false},

	{"line poly5", []matrix.Steric{
		matrix.LineMatrix{{111.98638916015625, 38.50357937743225},
			{111.96372985839844, 38.42723559654225},
			{111.91085815429688, 38.344887442462806},
			{111.87309265136717, 38.24680876017446},
			{111.84906005859375, 38.15075747130226}},
		matrix.PolygonMatrix{{{112.01248168945312, 38.4519755295767},
			{111.7096710205078, 38.39441521865825},
			{111.90948486328125, 38.176671418717746},
			{112.01248168945312, 38.4519755295767}}}},
		matrix.LineMatrix{{111.96859688870948, 38.44363360825604},
			{111.96372985839844, 38.42723559654225},
			{111.91085815429688, 38.344887442462806},
			{111.87309265136717, 38.24680876017446},
			{111.8671003424276, 38.22285924285058},
		},
		false,
	},
}

// TestsLineDifference ...
var TestsLineDifference = []TestCase{
	{"line line0", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{50, 50}, {50, 150}}},
		matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}, false},
	{"line line1", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{50, 120}, {50, 150}}},
		matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 120}},
			matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}}, false},
	{"line line2", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{50, 150}, {50, 250}}},
		matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}},
			matrix.LineMatrix{{50, 200}, {60, 200}}}, false},
	{"line line3", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{50, 100}, {50, 150}}},
		matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}, false},
	{"line line4", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{50, 150}, {50, 200}}},
		matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}},
			matrix.LineMatrix{{50, 200}, {60, 200}}}, false},
	{"line line5", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}},
		matrix.LineMatrix{{50, 50}, {50, 250}}},
		matrix.Collection{}, false},
	{"line line6", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{50, 50}, {50, 250}}},
		matrix.LineMatrix{{50, 200}, {60, 200}}, false},

	{"line line7", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}},
		matrix.LineMatrix{{30, 30}, {30, 150}}},
		matrix.LineMatrix{{50, 100}, {50, 200}}, false},
	{"line line8", []matrix.Steric{matrix.LineMatrix{{50, 100}, {50, 200}, {60, 200}},
		matrix.LineMatrix{{30, 150}, {60, 150}}},
		matrix.Collection{matrix.LineMatrix{{50, 100}, {50, 150}},
			matrix.LineMatrix{{50, 150}, {50, 200}, {60, 200}}}, false},

	{"line poly1", []matrix.Steric{
		matrix.LineMatrix{{111.98638916015625, 38.50357937743225},
			{111.96372985839844, 38.42723559654225},
			{111.91085815429688, 38.344887442462806},
			{111.87309265136717, 38.24680876017446},
			{111.84906005859375, 38.15075747130226}},
		matrix.PolygonMatrix{{{112.01248168945312, 38.4519755295767},
			{111.7096710205078, 38.39441521865825},
			{111.90948486328125, 38.176671418717746},
			{112.01248168945312, 38.4519755295767}}},
	},
		matrix.Collection{
			matrix.LineMatrix{{111.98638916015625, 38.50357937743225},
				{111.96859688870948, 38.44363360825604}},
			matrix.LineMatrix{{111.96859688870948, 38.44363360825604},
				{111.96372985839844, 38.42723559654225},
				{111.91085815429688, 38.344887442462806},
				{111.87309265136717, 38.24680876017446},
				{111.8671003424276, 38.22285924285058}},
			matrix.LineMatrix{{111.8671003424276, 38.22285924285058},
				{111.84906005859375, 38.15075747130226}},
		}, false},
	{Name: "line poly6", Fields: []matrix.Steric{
		matrix.LineMatrix{{200, 300}, {500, 300}, {500, 600}, {800, 900}},
		matrix.PolygonMatrix{{{300, 300}, {500, 300}, {500, 500}, {300, 500}, {300, 300}}},
	},
		Want: matrix.Collection{matrix.LineMatrix{{500, 500}, {500, 600}, {800, 900}},
			matrix.LineMatrix{{200, 300}, {300, 300}}},
	},
}
