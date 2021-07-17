package grid

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos/space"
)

func TestHexagonGrid(t *testing.T) {
	bound := space.Bound{Min: space.Point{1, 1}, Max: space.Point{1.5, 1.5}}
	wantGrids := [][]Grid{
		{
			Grid{
				space.Polygon{
					space.Ring{
						space.Point{1.2248646496725726, 1.3894176784497805},
						space.Point{1.4497292993451452, 1},
						space.Point{1.2248646496725726, 0.6105823215502195},
						space.Point{0.7751353503274274, 0.6105823215502195},
						space.Point{0.550270700654855, 1},
						space.Point{0.7751353503274274, 1.3894176784497805},
						space.Point{1.2248646496725726, 1.3894176784497805},
					},
				},
			},
			Grid{
				space.Polygon{
					space.Ring{
						space.Point{1.2248646496725726, 2.168253035349341},
						space.Point{1.4497292993451452, 1.778835356899561},
						space.Point{1.2248646496725726, 1.3894176784497805},
						space.Point{0.7751353503274274, 1.3894176784497805},
						space.Point{0.550270700654855, 1.778835356899561},
						space.Point{0.7751353503274274, 2.168253035349341},
						space.Point{1.2248646496725726, 2.168253035349341},
					},
				},
			},
		},
		{
			Grid{
				space.Polygon{
					space.Ring{
						space.Point{1.8994585986902903, 1.778835356899561},
						space.Point{2.1243232483628627, 1.3894176784497805},
						space.Point{1.8994585986902903, 1},
						space.Point{1.4497292993451452, 1},
						space.Point{1.2248646496725728, 1.3894176784497805},
						space.Point{1.4497292993451452, 1.778835356899561},
						space.Point{1.8994585986902903, 1.778835356899561},
					},
				},
			},
			Grid{
				space.Polygon{
					space.Ring{
						space.Point{1.8994585986902903, 2.5576707137991215},
						space.Point{2.1243232483628627, 2.168253035349341},
						space.Point{1.8994585986902903, 1.7788353568995607},
						space.Point{1.4497292993451452, 1.7788353568995607},
						space.Point{1.2248646496725728, 2.168253035349341},
						space.Point{1.4497292993451452, 2.5576707137991215},
						space.Point{1.8994585986902903, 2.5576707137991215},
					},
				},
			},
		},
	}
	var cellSize float64 = 50000
	type args struct {
		bound    space.Bound
		cellSize float64
	}
	tests := []struct {
		name          string
		args          args
		wantGridGeoms [][]Grid
	}{
		{name: "hexagonGrid", args: args{bound: bound, cellSize: cellSize}, wantGridGeoms: wantGrids},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGridGeoms := HexagonGrid(tt.args.bound, tt.args.cellSize); !reflect.DeepEqual(gotGridGeoms, tt.wantGridGeoms) {
				t.Errorf("HexagonGrid() = %v, want %v", gotGridGeoms, tt.wantGridGeoms)
			}
		})
	}
}
