package space

type args struct {
	g Geometry
}

var TestsCentroid = []struct {
	name    string
	args    args
	want    Geometry
	wantErr bool
}{

	{name: "P - empty", args: args{g: nil}, want: nil, wantErr: false},
	{name: "P - single point", args: args{g: Point{10, 10}}, want: Point{10, 10}, wantErr: false},
	{name: "mP - two points", args: args{g: MultiPoint{{10, 10}, {20, 20}}}, want: Point{15, 15}, wantErr: false},
	{name: "mP - 4 points", args: args{g: MultiPoint{{10, 10}, {20, 20}, {10, 20}, {20, 10}}}, want: Point{15, 15}, wantErr: false},
	{name: "mP - repeated points", args: args{g: MultiPoint{{10, 10}, {10, 10}, {10, 10}, {18, 18}}}, want: Point{12, 12}, wantErr: false},
	{name: "L - single segment", args: args{g: LineString{{10, 10}, {20, 20}}}, want: Point{15, 15}, wantErr: false},
	{name: "L - zero length line", args: args{g: LineString{{10, 10}, {10, 10}}}, want: Point{10, 10}, wantErr: false},
	{name: "mL - zero length lines", args: args{g: MultiLineString{LineString{{10, 10}, {10, 10}},
		LineString{{20, 20}, {20, 20}},
	}}, want: Point{15, 15}, wantErr: false},

	{name: "L - two segments", args: args{g: LineString{{60, 180}, {120, 100}, {180, 180}}}, want: Point{120, 140}, wantErr: false},

	{name: "L - elongated horseshoe", args: args{g: LineString{{80.0, 0.0}, {80.0, 120.0}, {120.0, 120.0}, {120.0, 0.0}}}, want: Point{100, 68.57142857142857}, wantErr: false},
	{name: "mL - two single-segment lines", args: args{g: MultiLineString{{{0, 0}, {0, 100}}, {{100, 0}, {100, 100}}}}, want: Point{50, 50}, wantErr: false},
	{name: "mL - two concentric rings, offset", args: args{g: MultiLineString{{{0, 0}, {0, 200}, {200, 200}, {200, 0}, {0, 0}},
		{{60, 180}, {20, 180}, {20, 140}, {60, 140}, {60, 180}}}}, want: Point{90, 110}, wantErr: false},
	{name: "mL - complicated symmetrical collection of lines", args: args{g: MultiLineString{{{20, 20}, {60, 60}},
		{{20, -20}, {60, -60}},
		{{-20, -20}, {-60, -60}},
		{{-20, 20}, {-60, 60}},
		{{-80, 0}, {0, 80}, {80, 0}, {0, -80}, {-80, 0}},
		{{-40, 20}, {-40, -20}},
		{{-20, 40}, {20, 40}},
		{{40, 20}, {40, -20}},
		{{20, -40}, {-20, -40}},
	}}, want: Point{0, 0}, wantErr: false},
	{name: "A - empty", args: args{g: Polygon{}}, want: nil, wantErr: false},
	{name: "A - box", args: args{g: Polygon{{{40, 160}, {160, 160}, {160, 40}, {40, 40}, {40, 160}}}}, want: Point{100, 100}, wantErr: false},
	{name: "A - box with hole", args: args{g: Polygon{{{0, 200}, {200, 200}, {200, 0}, {0, 0}, {0, 200}},
		{{20, 180}, {80, 180}, {80, 20}, {20, 20}, {20, 180}},
	}}, want: Point{115.78947368421052, 100}, wantErr: false},
	{name: "A - box with offset hole (showing difference between area and line centroid)",
		args: args{g: Polygon{{{0, 0}, {0, 200}, {200, 200}, {200, 0}, {0, 0}},
			{{60, 180}, {20, 180}, {20, 140}, {60, 140}, {60, 180}},
		}}, want: Point{102.5, 97.5}, wantErr: false},
	{name: "A - box  with 2 symmetric holes",
		args: args{g: Polygon{{{0, 0}, {0, 200}, {200, 200}, {200, 0}, {0, 0}},
			{{60, 180}, {20, 180}, {20, 140}, {60, 140}, {60, 180}},
			{{180, 60}, {140, 60}, {140, 20}, {180, 20}, {180, 60}},
		}}, want: Point{100, 100}, wantErr: false},
	{
		name: "A - invalid box ",
		args: args{g: Polygon{{{0, 0}, {0, 0}, {200, 0}, {200, 0}, {0, 0}}}},
		want: Point{100, 0}, wantErr: false,
	},

	{
		name: "A - invalid box - too few points",
		args: args{g: Polygon{{{0, 0}, {100, 100}, {0, 0}}}},
		want: Point{50, 50}, wantErr: false,
	},

	{
		name: "mA - symmetric angles",
		args: args{g: MultiPolygon{{{{0, 40}, {0, 140}, {140, 140}, {140, 120}, {20, 120}, {20, 40}, {0, 40}},
			{{0, 0}, {0, 20}, {120, 20}, {120, 100}, {140, 100}, {140, 0}, {0, 0}}}},
		},
		want: Point{70, 70}, wantErr: false,
	},

	{
		name: "GC - two adjacent Polygons (showing that centroids are additive) ",
		args: args{g: Collection{Polygon{{{0, 200}, {20, 180}, {20, 140}, {60, 140}, {200, 0}, {0, 0}, {0, 200}}},
			Polygon{{{200, 200}, {0, 200}, {20, 180}, {60, 180}, {60, 140}, {200, 0}, {200, 200}}},
		},
		},
		want: Point{102.5, 97.5}, wantErr: false,
	},

	{
		name: "GC - heterogeneous collection of lines, points",
		args: args{g: Collection{LineString{{80, 0}, {80, 120}, {120, 120}, {120, 0}},
			MultiPoint{{20, 60}, {40, 80}, {60, 60}},
		}},
		want: Point{100, 68.57142857142857}, wantErr: false,
	},

	{
		name: "GC - heterogeneous collection of Polygons, line",
		args: args{g: Collection{Polygon{{{0, 40}, {40, 40}, {40, 0}, {0, 0}, {0, 40}}},
			LineString{{80, 0}, {80, 80}, {120, 40}},
		}},
		want: Point{20, 20}, wantErr: false,
	},

	{
		name: "GC - collection of Polygons, lines, points",
		args: args{g: Collection{Polygon{{{0, 40}, {40, 40}, {40, 0}, {0, 0}, {0, 40}}},
			LineString{{80, 0}, {80, 80}, {120, 40}},
			MultiPoint{{20, 60}, {40, 80}, {60, 60}},
		}},
		want: Point{20, 20}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons and lines",
		args: args{g: Collection{Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			LineString{{20, 20}, {30, 30}},
		}},
		want: Point{25, 25}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons and zero-length lines",
		args: args{g: Collection{Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			LineString{{20, 20}, {20, 20}},
		}},
		want: Point{15, 15}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons, zero-length lines, and points",
		args: args{g: Collection{Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			LineString{{20, 20}, {20, 20}},
			MultiPoint{{20, 10}, {10, 20}},
		}},
		want: Point{15, 15}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons, zero-length lines, and points",
		args: args{g: Collection{Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			LineString{{20, 20}, {20, 20}},
			Point{},
		}},
		want: Point{15, 15}, wantErr: false,
	},

	{
		name: "GC - collection of zero-area Polygons, zero-length lines, and points",
		args: args{g: Collection{Polygon{{{10, 10}, {10, 10}, {10, 10}, {10, 10}}},
			LineString{},
			Point{},
		}},
		want: Point{10, 10}, wantErr: false,
	},

	{
		name: "GC - collection with empty Polygon, line, and point",
		args: args{g: Collection{Polygon{},
			LineString{{20, 20}, {30, 30}, {40, 40}},
			MultiPoint{{20, 10}, {10, 20}},
		}},
		want: Point{30, 30}, wantErr: false,
	},

	{
		name: "GC - collection with empty Polygon, empty line, and point",
		args: args{g: Collection{Polygon{},
			LineString{},
			Point{10, 10},
		}},
		want: Point{10, 10}, wantErr: false,
	},

	{
		name: "GC - collection with empty Polygon, empty line, and empty point",
		args: args{g: Collection{Polygon{},
			LineString{},
			Point{},
		}},
		want: nil, wantErr: false,
	},

	{
		name: "GC - overlapping Polygons ",
		args: args{g: Collection{Polygon{{{20, 100}, {20, -20}, {60, -20}, {60, 100}, {20, 100}}},
			Polygon{{{-20, 60}, {100, 60}, {100, 20}, {-20, 20}, {-20, 60}}},
		}},
		want: Point{40, 40}, wantErr: false,
	},

	{
		name: "A - degenerate box",
		args: args{g: Polygon{{{40, 160}, {160, 160}, {160, 160}, {40, 160}, {40, 160}}}},
		want: Point{100, 160}, wantErr: false,
	},

	{
		name: "A - degenerate triangle",
		args: args{g: Polygon{{{10, 10}, {100, 100}, {100, 100}, {10, 10}}}},
		want: Point{55, 55}, wantErr: false,
	},

	{
		name: "A - almost degenerate triangle",
		args: args{g: Polygon{{{
			56.528666666700, 25.2101666667}, {
			56.529000000000, 25.2105000000}, {
			56.528833333300, 25.2103333333}, {
			56.528666666700, 25.2101666667}}},
		},
		want: Point{56.52883333335, 25.21033333335}, wantErr: false,
	},

	{
		name: "A - almost degenerate MultiPolygon",
		args: args{g: MultiPolygon{{{{
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
		want: Point{-92.6553838608954, 36.58695407733924}, wantErr: false,
	},
}
