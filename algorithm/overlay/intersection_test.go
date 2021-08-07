package overlay

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func Test_perform(t *testing.T) {
	//multiPolygon, _ := wkt.UnmarshalString(`MULTIPOLYGON(((0 0, 10 0, 10 10, 0 10, 0 0)), ((5 5, 15 5, 15 15, 5 15, 5 5)))`)
	//expectPolygon, _ := wkt.UnmarshalString(`POLYGON((10 0,0 0,0 10,5 10,5 15,15 15,15 5,10 5,10 0))`)

	type args struct {
		subject  matrix.PolygonMatrix
		clipping matrix.PolygonMatrix
	}

	subject := matrix.PolygonMatrix{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}}
	clipping := matrix.PolygonMatrix{{{5, 5}, {15, 5}, {15, 15}, {5, 15}, {5, 5}}}

	want := matrix.PolygonMatrix{{{5, 10}, {10, 10}, {10, 5}, {5, 5}, {5, 10}}}

	tests := []struct {
		name string
		args args
		want matrix.PolygonMatrix
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
			if got, err := Intersection(tt.args.subject, tt.args.clipping); err != nil || !got.Equals(tt.want) {
				t.Errorf("perform() = %v, want %v", got, tt.want)
			}
		})
	}
}
