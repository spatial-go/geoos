package overlay

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func Test_perform(t *testing.T) {
	//multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
	//expectPolygon, _ := wkt.UnmarshalString(`POLYGON((10 0,0 0,0 10,5 10,5 15,15 15,15 5,10 5,10 0))`)

	type args struct {
		subject  *Polygon
		clipping *Polygon
	}

	subject := &Polygon{}
	subject.AddPoint(&Point{matrix: matrix.Matrix{0, 0}})
	subject.AddPoint(&Point{matrix: matrix.Matrix{10, 0}})
	subject.AddPoint(&Point{matrix: matrix.Matrix{10, 10}})
	subject.AddPoint(&Point{matrix: matrix.Matrix{0, 10}})
	subject.CloseRing()
	subject.rank = MAIN

	clipping := &Polygon{}
	clipping.AddPoint(&Point{matrix: matrix.Matrix{5, 5}})
	clipping.AddPoint(&Point{matrix: matrix.Matrix{15, 5}})
	clipping.AddPoint(&Point{matrix: matrix.Matrix{15, 15}})
	clipping.AddPoint(&Point{matrix: matrix.Matrix{5, 15}})
	clipping.CloseRing()
	clipping.rank = CUT

	want := &Polygon{}
	want.AddPoint(&Point{matrix: matrix.Matrix{5, 10}})
	want.AddPoint(&Point{matrix: matrix.Matrix{10, 10}})
	want.AddPoint(&Point{matrix: matrix.Matrix{10, 5}})
	want.AddPoint(&Point{matrix: matrix.Matrix{5, 5}})
	want.CloseRing()

	tests := []struct {
		name string
		args args
		want *Polygon
	}{
		{name: "perform Polygon",
			args: args{
				subject: subject, clipping: clipping,
			},

			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Weiler(tt.args.subject, tt.args.clipping, Clip); !got.Equal(tt.want) {
				t.Errorf("perform() = %v, want %v", got.ToString(), tt.want.ToString())
			}
		})
	}
}
