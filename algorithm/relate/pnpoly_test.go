package relate

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

type args struct {
	point matrix.Matrix
	want  bool
}

var testData = []struct {
	name string
	poly matrix.LineMatrix
	args []args
}{
	// test cases from http://rosettacode.org/wiki/Ray-casting_algorithm#Go
	{"rc square",
		matrix.LineMatrix{{0, 0}, {10, 0}, {10, 10}, {0, 10}},
		[]args{
			{matrix.Matrix{5, 5}, true},
			{matrix.Matrix{5, 8}, true},
			{matrix.Matrix{-10, 5}, false},
			{matrix.Matrix{8, 5}, true},
			{matrix.Matrix{1, 2}, true},
			{matrix.Matrix{2, 1}, true},
			{matrix.Matrix{0, 0}, true},
		},
	},
	{"rc square hole", // (there's a 0-width isthmus)
		matrix.LineMatrix{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0},
			{2.5, 2.5}, {7.5, 2.5}, {7.5, 7.5}, {2.5, 7.5}, {2.5, 2.5}},
		[]args{
			{matrix.Matrix{5, 5}, false},
			{matrix.Matrix{5, 8}, true},
			{matrix.Matrix{-10, 5}, false},
			{matrix.Matrix{8, 5}, true},
			{matrix.Matrix{1, 2}, true},
			{matrix.Matrix{2, 1}, true},
		}},
	{"rc strange", // (there's a 0-width spit)
		matrix.LineMatrix{{0, 0}, {2.5, 2.5}, {0, 10}, {2.5, 7.5},
			{7.5, 7.5}, {10, 10}, {10, 0}, {2.5, 2.5}},
		[]args{
			{matrix.Matrix{5, 5}, true},
			{matrix.Matrix{5, 8}, false},
			{matrix.Matrix{-10, 5}, false},
			{matrix.Matrix{8, 5}, true},
			{matrix.Matrix{1, 2}, false},
			{matrix.Matrix{2, 1}, false},
		}},
	{"rc exagon", // (sic) "hexagon"
		matrix.LineMatrix{{3, 0}, {7, 0}, {10, 5}, {7, 10}, {3, 10}, {0, 5}},
		[]args{
			{matrix.Matrix{5, 5}, true},
			{matrix.Matrix{5, 8}, true},
			{matrix.Matrix{-10, 5}, false},
			{matrix.Matrix{8, 5}, true},
			{matrix.Matrix{10, 10}, false},
			{matrix.Matrix{1, 2}, false},
			{matrix.Matrix{2, 1}, false},
		}},

	// https://github.com/JamesMilnerUK/pip-go.
	{"jm rectangle",
		matrix.LineMatrix{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
		[]args{
			{matrix.Matrix{1.1, 1.1}, true},
			{matrix.Matrix{1.2, 1.2}, true},
			{matrix.Matrix{1.3, 1.3}, true},
			{matrix.Matrix{1.4, 1.4}, true},
			{matrix.Matrix{1.5, 1.5}, true},
			{matrix.Matrix{1.6, 1.6}, true},
			{matrix.Matrix{1.7, 1.7}, true},
			{matrix.Matrix{1.8, 1.8}, true},

			{matrix.Matrix{-4.9, 1.2}, false},
			{matrix.Matrix{10.0, 10.0}, false},
			{matrix.Matrix{-5.0, -6.0}, false},
			{matrix.Matrix{-13.0, 1.0}, false},
			{matrix.Matrix{4.9, -1.2}, false},
			{matrix.Matrix{10.0, -10.0}, false},
			{matrix.Matrix{5.0, 6.0}, false},
			{matrix.Matrix{-13.0, 1.0}, false},
		}},

	// https://github.com/substack/point-in-polygon
	{"ss box",
		matrix.LineMatrix{{1, 1}, {1, 2}, {2, 2}, {2, 1}},
		[]args{
			{matrix.Matrix{1.5, 1.5}, true},
			{matrix.Matrix{1.2, 1.9}, true},
			{matrix.Matrix{0, 1.9}, false},
			{matrix.Matrix{1.5, 2}, false},
			{matrix.Matrix{1.5, 2.2}, false},
			{matrix.Matrix{3, 5}, false},
		}},

	// https://github.com/sromku/polygon-contains-point
	{"sr simple",
		matrix.LineMatrix{{1, 3}, {2, 8}, {5, 4}, {5, 9}, {7, 5}, {6, 1}, {3, 1}},
		[]args{
			{matrix.Matrix{5.5, 7}, true},
			{matrix.Matrix{4.5, 7}, false},
		}},
	{"sr holes",
		matrix.LineMatrix{
			{1, 2}, {1, 6}, {8, 7}, {8, 1},
			{2, 3}, {5, 5}, {6, 2}, {8, 1},
			{6, 6}, {7, 6}, {7, 5}, {8, 1}},
		[]args{
			{matrix.Matrix{6, 5}, true},
			{matrix.Matrix{4, 3}, false},
			{matrix.Matrix{6.5, 5.8}, false},
		}},
}

func TestInPolygon(t *testing.T) {
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				if got := InPolygon(arg.point, tt.poly); got != arg.want {
					t.Errorf("InPolygon(%v , %v) = %v, want %v", arg.point, tt.poly, got, arg.want)
				}
			}
		})
	}
}
