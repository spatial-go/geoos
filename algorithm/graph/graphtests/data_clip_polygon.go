// package graphtests ...

package graphtests

import (
	"github.com/spatial-go/geoos/algorithm/matrix"
)

// TestCaseMore ...
type TestCaseMore struct {
	Name    string
	Fields  []matrix.Steric
	Want    []matrix.Steric
	WantErr bool
}

// TestsPolygonIntersecation ...
var TestsPolygonIntersecation = []TestCase{
	{"poly point0", []matrix.Steric{
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		matrix.Matrix{100, 100}},
		matrix.Matrix{100, 100}, false},
	{"poly line0", []matrix.Steric{
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		matrix.LineMatrix{{100, 100}, {101, 101}}},
		matrix.LineMatrix{{100, 100}, {101, 101}}, false},
	{"poly line1", []matrix.Steric{
		matrix.PolygonMatrix{{{105, 105}, {105, 101}, {101, 101}, {101, 105}, {105, 105}}},
		matrix.LineMatrix{{100, 100}, {90, 101}}},
		matrix.Collection{}, false},

	{"poly poly1", []matrix.Steric{
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
	},
		matrix.PolygonMatrix{{{101, 100}, {100, 100}, {100, 101}, {101, 101}, {101, 100}}}, false},

	{"poly poly2", []matrix.Steric{
		matrix.PolygonMatrix{{{105, 105}, {105, 103}, {103, 103}, {103, 105}, {105, 105}}},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
	},
		matrix.PolygonMatrix{}, false},
	{"poly poly3", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
	},
		matrix.PolygonMatrix{{{5, 10}, {10, 10}, {10, 5}, {5, 5}, {5, 10}}}, false},
	{"poly poly3-1", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
		matrix.PolygonMatrix{{{5, 5}, {5, 15}, {15, 15}, {15, 5}, {5, 5}}},
	},
		matrix.PolygonMatrix{{{5, 10}, {10, 10}, {10, 5}, {5, 5}, {5, 10}}}, false},
	{"poly poly4", []matrix.Steric{
		matrix.PolygonMatrix{{{111.30523681640625, 38.117271658305},
			{112.34344482421875, 38.11727165830543},
			{112.34344482421875, 38.89103282648846},
			{111.30523681640625, 38.89103282648846},
			{111.30523681640625, 38.117271658305}}},
		matrix.PolygonMatrix{{{111.50848388671875, 37.6359849542696},
			{112.64007568359375, 37.6359849542696},
			{112.64007568359375, 38.35027253825765},
			{111.50848388671875, 38.35027253825765},
			{111.50848388671875, 37.6359849542696}}},
	},
		matrix.PolygonMatrix{{{112.34344482421875, 38.35027253825765},
			{112.34344482421875, 38.11727165830543},
			{111.50848388671875, 38.11727165830543},
			{111.50848388671875, 38.35027253825765},
			{112.34344482421875, 38.35027253825765}}}, false},

	{"poly poly5", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}},
			{{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
	},
		matrix.PolygonMatrix{{{9, 9}, {9, 5}, {10, 5}, {10, 10}, {5, 10}, {5, 9}, {9, 9}}}, false},
	{"poly poly5-1", []matrix.Steric{
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}},
			{{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
	},
		matrix.PolygonMatrix{{{9, 9}, {9, 5}, {10, 5}, {10, 10}, {5, 10}, {5, 9}, {9, 9}}}, false},
}

// TestsPolygonUnion ...
var TestsPolygonUnion = []TestCaseMore{
	{"poly poly", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
	},
		[]matrix.Steric{matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5},
			{15, 5}, {15, 15}, {5, 15}, {5, 10}}}}, false},

	{"poly poly01", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
		matrix.PolygonMatrix{{{5, 5}, {5, 15}, {15, 15}, {15, 5}, {5, 5}}},
	},
		[]matrix.Steric{
			matrix.PolygonMatrix{{{10, 5}, {10, 0}, {0, 0}, {0, 10}, {5, 10}, {5, 15}, {15, 15}, {15, 5}, {10, 5}}},
			matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}}},
		}, false},

	{"poly poly02", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}},
			{{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
	},
		[]matrix.Steric{
			matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}},
				{{5, 9}, {1, 9}, {1, 1}, {9, 1}, {9, 5}, {5, 5}, {5, 9}}}}, false},

	{"poly poly03", []matrix.Steric{
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{1, 1}, {9, 1}, {9, 9}, {1, 9}, {1, 1}}},
	},
		[]matrix.Steric{
			matrix.PolygonMatrix{{{10, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}},
				{{9, 5}, {5, 5}, {5, 9}, {1, 9}, {1, 1}, {9, 1}, {9, 5}}}}, false},

	{"poly poly1", []matrix.Steric{
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
	},
		[]matrix.Steric{
			matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}}}, false},

	{Name: "poly 1",
		Fields: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}},
		}, Want: []matrix.Steric{
			matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {1, 2}, {2, 2}, {2, 1}, {1, 1}}},
				matrix.PolygonMatrix{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
			matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{3, 1}, {5, 1}, {5, 2}, {3, 2}, {3, 1}}}},
		},
		WantErr: false},

	{Name: "poly 2",
		Fields: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}},
		}, Want: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {1, 2}, {1, 1}, {2, 1}, {5, 1}, {5, 2}, {2, 2}}},
		},
		WantErr: false},

	{Name: "poly 3",
		Fields: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}},
		}, Want: []matrix.Steric{
			matrix.Collection{matrix.PolygonMatrix{{{2, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}}},
				matrix.PolygonMatrix{{{2, 2}, {2, 3}, {5, 3}, {5, 2}, {2, 2}}}},
			matrix.Collection{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
				matrix.PolygonMatrix{{{2, 2}, {5, 2}, {5, 3}, {2, 3}, {2, 2}}}},
		},
		WantErr: false},

	{Name: "poly 4",
		Fields: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 2}, {3, 2}, {3, 3}, {1, 3}, {1, 2}}},
			matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 1}}},
		}, Want: []matrix.Steric{
			matrix.PolygonMatrix{{{2, 2}, {1, 2}, {1, 3}, {2, 3}, {2, 5}, {5, 5}, {5, 1}, {2, 1}, {2, 2}}},
			matrix.PolygonMatrix{{{2, 3}, {1, 3}, {1, 2}, {2, 2}, {2, 1}, {5, 1}, {5, 5}, {2, 5}, {2, 3}}},
		},
		WantErr: false},

	{Name: "poly 5",
		Fields: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
			matrix.PolygonMatrix{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}},
		}, Want: []matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {1, 5}, {5, 5}, {5, 1}, {1, 1}}},
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
		},
		WantErr: false},

	{Name: "poly 6",
		Fields: []matrix.Steric{matrix.PolygonMatrix{{{1, 1}, {2, 1}, {2, 2}, {1, 2}, {1, 1}}},
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
		}, Want: []matrix.Steric{
			matrix.PolygonMatrix{{{2, 1}, {1, 1}, {1, 2}, {1, 3}, {5, 3}, {5, 1}, {2, 1}}},
			matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 3}, {1, 3}, {1, 1}}},
		},
		WantErr: false},

	{Name: "poly g1",
		Fields: []matrix.Steric{
			matrix.PolygonMatrix{{{113.58272043315301, 34.737572973722635},
				{113.60467185859878, 34.7375632566498},
				{113.60467414450156, 34.75955780700805},
				{113.58272272415182, 34.75956752361057},
				{113.58272043315301, 34.737572973722635}}},
			matrix.PolygonMatrix{{{113.58270814000005, 34.71557888104119},
				{113.58271043000005, 34.73757297104119},
				{113.58271336086447, 34.737580042999575},
				{113.5827204344234, 34.737582969999025},
				{113.6046718644234, 34.737573259999024},
				{113.60467893299956, 34.73757032913553},
				{113.60468185999994, 34.737563258958815},
				{113.60468185999994, 34.737563258958815},
				{113.60467663913391, 34.715562086998816},
				{113.60466956557204, 34.71555916000098},
				{113.58271813557204, 34.71556888000098},
				{113.58271106699881, 34.71557181086609},
				{113.58270814000005, 34.71557888104119}}},
		}, Want: []matrix.Steric{
			matrix.PolygonMatrix{{{113.60467185963843, 34.73757326000114},
				{113.60467414450156, 34.75955780700805},
				{113.58272272415182, 34.75956752361057},
				{113.58272043419424, 34.7375829699042},
				{113.58271336086447, 34.737580042999575},
				{113.58271043000005, 34.73757297104119},
				{113.58270814000005, 34.71557888104119},
				{113.58271106699881, 34.71557181086609},
				{113.58271813557204, 34.71556888000098},
				{113.60466956557204, 34.71555916000098},
				{113.60467663913391, 34.715562086998816},
				{113.60468185999994, 34.737563258958815},
				{113.60467893299956, 34.73757032913553},
				{113.6046718644234, 34.737573259999024},
				{113.60467185963843, 34.73757326000114}}}},
		WantErr: false},
	{"poly poly_xx", []matrix.Steric{
		matrix.PolygonMatrix{{{5, 5}, {10, 5}, {10, 10}, {5, 15}, {5, 10}, {5, 5}}},
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
	},
		[]matrix.Steric{
			matrix.PolygonMatrix{{{0, 0}, {10, 0}, {5, 10}, {10, 10}, {5, 15}, {5, 10}, {0, 10}, {0, 0}}}}, false},
	{"poly x1", []matrix.Steric{
		matrix.PolygonMatrix{{{1, 1}, {2, 1}, {5, 1}, {5, 2}, {2, 2}, {1, 2}, {1, 1}}},
		matrix.PolygonMatrix{{{2, 1}, {5, 1}, {5, 2}, {2, 2}, {2, 1}}},
	},
		[]matrix.Steric{
			matrix.PolygonMatrix{{{1, 1}, {2, 1}, {5, 1}, {5, 2}, {2, 2}, {1, 2}, {1, 1}}}}, false},
}

// TestsPolygonDifference ...
var TestsPolygonDifference = []TestCase{
	{"poly poly", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
	},
		matrix.PolygonMatrix{{{5, 10}, {0, 10}, {0, 0}, {10, 0}, {10, 5}, {5, 5}, {5, 10}}}, false},

	{"poly poly1", []matrix.Steric{
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
	},
		matrix.PolygonMatrix{}, false},

	{"poly poly2", []matrix.Steric{
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
	},
		matrix.PolygonMatrix{{{100, 101}, {90, 101}, {90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}}}, false},
	{"poly poly3", []matrix.Steric{
		matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}}},
		matrix.PolygonMatrix{{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}},
	},
		matrix.PolygonMatrix{{{1, 1}, {5, 1}, {5, 5}, {1, 5}, {1, 1}},
			{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {2, 2}}}, false},
}

// TestsPolygonSymDifference ...
var TestsPolygonSymDifference = []TestCase{
	{"poly poly", []matrix.Steric{
		matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}},
		matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}},
	},
		matrix.Collection{
			matrix.PolygonMatrix{{{0, 0}, {0, 10}, {5, 10}, {5, 5}, {10, 5}, {10, 0}, {0, 0}}},
			matrix.PolygonMatrix{{{5, 10}, {5, 15}, {15, 15}, {15, 5}, {10, 5}, {10, 10}, {5, 10}}},
		}, false},
	{"poly poly0", []matrix.Steric{
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
	},

		matrix.PolygonMatrix{{{90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}, {90, 101}, {90, 90}}},
		false},
	{"poly poly1", []matrix.Steric{
		matrix.PolygonMatrix{{{90, 90}, {90, 101}, {101, 101}, {101, 90}, {90, 90}}},
		matrix.PolygonMatrix{{{100, 100}, {100, 101}, {101, 101}, {101, 100}, {100, 100}}},
	},
		matrix.PolygonMatrix{{{90, 90}, {101, 90}, {101, 100}, {100, 100}, {100, 101}, {90, 101}, {90, 90}}},
		false},
}
