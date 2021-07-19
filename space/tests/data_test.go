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
}

/*


<case>
  <desc>A - invalid box </desc>
  <a>    POLYGON ((0 0, 0 0, 200 0, 200 0, 0 0))
   </a>
<test><op name="getCentroid" arg1="A" >   POINT (100 0)   </op></test>
</case>

<case>
  <desc>A - invalid box - too few points</desc>
  <a>    POLYGON ((0 0, 100 100, 0 0))
   </a>
<test><op name="getCentroid" arg1="A" >   POINT (50 50)   </op></test>
</case>

<case>
  <desc>mA - symmetric angles</desc>
  <a>    MULTIPOLYGON (((0 40, 0 140, 140 140, 140 120, 20 120, 20 40, 0 40)),
  ((0 0, 0 20, 120 20, 120 100, 140 100, 140 0, 0 0)))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (70 70)   </op></test>
</case>

<case>
  <desc>GC - two adjacent polygons (showing that centroids are additive) </desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((0 200, 20 180, 20 140, 60 140, 200 0, 0 0, 0 200)),
  POLYGON ((200 200, 0 200, 20 180, 60 180, 60 140, 200 0, 200 200)))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (102.5 97.5)   </op></test>
</case>

<case>
  <desc>GC - heterogeneous collection of lines, points</desc>
  <a>    GEOMETRYCOLLECTION (LINESTRING (80 0, 80 120, 120 120, 120 0),
  MULTIPOINT ((20 60), (40 80), (60 60)))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (100 68.57142857142857)   </op></test>
</case>

<case>
  <desc>GC - heterogeneous collection of polygons, line</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((0 40, 40 40, 40 0, 0 0, 0 40)),
  LINESTRING (80 0, 80 80, 120 40))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (20 20)   </op></test>
</case>

<case>
  <desc>GC - collection of polygons, lines, points</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((0 40, 40 40, 40 0, 0 0, 0 40)),
  LINESTRING (80 0, 80 80, 120 40),
  MULTIPOINT ((20 60), (40 80), (60 60)))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (20 20)   </op></test>
</case>

<case>
  <desc>GC - collection of zero-area polygons and lines</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((10 10, 10 10, 10 10, 10 10)),
  LINESTRING (20 20, 30 30))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (25 25)   </op></test>
</case>

<case>
  <desc>GC - collection of zero-area polygons and zero-length lines</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((10 10, 10 10, 10 10, 10 10)),
  LINESTRING (20 20, 20 20))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (15 15)   </op></test>
</case>

<case>
  <desc>GC - collection of zero-area polygons, zero-length lines, and points</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((10 10, 10 10, 10 10, 10 10)),
  LINESTRING (20 20, 20 20),
  MULTIPOINT ((20 10), (10 20)) )
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (15 15)   </op></test>
</case>

<case>
  <desc>GC - collection of zero-area polygons, zero-length lines, and points</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((10 10, 10 10, 10 10, 10 10)),
  LINESTRING (20 20, 20 20),
  POINT EMPTY )
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (15 15)   </op></test>
</case>

<case>
  <desc>GC - collection of zero-area polygons, zero-length lines, and points</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((10 10, 10 10, 10 10, 10 10)),
  LINESTRING EMPTY,
  POINT EMPTY )
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (10 10)   </op></test>
</case>

<case>
  <desc>GC - collection with empty polygon, line, and point</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON EMPTY,
  LINESTRING (20 20, 30 30, 40 40),
  MULTIPOINT ((20 10), (10 20)) )
  </a>
<test><op name="getCentroid" arg1="A" >    POINT (30 30)   </op></test>
</case>

<case>
  <desc>GC - collection with empty polygon, empty line, and point</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON EMPTY,
  LINESTRING EMPTY,
  POINT (10 10) )
  </a>
<test><op name="getCentroid" arg1="A" >    POINT (10 10)   </op></test>
</case>

<case>
  <desc>GC - collection with empty polygon, empty line, and empty point</desc>
  <a>    GEOMETRYCOLLECTION (POLYGON EMPTY,
  LINESTRING EMPTY,
  POINT EMPTY )
  </a>
<test><op name="getCentroid" arg1="A" >    POINT EMPTY  </op></test>
</case>

<case>
  <desc>GC - overlapping polygons </desc>
  <a>    GEOMETRYCOLLECTION (POLYGON ((20 100, 20 -20, 60 -20, 60 100, 20 100)),
  POLYGON ((-20 60, 100 60, 100 20, -20 20, -20 60)))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (40 40)   </op></test>
</case>

<case>
  <desc>A - degenerate box</desc>
  <a>    POLYGON ((40 160, 160 160, 160 160, 40 160, 40 160))  </a>
<test><op name="getCentroid" arg1="A" >    POINT (100 160)   </op></test>
</case>

<case>
  <desc>A - degenerate triangle</desc>
  <a>    POLYGON ((10 10, 100 100, 100 100, 10 10))  </a>
<test><op name="getCentroid" arg1="A" >    POINT (55 55)   </op></test>
</case>

<case>
  <desc>A - almost degenerate triangle</desc>
  <a>    POLYGON((
56.528666666700 25.2101666667,
56.529000000000 25.2105000000,
56.528833333300 25.2103333333,
56.528666666700 25.2101666667))
	</a>
<test><op name="getCentroid" arg1="A" >    POINT (56.52883333335 25.21033333335)  </op></test>
</case>

<case>
  <desc>A - almost degenerate MultiPolygon</desc>
  <a>
    MULTIPOLYGON (((
     -92.661322 36.58994900000003,
     -92.66132199999993 36.58994900000005,
     -92.66132199999993 36.589949000000004,
     -92.661322 36.589949,
     -92.661322 36.58994900000003)),
    ((
     -92.65560500000008 36.58708800000005,
     -92.65560499999992 36.58708800000005,
     -92.65560499998745 36.587087999992576,
     -92.655605 36.587088,
     -92.65560500000008 36.58708800000005
    )),
    ((
     -92.65512450000065 36.586800000000466,
      -92.65512449999994 36.58680000000004,
     -92.65512449998666 36.5867999999905,
      -92.65512450000065 36.586800000000466
    )))
  </a>
  <test><op name="getCentroid" arg1="A" >POINT (-92.6553838608954 36.58695407733924)</op></test>
</case>
*/
