package chain

import (
	"testing"

	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestSegmentMutualIntersector_Process(t *testing.T) {
	type fields struct {
		SegmentMutual matrix.LineMatrix
	}
	type args struct {
		segStrings []*matrix.LineSegment
		segInt     Intersector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   matrix.Collection
	}{
		{"test ", fields{matrix.LineMatrix{{50, 100}, {50, 200}}},
			args{matrix.LineMatrix{{50, 50}, {50, 150}}.ToLineArray(), &IntersectionCollinearDifference{}},
			matrix.Collection{matrix.LineMatrix{{50, 150}, {50, 200}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SegmentMutualIntersector{
				SegmentMutual: tt.fields.SegmentMutual,
			}
			s.Process(tt.args.segStrings, tt.args.segInt)
			if !tt.args.segInt.Result().(matrix.Collection).Equals(tt.want) {
				t.Errorf("Process %v = %v, want %v", tt.name, tt.args.segInt.Result(), tt.want)
			}
		})
	}
}
