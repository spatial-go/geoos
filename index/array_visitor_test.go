package index

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestLineSegmentVisitor_VisitItem(t *testing.T) {
	type fields struct {
		QuerySeg *matrix.LineSegment
		Items    []*matrix.LineSegment
	}
	type args struct {
		item interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: " visit item true", fields: fields{QuerySeg: &matrix.LineSegment{P0: matrix.Matrix{1, 1}, P1: matrix.Matrix{2, 2}}, Items: []*matrix.LineSegment{}},
			args: args{&matrix.LineSegment{P0: matrix.Matrix{1, 1}, P1: matrix.Matrix{2, 2}}}, want: true},
		{name: " visit item false", fields: fields{QuerySeg: &matrix.LineSegment{P0: matrix.Matrix{1, 1}, P1: matrix.Matrix{2, 2}}, Items: []*matrix.LineSegment{}},
			args: args{&matrix.LineSegment{P0: matrix.Matrix{1, 1}, P1: matrix.Matrix{3, 2}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LineSegmentVisitor{
				QuerySeg:          tt.fields.QuerySeg,
				ItemsArrayLineSeg: tt.fields.Items,
			}
			l.VisitItem(tt.args.item)
			has := false
			for _, v := range l.ItemsArrayLineSeg {
				if v.P0.Equals(l.QuerySeg.P0) && v.P1.Equals(l.QuerySeg.P1) {
					has = true

				}
			}
			if has != tt.want {
				t.Errorf("%v : %v ,want:%s", tt.name, l.Items(), tt.args)
			}
		})
	}
}
