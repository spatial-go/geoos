package grid

import (
	"reflect"
	"testing"

	"github.com/spatial-go/geoos"
)

func TestSquareGrid(t *testing.T) {
	bound := geoos.Bound{Min: geoos.Point{1, 1}, Max: geoos.Point{1.5, 1.5}}
	wantGrids := [][]Grid{
		{
			Grid{
				geoos.Polygon{
					geoos.Ring{
						geoos.Point{0.9801624203929129, 0.980203518223959},
						geoos.Point{0.9801624203929129, 1.25},
						geoos.Point{1.25, 1.25},
						geoos.Point{1.25, 0.980203518223959},
						geoos.Point{0.9801624203929129, 0.980203518223959},
					},
				},
			},
			Grid{
				geoos.Polygon{
					geoos.Ring{
						geoos.Point{0.9801624203929129, 1.25},
						geoos.Point{0.9801624203929129, 1.519796481776041},
						geoos.Point{1.25, 1.519796481776041},
						geoos.Point{1.25, 1.25},
						geoos.Point{0.9801624203929129, 1.25},
					},
				},
			},
		},
		{
			Grid{
				geoos.Polygon{
					geoos.Ring{
						geoos.Point{1.25, 0.980203518223959},
						geoos.Point{1.25, 1.25},
						geoos.Point{1.519837579607087, 1.25},
						geoos.Point{1.519837579607087, 0.980203518223959},
						geoos.Point{1.25, 0.980203518223959},
					},
				},
			},
			Grid{
				geoos.Polygon{
					geoos.Ring{
						geoos.Point{1.25, 1.25},
						geoos.Point{1.25, 1.519796481776041},
						geoos.Point{1.519837579607087, 1.519796481776041},
						geoos.Point{1.519837579607087, 1.25},
						geoos.Point{1.25, 1.25},
					},
				},
			},
		},
	}

	var cellSize float64 = 30000

	type args struct {
		bound    geoos.Bound
		cellSize float64
	}
	tests := []struct {
		name      string
		args      args
		wantGrids [][]Grid
	}{
		{name: "squareGrid", args: args{bound: bound, cellSize: cellSize}, wantGrids: wantGrids},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGrids := SquareGrid(tt.args.bound, tt.args.cellSize)
			if !reflect.DeepEqual(gotGrids, tt.wantGrids) {
				t.Errorf("SquareGrid() gotGrids = %v, want %v", gotGrids, tt.wantGrids)
			}
		})
	}
}
