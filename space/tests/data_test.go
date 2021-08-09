package test

import (
	"github.com/spatial-go/geoos/space"
)

type args struct {
	g space.Geometry
}

var TestsCentroid = []struct {
	name    string
	args    args
	want    space.Geometry
	wantErr bool
}{

	{name: "P - empty", args: args{g: nil}, want: nil, wantErr: false},
	{name: "P - single point", args: args{g: space.Point{10, 10}}, want: space.Point{10, 10}, wantErr: false},
	{name: "mP - two points", args: args{g: space.MultiPoint{{10, 10}, {20, 20}}}, want: space.Point{15, 15}, wantErr: false},
	{name: "mP - 4 points", args: args{g: space.MultiPoint{{10, 10}, {20, 20}, {10, 20}, {20, 10}}}, want: space.Point{15, 15}, wantErr: false},
	{name: "mP - repeated points", args: args{g: space.MultiPoint{{10, 10}, {10, 10}, {10, 10}, {18, 18}}}, want: space.Point{12, 12}, wantErr: false},
	{name: "L - single segment", args: args{g: space.LineString{{10, 10}, {20, 20}}}, want: space.Point{15, 15}, wantErr: false},
	{name: "L - zero length line", args: args{g: space.LineString{{10, 10}, {10, 10}}}, want: space.Point{10, 10}, wantErr: false},
	{name: "mL - zero length lines", args: args{g: space.MultiLineString{space.LineString{{10, 10}, {10, 10}},
		space.LineString{{20, 20}, {20, 20}},
	}}, want: space.Point{15, 15}, wantErr: false},

	{name: "L - two segments", args: args{g: space.LineString{{60, 180}, {120, 100}, {180, 180}}}, want: space.Point{120, 140}, wantErr: false},

	{name: "L - elongated horseshoe", args: args{g: space.LineString{{80.0, 0.0}, {80.0, 120.0}, {120.0, 120.0}, {120.0, 0.0}}}, want: space.Point{100, 68.57142857142857}, wantErr: false},
	{name: "mL - two single-segment lines", args: args{g: space.MultiLineString{{{0, 0}, {0, 100}}, {{100, 0}, {100, 100}}}}, want: space.Point{50, 50}, wantErr: false},
	{name: "mL - two concentric rings, offset", args: args{g: space.MultiLineString{{{0, 0}, {0, 200}, {200, 200}, {200, 0}, {0, 0}},
		{{60, 180}, {20, 180}, {20, 140}, {60, 140}, {60, 180}}}}, want: space.Point{90, 110}, wantErr: false},
	{name: "mL - complicated symmetrical collection of lines", args: args{g: space.MultiLineString{{{20, 20}, {60, 60}},
		{{20, -20}, {60, -60}},
		{{-20, -20}, {-60, -60}},
		{{-20, 20}, {-60, 60}},
		{{-80, 0}, {0, 80}, {80, 0}, {0, -80}, {-80, 0}},
		{{-40, 20}, {-40, -20}},
		{{-20, 40}, {20, 40}},
		{{40, 20}, {40, -20}},
		{{20, -40}, {-20, -40}},
	}}, want: space.Point{0, 0}, wantErr: false},
	{name: "A - empty", args: args{g: space.Polygon{}}, want: nil, wantErr: false},
	{name: "A - box", args: args{g: space.Polygon{{{40, 160}, {160, 160}, {160, 40}, {40, 40}, {40, 160}}}}, want: space.Point{100, 100}, wantErr: false},
	{name: "A - box with hole", args: args{g: space.Polygon{{{0, 200}, {200, 200}, {200, 0}, {0, 0}, {0, 200}},
		{{20, 180}, {80, 180}, {80, 20}, {20, 20}, {20, 180}},
	}}, want: space.Point{115.78947368421052, 100}, wantErr: false},
	{name: "A - box with offset hole (showing difference between area and line centroid)",
		args: args{g: space.Polygon{{{0, 0}, {0, 200}, {200, 200}, {200, 0}, {0, 0}},
			{{60, 180}, {20, 180}, {20, 140}, {60, 140}, {60, 180}},
		}}, want: space.Point{102.5, 97.5}, wantErr: false},
	{name: "A - box  with 2 symmetric holes",
		args: args{g: space.Polygon{{{0, 0}, {0, 200}, {200, 200}, {200, 0}, {0, 0}},
			{{60, 180}, {20, 180}, {20, 140}, {60, 140}, {60, 180}},
			{{180, 60}, {140, 60}, {140, 20}, {180, 20}, {180, 60}},
		}}, want: space.Point{100, 100}, wantErr: false},
	{
		name: "A - invalid box ",
		args: args{g: space.Polygon{{{0, 0}, {0, 0}, {200, 0}, {200, 0}, {0, 0}}}},
		want: space.Point{100, 0}, wantErr: false,
	},

	{
		name: "A - invalid box - too few points",
		args: args{g: space.Polygon{{{0, 0}, {100, 100}, {0, 0}}}},
		want: space.Point{50, 50}, wantErr: false,
	},

	{
		name: "mA - symmetric angles",
		args: args{g: space.MultiPolygon{{{{0, 40}, {0, 140}, {140, 140}, {140, 120}, {20, 120}, {20, 40}, {0, 40}},
			{{0, 0}, {0, 20}, {120, 20}, {120, 100}, {140, 100}, {140, 0}, {0, 0}}}},
		},
		want: space.Point{70, 70}, wantErr: false,
	},

	{
		name: "GC - two adjacent Polygons (showing that centroids are additive) ",
		args: args{g: space.Collection{space.Polygon{{{0, 200}, {20, 180}, {20, 140}, {60, 140}, {200, 0}, {0, 0}, {0, 200}}},
			space.Polygon{{{200, 200}, {0, 200}, {20, 180}, {60, 180}, {60, 140}, {200, 0}, {200, 200}}},
		},
		},
		want: space.Point{102.5, 97.5}, wantErr: false,
	},

	{
		name: "GC - heterogeneous collection of lines, points",
		args: args{g: space.Collection{space.LineString{{80, 0}, {80, 120}, {120, 120}, {120, 0}},
			space.MultiPoint{{20, 60}, {40, 80}, {60, 60}},
		}},
		want: space.Point{100, 68.57142857142857}, wantErr: false,
	},

	{
		name: "GC - heterogeneous collection of Polygons, line",
		args: args{g: space.Collection{space.Polygon{{{0, 40}, {40, 40}, {40, 0}, {0, 0}, {0, 40}}},
			space.LineString{{80, 0}, {80, 80}, {120, 40}},
		}},
		want: space.Point{20, 20}, wantErr: false,
	},

	{
		name: "GC - collection of Polygons, lines, points",
		args: args{g: space.Collection{space.Polygon{{{0, 40}, {40, 40}, {40, 0}, {0, 0}, {0, 40}}},
			space.LineString{{80, 0}, {80, 80}, {120, 40}},
			space.MultiPoint{{20, 60}, {40, 80}, {60, 60}},
		}},
		want: space.Point{20, 20}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons and lines",
		args: args{g: space.Collection{space.Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			space.LineString{{20, 20}, {30, 30}},
		}},
		want: space.Point{25, 25}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons and zero-length lines",
		args: args{g: space.Collection{space.Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			space.LineString{{20, 20}, {20, 20}},
		}},
		want: space.Point{15, 15}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons, zero-length lines, and points",
		args: args{g: space.Collection{space.Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			space.LineString{{20, 20}, {20, 20}},
			space.MultiPoint{{20, 10}, {10, 20}},
		}},
		want: space.Point{15, 15}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons, zero-length lines, and points",
		args: args{g: space.Collection{space.Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			space.LineString{{20, 20}, {20, 20}},
			space.Point{},
		}},
		want: space.Point{15, 15}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons, zero-length lines, and points",
		args: args{g: space.Collection{space.Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			space.LineString{},
			space.Point{},
		}},
		want: space.Point{10, 10}, wantErr: false,
	},

	{
		name: "GC - collection with empty Polygon, line, and point",
		args: args{g: space.Collection{space.Polygon{},
			space.LineString{{20, 20}, {30, 30}, {40, 40}},
			space.MultiPoint{{20, 10}, {10, 20}},
		}},
		want: space.Point{30, 30}, wantErr: false,
	},

	{
		name: "GC - collection with empty Polygon, empty line, and point",
		args: args{g: space.Collection{space.Polygon{},
			space.LineString{},
			space.Point{10, 10},
		}},
		want: space.Point{10, 10}, wantErr: false,
	},

	{
		name: "GC - collection with empty Polygon, empty line, and empty point",
		args: args{g: space.Collection{space.Polygon{},
			space.LineString{},
			space.Point{},
		}},
		want: nil, wantErr: false,
	},

	{
		name: "GC - overlapping Polygons ",
		args: args{g: space.Collection{space.Polygon{{{20, 100}, {20, -20}, {60, -20}, {60, 100}, {20, 100}}},
			space.Polygon{{{-20, 60}, {100, 60}, {100, 20}, {-20, 20}, {-20, 60}}},
		}},
		want: space.Point{40, 40}, wantErr: false,
	},

	{
		name: "A - degenerate box",
		args: args{g: space.Polygon{{{40, 160}, {160, 160}, {160, 160}, {40, 160}, {40, 160}}}},
		want: space.Point{100, 160}, wantErr: false,
	},

	{
		name: "A - degenerate triangle",
		args: args{g: space.Polygon{{{10, 10}, {100, 100}, {100, 100}, {10, 10}}}},
		want: space.Point{55, 55}, wantErr: false,
	},

	{
		name: "A - almost degenerate triangle",
		args: args{g: space.Polygon{{{
			56.528666666700, 25.2101666667}, {
			56.529000000000, 25.2105000000}, {
			56.528833333300, 25.2103333333}, {
			56.528666666700, 25.2101666667}}},
		},
		want: space.Point{56.52883333335, 25.21033333335}, wantErr: false,
	},

	{
		name: "A - almost degenerate MultiPolygon",
		args: args{g: space.MultiPolygon{{{{
			-92.661322, 36.58994900000003},
			{-92.66132199999993, 36.58994900000005},
			{-92.66132199999993, 36.589949000000004},
			{-92.661322, 36.589949},
			{-92.661322, 36.58994900000003}}},
			{{{
				-92.65560500000008, 36.58708800000005},
				{-92.65560499999992, 36.58708800000005},
				{-92.65560499998745, 36.587087999992576},
				{-92.655605, 36.587088},
				{-92.65560500000008, 36.58708800000005}}},
			{{{
				-92.65512450000065, 36.586800000000466},
				{-92.65512449999994, 36.58680000000004},
				{-92.65512449998666, 36.5867999999905},
				{-92.65512450000065, 36.586800000000466}}},
		}},
		want: space.Point{-92.6553838608954, 36.58695407733924}, wantErr: false,
	},
}
